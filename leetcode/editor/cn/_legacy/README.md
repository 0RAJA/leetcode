# Legacy LeetCode Editor 存量区

这里保存历史上由 GoLand LeetCode Editor 插件生成或手动维护的旧题解。

## 目录含义

- `root/`：原 `leetcode/editor/cn/` 根目录下的散落 `.go` / `*_test.go` 文件。
- 其他子目录：原有非标准旧目录，按原名保留。

## 维护规则

- 该目录提交到 Git，用于保留历史题解记录。
- 不在这里继续新增或调试题解。
- 需要维护旧题时，迁移到 `leetcode/editor/cn/pXXXX/solution.go`，再补 `solution_test.go`。
- 目录名前缀 `_legacy` 是有意设计：Go 的 `./...` 包发现会跳过该目录，避免旧题之间的 `TreeNode`、`ListNode`、`Constructor`、`main` 等符号冲突。
