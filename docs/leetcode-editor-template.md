# GoLand LeetCode 插件模板约定

目标：新生成的 Go 题解按“一题一个目录、一个独立 package”落盘，避免同目录下 `package main` 汇总编译导致的测试失败和命名冲突。

## 当前插件配置位置

GoLand/JetBrains 插件的真实配置在用户级目录，不提交到 Git：

- `~/Library/Application Support/JetBrains/GoLand2025.2/options/leetcode-config.xml`

该文件可能包含 `loginName`、`userCookie`、`LEETCODE_SESSION`、`csrftoken` 等登录态，不能提交。

## 推荐配置

在插件设置中使用以下非密钥配置：

- `Code Type`: `Go`
- `File Path`: 当前仓库根目录，例如 `/Users/raja/Desktop/space/code/study/algorithm/leetcode`
- `Custom Template`: enabled
- `CodeFileName`:

```velocity
p$!velocityTool.leftPadZeros(${question.questionId},4)/solution
```

- `CodeTemplate`:

```velocity
${question.content}

package p$!velocityTool.leftPadZeros(${question.questionId},4)

${question.code}
```

插件会在 `File Path` 后自动拼接 `leetcode/editor/cn/`，并根据语言自动追加 `.go` 后缀。因此 `CodeFileName` 只写 `pXXXX/solution`，不要写 `leetcode/editor/cn/pXXXX/solution.go`。

## 设计依据

- 使用 `question.questionId`：它是 LeetCode 内部数字 ID，能覆盖普通题、LCR、剑指 Offer 等非纯数字前端题号。
- 使用 `p0001` 形式：Go package 必须是合法标识符，不能直接使用 `[1]两数之和`、`LCR 006`、中文题名或带空格路径。
- 每题独立目录：`go test ./leetcode/editor/cn/p0001` 只编译该题，避免 `TreeNode`、`ListNode`、`Constructor`、`main` 等符号冲突。

## 新题目录示例

```text
leetcode/editor/cn/p0001/
  solution.go
  solution_test.go
```

`solution_test.go` 由本地手写，不建议放进插件模板；原因是不同题目的断言、辅助构造函数和输入结构差异较大。

## 迁移注意

- 已生成的旧文件不会因模板修改自动迁移。
- 后续迁移旧题时，优先移动到 `pNNNN/solution.go`，再补 `solution_test.go`。
- 迁移提交应与插件模板提交分开，便于回滚和 review。
