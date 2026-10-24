#!/usr/bin/env python3
"""Build the review catalog from versioned solutions; Python standard library only."""

import argparse
import collections
import json
from pathlib import Path, PurePosixPath
import re
import subprocess
import sys
import textwrap
from urllib.parse import quote


ROOT = Path(__file__).resolve().parents[1]
LEVELS = {1: "简单", 2: "中等", 3: "困难"}


def normalized(value):
    return re.sub(r"[\W_]+", "", value.casefold())


def read(root, path):
    return (root / path).read_text(encoding="utf-8-sig")


def read_json(root, path):
    """Read catalog metadata, retaining the source path and position on JSON errors."""
    try:
        return json.loads(read(root, path))
    except json.JSONDecodeError as error:
        raise ValueError(f"{path}:{error.lineno}:{error.colno} JSON 格式错误：{error.msg}") from error


def tracked_paths(root):
    # Only inspect versioned files: ignored caches, cookies and IDE state stay out.
    result = subprocess.run(
        ["git", "ls-files", "-z"], cwd=root, check=True, capture_output=True
    )
    return sorted(set(result.stdout.decode("utf-8").split("\0")) - {""})


def metadata_indexes(records):
    internal = {str(row["questionId"]): row for row in records}
    frontend = {str(row["frontendQuestionId"]): row for row in records}
    names = {}
    titles = collections.defaultdict(list)
    for row in records:
        key = str(row["frontendQuestionId"])
        names[normalized(key + row["title"])] = row
        titles[normalized(row["title"])].append(row)
    return internal, frontend, names, titles


def filename_identity(path, frontend, names, titles):
    stem = PurePosixPath(path).stem.strip()
    exact = names.get(normalized(stem))
    if exact:
        return str(exact["frontendQuestionId"]), exact["title"], exact
    match = re.match(r"^\[([^\]]+)\]\s*(.*)$", stem)
    if not match:
        match = re.match(
            r"^(剑指 Offer II \d+|剑指 Offer \d+(?:\s*-\s*[IV]+)?|面试题 \d+\.\d+|LCR \d+|LCP \d+)(.*)$",
            stem,
        )
    if not match:
        match = re.match(r"^(\d+)[\].\s]*(.*)$", stem)
    if match:
        identity, title = match.groups()
        record = frontend.get(identity.strip())
        return identity.strip(), (record or {}).get("title", title.strip() or stem), record
    title_matches = titles.get(normalized(stem), [])
    if len(title_matches) == 1:
        record = title_matches[0]
        return str(record["frontendQuestionId"]), record["title"], record
    parent = PurePosixPath(path).parent.name
    if parent.isdigit() and parent in frontend:
        record = frontend[parent]
        return parent, record["title"], record
    return None


def extract_topics(source):
    match = re.search(r"Related Topics\s+(.+?)(?:👍|👎|\r?$)", source, re.MULTILINE)
    if not match:
        return []
    topics = match.group(1).strip()
    # The editor separates tags with spaces, including some multiword tag names.
    for compound in ("0-1 背包", "树形 DP", "Floyd 判圈算法", "Floyd 算法", "Brute-Force Search", "Bellman–Ford 算法"):
        topics = topics.replace(compound, compound.replace(" ", "\0"))
    return list(dict.fromkeys(topic.replace("\0", " ") for topic in topics.split()))


def extract_code(source):
    # Drop the statement comment before package, preserving imports and solution comments.
    match = re.search(r"^package\s+\w+", source, re.MULTILINE)
    code = source[match.start():] if match else source
    code = re.sub(r"^//\s*leetcode (?:submit|end).*\n?", "", code, flags=re.MULTILINE)
    return code.strip()


def note_entries(markdown):
    """Read top-level numbered notes, leaving nested lists/code blocks intact."""
    section = "刷题笔记"
    lines = []
    entry_section = section
    fenced = False
    for line in markdown.splitlines() + ["# END"]:
        if re.match(r"\s*```", line):
            fenced = not fenced
        heading = re.match(r"^#{1,3}\s+(.+)$", line) if not fenced else None
        numbered = re.match(r"^\d+\.\s+(.+)$", line) if not fenced else None
        if heading or numbered:
            if lines:
                yield entry_section, textwrap.dedent("\n".join(lines)).strip()
            lines = []
        if heading:
            section = heading.group(1)
        elif numbered:
            entry_section = section
            lines = [numbered.group(1)]
        elif lines:
            lines.append(line)


def entry_heading(text):
    first = text.splitlines()[0]
    # Some original notes use "数组：704.二分查找" or "第454题.四数相加II".
    return re.sub(r"^(?:数组|链表|哈希表|字符串|栈与队列|二叉树|回溯|贪心算法|单调栈)[：:]\s*", "", first)


def entry_identity(text, frontend, titles):
    first = entry_heading(text)
    number = re.match(r"^(?:第)?(\d+)(?:[.、\s]|题)", first)
    if number:
        return number.group(1)
    beginning = normalized(re.split(r"[：:]", first, maxsplit=1)[0])
    matches = []
    for title, rows in titles.items():
        if beginning.startswith(title) and len(rows) == 1:
            matches.append((len(title), str(rows[0]["frontendQuestionId"])))
    return max(matches, default=(0, None))[1]


def configured_note_ids(text, problems, identities):
    heading = re.sub(r"^第(\d+)题", r"\1", entry_heading(text))
    beginning = normalized(re.split(r"[：:]", heading, maxsplit=1)[0])
    matches = []
    for identity in identities:
        if identity not in problems:
            continue
        title = problems[identity]["title"]
        number = re.search(r"\d+$", identity)
        prefixes = [title, identity + title]
        if number:
            prefixes.append(number.group() + title)
        for prefix in prefixes:
            key = normalized(prefix)
            if key and beginning.startswith(key):
                matches.append((len(key), identity))
    longest = max((length for length, _ in matches), default=0)
    return {identity for length, identity in matches if length == longest}


def is_todo_only(text):
    """Recognize TODO-only placeholders, including comment and Markdown prefixes."""
    lines = [line.strip() for line in text.splitlines() if line.strip()]
    return bool(lines) and all(re.fullmatch(
        r"(?:(?://|[#>*-])\s*)?TODO[.!。！]?", line, re.IGNORECASE
    ) for line in lines)


def has_standalone_note_content(text):
    """A standalone note needs nonempty text beyond a TODO placeholder."""
    return bool(text.strip()) and not is_todo_only(text)


def has_note_content(text):
    # A bare numbered title is an entry in a route, not a written explanation.
    lines = text.splitlines()
    if not lines:
        return False
    lines[0] = entry_heading(text)
    if re.search(r"[：:]\s*TODO\s*$", "\n".join(lines), re.IGNORECASE):
        return False
    inline = re.search(r"[：:]\s*(.*)$", lines[0])
    if is_todo_only("\n".join(lines[1:])) and (
            not inline or not inline.group(1).strip() or is_todo_only(inline.group(1))):
        return False
    return len(lines) > 1 or bool(re.search(r"[：:]\s*\S", lines[0]))


def build_catalog(root, paths=None):
    paths = tracked_paths(root) if paths is None else sorted(set(paths))
    paths = [path for path in paths if (root / path).is_file()]
    path_set = set(paths)
    config = read_json(root, "catalog.config.json")
    records = read_json(root, "leetcode/editor/cn/all.json")
    internal, frontend, names, titles = metadata_indexes(records)
    problems = {}
    unindexed = []
    for path in paths:
        if not path.endswith(".go") or path.endswith("_test.go"):
            continue
        current = re.fullmatch(r"leetcode/editor/cn/p(\d+)/solution\.go", path)
        archived = "/_legacy/" in path and path.startswith("leetcode/")
        custom = config.get("sources", {}).get(path)
        if not (current or archived or custom):
            continue
        source = read(root, path)
        if custom:
            identity = custom["id"]
            record = frontend.get(identity)
            title = (record or {}).get("title", custom["title"])
        elif current:
            record = internal.get(str(int(current.group(1))))
            if not record:
                # Do not silently pretend internal IDs are frontend IDs.
                unindexed.append(path)
                continue
            identity, title = str(record["frontendQuestionId"]), record["title"]
        else:
            found = filename_identity(path, frontend, names, titles)
            if not found:
                unindexed.append(path)
                continue
            identity, title, record = found
        slug = (record or {}).get("titleSlug")
        url = "https://leetcode.cn/problems/" + slug + "/" if slug else ""
        if custom:
            url = custom.get("url", url)
        problem = problems.setdefault(identity, {
            "id": identity, "title": title,
            "difficulty": LEVELS.get((record or {}).get("level"), "未标注"),
            "topics": [], "url": url, "active": False, "sources": [], "notes": [],
        })
        kind = "archive" if archived else "current"
        problem["active"] = problem["active"] or kind == "current"
        topics = extract_topics(source) + (custom or {}).get("topics", [])
        problem["topics"] = sorted(set(problem["topics"] + topics))
        parent = str(PurePosixPath(path).parent) + "/"
        tests = sorted(p for p in path_set if p.startswith(parent)
                       and "/" not in p[len(parent):] and p.endswith("_test.go"))
        # In legacy/root each file is an independent old exercise, not a Go package.
        if archived:
            tests = [p for p in tests if p == path[:-3] + "_test.go"]
        problem["sources"].append({
            "path": path, "kind": kind, "code": extract_code(source), "tests": tests,
        })

    for identity, patch in config.get("overrides", {}).items():
        if identity in problems:
            for key in ("title", "topics", "url", "difficulty"):
                if key in patch:
                    problems[identity][key] = patch[key]

    route = "2026/README.md"
    configured_ids = {source["id"] for source in config.get("sources", {}).values()}
    if route in path_set:
        for section, text in note_entries(read(root, route)):
            # Custom platforms can reuse a LeetCode number. Match their configured
            # titles before numeric fallback; new mappings need no code changes.
            local_ids = configured_note_ids(text, problems, configured_ids)
            if local_ids:
                if len(local_ids) == 1 and has_note_content(text):
                    problems[next(iter(local_ids))]["notes"].append({
                        "title": section, "text": text, "path": route,
                        "anchor": section.lower().replace(" ", "-"),
                    })
                continue
            identity = entry_identity(text, frontend, titles)
            if identity in problems and has_note_content(text):
                problems[identity]["notes"].append({
                    "title": section, "text": text, "path": route,
                    "anchor": section.lower().replace(" ", "-"),
                })

    for path in paths:
        if not path.startswith("leetcode/editor/cn/doc/note/") or not path.endswith(".md"):
            continue
        text = read(root, path).strip()
        local_ids = configured_note_ids(PurePosixPath(path).stem, problems, configured_ids)
        if local_ids:
            if has_standalone_note_content(text) and len(local_ids) == 1:
                problems[next(iter(local_ids))]["notes"].append({
                    "title": "独立笔记", "text": text, "path": path, "anchor": "",
                })
            continue
        found = filename_identity(path, frontend, names, titles)
        if has_standalone_note_content(text) and found and found[0] in problems:
            problems[found[0]]["notes"].append({
                "title": "独立笔记", "text": text, "path": path, "anchor": "",
            })

    rows = sorted(problems.values(), key=lambda p: sort_id(p["id"]))
    for row in rows:
        row["sources"].sort(key=lambda s: (s["kind"] != "current", s["path"]))
        if not row["topics"]:
            row["topics"] = ["未分类"]
    return {"schemaVersion": 1, "repository": config["repository"],
            "branch": config["branch"], "problems": rows,
            "unindexed": sorted(unindexed)}


def catalog_stats(data):
    rows = data["problems"]
    return {
        "total": len(rows), "active": sum(p["active"] for p in rows),
        "archived": sum(not p["active"] for p in rows),
        "withNotes": sum(bool(p["notes"]) for p in rows),
        "withTests": sum(any(s["tests"] for s in p["sources"]) for p in rows),
    }


def sort_id(identity):
    return (0, int(identity)) if identity.isdigit() else (1, identity)


def file_link(data, path, anchor=""):
    return (data["repository"] + "/blob/" + quote(data["branch"], safe="") + "/"
            + quote(path, safe="/") + ("#" + quote(anchor) if anchor else ""))


def markdown_row(data, problem):
    p = problem
    title = p["title"].replace("|", "\\|")
    title = f'[{title}]({p["url"]})' if p["url"] else title
    sources = [f'[{"题解" if s["kind"] == "current" else "历史"}]({file_link(data, s["path"])})'
               for s in p["sources"]]
    notes = [f'[{n["title"]}]({file_link(data, n["path"], n["anchor"])})' for n in p["notes"]]
    tests = sorted({t for s in p["sources"] for t in s["tests"]})
    tests = [f'[测试]({file_link(data, t)})' for t in tests]
    return f'| {p["id"]} | {title} | {p["difficulty"]} | {"、".join(p["topics"])} | {" · ".join(sources)} | {" · ".join(notes) or "—"} | {" · ".join(tests) or "—"} |'


def catalog_markdown(data):
    lines = ["# 题解索引", "", "> 自动从已跟踪的题解和笔记生成，请勿直接编辑此文件。", "",
             "题号按页面展示编号映射；同题多版本合并。收录表示有本地代码记录，不等于 LeetCode 已通过。测试链接表示存在测试文件。", "",
             "[导航使用说明](catalog-guide.md) · [维护指南](maintenance.md) · [原始复习笔记](../2026/README.md)", ""]
    for active, heading in [(True, "当前维护"), (False, "历史归档")]:
        lines += [f"## {heading}", "", "| 题号 | 题目 | 难度 | 题型 | 代码 | 笔记 | 测试 |",
                  "| --- | --- | --- | --- | --- | --- | --- |"]
        lines += [markdown_row(data, p) for p in data["problems"] if p["active"] == active]
        lines.append("")
    if data["unindexed"]:
        lines += ["## 待补映射", "", "以下记录保留在仓库中，暂未能可靠识别题目。可在 catalog.config.json 中补充映射。", ""]
        lines += [f'- [{path}]({file_link(data, path)})' for path in data["unindexed"]]
    return "\n".join(lines) + "\n"


def outputs(data):
    return {
        "docs/explorer/data.js": "// Generated by scripts/build_catalog.py; do not edit.\nwindow.LEETCODE_DATA = " + json.dumps(data, ensure_ascii=False, separators=(",", ":")).replace("<", "\\u003c") + ";\n",
        "docs/CATALOG.md": catalog_markdown(data),
    }


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--check", action="store_true", help="fail when generated files are stale")
    args = parser.parse_args()
    data = build_catalog(ROOT)
    stale = []
    for path, content in outputs(data).items():
        target = ROOT / path
        if not target.exists() or target.read_text(encoding="utf-8") != content:
            stale.append(path)
            if not args.check:
                target.parent.mkdir(parents=True, exist_ok=True)
                target.write_text(content, encoding="utf-8", newline="\n")
    print(json.dumps({"stats": catalog_stats(data), "unindexed": data["unindexed"], "changed": stale}, ensure_ascii=False))
    return int(args.check and bool(stale))


if __name__ == "__main__":
    sys.exit(main())
