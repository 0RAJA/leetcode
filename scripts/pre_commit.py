#!/usr/bin/env python3
"""Validate staged catalog inputs and stage derived files without reading unstaged inputs."""

import argparse
from collections import Counter
from dataclasses import dataclass
import importlib.util
import io
from pathlib import Path
import os
import re
import stat
import subprocess
import sys
import tempfile


GENERATED = ("docs/CATALOG.md", "docs/explorer/data.js")
ROUTE = "2026/README.md"
GENERATOR = "scripts/build_catalog.py"


def git(root, *args, input=None, check=True):
    """Run Git with the hook's index environment, including alternate commit indexes."""
    return subprocess.run(["git", "-C", str(root), *args], input=input,
                          capture_output=True, check=check)


def entries(root, revision=None):
    """Return commit-visible blobs, excluding intent-to-add entries; reject conflicts."""
    if revision:
        result = git(root, "ls-tree", "-rz", revision, check=False)
        if result.returncode:
            return {}
    else:
        result = git(root, "ls-files", "--stage", "-z")
    found = {}
    for record in result.stdout.split(b"\0"):
        if not record:
            continue
        info, path = record.split(b"\t", 1)
        mode, middle, last = info.decode().split()
        if revision:
            blob = last
        else:
            blob = middle
            if last != "0":
                raise ValueError("存在未解决的合并冲突，请先解决并 git add。")
        found[os.fsdecode(path)] = (mode, blob)
    if not revision:
        # ls-files also lists git add -N placeholders. An actual new empty file
        # appears in the cached diff; a placeholder does not enter the commit.
        visible = {os.fsdecode(path) for path in git(
            root, "diff", "--cached", "--name-only", "--no-renames",
            "--ita-invisible-in-index", "-z").stdout.split(b"\0") if path}
        visible.update(entries(root, "HEAD"))
        found = {path: value for path, value in found.items() if path in visible}
    return found


def is_input(path):
    """Identify only the versioned inputs that the catalog generator consumes."""
    return path in (ROUTE, GENERATOR, "catalog.config.json", "leetcode/editor/cn/all.json") or (
        path.endswith(".go") and path.startswith(("leetcode/", "2026/"))
    ) or (path.startswith("leetcode/editor/cn/doc/note/") and path.endswith(".md"))


def materialize(root, records, destination):
    """Read raw Git blobs into a snapshot; never follow working-tree symlinks or filters."""
    selected = [(path, mode, blob) for path, (mode, blob) in records.items()
                if is_input(path) or path in GENERATED]
    for path, mode, _ in selected:
        if mode not in ("100644", "100755"):
            raise ValueError(f"{path}: 索引输入必须是普通文件。")
    if not selected:
        return
    payload = "".join(blob + "\n" for _, _, blob in selected).encode()
    stream = io.BytesIO(git(root, "cat-file", "--batch", input=payload).stdout)
    for path, _, _ in selected:
        _, kind, size = stream.readline().split()
        if kind != b"blob":
            raise ValueError(f"{path}: 无法读取暂存文件。")
        content = stream.read(int(size))
        stream.read(1)
        target = destination / path
        target.parent.mkdir(parents=True, exist_ok=True)
        target.write_bytes(content)


def build(snapshot, paths):
    """Load the generator from the same snapshot as its inputs, including staged edits."""
    spec = importlib.util.spec_from_file_location("catalog_snapshot", snapshot / GENERATOR)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module, module.build_catalog(snapshot, paths)


@dataclass
class Diagnostic:
    level: str
    path: str
    line: int
    message: str
    token: str

    def key(self):
        """Ignore shifted line numbers when comparing existing issues with HEAD."""
        return self.level, self.path, self.message, self.token


def route_records(module, markdown):
    """Pair the generator's own note entries with their original line numbers."""
    starts = []
    syntax = []
    fence = None
    for line_number, line in enumerate(markdown.splitlines(), 1):
        if re.match(r"\s*```", line):
            fence = None if fence else line_number
        elif fence is None:
            if re.match(r"^\d+\.\s+.+$", line):
                starts.append(line_number)
            elif re.match(r"^\d+(?:[.．、)]|\s+\S)", line):
                syntax.append(Diagnostic("ERROR", ROUTE, line_number,
                                         "笔记须以“序号. 题号 题名”开头，点后保留空格。", line))
            elif re.match(r"^\s+\d+\.\s+(?:第)?\d+(?:[.、\s]|题)", line):
                syntax.append(Diagnostic("ERROR", ROUTE, line_number,
                                         "独立题目笔记须顶格编号，当前缩进会并入上一条笔记。", line))
    if fence is not None:
        syntax.append(Diagnostic("ERROR", ROUTE, fence, "代码块缺少结束的 ```。",
                                 "\n".join(markdown.splitlines()[fence - 1:])))
    notes = [(line, section, text) for line, (section, text)
             in zip(starts, module.note_entries(markdown))]
    return notes, syntax


def diagnose(snapshot, records, module, data):
    """Report actual attachment results, bad headers and intentional note placeholders."""
    results = []
    sources = {s["path"]: p for p in data["problems"] for s in p["sources"]}
    attached = {(n["path"], n["text"]): p for p in data["problems"] for n in p["notes"]}
    for path in records:
        if not is_input(path) or not path.endswith(".go") or path.endswith("_test.go"):
            continue
        if path in sources:
            problem = sources[path]
            results.append(Diagnostic("OK", path, 1,
                                      f"题解已收录：{problem['id']} {problem['title']}", records[path][1]))
        else:
            results.append(Diagnostic("ERROR", path, 1,
                                      "题解未收录：检查标准目录、all.json 内部题号或 sources 映射。", records[path][1]))

    if ROUTE in records:
        markdown = (snapshot / ROUTE).read_text(encoding="utf-8-sig")
        notes, syntax = route_records(module, markdown)
        results.extend(syntax)
        for line, _, text in notes:
            if not module.has_note_content(text):
                results.append(Diagnostic("WARN", ROUTE, line,
                                          "未收录笔记：仅标题或 TODO，补充正文后会自动收录。", text))
                continue
            problem = attached.get((ROUTE, text))
            if not problem:
                results.append(Diagnostic("ERROR", ROUTE, line,
                                          "笔记未收录：题号/题名无法关联已有题解，请检查或补充 sources 映射。", text))
                continue
            heading = module.entry_heading(text)
            number = re.match(r"^(?:第)?(\d+)(?:[.、\s]|题)", heading)
            title = re.sub(r"^(?:第)?\d+(?:题)?[.、\s]*", "", heading)
            if number and problem["id"].isdigit() and not module.normalized(title).startswith(
                    module.normalized(problem["title"])):
                results.append(Diagnostic("ERROR", ROUTE, line,
                                          f"题号与题名不一致：{problem['id']} 应为“{problem['title']}”。", text))
            else:
                results.append(Diagnostic("OK", ROUTE, line,
                                          f"笔记已收录：{problem['id']} {problem['title']}", text))

    for path in records:
        if not (path.startswith("leetcode/editor/cn/doc/note/") and path.endswith(".md")):
            continue
        text = (snapshot / path).read_text(encoding="utf-8-sig").strip()
        problem = attached.get((path, text))
        # Older HEAD generators accepted any nonempty standalone text. Keep the
        # historical comparison faithful while using the staged generator's rule.
        written = getattr(module, "has_standalone_note_content", bool)(text)
        if not written:
            level, message = "WARN", "未收录笔记：文件为空或仅有 TODO。"
        elif not problem:
            level, message = "ERROR", "笔记未收录：文件名须包含可识别的题号、题名，并有对应题解。"
        elif problem["id"].isdigit() and not any(
                module.normalized(Path(path).stem).startswith(module.normalized(prefix))
                for prefix in (problem["id"] + problem["title"], problem["title"])):
            level, message = "ERROR", f"文件名题号与题名不一致，应为“{problem['id']} {problem['title']}”。"
        else:
            level, message = "OK", f"笔记已收录：{problem['id']} {problem['title']}"
        results.append(Diagnostic(level, path, 1, message, text))
    return results


def install(root):
    """Enable the versioned hooks without disabling a user's existing active hooks."""
    hook = root / ".githooks/pre-commit"
    if not hook.is_file():
        raise ValueError("缺少 .githooks/pre-commit，请先获取完整的 hook 文件。")
    current = git(root, "config", "--get", "core.hooksPath", check=False).stdout.decode().strip()
    if current and (root / current).resolve() != (root / ".githooks").resolve():
        raise ValueError(f"已配置 core.hooksPath={current}，请手工合并现有 hook，不会覆盖。")
    if not current:
        default = Path(os.fsdecode(git(root, "rev-parse", "--git-path", "hooks").stdout).strip())
        if not default.is_absolute():
            default = root / default
        active = [p.name for p in default.iterdir() if not p.name.endswith(".sample")
                  and p.is_file() and os.access(p, os.X_OK)] if default.is_dir() else []
        if active:
            raise ValueError(f"已有活动 hook：{', '.join(active)}，请先手工合并，不会覆盖。")
    hook.chmod(hook.stat().st_mode | stat.S_IXUSR | stat.S_IXGRP | stat.S_IXOTH)
    git(root, "config", "--local", "core.hooksPath", ".githooks")
    print("已启用本仓库的 .githooks/pre-commit；后续 git commit 自动校验并生成索引。")


def update_outputs(root, snapshot, staged, generated, check_only):
    """Protect unstaged output edits, then write and stage all derived changes together."""
    changed = []
    before = {}
    for path, content in generated.items():
        target = root / path
        expected = (snapshot / path).read_bytes() if path in staged else None
        if target.is_symlink() or (target.exists() and not target.is_file()):
            raise ValueError(f"{path}: 生成文件必须是普通文件。")
        actual = target.read_bytes() if target.exists() else None
        if actual != expected:
            raise ValueError(f"{path}: 存在未暂存的生成文件改动，请先保存并处理；不会覆盖。")
        payload = content.encode("utf-8")
        if payload != expected:
            changed.append(path)
            before[path] = actual
    if not changed:
        print("生成索引已与暂存区一致。")
        return
    if check_only:
        print("提交时将自动生成并暂存：" + "、".join(changed))
        return
    if os.environ.get("GIT_INDEX_FILE"):
        env = dict(os.environ)
        env.pop("GIT_INDEX_FILE")
        default = subprocess.run(["git", "-C", str(root), "rev-parse", "--git-path", "index"],
                                 env=env, capture_output=True, check=True)
        default_path = root / os.fsdecode(default.stdout).strip()
        if (root / os.environ["GIT_INDEX_FILE"]).resolve() != default_path.resolve():
            raise ValueError("当前提交使用临时暂存区（如 git commit --only/指定文件名），"
                             "无法安全同步生成索引。请先 git add 所需文件，再使用普通 git commit。")
    if entries(root) != staged:
        raise ValueError("检查期间暂存区发生变化，请重新运行 git commit。")
    try:
        for path in changed:
            target = root / path
            target.parent.mkdir(parents=True, exist_ok=True)
            target.write_text(generated[path], encoding="utf-8", newline="\n")
        git(root, "add", "--", *changed)
    except (OSError, subprocess.CalledProcessError):
        for path, content in before.items():
            target = root / path
            if content is None:
                target.unlink(missing_ok=True)
            else:
                target.write_bytes(content)
        raise
    print("已生成并暂存：" + "、".join(changed))


def run(root, check_only):
    """Validate the active index against HEAD, ignoring unchanged historical issues."""
    staged = entries(root)
    head = entries(root, "HEAD")
    changed = {p for p in staged.keys() | head.keys() if staged.get(p) != head.get(p)}
    if not check_only and not any(is_input(p) or p in GENERATED or p in (
            "scripts/pre_commit.py", ".githooks/pre-commit") for p in changed):
        print("无算法手记相关暂存改动，跳过索引生成。")
        return 0
    with tempfile.TemporaryDirectory(prefix="leetcode-catalog-") as directory:
        snapshot = Path(directory) / "index"
        snapshot.mkdir()
        materialize(root, staged, snapshot)
        module, data = build(snapshot, list(staged))
        diagnostics = diagnose(snapshot, staged, module, data)
        previous = []
        if head:
            baseline = Path(directory) / "head"
            baseline.mkdir()
            materialize(root, head, baseline)
            old_module, old_data = build(baseline, list(head))
            previous = diagnose(baseline, head, old_module, old_data)
        remaining = Counter(d.key() for d in previous)
        new = []
        for diagnostic in diagnostics:
            if remaining[diagnostic.key()]:
                remaining[diagnostic.key()] -= 1
            else:
                new.append(diagnostic)
        for item in new:
            print(f"[{item.level}] {item.path}:{item.line} {item.message}")
        if any(item.level == "ERROR" for item in new):
            print("索引校验失败，修正后请重新 git add 再提交。", file=sys.stderr)
            return 1
        old_errors = sum(item.level == "ERROR" for item in diagnostics) - sum(
            item.level == "ERROR" for item in new)
        note_count = sum(len(p["notes"]) for p in data["problems"])
        print(f"索引校验通过：共 {len(data['problems'])} 题、{note_count} 条笔记；"
              f"历史未收录/格式问题 {old_errors} 项（本次未新增）。")
        update_outputs(root, snapshot, staged, module.outputs(data), check_only)
    return 0


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    mode = parser.add_mutually_exclusive_group()
    mode.add_argument("--install", action="store_true", help="enable this repository's hooks safely")
    mode.add_argument("--check", action="store_true", help="preview staged validation without changing files or index")
    args = parser.parse_args()
    try:
        root = Path(os.fsdecode(git(Path.cwd(), "rev-parse", "--show-toplevel").stdout).strip())
        if args.install:
            install(root)
            return 0
        return run(root, args.check)
    except subprocess.CalledProcessError as error:
        print("算法手记校验失败：" + os.fsdecode(error.stderr).strip(), file=sys.stderr)
    except (OSError, ValueError, KeyError, SyntaxError, ImportError) as error:
        print(f"算法手记校验失败：{error}", file=sys.stderr)
    return 1


if __name__ == "__main__":
    sys.exit(main())
