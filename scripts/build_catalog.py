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
START = "<!-- catalog:start -->"
END = "<!-- catalog:end -->"


def normalized(value):
    return re.sub(r"[\W_]+", "", value.casefold())


def read(root, path):
    return (root / path).read_text(encoding="utf-8-sig")


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


def entry_identity(text, frontend, titles):
    first = text.splitlines()[0]
    # Some original notes use "数组：704.二分查找" or "第454题.四数相加II".
    first = re.sub(r"^(?:数组|链表|哈希表|字符串|栈与队列|二叉树|回溯|贪心算法|单调栈)[：:]\s*", "", first)
    number = re.match(r"^(?:第)?(\d+)(?:[.、\s]|题)", first)
    if number:
        return number.group(1)
    beginning = normalized(re.split(r"[：:]", first, maxsplit=1)[0])
    matches = []
    for title, rows in titles.items():
        if beginning.startswith(title) and len(rows) == 1:
            matches.append((len(title), str(rows[0]["frontendQuestionId"])))
    return max(matches, default=(0, None))[1]


def has_note_content(text):
    # A bare numbered title is an entry in a route, not a written explanation.
    if re.search(r"[：:]\s*TODO\s*$", text, re.IGNORECASE):
        return False
    return len(text.splitlines()) > 1 or bool(re.search(r"[：:]\s*\S", text.splitlines()[0]))


def build_catalog(root, paths=None):
    paths = tracked_paths(root) if paths is None else sorted(set(paths))
    path_set = set(paths)
    config = json.loads(read(root, "catalog.config.json"))
    records = json.loads(read(root, "leetcode/editor/cn/all.json"))
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
    if route in path_set:
        for section, text in note_entries(read(root, route)):
            local_id = next((identity for identity, word in [
                ("卡码网 58", "区间和"), ("卡码网 44", "开发商购买土地")
            ] if word in text.splitlines()[0]), None)
            if local_id:
                if local_id in problems and has_note_content(text):
                    problems[local_id]["notes"].append({
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
        found = filename_identity(path, frontend, names, titles)
        if text and found and found[0] in problems:
            problems[found[0]]["notes"].append({
                "title": "独立笔记", "text": text, "path": path, "anchor": "",
            })

    rows = sorted(problems.values(), key=lambda p: sort_id(p["id"]))
    for row in rows:
        row["sources"].sort(key=lambda s: (s["kind"] != "current", s["path"]))
        if not row["topics"]:
            row["topics"] = ["未分类"]
    stats = {
        "total": len(rows), "active": sum(p["active"] for p in rows),
        "archived": sum(not p["active"] for p in rows),
        "withNotes": sum(bool(p["notes"]) for p in rows),
        "withTests": sum(any(s["tests"] for s in p["sources"]) for p in rows),
    }
    return {"schemaVersion": 1, "repository": config["repository"],
            "branch": config["branch"], "problems": rows, "stats": stats,
            "unindexed": sorted(unindexed)}


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
             "[交互导航的打开方式](catalog-guide.md) · [原始复习笔记](../2026/README.md)", ""]
    for active, heading in [(True, "当前维护"), (False, "历史归档")]:
        lines += [f"## {heading}", "", "| 题号 | 题目 | 难度 | 题型 | 代码 | 笔记 | 测试 |",
                  "| --- | --- | --- | --- | --- | --- | --- |"]
        lines += [markdown_row(data, p) for p in data["problems"] if p["active"] == active]
        lines.append("")
    if data["unindexed"]:
        lines += ["## 待补映射", "", "以下记录保留在仓库中，暂未能可靠识别题目。可在 catalog.config.json 中补充映射。", ""]
        lines += [f'- [{path}]({file_link(data, path)})' for path in data["unindexed"]]
    return "\n".join(lines) + "\n"


def readme_summary(data):
    stats = data["stats"]
    active = [p for p in data["problems"] if p["active"]]
    topics = collections.Counter(t for p in active for t in p["topics"])
    lines = [START, "", "![题解概览](docs/assets/catalog.svg)", "",
             f'当前维护 **{stats["active"]}** 题 · 历史归档 **{stats["archived"]}** 题 · 合并去重 **{stats["total"]}** 题', "",
             f'其中 **{stats["withNotes"]}** 题关联了思路笔记，**{stats["withTests"]}** 题有测试文件。', "",
             "| 当前题型 | 收录题数 | 当前题型 | 收录题数 |",
             "| --- | ---: | --- | ---: |"]
    frequent = topics.most_common(12)
    for i in range(0, len(frequent), 2):
        pair = frequent[i:i + 2]
        if len(pair) == 1:
            pair.append(("—", "—"))
        lines.append(f"| {pair[0][0]} | {pair[0][1]} | {pair[1][0]} | {pair[1][1]} |")
    lines += ["", "同题多版本只统计一次，题型可以重叠。以上为仓库代码记录，未读取或推断在线 AC 状态。", "", END]
    return "\n".join(lines)


def overview_svg(data):
    s = data["stats"]
    # Native SVG text and shapes keep the repository self contained.
    return f'''<svg xmlns="http://www.w3.org/2000/svg" width="1000" height="224" viewBox="0 0 1000 224" role="img" aria-labelledby="title desc">
<title id="title">LeetCode 题解与复习索引</title><desc id="desc">当前维护 {s['active']} 题，历史归档 {s['archived']} 题，关联笔记 {s['withNotes']} 题。</desc>
<rect width="1000" height="224" rx="12" fill="#f2f1e9"/><rect x="0" y="0" width="10" height="224" fill="#657347"/>
<g font-family="Georgia, 'Noto Serif SC', serif" fill="#273321"><text x="42" y="48" font-size="15" letter-spacing="3">RAJA / ALGORITHM NOTES</text>
<text x="42" y="93" font-size="31">把做过的题，连成可回顾的思路。</text></g>
<g font-family="'Microsoft YaHei', sans-serif" fill="#56604c" font-size="15"><text x="44" y="137">当前维护</text><text x="275" y="137">历史归档</text><text x="507" y="137">关联笔记</text><text x="740" y="137">带测试文件</text></g>
<g font-family="Georgia, serif" fill="#283421" font-size="39"><text x="42" y="187">{s['active']}</text><text x="273" y="187">{s['archived']}</text><text x="505" y="187">{s['withNotes']}</text><text x="738" y="187">{s['withTests']}</text></g></svg>
'''


def outputs(root, data):
    summary = readme_summary(data)
    original = read(root, "README.md")
    if START not in original or END not in original:
        raise ValueError("README.md must contain catalog:start and catalog:end markers")
    prefix, rest = original.split(START, 1)
    _, suffix = rest.split(END, 1)
    return {
        "docs/explorer/data.js": "// Generated by scripts/build_catalog.py; do not edit.\nwindow.LEETCODE_DATA = " + json.dumps(data, ensure_ascii=False, separators=(",", ":")).replace("<", "\\u003c") + ";\n",
        "docs/CATALOG.md": catalog_markdown(data),
        "docs/assets/catalog.svg": overview_svg(data),
        "README.md": prefix + summary + suffix,
    }


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--check", action="store_true", help="fail when generated files are stale")
    args = parser.parse_args()
    data = build_catalog(ROOT)
    stale = []
    for path, content in outputs(ROOT, data).items():
        target = ROOT / path
        if not target.exists() or target.read_text(encoding="utf-8") != content:
            stale.append(path)
            if not args.check:
                target.parent.mkdir(parents=True, exist_ok=True)
                target.write_text(content, encoding="utf-8", newline="\n")
    print(json.dumps({"stats": data["stats"], "unindexed": data["unindexed"], "changed": stale}, ensure_ascii=False))
    return int(args.check and bool(stale))


if __name__ == "__main__":
    sys.exit(main())
