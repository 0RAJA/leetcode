// 给出由小写字母组成的字符串 s，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
//
// 在 s 上反复执行重复项删除操作，直到无法继续删除。 
//
// 在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。 
//
// 
//
// 示例： 
//
// 
// 输入："abbaca"
// 输出："ca"
// 解释：
// 例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。之后我们得到字符串 "aaca"，其中又
// 只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 10⁵ 
// s 仅由小写英文字母组成。 
// 
//
// Related Topics 栈 字符串 👍 738 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(s string) string {
	// 维护一个栈，来一个新的字符，如果与栈顶相同则弹出，否则入栈
	stack := make([]rune, 0, len(s))
	for _, ch := range s {
		if len(stack) > 0 && stack[len(stack)-1] == ch {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, ch)
	}
	return string(stack)
}

// leetcode submit region end(Prohibit modification and deletion)
