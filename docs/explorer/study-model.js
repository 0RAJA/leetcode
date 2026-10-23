/* Pure study rules, shared by the offline UI and Node's built-in tests. */
(function (root, factory) {
  const model = factory();
  if (typeof module === "object" && module.exports) module.exports = model;
  else root.StudyModel = model;
})(typeof window === "object" ? window : globalThis, function () {
  "use strict";
  const broadTopics = new Set(["数组", "字符串", "数学", "排序", "未分类", "算法"]);
  const results = new Set(["needs-practice", "recalled"]);

  function relatedProblems(problem, problems, paths, limit = 4) {
    const byId = new Map(problems.map((item) => [item.id, item]));
    const related = new Map();
    for (const path of paths) {
      const index = path.steps.findIndex((step) => step.id === problem.id);
      if (index < 0) continue;
      for (const offset of [1, -1]) {
        const neighbor = path.steps[index + offset];
        if (!neighbor || !byId.has(neighbor.id) || related.has(neighbor.id)) continue;
        related.set(neighbor.id, {
          problem: byId.get(neighbor.id), kind: "path", path: path.title,
          direction: offset === 1 ? "下一步对比" : "回到基础",
          reason: offset === 1 ? neighbor.connection : path.steps[index].connection,
          shared: problem.topics.filter((topic) => byId.get(neighbor.id).topics.includes(topic)),
        });
      }
    }
    const frequency = new Map();
    for (const item of problems) for (const topic of new Set(item.topics)) {
      frequency.set(topic, (frequency.get(topic) || 0) + 1);
    }
    const candidates = [];
    for (const item of problems) {
      if (item.id === problem.id || related.has(item.id)) continue;
      const shared = item.topics.filter((topic) => problem.topics.includes(topic) && !broadTopics.has(topic));
      if (!shared.length) continue;
      const score = shared.reduce((total, topic) => total + Math.log((problems.length + 1) / (frequency.get(topic) + 1)), 0)
        + (item.active ? 1 : 0) + (item.notes.length ? 0.1 : 0);
      candidates.push({ problem: item, kind: "topic", direction: "相同题型", shared,
        reason: `共同题型：${shared.join("、")}。可对照两题的状态、边界和实现。`, score });
    }
    candidates.sort((a, b) => b.score - a.score || a.problem.id.localeCompare(b.problem.id, "zh-CN", { numeric: true }));
    return [...related.values(), ...candidates].slice(0, limit);
  }

  function chooseReview(problems, markedIds, attempts, random = Math.random) {
    if (!problems.length) return null;
    const marked = problems.filter((problem) => markedIds.has(problem.id));
    const candidates = marked.length ? marked : problems;
    const unseen = candidates.filter((problem) => !attempts[problem.id]);
    if (unseen.length) return unseen[Math.min(unseen.length - 1, Math.floor(random() * unseen.length))];
    return [...candidates].sort((a, b) => attempts[a.id].at - attempts[b.id].at)[0];
  }

  function validAttempts(value, knownIds) {
    const clean = Object.create(null);
    if (!value || typeof value !== "object" || Array.isArray(value)) return clean;
    for (const [id, attempt] of Object.entries(value)) {
      if (!knownIds.has(id) || !attempt || !results.has(attempt.result)
          || !Number.isFinite(attempt.at) || attempt.at <= 0) continue;
      clean[id] = { result: attempt.result, at: attempt.at };
    }
    return clean;
  }

  // A reading session keeps its order even when a self-rating changes the filter.
  function readingSequence(id, ids, label) {
    const unique = [...new Set(ids)];
    return unique.includes(id) ? { ids: unique, label } : { ids: [id], label: "单题阅读" };
  }

  return { relatedProblems, chooseReview, validAttempts, readingSequence };
});
