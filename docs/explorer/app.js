/* Repository-native, dependency-free review interface. data.js is generated. */
(() => {
  "use strict";
  const $ = (id) => document.getElementById(id);
  const data = window.LEETCODE_DATA;
  if (!data || !Array.isArray(data.problems)) {
    $("load-error").hidden = false;
    $("load-error").textContent = "题目索引尚未生成。请先运行仓库的索引生成脚本，再打开此页。";
    return;
  }
  const problems = data.problems;
  const byId = new Map(problems.map((problem) => [problem.id, problem]));
  const paths = window.LEETCODE_PATHS?.paths || [];
  const model = window.StudyModel;
  const summary = model.catalogSummary(problems);
  const storageKey = "raja-leetcode-review-v1";
  const attemptKey = "raja-leetcode-attempts-v1";
  const pageSize = 16;
  const scopes = {
    current: ["当前练习", "独立题解与本地练习，随时回到熟悉的思路。"],
    all: ["全部收录", "按题号去重，串起当前练习与历史归档。"],
    archive: ["历史归档", "仅展示尚未整理到当前练习的历史题目；代码保留当时的样子。"],
    review: ["待复习", "给值得再想一次的题目，留一个位置。标记仅保存在当前浏览器。"],
  };
  const query = new URLSearchParams(location.search);
  const state = {
    scope: Object.hasOwn(scopes, query.get("scope")) ? query.get("scope") : "current",
    q: query.get("q") || "",
    topic: query.get("topic") || "",
    difficulty: query.get("difficulty") || "",
    record: query.get("record") || "",
    view: query.get("view") === "map" ? "map" : "list",
    page: 1,
  };
  let reviews = new Set();
  let currentProblem = null;
  let currentTab = "recall";
  let problemTrail = [];
  let readingSequence = { ids: [], label: "" };
  let readerExpanded = false;
  let returnFocus = null;
  let attempts = Object.create(null);
  let filtered = [];
  let toastTimeout;
  try {
    const saved = JSON.parse(localStorage.getItem(storageKey) || "[]");
    if (Array.isArray(saved)) reviews = new Set(saved.filter((id) => typeof id === "string" && byId.has(id)));
  } catch (_) {
    $("storage-note").textContent = "当前浏览器无法读取复习标记，可使用导入与导出";
  }
  try { attempts = model.validAttempts(JSON.parse(localStorage.getItem(attemptKey) || "{}"), byId); }
  catch (_) { /* Keep the existing bookmark queue usable when history cannot load. */ }

  const element = (tag, className, text) => {
    const node = document.createElement(tag);
    if (className) node.className = className;
    if (text !== undefined) node.textContent = text;
    return node;
  };
  const repo = /^https:\/\/github\.com\/[^/]+\/[^/]+\/?$/.test(data.repository) ? data.repository.replace(/\/$/, "") : "https://github.com/0RAJA/leetcode";
  const repoURL = (path, anchor = "") => `${repo}/blob/${encodeURIComponent(data.branch || "master")}/${path.split("/").map(encodeURIComponent).join("/")}${anchor ? `#${encodeURIComponent(anchor)}` : ""}`;
  const safeExternal = (url) => { try { const parsed = new URL(url); return parsed.protocol === "https:" ? parsed.href : ""; } catch (_) { return ""; } };
  const link = (label, href, className = "") => {
    const node = element("a", className, label);
    node.href = href;
    node.target = "_blank";
    node.rel = "noopener noreferrer";
    return node;
  };
  const hasTests = (problem) => problem.sources.some((source) => source.tests.length > 0);
  const difficultyClass = (value) => ({ "简单": "easy", "中等": "medium", "困难": "hard" }[value] || "");
  const scopeBase = () => problems.filter((problem) => state.scope === "current" ? problem.active : state.scope === "archive" ? !problem.active : state.scope === "review" ? reviews.has(problem.id) : true);
  const ordered = (items) => [...items].sort((a, b) => a.id.localeCompare(b.id, "zh-CN", { numeric: true }));
  const showToast = (text) => {
    clearTimeout(toastTimeout);
    const target = $("problem-dialog").open ? $("reader-toast") : $("toast");
    $("toast").hidden = true;
    $("reader-toast").hidden = true;
    target.textContent = text;
    target.hidden = false;
    toastTimeout = setTimeout(() => { target.hidden = true; }, 3600);
  };
  const saveReviews = () => {
    try { localStorage.setItem(storageKey, JSON.stringify([...reviews].sort())); }
    catch (_) { $("storage-note").textContent = "标记暂存于本页，请导出以免丢失"; showToast("浏览器未允许保存；可用“导出标记”保留这次记录。"); }
  };
  const setBookmarkButton = (button, problem, detail = false) => {
    const marked = reviews.has(problem.id);
    button.textContent = detail ? marked ? "★ 已加入待复习" : "☆ 加入待复习" : marked ? "★" : "☆";
    button.classList.toggle("marked", marked);
    button.setAttribute("aria-pressed", String(marked));
    button.setAttribute("aria-label", `${marked ? "将" : "把"}${problem.title}${marked ? "移出" : "加入"}待复习`);
    button.title = marked ? "移出待复习" : "加入待复习";
  };
  const toggleReview = (id) => {
    if (!byId.has(id)) return;
    if (reviews.has(id)) reviews.delete(id); else reviews.add(id);
    saveReviews();
    render();
    if (currentProblem) setBookmarkButton($("detail-review"), currentProblem, true);
    if (currentProblem) renderAttempt();
  };
  const updateURL = () => {
    const url = new URL(location.href);
    url.search = "";
    for (const [key, value] of Object.entries(state)) {
      if (key !== "page" && value && !(key === "scope" && value === "current") && !(key === "view" && value === "list")) url.searchParams.set(key, value);
    }
    try { history.replaceState(null, "", url); } catch (_) { /* Local file URLs can restrict history changes. */ }
  };
  function renderTopics() {
    const counts = new Map();
    for (const problem of scopeBase()) for (const topic of problem.topics) counts.set(topic, (counts.get(topic) || 0) + 1);
    const topics = [...counts].sort((a, b) => b[1] - a[1] || a[0].localeCompare(b[0], "zh-CN"));
    const max = topics[0]?.[1] || 1;
    const nav = $("topic-nav");
    nav.replaceChildren();
    for (const [topic, count] of topics.slice(0, 10)) {
      const button = element("button", "topic-button");
      button.type = "button";
      button.dataset.topic = topic;
      button.style.setProperty("--coverage", `${Math.round(count / max * 80)}%`);
      button.classList.toggle("selected", state.topic === topic);
      button.setAttribute("aria-pressed", String(state.topic === topic));
      button.append(element("span", "", topic), element("span", "", String(count)));
      nav.append(button);
    }
    const select = $("topic-filter");
    select.replaceChildren(new Option("全部题型", ""));
    for (const [topic, count] of topics) select.add(new Option(`${topic} · ${count}`, topic));
    if (state.topic && !counts.has(state.topic)) select.add(new Option(state.topic, state.topic));
    select.value = state.topic;
    $("clear-topic").hidden = !state.topic;
  }
  function renderRows() {
    const fragment = document.createDocumentFragment();
    for (const problem of filtered.slice((state.page - 1) * pageSize, state.page * pageSize)) {
      const row = element("tr", "question-row");
      row.dataset.problem = problem.id;
      const titleCell = element("td");
      const title = element("button", "question-title");
      title.type = "button";
      title.dataset.open = problem.id;
      title.append(element("span", "problem-id", problem.id), element("span", "problem-name", problem.title));
      const primary = problem.sources.find((source) => source.kind === "current") || problem.sources[0];
      const sourceLabel = primary.kind === "current" ? "当前练习" : "历史归档";
      const subtitle = element("p", "question-subtitle", `${sourceLabel}${problem.sources.length > 1 ? ` · ${problem.sources.length} 份题解` : ""}${problem.notes.length ? ` · ${problem.notes.length} 处笔记` : ""}`);
      if (attempts[problem.id]) subtitle.append(document.createTextNode(` · 上次自评：${attempts[problem.id].result === "recalled" ? "能独立写出" : "需要再练"}`));
      titleCell.append(title, subtitle);
      const topicsCell = element("td");
      for (const topic of problem.topics.slice(0, 3)) {
        const button = element("button", "topic-tag", topic);
        button.type = "button";
        button.dataset.topic = topic;
        button.title = `按${topic}筛选`;
        topicsCell.append(button);
      }
      if (problem.topics.length > 3) {
        const more = element("span", "topic-more", `+${problem.topics.length - 3}`);
        more.title = problem.topics.slice(3).join("、");
        topicsCell.append(more);
      }
      const difficultyCell = element("td");
      difficultyCell.append(element("span", `difficulty ${difficultyClass(problem.difficulty)}`, problem.difficulty));
      const recordCell = element("td");
      recordCell.append(element("span", `record-label${problem.notes.length ? "" : " missing"}`, problem.notes.length ? "↳ 思路笔记" : "待补笔记"));
      if (hasTests(problem)) recordCell.append(element("span", "record-label", "◦ 测试文件"));
      const reviewCell = element("td");
      const review = element("button", "review-toggle");
      review.type = "button";
      review.dataset.review = problem.id;
      setBookmarkButton(review, problem);
      reviewCell.append(review);
      row.append(titleCell, topicsCell, difficultyCell, recordCell, reviewCell);
      fragment.append(row);
    }
    $("questions").replaceChildren(fragment);
  }
  function renderMap() {
    const container = $("relationship-map");
    container.replaceChildren();
    const matches = new Set(filtered.map((problem) => problem.id));
    const visible = paths.filter((path) => path.steps.some((step) => matches.has(step.id)));
    const intro = element("div", "map-intro");
    intro.append(element("h3", "", "沿着一个思路，比较它的变化"), element("p", "", "精选路线只串起适合比较的题。筛选时保留前后文，高亮匹配题；点击题目，查看思路、代码和相邻题的对比理由。"));
    container.append(intro);
    if (!visible.length) {
      const empty = element("div", "empty-state");
      empty.append(element("h3", "", "当前筛选没有覆盖精选路线"), element("p", "", "题目列表仍收录全部记录。没有手工路线的题，也会在详情中按具体题型查找关联。"));
      const button = element("button", "secondary-button", "回到题目列表");
      button.type = "button";
      button.dataset.view = "list";
      empty.append(button); container.append(empty); return;
    }
    for (const path of visible) {
      const lane = element("article", "map-lane");
      const heading = element("div", "lane-heading");
      heading.append(element("h3", "", path.title), element("span", "count-pill", `${path.steps.length} 道对比题`));
      lane.append(heading, element("p", "lane-description", path.description));
      const board = element("div", "map-board");
      const steps = element("div", "route-steps");
      path.steps.forEach((step, index) => {
        const problem = byId.get(step.id);
        if (!problem) return;
        if (index) {
          const arrow = element("span", "route-line", "→");
          arrow.title = step.connection;
          arrow.setAttribute("aria-hidden", "true");
          steps.append(arrow);
        }
        const node = element("button", "route-step");
        node.type = "button";
        node.dataset.problem = step.id;
        node.dataset.path = path.id;
        node.classList.toggle("muted", !matches.has(step.id));
        node.title = step.connection || step.focus;
        node.append(element("span", "step-number", String(index + 1).padStart(2, "0")),
          element("strong", "step-title", `${problem.id} ${problem.title}`),
          element("span", "step-detail", step.focus),
          element("span", "step-status", reviews.has(step.id) ? "☆ 待复习" : attempts[step.id] ? "已自评" : problem.notes.length ? "有思路笔记" : "有 Go 题解"));
        steps.append(node);
      });
      board.append(steps); lane.append(board); container.append(lane);
    }
  }
  function render() {
    const q = state.q.trim().toLocaleLowerCase();
    filtered = ordered(scopeBase().filter((problem) => {
      const haystack = `${problem.id} ${problem.title} ${problem.topics.join(" ")} ${problem.notes.map((note) => note.text).join(" ")}`.toLocaleLowerCase();
      return (!q || haystack.includes(q)) && (!state.topic || problem.topics.includes(state.topic)) && (!state.difficulty || problem.difficulty === state.difficulty) && (!state.record || (state.record === "notes" ? problem.notes.length > 0 : state.record === "tests" ? hasTests(problem) : problem.notes.length === 0));
    }));
    const pageCount = Math.max(1, Math.ceil(filtered.length / pageSize));
    state.page = Math.min(state.page, pageCount);
    $("count-current").textContent = String(summary.active);
    $("count-all").textContent = String(summary.total);
    $("count-archive").textContent = String(summary.archived);
    $("count-review").textContent = String(reviews.size);
    $("list-title").textContent = scopes[state.scope][0];
    $("scope-description").textContent = scopes[state.scope][1];
    $("result-count").textContent = `${filtered.length} 道题`;
    for (const button of document.querySelectorAll("[data-scope]")) {
      const selected = button.dataset.scope === state.scope;
      button.classList.toggle("selected", selected);
      button.setAttribute("aria-pressed", String(selected));
    }
    $("search").value = state.q;
    $("difficulty-filter").value = state.difficulty;
    $("record-filter").value = state.record;
    renderTopics();
    renderRows();
    $("list-view").hidden = state.view !== "list";
    $("relationship-map").hidden = state.view !== "map";
    for (const button of document.querySelectorAll("[data-view]")) button.setAttribute("aria-pressed", String(button.dataset.view === state.view));
    if (state.view === "map") renderMap();
    const chips = $("active-filters");
    chips.replaceChildren();
    const labels = { topic: state.topic, difficulty: state.difficulty, record: { notes: "有思路笔记", tests: "有测试文件", "no-notes": "待补笔记" }[state.record] };
    for (const [key, text] of Object.entries(labels)) if (text) {
      const button = element("button", "filter-chip", text);
      button.type = "button";
      button.dataset.clear = key;
      button.setAttribute("aria-label", `移除${text}筛选`);
      button.append(element("span", "", "×"));
      chips.append(button);
    }
    $("reset-filters").hidden = !(state.q || state.topic || state.difficulty || state.record);
    $("empty-state").hidden = filtered.length > 0;
    $("empty-description").textContent = state.scope === "review" && !reviews.size ? "点击题目旁的 ☆，把想再做一遍的题目放在这里。" : "试着换一个关键词，或者放宽筛选条件。";
    $("empty-reset").textContent = state.scope === "review" && !reviews.size ? "去当前练习看看" : "重置筛选";
    $("pagination-description").textContent = filtered.length ? `显示 ${(state.page - 1) * pageSize + 1}–${Math.min(state.page * pageSize, filtered.length)} / ${filtered.length} 道题` : "0 道题";
    $("page-label").textContent = `${state.page} / ${pageCount}`;
    $("previous-page").disabled = state.page <= 1;
    $("next-page").disabled = state.page >= pageCount;
    $("random-review").disabled = filtered.length === 0;
    updateURL();
  }
  function resetFilters() {
    Object.assign(state, { q: "", topic: "", difficulty: "", record: "", page: 1 });
    if (state.scope === "review" && !reviews.size) state.scope = "current";
    render();
  }

  // Notes stay actual repository text. Only inline code and fenced blocks are formatted.
  function appendInline(parent, text) {
    for (const part of text.split(/(`[^`]+`)/g)) parent.append(part.startsWith("`") && part.endsWith("`") ? element("code", "", part.slice(1, -1)) : document.createTextNode(part));
  }
  function highlightCode(node, text, language = "go") {
    node.textContent = text;
    node.className = `language-${language}`;
    node.removeAttribute("data-highlighted");
    if (window.hljs?.getLanguage(language)) window.hljs.highlightElement(node);
  }
  function noteContent(text) {
    const container = element("div", "note-prose");
    let inCode = false;
    let codeLanguage = "go";
    let code = [];
    let paragraph = [];
    const flushParagraph = () => {
      if (!paragraph.length) return;
      const p = element("p");
      appendInline(p, paragraph.join("\n"));
      container.append(p);
      paragraph = [];
    };
    const flushCode = () => {
      const pre = element("pre");
      pre.tabIndex = 0;
      const codeNode = element("code");
      highlightCode(codeNode, code.join("\n"), codeLanguage);
      pre.append(codeNode);
      container.append(pre);
      code = [];
    };
    for (let line of text.replace(/\r\n/g, "\n").split("\n")) {
      if (/^\s*```/.test(line)) {
        flushParagraph();
        if (inCode) flushCode();
        else {
          const language = line.trim().slice(3).trim().toLowerCase();
          codeLanguage = !language ? "go" : /^[\w+-]+$/.test(language) ? language : "plaintext";
        }
        inCode = !inCode; continue;
      }
      if (inCode) { code.push(line.replace(/^ {4}/, "")); continue; }
      line = line.trim().replace(/^\/\/\s?/, "");
      if (!line) { flushParagraph(); continue; }
      if (/^#{1,5}\s/.test(line)) { flushParagraph(); container.append(element("h4", "", line.replace(/^#+\s*/, ""))); }
      else paragraph.push(line);
    }
    flushParagraph();
    if (inCode || code.length) flushCode();
    return container;
  }
  function showTab(tab) {
    currentTab = tab;
    for (const name of ["recall", "notes", "code"]) {
      const selected = name === tab;
      $(`${name}-tab`).setAttribute("aria-selected", String(selected));
      $(`${name}-tab`).tabIndex = selected ? 0 : -1;
      $(`${name}-panel`).hidden = !selected;
    }
    $("reader-scroll").scrollTop = 0;
    $("problem-dialog").classList.remove("outline-open");
    $("toggle-reader-outline").setAttribute("aria-expanded", "false");
  }
  function resource(label, path, suffix, anchor) {
    const node = link("", repoURL(path, anchor), "resource-link");
    const title = element("div", "", label);
    if (suffix) title.append(element("small", "", suffix));
    node.append(title, element("span", "", "↗"));
    return node;
  }
  function renderSource(index) {
    const source = currentProblem.sources[index] || currentProblem.sources[0];
    highlightCode($("source-code"), source.code || "// 此版本暂无可显示的代码，请通过右侧链接查看原始文件。");
    $("source-code").parentElement.scrollTop = 0;
    $("source-code").parentElement.scrollLeft = 0;
    $("source-path").textContent = source.path;
    $("reader-scroll").scrollTop = 0;
  }
  function sequenceFor(id, pathId) {
    const path = paths.find((item) => item.id === pathId);
    const ids = path ? path.steps.map((step) => step.id).filter((stepId) => byId.has(stepId)) : filtered.map((problem) => problem.id);
    const filteredLabel = [scopes[state.scope][0], state.topic, state.difficulty, state.q ? `搜索“${state.q}”` : "", { notes: "有思路笔记", tests: "有测试文件", "no-notes": "待补笔记" }[state.record]].filter(Boolean).join(" · ");
    return { ...model.readingSequence(id, ids, path ? path.title : filteredLabel), pathId: path?.id };
  }
  function renderReaderNavigation() {
    const index = readingSequence.ids.indexOf(currentProblem.id);
    const total = readingSequence.ids.length;
    const position = `${index + 1} / ${total}`;
    $("reader-position").textContent = position;
    $("outline-count").textContent = String(total);
    $("reader-context").textContent = readingSequence.label;
    $("reader-progress").textContent = `第 ${index + 1} / ${total} 题 · ${readingSequence.label}`;
    for (const [direction, offset] of [["previous", -1], ["next", 1]]) {
      const neighbor = byId.get(readingSequence.ids[index + offset]);
      const button = $(`${direction}-problem`);
      button.disabled = !neighbor;
      $(`${direction}-problem-title`).textContent = neighbor ? `${neighbor.id} ${neighbor.title}` : direction === "previous" ? "已经是第一题" : "已经是最后一题";
      button.title = neighbor ? `${direction === "previous" ? "上一题" : "下一题"}：${neighbor.title}` : "";
    }
    const queue = $("reader-queue");
    queue.replaceChildren();
    readingSequence.ids.forEach((id, queueIndex) => {
      const problem = byId.get(id);
      const button = element("button", "reader-queue-item");
      button.type = "button";
      button.dataset.reading = id;
      button.setAttribute("aria-current", id === currentProblem.id ? "step" : "false");
      button.append(element("span", "queue-number", String(queueIndex + 1).padStart(2, "0")), element("span", "", `${problem.id} ${problem.title}`));
      queue.append(button);
    });
    $("back-problem").hidden = !problemTrail.length;
  }
  function currentReadingPosition() {
    return { id: currentProblem.id, tab: currentTab, sequence: readingSequence,
      sourceIndex: Number($("source-select").value), scrollTop: $("reader-scroll").scrollTop };
  }
  function scrollQueueToCurrent() {
    const selected = $("reader-queue").querySelector('[aria-current="step"]');
    if (selected) $("reader-queue").scrollTop = selected.offsetTop - $("reader-queue").offsetTop - 56;
  }
  function navigateReading(id) {
    if (!id || id === currentProblem?.id) return;
    problemTrail.push(currentReadingPosition());
    openProblem(id, { tab: currentTab, keepTrail: true, sequence: readingSequence });
  }
  function adjacentProblem(offset) {
    if (!currentProblem) return;
    const index = readingSequence.ids.indexOf(currentProblem.id);
    navigateReading(readingSequence.ids[index + offset]);
  }
  function renderRelated(problem) {
    const container = $("related-problems");
    container.replaceChildren();
    const related = model.relatedProblems(problem, problems, paths);
    if (!related.length) container.append(element("p", "", "暂时没有明确关联，可以按题型继续探索。"));
    for (const item of related) {
      const button = element("button", "related-card");
      button.type = "button";
      button.dataset.related = item.problem.id;
      button.append(element("span", "related-direction", item.direction), element("strong", "", `${item.problem.id} ${item.problem.title}`),
        element("p", "related-reason", item.reason), element("small", "related-kind", item.kind === "path" ? `精选对比 · ${item.path}` : "自动关联 · 共同题型"));
      container.append(button);
    }
  }
  function renderAttempt() {
    const attempt = currentProblem && attempts[currentProblem.id];
    $("rate-again").setAttribute("aria-pressed", String(attempt?.result === "needs-practice"));
    $("rate-recalled").setAttribute("aria-pressed", String(attempt?.result === "recalled"));
    $("clear-rating").hidden = !attempt;
    $("review-result").textContent = attempt ? `最近自评：${attempt.result === "recalled" ? "能独立写出" : "需要再练"}。当前${reviews.has(currentProblem.id) ? "已加入" : "未加入"}待复习，仅保存在当前浏览器。` : "还没有自评记录；查阅代码不会自动标为掌握。";
  }
  function rateAttempt(result) {
    if (!currentProblem) return;
    attempts[currentProblem.id] = { result, at: Date.now() };
    if (result === "needs-practice") reviews.add(currentProblem.id); else reviews.delete(currentProblem.id);
    try { localStorage.setItem(attemptKey, JSON.stringify(attempts)); }
    catch (_) { showToast("本次自评暂存于页面，请导出备份。"); }
    saveReviews(); render(); renderAttempt();
    setBookmarkButton($("detail-review"), currentProblem, true);
  }
  function openProblem(id, { updateHash = true, tab = "recall", keepTrail = false, sequence = null, pathId, sourceIndex = 0, scrollTop = 0 } = {}) {
    const problem = byId.get(id);
    if (!problem) return;
    if (!keepTrail) problemTrail = [];
    if (!$("problem-dialog").open) returnFocus = document.activeElement;
    readingSequence = sequence || sequenceFor(id, pathId);
    currentProblem = problem;
    $("back-problem").hidden = !problemTrail.length;
    $("detail-title").textContent = `${problem.id} · ${problem.title}`;
    $("detail-eyebrow").textContent = problem.active ? "当前练习 / 题目手记" : "历史归档 / 题目手记";
    const tags = $("detail-tags");
    tags.replaceChildren(element("span", `difficulty ${difficultyClass(problem.difficulty)}`, problem.difficulty));
    for (const topic of problem.topics) tags.append(element("span", "topic-tag", topic));
    setBookmarkButton($("detail-review"), problem, true);
    $("detail-note-count").textContent = String(problem.notes.length);
    $("detail-source-count").textContent = String(problem.sources.length);
    const notes = $("notes-panel");
    notes.replaceChildren();
    if (!problem.notes.length) {
      const empty = element("div", "note-empty");
      empty.append(element("h3", "", "思路，等下次回顾时补上。"), element("p", "", "这道题已有代码，暂时没有找到关联的思路笔记。可以先读自己的实现，再把关键思路和易错点记回仓库。"));
      const button = element("button", "secondary-button", "先看 Go 题解 →");
      button.type = "button";
      button.addEventListener("click", () => { showTab("code"); $("code-tab").focus(); });
      empty.append(button);
      notes.append(empty);
    }
    for (const note of problem.notes) {
      const section = element("article", "note-entry");
      const header = element("div", "note-entry-header");
      header.append(element("h3", "", note.title), link("查看原始笔记 ↗", repoURL(note.path, note.anchor)));
      section.append(header, noteContent(note.text));
      notes.append(section);
    }
    const select = $("source-select");
    select.replaceChildren();
    problem.sources.forEach((source, index) => select.add(new Option(`${source.kind === "current" ? "当前练习" : "历史版本"} · ${source.path}`, String(index))));
    select.value = String(problem.sources[sourceIndex] ? sourceIndex : 0);
    renderSource(Number(select.value));
    const external = safeExternal(problem.url);
    $("detail-problem-link").hidden = !external;
    if (external) $("detail-problem-link").href = external; else $("detail-problem-link").removeAttribute("href");
    $("recall-problem-link").hidden = !external;
    if (external) $("recall-problem-link").href = external; else $("recall-problem-link").removeAttribute("href");
    const route = paths.find((path) => path.steps.some((step) => step.id === id));
    $("recall-focus").textContent = route ? route.steps.find((step) => step.id === id).focus : "先说清楚：为什么选这个解法，边界在哪里？";
    $("recall-route").textContent = route ? `本次复习线索：${route.title}` : `可从 ${problem.topics.join("、")} 的角度识别题型，再检查这个方法成立的条件。`;
    const resources = $("detail-resources");
    resources.replaceChildren();
    problem.sources.forEach((source, index) => {
      resources.append(resource(problem.sources.length > 1 ? `Go 题解 · ${index + 1}` : "Go 题解", source.path, source.kind === "current" ? "当前练习" : "历史归档"));
      source.tests.forEach((path) => resources.append(resource("测试文件", path, path.split("/").pop())));
    });
    const seenNotes = new Set();
    for (const note of problem.notes) {
      const key = `${note.path}#${note.anchor}`;
      if (!seenNotes.has(key)) { resources.append(resource("原始思路", note.path, note.title, note.anchor)); seenNotes.add(key); }
    }
    renderRelated(problem);
    renderAttempt();
    renderReaderNavigation();
    showTab(tab);
    if (!$("problem-dialog").open) $("problem-dialog").showModal();
    $("detail-title").focus({ preventScroll: true });
    $("reader-scroll").scrollTop = scrollTop;
    scrollQueueToCurrent();
    $("reader-toast").hidden = true;
    if (updateHash) {
      const url = new URL(location.href);
      const hash = new URLSearchParams({ problem: id });
      if (readingSequence.pathId) hash.set("path", readingSequence.pathId);
      url.hash = hash.toString();
      try { history.replaceState(null, "", url); } catch (_) { /* Details remain usable on restrictive file origins. */ }
    }
  }
  $("problem-dialog").addEventListener("close", () => {
    currentProblem = null;
    problemTrail = [];
    $("problem-dialog").classList.remove("outline-open");
    $("toggle-reader-outline").setAttribute("aria-expanded", "false");
    $("reader-toast").hidden = true;
    const url = new URL(location.href);
    url.hash = "";
    try { history.replaceState(null, "", url); } catch (_) { /* No navigation needed. */ }
    (returnFocus?.isConnected ? returnFocus : $("search")).focus({ preventScroll: true });
  });
  $("back-problem").addEventListener("click", () => {
    const previous = problemTrail.pop();
    if (previous) openProblem(previous.id, { ...previous, keepTrail: true });
  });
  $("previous-problem").addEventListener("click", () => adjacentProblem(-1));
  $("next-problem").addEventListener("click", () => adjacentProblem(1));
  $("reader-queue").addEventListener("click", (event) => {
    const button = event.target.closest("[data-reading]");
    if (!button) return;
    navigateReading(button.dataset.reading);
    $("problem-dialog").classList.remove("outline-open");
    $("toggle-reader-outline").setAttribute("aria-expanded", "false");
  });
  $("toggle-reader-size").addEventListener("click", () => {
    readerExpanded = !readerExpanded;
    $("problem-dialog").classList.toggle("reader-expanded", readerExpanded);
    $("toggle-reader-size").setAttribute("aria-pressed", String(readerExpanded));
    $("reader-size-label").textContent = readerExpanded ? "退出全屏" : "全屏阅读";
  });
  $("toggle-reader-outline").addEventListener("click", () => {
    const expanded = $("problem-dialog").classList.toggle("outline-open");
    $("toggle-reader-outline").setAttribute("aria-expanded", String(expanded));
    if (expanded) scrollQueueToCurrent();
  });
  $("return-index").addEventListener("click", () => $("problem-dialog").close());
  $("close-dialog").addEventListener("click", () => $("problem-dialog").close());
  $("problem-dialog").addEventListener("click", (event) => {
    if (event.target !== $("problem-dialog")) return;
    const bounds = $("problem-dialog").getBoundingClientRect();
    if (event.clientX < bounds.left || event.clientX > bounds.right || event.clientY < bounds.top || event.clientY > bounds.bottom) $("problem-dialog").close();
  });
  $("recall-tab").addEventListener("click", () => showTab("recall"));
  $("notes-tab").addEventListener("click", () => showTab("notes"));
  $("code-tab").addEventListener("click", () => showTab("code"));
  $("reveal-notes").addEventListener("click", () => { showTab("notes"); $("notes-tab").focus(); });
  $("reveal-code").addEventListener("click", () => { showTab("code"); $("code-tab").focus(); });
  $("rate-again").addEventListener("click", () => rateAttempt("needs-practice"));
  $("rate-recalled").addEventListener("click", () => rateAttempt("recalled"));
  $("clear-rating").addEventListener("click", () => {
    if (!currentProblem) return;
    delete attempts[currentProblem.id];
    try { localStorage.setItem(attemptKey, JSON.stringify(attempts)); } catch (_) { /* In-memory state remains usable. */ }
    render(); renderAttempt(); showToast("此题自评已清除，待复习标记保持原状。");
  });
  document.querySelector(".detail-tabs").addEventListener("keydown", (event) => {
    if (event.altKey || event.ctrlKey || event.metaKey || event.shiftKey) return;
    if (!["ArrowLeft", "ArrowRight", "Home", "End"].includes(event.key)) return;
    event.preventDefault();
    const tabs = ["recall", "notes", "code"];
    const tab = event.key === "Home" ? tabs[0] : event.key === "End" ? tabs[2] : tabs[(tabs.indexOf(currentTab) + (event.key === "ArrowLeft" ? 2 : 1)) % 3];
    showTab(tab);
    $(`${tab}-tab`).focus();
  });
  $("source-select").addEventListener("change", (event) => renderSource(Number(event.target.value)));
  $("wrap-code").addEventListener("click", () => {
    const wrap = $("source-code").parentElement.classList.toggle("wrap-code");
    $("wrap-code").setAttribute("aria-pressed", String(wrap));
  });
  $("detail-review").addEventListener("click", () => { if (currentProblem) toggleReview(currentProblem.id); });
  $("copy-code").addEventListener("click", async () => {
    try { await navigator.clipboard.writeText($("source-code").textContent); showToast("代码已复制。"); }
    catch (_) { showToast("浏览器未允许复制，请选中代码后手动复制。"); }
  });
  $("filter-form").addEventListener("submit", (event) => event.preventDefault());
  $("search").addEventListener("input", (event) => { state.q = event.target.value; state.page = 1; render(); });
  for (const [control, key] of [["topic-filter", "topic"], ["difficulty-filter", "difficulty"], ["record-filter", "record"]]) $(control).addEventListener("change", (event) => { state[key] = event.target.value; state.page = 1; render(); });
  document.addEventListener("click", (event) => {
    const view = event.target.closest("[data-view]");
    if (view) { state.view = view.dataset.view; render(); return; }
    const related = event.target.closest("[data-related]");
    if (related && currentProblem) {
      problemTrail.push(currentReadingPosition());
      const id = related.dataset.related;
      const sequence = readingSequence.ids.includes(id) ? readingSequence : sequenceFor(id);
      openProblem(id, { keepTrail: true, sequence }); return;
    }
    const scope = event.target.closest("[data-scope]");
    if (scope) { state.scope = scope.dataset.scope; state.topic = ""; state.page = 1; render(); return; }
    const topic = event.target.closest("[data-topic]");
    if (topic) { state.topic = state.topic === topic.dataset.topic ? "" : topic.dataset.topic; state.page = 1; render(); return; }
    const clear = event.target.closest("[data-clear]");
    if (clear) { state[clear.dataset.clear] = ""; state.page = 1; render(); return; }
    const review = event.target.closest("[data-review]");
    if (review) { toggleReview(review.dataset.review); return; }
    const row = event.target.closest("[data-problem]");
    if (row) openProblem(row.dataset.problem, { pathId: row.dataset.path });
  });
  $("clear-topic").addEventListener("click", () => { state.topic = ""; state.page = 1; render(); });
  $("reset-filters").addEventListener("click", resetFilters);
  $("empty-reset").addEventListener("click", resetFilters);
  $("previous-page").addEventListener("click", () => { state.page--; render(); });
  $("next-page").addEventListener("click", () => { state.page++; render(); });
  $("random-review").addEventListener("click", () => {
    const problem = model.chooseReview(filtered, reviews, attempts);
    if (problem) openProblem(problem.id);
  });
  document.addEventListener("keydown", (event) => {
    if ($("problem-dialog").open && event.altKey && !event.ctrlKey && !event.metaKey && !event.shiftKey && ["ArrowLeft", "ArrowRight"].includes(event.key) && !["INPUT", "TEXTAREA", "SELECT"].includes(event.target.tagName)) {
      event.preventDefault();
      adjacentProblem(event.key === "ArrowLeft" ? -1 : 1);
    }
    if (event.key === "/" && !$("problem-dialog").open && !event.ctrlKey && !event.metaKey && !event.altKey && !["INPUT", "TEXTAREA", "SELECT"].includes(document.activeElement.tagName)) { event.preventDefault(); $("search").focus(); }
  });
  $("export-review").addEventListener("click", () => {
    const payload = JSON.stringify({ schemaVersion: 1, repository: repo, reviewIds: [...reviews].sort(), attempts }, null, 2);
    const url = URL.createObjectURL(new Blob([payload], { type: "application/json" }));
    const download = element("a");
    download.href = url;
    download.download = "leetcode-review.json";
    document.body.append(download);
    download.click();
    download.remove();
    setTimeout(() => URL.revokeObjectURL(url), 1000);
    showToast("复习标记已导出，可在另一台设备上导入。");
  });
  $("import-review").addEventListener("click", () => $("review-file").click());
  $("review-file").addEventListener("change", async (event) => {
    const file = event.target.files[0];
    if (!file) return;
    try {
      if (file.size > 1024 * 1024) throw new Error("文件过大");
      const payload = JSON.parse(await file.text());
      if (payload.schemaVersion !== 1 || payload.repository !== repo || !Array.isArray(payload.reviewIds) || payload.reviewIds.some((id) => typeof id !== "string")) throw new Error("文件格式不符");
      let added = 0;
      for (const id of payload.reviewIds) if (byId.has(id) && !reviews.has(id)) { reviews.add(id); added++; }
      const imported = model.validAttempts(payload.attempts, byId);
      for (const [id, attempt] of Object.entries(imported)) if (!attempts[id] || attempt.at > attempts[id].at) attempts[id] = attempt;
      try { localStorage.setItem(attemptKey, JSON.stringify(attempts)); } catch (_) { /* Bookmark fallback already explains unavailable storage. */ }
      saveReviews(); render();
      showToast(`已合并 ${added} 道题的复习标记。`);
    } catch (_) { showToast("无法导入，请选择此仓库导出的复习标记 JSON 文件。"); }
    event.target.value = "";
  });
  $("repo-link").href = repo;
  $("readme-link").href = `${repo}#readme`;
  $("stat-active").textContent = String(summary.active);
  $("stat-total").textContent = String(summary.total);
  $("stat-notes").textContent = String(summary.withNotes);
  $("stat-tests").textContent = String(summary.withTests);
  if (Array.isArray(data.unindexed) && data.unindexed.length) {
    const notice = $("mapping-notice");
    notice.append(document.createTextNode(`另有 ${data.unindexed.length} 份历史练习待补题目映射，原始文件均已保留。`), link("查看待补映射清单 ↗", repoURL("docs/CATALOG.md", "待补映射")));
    notice.hidden = false;
  }
  render();
  function openHash() {
    const hash = new URLSearchParams(location.hash.slice(1));
    if (hash.has("problem")) { try { openProblem(hash.get("problem"), { updateHash: false, pathId: hash.get("path") }); } catch (_) { /* Ignore invalid deep links. */ } }
  }
  window.addEventListener("hashchange", openHash);
  openHash();
})();
