# LeetCode Editor 无密钥配置示例

目标：记录可复用的插件配置，不提交真实 `leetcode-config.xml`、登录态、cookie 或项目本地映射。

## 不提交的真实配置

以下文件属于本机状态，不进入 Git：

- `~/Library/Application Support/JetBrains/GoLand*/options/leetcode-config.xml`
- `~/Library/Application Support/JetBrains/GoLand*/options/leetcode-pro-config.xml`
- `.idea/leetcode/editor.xml`

原因：

- 用户级 XML 可能包含 `loginName`、`userCookie`、`LEETCODE_SESSION`、`csrftoken`。
- `.idea/leetcode/editor.xml` 是本地题目路径映射，包含 `$PROJECT_DIR$`、旧题缓存和插件点击记录，迁移到其他机器不稳定。

## 可复用配置

在 GoLand 中打开 LeetCode Editor 插件设置，按以下值配置：

| 配置项 | 值 |
| --- | --- |
| Code Type | `Go` |
| File Path | 仓库根目录，例如 `/Users/raja/Desktop/space/code/study/algorithm/leetcode` |
| Custom Template | enabled |
| CodeFileName | `p$!velocityTool.leftPadZeros(${question.questionId},4)/solution` |

CodeTemplate：

```velocity
${question.content}

package p$!velocityTool.leftPadZeros(${question.questionId},4)

${question.code}
```

## 生成结果

插件会自动在 `File Path` 后拼接 `leetcode/editor/cn/`，并为 Go 语言追加 `.go` 后缀。

期望新题生成结果：

```text
leetcode/editor/cn/p0001/solution.go
leetcode/editor/cn/p0003/solution.go
```

不要把 `CodeFileName` 写成：

```velocity
leetcode/editor/cn/p$!velocityTool.leftPadZeros(${question.questionId},4)/solution.go
```

这样会导致路径重复或插件继续落到旧格式。

注意：这里使用插件内部 `questionId`，优先保证 GoLand LeetCode 插件点击题目、打开文件、提交记录和本地文件路径一致。commit 正文中再使用 LeetCode 页面展示题号记录当天完成的题。

## 新题维护流程

1. 插件生成 `leetcode/editor/cn/pXXXX/solution.go`。
2. 本地手写 `leetcode/editor/cn/pXXXX/solution_test.go`。
3. 运行 `go test ./leetcode/editor/cn/pXXXX`。
4. 确认无密钥和无 `.idea/` 变更后再提交。
