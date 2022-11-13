// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//
//
//
//
// 示例 1：
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
// Related Topics 栈 字符串 动态规划 👍 1565 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
// 把所有不能匹配的括号位置置为1,则题目转换为求最长的0的长度。
func longestValidParentheses(s string) (ret int) {
	isOK := make([]bool, len(s))
	stack := make([]int, 0, len(s)/2)
	// 用栈校验括号是否正确
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 { // 无法匹配
				isOK[i] = true
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	// 标记无法匹配的括号
	for i := range stack {
		isOK[stack[i]] = true
	}
	// 计算最长连续0的长度
	cnt := 0
	for i := range isOK {
		if !isOK[i] {
			cnt++
		} else {
			if cnt > ret {
				ret = cnt
			}
			cnt = 0
		}
	}
	if cnt > ret {
		ret = cnt
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
