//给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
//
// '?' 可以匹配任何单个字符。
//'*' 可以匹配任意字符串（包括空字符串）。
//
//
// 两个字符串完全匹配才算匹配成功。
//
// 说明:
//
//
// s 可能为空，且只包含从 a-z 的小写字母。
// p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
//
//
// 示例 1:
//
// 输入:
//s = "aa"
//p = "a"
//输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。
//
// 示例 2:
//
// 输入:
//s = "aa"
//p = "*"
//输出: true
//解释: '*' 可以匹配任意字符串。
//
//
// 示例 3:
//
// 输入:
//s = "cb"
//p = "?a"
//输出: false
//解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
//
//
// 示例 4:
//
// 输入:
//s = "adceb"
//p = "*a*b"
//输出: true
//解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
//
//
// 示例 5:
//
// 输入:
//s = "acdcb"
//p = "a*c?b"
//输出: false
// Related Topics 贪心 递归 字符串 动态规划 👍 820 👎 0

package main

import "fmt"

func main() {
	s := "acdcb"
	p := "a*c?b"
	fmt.Println(isMatch2(s, p))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isMatch2(s string, p string) bool {
	var dfs func(i, j int) bool
	cache := make([]map[int]bool, len(s))
	for i := range cache {
		cache[i] = make(map[int]bool)
	}
	dfs = func(i, j int) (ret bool) {
		if j == len(p) {
			return i == len(s)
		}
		if i == len(s) {
			for k := j; k < len(p); k++ {
				if p[k] != '*' {
					return false
				}
			}
			return true
		}
		if v, ok := cache[i][j]; ok {
			return v
		}
		defer func() { cache[i][j] = ret }()
		switch p[j] {
		case '?':
			return dfs(i+1, j+1)
		case '*':
			return dfs(i+1, j) || dfs(i, j+1) || dfs(i+1, j+1)
		default:
			if s[i] == p[j] {
				return dfs(i+1, j+1)
			}
			return false
		}
	}
	return dfs(0, 0)
}

/*
	dp[i][j] = {
		if p[j] == '*' :
			dp[i][j] = dp[i-1][j] || dp[i][j-1]
		else if p[j] == '？' :
			dp[i][j] = dp[i-1][j-1]
		else :
			dp[i][j] = s[i]==p[j] && dp[i-1][j-1]
	}
*/
func isMatch2DP(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			switch p[j-1] {
			case '*':
				dp[i][j] = dp[i][j-1] || i >= 1 && dp[i-1][j]
			case '?':
				dp[i][j] = i >= 1 && dp[i-1][j-1]
			default:
				dp[i][j] = i >= 1 && s[i-1] == p[j-1] && dp[i-1][j-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}

//leetcode submit region end(Prohibit modification and deletion)
