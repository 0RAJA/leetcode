# leetcode 刷题之旅

这个仓库用于保存 Go 语言 LeetCode 刷题记录、本地练习记录和少量学习笔记。

## 目录约定

- `leetcode/editor/cn/`：GoLand LeetCode Editor 插件生成的题解、题面缓存和题库索引。
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

如果历史上已经提交过 IDE 文件，单靠 `.gitignore` 不会取消跟踪，需要执行：

```bash
git rm -r --cached .idea
```

如果决定公开仓库不提交题面缓存，可取消跟踪已入库缓存但保留本地文件：

```bash
git rm -r --cached leetcode/editor/cn/doc/content leetcode/editor/cn/doc/solution leetcode/editor/cn/doc/submission
```

## 插件模板

GoLand LeetCode 插件的新题生成模板见 `docs/leetcode-editor-template.md`。该模板约定新 Go 题解生成到独立目录和独立 package，插件登录态配置仍保留在本机用户级目录，不提交到 Git。

## Legacy 策略

`leetcode/editor/cn` 根目录下的历史散落题解属于 legacy 存量。legacy 文件不作为新增题模板，也不作为日常 `go test ./...` 的完成标准。需要复习、调试或补测试的旧题，应迁移到 `leetcode/editor/cn/pXXXX/` 后再维护。

常用验证命令：

```bash
go test ./2026/...
go test ./leetcode/editor/cn/pXXXX
```

在 legacy 完成隔离前，不要求 `go test ./...` 全量通过。

## 验证策略

当前阶段以目标 package 测试为准：

```bash
go test ./2026/...
go test ./leetcode/editor/cn/pXXXX
```

当 legacy 散落文件完成隔离或迁移后，再把 `go test ./...` 作为提交前必跑命令。
