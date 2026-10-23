"""Keep curated comparisons connected to real catalog records as files evolve."""

import json
from pathlib import Path
import unittest


ROOT = Path(__file__).resolve().parents[1]


def load_assignment(path, variable):
    source = path.read_text(encoding="utf-8")
    payload = source.split(variable + " = ", 1)[1].strip()
    return json.loads(payload.removesuffix(";"))


class LearningPathTests(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.routes = load_assignment(
            ROOT / "docs/explorer/learning-paths.js", "window.LEETCODE_PATHS"
        )["paths"]
        catalog = load_assignment(
            ROOT / "docs/explorer/data.js", "window.LEETCODE_DATA"
        )
        cls.problems = {problem["id"]: problem for problem in catalog["problems"]}

    def test_small_routes_have_unique_ids_and_no_repeated_step(self):
        self.assertGreater(len(self.routes), 0)
        self.assertLessEqual(len(self.routes), 6)
        self.assertEqual(len({route["id"] for route in self.routes}), len(self.routes))
        for route in self.routes:
            with self.subTest(route=route["id"]):
                ids = [step["id"] for step in route["steps"]]
                self.assertGreaterEqual(len(ids), 3)
                self.assertLessEqual(len(ids), 6)
                self.assertEqual(len(set(ids)), len(ids))

    def test_every_step_links_to_an_actual_current_solution(self):
        for route in self.routes:
            for step in route["steps"]:
                with self.subTest(route=route["id"], question=step["id"]):
                    self.assertIn(step["id"], self.problems)
                    problem = self.problems[step["id"]]
                    self.assertTrue(problem["active"])
                    self.assertTrue(any(
                        source["kind"] == "current" and (ROOT / source["path"]).is_file()
                        for source in problem["sources"]
                    ))

    def test_each_transition_explains_the_change_in_the_next_problem(self):
        for route in self.routes:
            self.assertTrue(route["description"].strip())
            for index, step in enumerate(route["steps"]):
                with self.subTest(route=route["id"], question=step["id"]):
                    self.assertTrue(step["focus"].strip())
                    if index:
                        self.assertTrue(step.get("connection", "").strip())


if __name__ == "__main__":
    unittest.main()
