//给定两个单词 word1 和 word2 ，返回使得 word1 和 word2 相同所需的最小步数。
//
// 每步 可以删除任意一个字符串中的一个字符。
//
//
//
// 示例 1：
//
//
//输入: word1 = "sea", word2 = "eat"
//输出: 2
//解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
//
//
// 示例 2:
//
//
//输入：word1 = "leetcode", word2 = "etco"
//输出：4
//
//
//
//
// 提示：
//
//
//
// 1 <= word1.length, word2.length <= 500
// word1 和 word2 只包含小写英文字母
//
// Related Topics 字符串 动态规划 👍 389 👎 0

package main

import "fmt"

func main() {
	fmt.Println(minDistance("1", "1"))
}

//leetcode submit region begin(Prohibit modification and deletion)
/*
	其实还是能不能选的问题，能选就看看选了会不会更高，不能选就是保持当前最长 max([i-1][j],[i][j-1])的情况。
*/
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}
	return len(word1) + len(word2) - 2*dp[len(word1)][len(word2)]
}

//leetcode submit region end(Prohibit modification and deletion)
