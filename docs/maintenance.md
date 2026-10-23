# leetcode 刷题之旅

这个仓库用于保存 Go 语言 LeetCode 刷题记录、本地练习记录和少量学习笔记。

## 目录约定

- `leetcode/editor/cn/pXXXX/`：GoLand LeetCode Editor 插件按标准模板生成的新题解目录。
- `leetcode/editor/cn/_legacy/`：历史存量题解隔离区，提交到 Git，但被 Go 工具链忽略。
- `leetcode/_legacy/`：更早期的根目录散落练习隔离区，提交到 Git，但被 Go 工具链忽略。
- `leetcode/editor/cn/doc/content/`：题面 Markdown 缓存；公开 GitHub 仓库默认不提交。
- `leetcode/editor/cn/doc/note/`：本地笔记；有内容时适合提交。
- `2026/`：按刷题路线重新整理的题解目录。

## 提交边界

建议提交：

- `*.go` 题解和对应 `*_test.go`。
- `README.md`、专题笔记、刷题路线记录。
- 插件生成的非密钥题库索引，例如 `leetcode/editor/cn/all.json`、`translation.json`；如果嫌变更噪声大，也可以不提交。
- 本地笔记：`leetcode/editor/cn/doc/note/`，但空占位文件不建议提交。

不建议提交：

- `.idea/`、`.vscode/`、`*.iml` 等本机 IDE 状态。
- `.env*`、cookie、token、LeetCode 登录态。
- `.serena/`、`task_plan.md`、`findings.md`、`progress.md` 等本地 agent 工作状态。
- Go 测试/覆盖率/二进制产物。
- 插件生成的题面缓存：`leetcode/editor/cn/doc/content/`、`doc/solution/`、`doc/submission/`。

## 提交规范

提交信息采用 Conventional Commits 风格：

```text
<type>: <summary>

- <problem-id> <problem-title>
- <problem-id> <problem-title>
```

常用 `type`：

- `feat`：新增题解、测试或刷题记录。
- `fix`：修正已有题解、测试或目录命名问题。
- `docs`：只更新文档、笔记、模板说明。
- `refactor`：目录迁移、包拆分、结构调整，不改变题解行为。
- `chore`：仓库配置、忽略规则、非题解维护。

每日题解提交建议使用 `feat`，并在正文列出当天完成的题目，例如：

```text
feat: add daily leetcode solutions

- 435 无重叠区间
- 452 用最少数量的箭引爆气球
- 763 划分字母区间
```

如果同一天同时包含文档或结构调整，优先拆成独立提交；不方便拆分时，在正文继续列出非题解变更。

如果历史上已经提交过 IDE 文件，单靠 `.gitignore` 不会取消跟踪，需要执行：

```bash
git rm -r --cached .idea
```

如果决定公开仓库不提交题面缓存，可取消跟踪已入库缓存但保留本地文件：

```bash
git rm -r --cached leetcode/editor/cn/doc/content leetcode/editor/cn/doc/solution leetcode/editor/cn/doc/submission
```

## 插件模板

GoLand LeetCode 插件的新题生成模板见 `docs/leetcode-editor-template.md`，可复用配置清单见 `docs/leetcode-editor-config.example.md`。这些文档只记录非密钥配置；插件登录态配置仍保留在本机用户级目录，不提交到 Git。

## Legacy 策略

历史散落题解已统一移动到 `_legacy` 目录：

- `_legacy/root/`：原 `leetcode/editor/cn/` 根目录下的散落 `.go` / `*_test.go`。
- `_legacy/<old-dir>/`：原有非标准旧目录，例如 `106/`、`212/`、`subsetsWithDup/`。
- `leetcode/_legacy/root/`：更早期的 `leetcode/` 根目录散落 `.go` 文件。
- `_legacy/` 目录名前缀为 `_`，Go 的 `./...` 包发现会跳过它，避免旧文件继续造成命名冲突和编译失败。

legacy 文件只作为历史记录保留。需要复习、调试或补测试的旧题，应迁移到 `leetcode/editor/cn/pXXXX/` 后再维护。

常用验证命令：

```bash
go test ./2026/...
go test ./leetcode/editor/cn/...
```

提交前先跑目标目录测试，再视变更范围跑 `go test ./...`。

## 验证策略

当前阶段以目标 package 测试为准：

```bash
go test ./2026/...
go test ./leetcode/editor/cn/...
```

如果改动跨越仓库根目录或非 LeetCode Editor 目录，再追加：

```bash
go test ./...
```
