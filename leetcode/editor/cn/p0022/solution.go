// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
//
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
//
//
// 示例 2：
//
//
// 输入：n = 1
// 输出：["()"]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
//
// Related Topics 字符串 动态规划 回溯 括号序列 👍 4169 👎 0

package p0022

// leetcode submit region begin(Prohibit modification and deletion)
// dfs
// 1. 确保生成的组合是有效的组合：拼装过程中 left 数 >= right 数；最终 left 数 == right 数
// 2. 控制单个类型括号数不超出 N
// 3. 注意每种组合记录 map 避免走重复路径
func generateParenthesis(n int) (res []string) {
	m := make(map[string]bool)
	var dfs func(leftNum, rightNum int, str string)
	dfs = func(leftNum, rightNum int, str string) {
		if leftNum == n && rightNum == n {
			res = append(res, str)
			return
		}
		if leftNum < n {
			nextStr := str + "("
			if !m[nextStr] {
				m[nextStr] = true
				dfs(leftNum+1, rightNum, nextStr)
			}
		}
		if rightNum < n && leftNum-rightNum >= 1 {
			nextStr := str + ")"
			if !m[nextStr] {
				m[nextStr] = true
				dfs(leftNum, rightNum+1, nextStr)
			}
		}
	}
	dfs(0, 0, "")
	return
}

// leetcode submit region end(Prohibit modification and deletion)
