"""Exercise the catalog hook against real indexes and commits in isolated repos."""

import json
import os
from pathlib import Path
import shutil
import subprocess
import sys
import tempfile
import unittest


ROOT = Path(__file__).resolve().parents[1]
ENTRY = ROOT / "scripts/pre_commit.py"
ROUTE = "2026/README.md"
SOURCE = "leetcode/editor/cn/p0792/solution.go"
NOTE = "1. 704 二分查找\n   维护左闭右开区间。\n"


class PreCommitTests(unittest.TestCase):
    def setUp(self):
        parent = ROOT / "tmp"
        parent.mkdir(exist_ok=True)
        self.temp = tempfile.TemporaryDirectory(dir=parent, prefix="hook-test-")
        self.addCleanup(self.temp.cleanup)
        self.root = Path(self.temp.name)
        self.env = {key: value for key, value in os.environ.items()
                    if not key.startswith("GIT_")}
        self.env.update(GIT_CONFIG_GLOBAL=os.devnull, GIT_CONFIG_NOSYSTEM="1")
        self.git("init", "-q")
        self.git("config", "user.name", "Hook Test")
        self.git("config", "user.email", "hook@example.com")
        self.git("config", "commit.gpgsign", "false")
        for path in ("scripts/build_catalog.py", "scripts/pre_commit.py",
                     ".githooks/pre-commit"):
            if (ROOT / path).exists():
                self.write(path, (ROOT / path).read_text())
                shutil.copymode(ROOT / path, self.root / path)
        self.write("leetcode/editor/cn/all.json", json.dumps([
            {"questionId": "792", "frontendQuestionId": "704", "title": "二分查找",
             "titleSlug": "binary-search", "level": 1},
            {"questionId": "768", "frontendQuestionId": "763", "title": "划分字母区间",
             "titleSlug": "partition-labels", "level": 2},
        ]))
        self.write("catalog.config.json", json.dumps({
            "repository": "https://github.com/example/practice", "branch": "master",
            "sources": {}, "overrides": {},
        }))
        self.write(SOURCE, "package p0792\nfunc search() int { return 1 }\n")
        self.write(ROUTE, "# 数组\n" + NOTE)
        self.git("add", ".")
        self.command([sys.executable, "scripts/build_catalog.py"])
        self.git("add", ".")
        self.commit("fixture")

    def command(self, args, check=True):
        result = subprocess.run(args, cwd=self.root, env=self.env,
                                capture_output=True, text=True)
        if check:
            self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        return result

    def git(self, *args, check=True):
        return self.command(["git", *args], check=check)

    def commit(self, message):
        return self.git("-c", "core.hooksPath=/dev/null", "commit", "-qm", message)

    def write(self, name, text):
        path = self.root / name
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(text, encoding="utf-8")

    def run_hook(self, *args):
        return self.command([sys.executable, str(ENTRY), *args], check=False)

    def stage_note(self, text):
        self.write(ROUTE, "# 数组\n" + text)
        self.git("add", ROUTE)

    def data(self, revision=":"):
        raw = self.git("show", revision + "docs/explorer/data.js").stdout
        return json.loads(raw.split("window.LEETCODE_DATA = ", 1)[1].rstrip(";\n"))

    def test_commit_generates_catalog_from_index_and_preserves_unstaged_notes(self):
        installed = self.run_hook("--install")
        self.assertEqual(installed.returncode, 0, installed.stderr)
        self.stage_note(NOTE + "2. 704 二分查找：使用二分定位边界。\n")
        unstaged = (self.root / ROUTE).read_text() + "3. 99999 未暂存的错误\n   不得收录。\n"
        self.write(ROUTE, unstaged)
        result = self.git("commit", "-m", "practice", check=False)
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        notes = self.data("HEAD:")["problems"][0]["notes"]
        self.assertEqual(len(notes), 2)
        self.assertIn("使用二分定位边界", notes[1]["text"])
        self.assertNotIn("未暂存", json.dumps(notes, ensure_ascii=False))
        self.assertEqual((self.root / ROUTE).read_text(), unstaged)
        self.assertIn("docs/CATALOG.md", self.git("show", "--format=", "--name-only", "HEAD").stdout)

    def test_unknown_internal_id_blocks_without_changing_outputs_or_index(self):
        path = "leetcode/editor/cn/p9999/solution.go"
        self.write(path, "package p9999\nfunc example() {}\n")
        self.git("add", path)
        before = self.git("diff", "--cached", "--binary").stdout
        output = (self.root / "docs/CATALOG.md").read_bytes()
        result = self.run_hook()
        self.assertNotEqual(result.returncode, 0)
        self.assertIn(path, result.stdout + result.stderr)
        self.assertEqual(self.git("diff", "--cached", "--binary").stdout, before)
        self.assertEqual((self.root / "docs/CATALOG.md").read_bytes(), output)

    def test_bad_note_headers_are_reported_with_path_and_line(self):
        for header in ("2. 99999 未知题目", "2. 704 划分字母区间",
                       "2.704 二分查找", "2、704 二分查找"):
            with self.subTest(header=header):
                self.stage_note(NOTE + header + "\n   新增思路。\n")
                result = self.run_hook("--check")
                self.assertNotEqual(result.returncode, 0, result.stdout)
                self.assertIn(ROUTE + ":4", result.stdout + result.stderr)

    def test_unclosed_code_fence_is_not_silently_lost(self):
        self.stage_note(NOTE + "2. 704 二分查找\n   ```go\n   code\n")
        result = self.run_hook("--check")
        self.assertNotEqual(result.returncode, 0)
        self.assertIn(ROUTE + ":5", result.stdout + result.stderr)

    def test_indented_problem_entry_does_not_silently_join_previous_note(self):
        self.stage_note(NOTE + "   2. 704 二分查找\n      这条应该顶格书写。\n")
        result = self.run_hook("--check")
        self.assertNotEqual(result.returncode, 0)
        self.assertIn(ROUTE + ":4", result.stdout + result.stderr)

    def test_nested_explanation_lists_and_numbered_code_are_allowed(self):
        self.stage_note(NOTE + "   1. 维护左边界\n   2. 检查右边界\n"
                        "   ```text\n   3. 704 这是代码示例\n   ```\n")
        result = self.run_hook("--check")
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)

    def test_commit_only_is_blocked_before_leaving_reverse_generated_changes(self):
        self.assertEqual(self.run_hook("--install").returncode, 0)
        bad = "leetcode/editor/cn/p9999/solution.go"
        self.write(bad, "package p9999\n")
        self.git("add", bad)
        self.write(ROUTE, "# 数组\n" + NOTE + "2. 704 二分查找：单文件提交。\n")
        before = self.git("rev-parse", "HEAD").stdout
        output = (self.root / "docs/explorer/data.js").read_bytes()
        result = self.git("commit", "--only", ROUTE, "-m", "selected note", check=False)
        self.assertNotEqual(result.returncode, 0)
        self.assertIn("git add", result.stdout + result.stderr)
        self.assertEqual(self.git("rev-parse", "HEAD").stdout, before)
        self.assertEqual((self.root / "docs/explorer/data.js").read_bytes(), output)
        self.assertEqual(self.git("diff", "--cached", "--name-only").stdout, bad + "\n")

    def test_title_only_and_todo_entries_warn_instead_of_claiming_notes(self):
        self.stage_note(NOTE + "2. 763 划分字母区间\n3. 704 二分查找：TODO\n")
        result = self.run_hook()
        self.assertEqual(result.returncode, 0, result.stderr)
        self.assertIn("未收录", result.stdout)
        self.assertEqual(len(self.data()["problems"][0]["notes"]), 1)

    def test_separate_todo_bodies_warn_and_are_not_published(self):
        self.stage_note(NOTE + "2. 704 二分查找\n   TODO\n")
        path = "leetcode/editor/cn/doc/note/704二分查找.md"
        self.write(path, "TODO\n")
        self.git("add", path)
        result = self.run_hook()
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        self.assertIn("[WARN]", result.stdout)
        self.assertEqual(len(self.data()["problems"][0]["notes"]), 1)

    def test_unstaged_generated_file_is_never_overwritten(self):
        self.stage_note(NOTE + "2. 704 二分查找：新的边界说明。\n")
        self.write("docs/CATALOG.md", "USER UNSAVED OUTPUT\n")
        before = self.git("diff", "--cached", "--binary").stdout
        data_before = (self.root / "docs/explorer/data.js").read_bytes()
        result = self.run_hook()
        self.assertNotEqual(result.returncode, 0)
        self.assertIn("docs/CATALOG.md", result.stdout + result.stderr)
        self.assertEqual((self.root / "docs/CATALOG.md").read_text(), "USER UNSAVED OUTPUT\n")
        self.assertEqual((self.root / "docs/explorer/data.js").read_bytes(), data_before)
        self.assertEqual(self.git("diff", "--cached", "--binary").stdout, before)

    def test_unstaged_metadata_and_generator_do_not_affect_staged_result(self):
        self.stage_note(NOTE + "2. 704 二分查找：来自暂存区。\n")
        self.write("leetcode/editor/cn/all.json", "INVALID UNSTAGED JSON\n")
        self.write("scripts/build_catalog.py", "raise RuntimeError('UNSTAGED GENERATOR')\n")
        result = self.run_hook()
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        self.assertEqual(len(self.data()["problems"][0]["notes"]), 2)

    def test_invalid_staged_json_reports_its_file_and_line(self):
        for path in ("catalog.config.json", "leetcode/editor/cn/all.json"):
            with self.subTest(path=path):
                original = (self.root / path).read_text()
                self.write(path, "{broken\n")
                self.git("add", path)
                result = self.run_hook("--check")
                self.assertNotEqual(result.returncode, 0)
                self.assertIn(path + ":1", result.stdout + result.stderr)
                self.write(path, original)
                self.git("add", path)

    def test_old_unmapped_note_does_not_block_unrelated_new_note(self):
        self.stage_note(NOTE + "2. 99999 历史未知\n   旧思路。\n")
        self.commit("legacy note")
        self.stage_note(NOTE + "2. 99999 历史未知\n   旧思路。\n3. 704 二分查找：新的说明。\n")
        result = self.run_hook()
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        self.assertEqual(len(self.data()["problems"][0]["notes"]), 2)

    def test_removing_only_source_blocks_newly_orphaned_note(self):
        self.git("rm", SOURCE)
        result = self.run_hook("--check")
        self.assertNotEqual(result.returncode, 0)
        self.assertIn(ROUTE + ":2", result.stdout + result.stderr)

    def test_standalone_note_is_linked_and_empty_note_warns(self):
        path = "leetcode/editor/cn/doc/note/704二分查找.md"
        self.write(path, "维护左右边界，逐步缩小搜索范围。\n")
        self.git("add", path)
        result = self.run_hook()
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        self.assertEqual(self.data()["problems"][0]["notes"][-1]["path"], path)
        self.write(path, "\n")
        self.git("add", path)
        result = self.run_hook()
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        self.assertIn("未收录", result.stdout)

    def test_unknown_standalone_note_is_an_error(self):
        path = "leetcode/editor/cn/doc/note/写错的题目.md"
        self.write(path, "无法关联的思路。\n")
        self.git("add", path)
        result = self.run_hook("--check")
        self.assertNotEqual(result.returncode, 0)
        self.assertIn(path + ":1", result.stdout + result.stderr)

    def test_custom_platform_note_uses_source_mapping(self):
        path = "2026/intervalsum/solution.go"
        self.write(path, "package intervalsum\n")
        config = json.loads((self.root / "catalog.config.json").read_text())
        config["sources"][path] = {"id": "卡码网 58", "title": "区间和"}
        self.write("catalog.config.json", json.dumps(config))
        self.stage_note(NOTE + "2. 58 区间和：利用前缀和。\n")
        self.git("add", path, "catalog.config.json")
        result = self.run_hook()
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        record = next(p for p in self.data()["problems"] if p["id"] == "卡码网 58")
        self.assertEqual(len(record["notes"]), 1)

    def test_read_only_check_reports_pending_generation_without_mutation(self):
        self.stage_note(NOTE + "2. 704 二分查找：只读预览。\n")
        before = self.git("diff", "--cached", "--binary").stdout
        output = (self.root / "docs/CATALOG.md").read_bytes()
        result = self.run_hook("--check")
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        self.assertIn("docs/CATALOG.md", result.stdout)
        self.assertEqual(self.git("diff", "--cached", "--binary").stdout, before)
        self.assertEqual((self.root / "docs/CATALOG.md").read_bytes(), output)

    def test_first_commit_without_head_is_supported(self):
        self.git("update-ref", "-d", "HEAD")
        result = self.run_hook()
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        self.assertEqual(self.data()["problems"][0]["id"], "704")

    def test_unrelated_changes_skip_generation(self):
        self.write("other.txt", "unrelated\n")
        self.git("add", "other.txt")
        result = self.run_hook()
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        self.assertEqual(self.git("diff", "--cached", "--name-only").stdout, "other.txt\n")

    def test_intent_to_add_solution_does_not_block_unrelated_commit(self):
        path = "leetcode/editor/cn/p9999/solution.go"
        self.write(path, "package p9999\n")
        self.git("add", "-N", path)
        self.write("other.txt", "unrelated\n")
        self.git("add", "other.txt")
        self.assertEqual(self.run_hook("--install").returncode, 0)
        result = self.git("commit", "-m", "unrelated change", check=False)
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        self.assertNotEqual(self.git("cat-file", "-e", "HEAD:" + path, check=False).returncode, 0)
        self.assertEqual((self.root / path).read_text(), "package p9999\n")

    def test_intent_to_add_is_excluded_from_generated_snapshot(self):
        path = "leetcode/editor/cn/p9999/solution.go"
        self.write(path, "package p9999\n")
        self.git("add", "-N", path)
        self.stage_note(NOTE + "2. 704 二分查找：真实暂存的思路。\n")
        result = self.run_hook()
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        self.assertEqual(len(self.data()["problems"][0]["notes"]), 2)
        self.assertNotIn(path, self.git("diff", "--cached", "--name-only").stdout)
        self.assertIn(path, self.git("diff", "--name-only").stdout)

    def test_actually_staged_empty_solution_still_requires_mapping(self):
        path = "leetcode/editor/cn/p9999/solution.go"
        self.write(path, "")
        self.git("add", path)
        result = self.run_hook()
        self.assertNotEqual(result.returncode, 0)
        self.assertIn(path, result.stdout + result.stderr)

    def test_install_is_idempotent_and_does_not_override_existing_hooks_path(self):
        first = self.run_hook("--install")
        self.assertEqual(first.returncode, 0, first.stdout + first.stderr)
        self.assertEqual(self.run_hook("--install").returncode, 0)
        self.assertEqual(self.git("config", "--local", "--get", "core.hooksPath").stdout.strip(), ".githooks")
        self.git("config", "core.hooksPath", "custom-hooks")
        result = self.run_hook("--install")
        self.assertNotEqual(result.returncode, 0)
        self.assertEqual(self.git("config", "--get", "core.hooksPath").stdout.strip(), "custom-hooks")

    def test_install_protects_existing_default_hook(self):
        path = self.root / ".git/hooks/pre-commit"
        path.write_text("#!/bin/sh\nexit 0\n")
        path.chmod(0o755)
        result = self.run_hook("--install")
        self.assertNotEqual(result.returncode, 0)
        self.assertIn("pre-commit", result.stdout + result.stderr)
        self.assertNotEqual(self.git("config", "--get", "core.hooksPath", check=False).returncode, 0)

    def test_install_does_not_hide_other_default_hooks(self):
        path = self.root / ".git/hooks/pre-push"
        path.write_text("#!/bin/sh\nexit 0\n")
        path.chmod(0o755)
        result = self.run_hook("--install")
        self.assertNotEqual(result.returncode, 0)
        self.assertIn("pre-push", result.stdout + result.stderr)


if __name__ == "__main__":
    unittest.main()
