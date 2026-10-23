const test = require('node:test');
const assert = require('node:assert/strict');
const vm = require('node:vm');
const fs = require('node:fs');
const path = require('node:path');
const model = require('../docs/explorer/study-model.js');
const explorer = path.join(__dirname, '../docs/explorer');
const context = { window: {} };
vm.runInNewContext(fs.readFileSync(path.join(explorer, 'data.js'), 'utf8'), context);
vm.runInNewContext(fs.readFileSync(path.join(explorer, 'learning-paths.js'), 'utf8'), context);
const data = context.window.LEETCODE_DATA;
const paths = context.window.LEETCODE_PATHS.paths;
const problems = data.problems;
const byId = new Map(problems.map((problem) => [problem.id, problem]));

test('summary counts question records, not versions, and follows changed records', () => {
  const rows = [
    {active: true, notes: ['note'], sources: [{tests: ['a_test.go']}, {tests: ['b_test.go']}]},
    {active: false, notes: [], sources: [{tests: []}]},
  ];
  assert.deepEqual(model.catalogSummary(rows), {total: 2, active: 1, archived: 1, withNotes: 1, withTests: 1});
  rows[1].active = true;
  rows[1].notes.push('new note');
  assert.deepEqual(model.catalogSummary(rows), {total: 2, active: 2, archived: 0, withNotes: 2, withTests: 1});
  assert.deepEqual(model.catalogSummary([]), {total: 0, active: 0, archived: 0, withNotes: 0, withTests: 0});
});

test('curated binary-search neighbor explains the boundary change', () => {
  const next = model.relatedProblems(byId.get('704'), problems, paths);
  assert.equal(next[0].problem.id, '34');
  assert.equal(next[0].kind, 'path');
  assert.match(next[0].reason, /边界|左|右/);
  assert.ok(next.every((item) => item.problem.id !== '704' && item.reason));
  assert.equal(new Set(next.map((item) => item.problem.id)).size, next.length);
});

test('generic array overlap alone does not create a misleading relationship', () => {
  const rows = ['a', 'b'].map((id) => ({ id, topics: ['数组'], notes: [], active: true }));
  assert.deepEqual(model.relatedProblems(rows[0], rows, []), []);
});

test('specific-topic relationships are explicitly labeled automatic', () => {
  const next = model.relatedProblems(byId.get('704'), problems, []);
  assert.ok(next.length);
  assert.ok(next.every((item) => item.kind === 'topic' && item.shared.includes('二分查找')));
});

test('review choice prioritizes marked questions in the visible set', () => {
  const visible = [byId.get('704'), byId.get('34')];
  assert.equal(model.chooseReview(visible, new Set(['34']), {}, () => 0).id, '34');
  assert.equal(model.chooseReview(visible, new Set(['999']), {}, () => 0).id, '704');
  assert.equal(model.chooseReview([], new Set(), {}), null);
});

test('review choice uses self-report history without treating opened code as mastery', () => {
  const visible = [byId.get('704'), byId.get('34')];
  assert.equal(model.chooseReview(visible, new Set(), {'704': {at: 2}}, () => 0).id, '34');
  assert.equal(model.chooseReview(visible, new Set(), {'704': {at: 2}, '34': {at: 1}}).id, '34');
});

test('import validation drops unknown ids, invalid status and malformed timestamps', () => {
  const clean = model.validAttempts({
    '704': {result: 'recalled', at: 10}, '34': {result: 'AC', at: 20},
    '33': {result: 'needs-practice', at: 'today'}, 'unknown': {result: 'recalled', at: 2}
  }, new Set(['704', '34', '33']));
  assert.deepEqual(Object.keys(clean), ['704']);
  assert.deepEqual(clean['704'], {result: 'recalled', at: 10});
});

test('vendored highlighter recognizes Go and safely escapes source content', () => {
  const highlighter = {};
  vm.runInNewContext(fs.readFileSync(path.join(explorer, 'vendor/highlight/highlight.min.js'), 'utf8'), highlighter);
  assert.ok(highlighter.hljs.getLanguage('go'));
  const html = highlighter.hljs.highlight('func add(a int) int { return a + 1 } // <script>', {language: 'go'}).value;
  assert.match(html, /hljs-keyword/);
  assert.match(html, /hljs-number/);
  assert.ok(!html.includes('<script>'));
  assert.ok(html.includes('&lt;script&gt;'));
});

test('reading sequence follows a curated path instead of numeric question order', () => {
  const route = paths.find((item) => item.id === 'binary-search');
  const sequence = model.readingSequence('704', route.steps.map((step) => step.id), route.title);
  assert.deepEqual(sequence.ids, ['704', '34', '33']);
  assert.equal(sequence.label, route.title);
});

test('self-rating cannot remove the current question from an existing reading session', () => {
  const visibleIds = ['1', '15', '18'];
  const sequence = model.readingSequence('15', visibleIds, '待复习 · 双指针');
  visibleIds.splice(1, 1);
  assert.deepEqual(sequence.ids, ['1', '15', '18']);
  assert.equal(sequence.ids[sequence.ids.indexOf('15') + 1], '18');
});

test('a deep link outside the current filter opens an explicit single-question session', () => {
  const sequence = model.readingSequence('704', ['1', '15', '18'], '当前练习');
  assert.deepEqual(sequence, {ids: ['704'], label: '单题阅读'});
});
