"""Regression checks for linking local exercises, metadata and review notes."""

import json
from pathlib import Path
import tempfile
import unittest

import build_catalog as catalog


def question(internal, frontend, title, level=2, slug="example"):
    return {
        "questionId": str(internal),
        "frontendQuestionId": str(frontend),
        "title": title,
        "titleSlug": slug,
        "level": level,
        # A stale editor snapshot must not become a claim of online acceptance.
        "status": "ac",
    }


METADATA = [
    question(768, 763, "划分字母区间"),
    question(792, 704, "二分查找", 1, "binary-search"),
    question(1000026, "面试题 16.25", "LRU 缓存"),
    question(342, 342, "4的幂", 1),
    question(58, 58, "最后一个单词的长度", 1),
]


class CatalogIdentityTests(unittest.TestCase):
    def test_compound_topic_names_stay_together(self):
        self.assertEqual(
            catalog.extract_topics("// Related Topics 数组 0-1 背包 树形 DP Floyd 算法 Brute-Force Search 👍 100"),
            ["数组", "0-1 背包", "树形 DP", "Floyd 算法", "Brute-Force Search"],
        )

    def test_todo_placeholder_is_not_a_written_explanation(self):
        self.assertFalse(catalog.has_note_content("459.重复的子字符串: TODO"))

    def test_topic_prefix_does_not_turn_a_bare_title_into_a_note(self):
        self.assertFalse(catalog.has_note_content("数组：704.二分查找"))
        self.assertFalse(catalog.has_note_content("数组：704.二分查找：TODO"))
        self.assertTrue(catalog.has_note_content("数组：704.二分查找：维护左闭右开区间"))

    def setUp(self):
        self.internal, self.frontend, self.names, self.titles = (
            catalog.metadata_indexes(METADATA)
        )

    def identify(self, path):
        return catalog.filename_identity(
            path, self.frontend, self.names, self.titles
        )

    def test_editor_internal_ids_are_not_display_ids(self):
        self.assertEqual(self.internal["768"]["frontendQuestionId"], "763")
        self.assertEqual(self.internal["792"]["frontendQuestionId"], "704")
        self.assertEqual(
            self.internal["1000026"]["frontendQuestionId"], "面试题 16.25"
        )
        self.assertNotIn("768", self.frontend)

    def test_numeric_title_is_not_consumed_as_part_of_problem_number(self):
        identity, title, record = self.identify(
            "leetcode/editor/cn/_legacy/root/3424的幂.go"
        )
        self.assertEqual((identity, title), ("342", "4的幂"))
        self.assertEqual(record["questionId"], "342")

    def test_bracket_ids_survive_missing_metadata(self):
        identity, title, record = self.identify(
            "leetcode/editor/cn/_legacy/root/[LCR 006]两数之和 II.go"
        )
        self.assertEqual((identity, title), ("LCR 006", "两数之和 II"))
        self.assertIsNone(record)

    def test_other_platform_p_prefix_is_not_a_leetcode_id(self):
        self.assertIsNone(
            self.identify("leetcode/editor/cn/_legacy/root/P1387.go")
        )

    def test_note_number_is_the_problem_id_after_the_list_ordinal(self):
        markdown = "# 数组\n\n1. 数组：704.二分查找: 维护区间\n2. 第763题 划分字母区间：合并区间\n"
        entries = list(catalog.note_entries(markdown))
        identities = [
            catalog.entry_identity(text, self.frontend, self.titles)
            for _, text in entries
        ]
        self.assertEqual(identities, ["704", "763"])
        self.assertEqual([section for section, _ in entries], ["数组", "数组"])

    def test_numbers_in_an_explanation_do_not_replace_the_problem_id(self):
        self.assertEqual(
            catalog.entry_identity("704 二分查找：2 个边界", self.frontend, self.titles),
            "704",
        )

    def test_code_fences_do_not_create_false_notes_or_sections(self):
        markdown = (
            "# 数组\n\n1. 704 二分查找\n"
            "   维护左右边界\n"
            "   ```text\n"
            "2. 763 这行只是代码内容\n"
            "# 这个标题也在代码里面\n"
            "   ```\n\n"
            "2. 763 划分字母区间\n"
            "   合并字符区间\n"
        )
        entries = list(catalog.note_entries(markdown))
        self.assertEqual(len(entries), 2)
        self.assertEqual([section for section, _ in entries], ["数组", "数组"])
        self.assertIn("# 这个标题也在代码里面", entries[0][1])
        self.assertIn("2. 763 这行只是代码内容", entries[0][1])
        self.assertTrue(entries[1][1].startswith("763 划分字母区间"))


class CatalogIntegrationTests(unittest.TestCase):
    def setUp(self):
        # Use the ignored workspace temp directory so restricted desktop runners
        # can create nested fixtures as well as CI's ordinary Python runtimes.
        temp_parent = catalog.ROOT / "tmp"
        temp_parent.mkdir(exist_ok=True)
        self.workspace = tempfile.TemporaryDirectory(
            dir=temp_parent, prefix="catalog-test-"
        )
        self.addCleanup(self.workspace.cleanup)
        self.root = Path(self.workspace.name)
        self.paths = []
        self.write("leetcode/editor/cn/all.json", json.dumps(METADATA))
        self.write(
            "catalog.config.json",
            json.dumps({
                "repository": "https://github.com/example/exercises",
                "branch": "master",
                "sources": {
                    "2026/intervalsum/solution.go": {
                        "id": "卡码网 58", "title": "区间和",
                        "topics": ["数组", "前缀和"],
                    },
                },
            }),
        )

    def write(self, name, content):
        path = self.root / name
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(content, encoding="utf-8")
        self.paths.append(name)

    def build(self):
        return catalog.build_catalog(self.root, self.paths)

    def test_deduplicates_current_and_archive_and_links_space_named_tests(self):
        self.write(
            "leetcode/editor/cn/p0792/solution.go",
            "// Related Topics 数组 二分查找 👍 10\npackage p0792\nfunc search() {}\n",
        )
        self.write(
            "leetcode/editor/cn/p0792/solution _test.go",
            "package p0792\n",
        )
        self.write(
            "leetcode/editor/cn/_legacy/root/[704]二分查找.go",
            "package main\nfunc searchOld() {}\n",
        )
        self.write(
            "leetcode/editor/cn/_legacy/root/[704]二分查找_test.go",
            "package main\n",
        )
        # An unrelated test sharing legacy/root is not this problem's test.
        self.write(
            "leetcode/editor/cn/_legacy/root/[763]划分字母区间_test.go",
            "package main\n",
        )
        self.write("2026/README.md", "# 数组\n1. 704 二分查找：维护左闭右开区间\n")
        data = self.build()
        self.assertEqual(catalog.catalog_stats(data)["total"], 1)
        self.assertEqual(catalog.catalog_stats(data)["active"], 1)
        self.assertEqual(catalog.catalog_stats(data)["archived"], 0)
        self.assertEqual(catalog.catalog_stats(data)["withTests"], 1)
        self.assertEqual(catalog.catalog_stats(data)["withNotes"], 1)
        problem = data["problems"][0]
        self.assertEqual(problem["id"], "704")
        self.assertEqual(len(problem["sources"]), 2)
        self.assertEqual(problem["sources"][0]["kind"], "current")
        self.assertEqual(problem["sources"][0]["tests"], [
            "leetcode/editor/cn/p0792/solution _test.go",
        ])
        self.assertEqual(problem["sources"][1]["tests"], [
            "leetcode/editor/cn/_legacy/root/[704]二分查找_test.go",
        ])
        self.assertEqual(problem["notes"][0]["anchor"], "数组")
        self.assertEqual(problem["url"], "https://leetcode.cn/problems/binary-search/")
        self.assertNotIn("status", problem)

    def test_kamacoder_note_58_does_not_attach_to_leetcode_58(self):
        self.write("2026/intervalsum/solution.go", "package intervalsum\n")
        self.write("leetcode/editor/cn/p0058/solution.go", "package p0058\n")
        self.write(
            "2026/README.md",
            "# 数组\n1. 58 区间和：计算前缀和，用区间差计算答案\n",
        )
        data = self.build()
        problems = {problem["id"]: problem for problem in data["problems"]}
        self.assertEqual(len(problems["卡码网 58"]["notes"]), 1)
        self.assertEqual(problems["58"]["notes"], [])
        self.assertIn("前缀和", problems["卡码网 58"]["topics"])

    def test_unknown_internal_id_is_reported_instead_of_fabricating_id(self):
        source = "leetcode/editor/cn/p9999/solution.go"
        self.write(source, "package p9999\n")
        data = self.build()
        self.assertEqual(data["problems"], [])
        self.assertEqual(data["unindexed"], [source])

    def test_new_platform_note_mapping_only_requires_config_not_generator_edits(self):
        source = "2026/another-platform/solution.go"
        config = json.loads((self.root / "catalog.config.json").read_text(encoding="utf-8"))
        config["sources"][source] = {"id": "其他平台 58", "title": "独立区间统计"}
        self.write("catalog.config.json", json.dumps(config))
        self.write(source, "package intervalstats\n")
        self.write("leetcode/editor/cn/p0058/solution.go", "package p0058\n")
        self.write("2026/README.md", "# 数组\n1. 58 独立区间统计：用前缀和计算\n")
        data = self.build()
        problems = {problem["id"]: problem for problem in data["problems"]}
        self.assertEqual(len(problems["其他平台 58"]["notes"]), 1)
        self.assertEqual(problems["58"]["notes"], [])

    def test_standalone_platform_note_uses_the_same_mapping_as_route_notes(self):
        self.write("2026/intervalsum/solution.go", "package intervalsum\n")
        self.write("leetcode/editor/cn/p0058/solution.go", "package p0058\n")
        self.write("leetcode/editor/cn/doc/note/58区间和.md", "利用前缀和计算区间和。\n")
        data = self.build()
        problems = {problem["id"]: problem for problem in data["problems"]}
        self.assertEqual(len(problems["卡码网 58"]["notes"]), 1)
        self.assertEqual(problems["58"]["notes"], [])

    def test_explicit_source_mapping_can_repair_a_missing_internal_id(self):
        source = "leetcode/editor/cn/p9999/solution.go"
        config_path = self.root / "catalog.config.json"
        config = json.loads(config_path.read_text(encoding="utf-8"))
        config["sources"][source] = {
            "id": "LCR 999", "title": "手工补充的题目",
        }
        self.write("catalog.config.json", json.dumps(config))
        self.write(source, "package p9999\n")
        data = self.build()
        self.assertEqual([p["id"] for p in data["problems"]], ["LCR 999"])
        self.assertEqual(data["unindexed"], [])

    def test_empty_standalone_notes_and_untracked_cache_are_not_published(self):
        self.write("leetcode/editor/cn/p0792/solution.go", "package p0792\n")
        self.write("leetcode/editor/cn/doc/note/704二分查找.md", "\n\n")
        cache = self.root / "leetcode/editor/cn/doc/content/private-cache.md"
        cache.parent.mkdir(parents=True, exist_ok=True)
        cache.write_text("DO NOT PUBLISH THIS CACHE", encoding="utf-8")
        data = self.build()
        self.assertEqual(catalog.catalog_stats(data)["withNotes"], 0)
        self.assertNotIn("DO NOT PUBLISH THIS CACHE", json.dumps(data))

    def test_statistics_follow_records_without_storing_duplicate_totals(self):
        self.write("leetcode/editor/cn/p0792/solution.go", "package p0792\n")
        data = self.build()
        self.assertNotIn("stats", data)
        self.assertEqual(catalog.catalog_stats(data)["total"], 1)
        self.write("leetcode/editor/cn/_legacy/root/[763]划分字母区间.go", "package main\n")
        updated = self.build()
        self.assertEqual(catalog.catalog_stats(updated)["total"], 2)
        self.assertEqual(catalog.catalog_stats(updated)["archived"], 1)
        # A deleted tracked path may still appear in git ls-files until staged.
        (self.root / "leetcode/editor/cn/_legacy/root/[763]划分字母区间.go").unlink()
        self.assertEqual(catalog.catalog_stats(self.build())["total"], 1)
        (self.root / "leetcode/editor/cn/p0792/solution.go").unlink()
        self.assertTrue(all(value == 0 for value in catalog.catalog_stats(self.build()).values()))

    def test_generation_does_not_require_or_rewrite_a_readme_or_stats_image(self):
        self.write("leetcode/editor/cn/p0792/solution.go", "package p0792\n")
        data = self.build()
        # No README or generated-section markers exist in this repository fixture.
        generated = catalog.outputs(data)
        self.assertEqual(set(generated), {"docs/explorer/data.js", "docs/CATALOG.md"})
        self.assertIn('"problems":', generated["docs/explorer/data.js"])
        self.assertNotIn('"stats":', generated["docs/explorer/data.js"])
        self.assertIn("704", generated["docs/CATALOG.md"])


if __name__ == "__main__":
    unittest.main()
