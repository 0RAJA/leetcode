// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号 子串 的长度。
//
// 左右括号匹配，即每个左括号都有对应的右括号将其闭合的字符串是格式正确的，比如 "(()())"。
//
//
//
//
//
// 示例 1：
//
//
//
// 输入：s = "(()"
// 输出：2
// 解释：最长有效括号子串是 "()"
//
//
// 示例 2：
//
//
// 输入：s = ")()())"
// 输出：4
// 解释：最长有效括号子串是 "()()"
//
//
// 示例 3：
//
//
// 输入：s = ""
// 输出：0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 3 * 10⁴
// s[i] 为 '(' 或 ')'
//
//
//
//
// Related Topics 栈 字符串 动态规划 括号序列 👍 2949 👎 0

package p0032

// leetcode submit region begin(Prohibit modification and deletion)
/*
栈里存：下标,先放一个：-1 作为虚拟边界；
遇到 ( 把下标入栈： stack = append(stack, i)
遇到 ) 先弹出一个下标；如果弹完后栈为空：说明当前右括号没有匹配对象，把当前下标当成新的边界：stack = append(stack, i)
否则：当前合法长度 = i - stackTop
*/
func longestValidParentheses(s string) int {
	stack := []int{-1}
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}
		// 当前是 ')'
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			// 当前 ')' 无法匹配，
			// 作为新的非法边界
			stack = append(stack, i)
		} else {
			length := i - stack[len(stack)-1]
			ans = max(ans, length)
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
