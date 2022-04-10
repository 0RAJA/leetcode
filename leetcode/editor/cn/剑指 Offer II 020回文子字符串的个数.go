//给定一个字符串 s ，请计算这个字符串中有多少个回文子字符串。
//
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
//
//
//
// 示例 1：
//
//
//输入：s = "abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
//
//
// 示例 2：
//
//
//输入：s = "aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 由小写英文字母组成
//
//
//
//
// 注意：本题与主站 647 题相同：https://leetcode-cn.com/problems/palindromic-substrings/
// Related Topics 字符串 动态规划 👍 40 👎 0

package main

import "fmt"

func main() {
	fmt.Println(countSubstrings("aaa"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countSubstrings(s string) (ret int) {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	for i := len(dp) - 1; i >= 0; i-- {
		for j := i; j < len(dp); j++ {
			dp[i][j] = i == j || s[i] == s[j] && (i+1 >= j-1 || dp[i+1][j-1])
			if dp[i][j] {
				ret++
			}
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
