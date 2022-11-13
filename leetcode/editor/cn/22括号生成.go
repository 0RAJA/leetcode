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
// Related Topics 字符串 动态规划 回溯 👍 2430 👎 0

package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

// leetcode submit region begin(Prohibit modification and deletion)
/*
只要保证l<=r即可
l == r 只能放左括号
l < r 既可以放左括号，也可以放右括号
左括号放完就可以放右括号
*/
func generateParenthesis(n int) (ret []string) {
	var dfs func(s string, l, r int) // true ( false )
	dfs = func(s string, l, r int) {
		if l == 0 && r == 0 {
			ret = append(ret, s)
			return
		}
		if l == r {
			dfs(s+"(", l-1, r)
		} else {
			if l > 0 {
				dfs(s+"(", l-1, r)
			}
			dfs(s+")", l, r-1)
		}
	}
	dfs("", n, n)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
