//给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
//
//
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
//
//
// 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
//
//
// 示例 1：
//
//
//输入：s = "aa", p = "a"
//输出：false
//解释："a" 无法匹配 "aa" 整个字符串。
//
//
// 示例 2:
//
//
//输入：s = "aa", p = "a*"
//输出：true
//解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
//
//
// 示例 3：
//
//
//输入：s = "ab", p = ".*"
//输出：true
//解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 20
// 1 <= p.length <= 30
// s 只包含从 a-z 的小写字母。
// p 只包含从 a-z 的小写字母，以及字符 . 和 *。
// 保证每次出现字符 * 时，前面都匹配到有效的字符
//
// Related Topics 递归 字符串 动态规划 👍 2714 👎 0

package main

import "fmt"

func main() {
	s := "a"
	p := "ab*"
	fmt.Println(isMatchDP(s, p))
}

//leetcode submit region begin(Prohibit modification and deletion)
//记忆化搜索
func isMatch(s string, p string) bool {
	parent := make([][2]byte, 0, len(p)) //把原有模板压缩一下，来将数字和其可以出现的次数来绑定在一起。
	for i := range p {
		if p[i] != '*' {
			var nextc byte //为了方便，直接把当前位和下一位的数字绑定在一起。
			if i != len(p)-1 {
				nextc = p[i+1]
			}
			parent = append(parent, [2]byte{p[i], nextc})
		}
	}
	cache := [21]map[int]bool{} //缓存
	for i := range cache {
		cache[i] = make(map[int]bool)
	}
	var dfs func(i, j int) bool
	dfs = func(i, j int) (ret bool) { // 深搜
		if j == len(parent) { // 模板串匹配完了，直接看目标串是否匹配完。都匹配完了就说明匹配成功。
			return i == len(s)
		}
		if i == len(s) { // 目标串匹配完了，不能直接看模板串是否匹配完，特殊情况例如： "a" "ab*b*" 这类。
			for k := j; k < len(parent); k++ {
				if parent[k][1] != '*' {
					return false
				}
			}
			return true
		}
		if v, ok := cache[i][j]; ok { //有缓存直接走缓存。
			return v
		}
		defer func() { cache[i][j] = ret }() // 记录缓存。
		a, b := parent[j][0], parent[j][1]   // a,b 用来判断 当前的两个字符串是否可以匹配，以及可以进行的匹配的次数（0,1,n）
		if b == '*' {                        // 如果可以进行多次匹配（0,1,n 次匹配）
			if a == '.' || a == s[i] { // 将0,1,n次匹配的情况都搜一遍。
				return dfs(i, j+1) || dfs(i+1, j+1) || dfs(i+1, j)
			}
			return dfs(i, j+1) //因为匹配不上，所以只能选择0次匹配
		}
		if a == '.' || a == s[i] { //如果可以匹配就进行1次匹配。
			return dfs(i+1, j+1)
		}
		return false // 连一次匹配都匹配不上就说明失败了。
	}
	return dfs(0, 0)
}

/*
dp[i][j] == {
				if dp[i][j] == '*' :
					dp[i-1][j] || dp[i][j-2] (match(s[i],p[j-1]))
					dp[i][j-2] (!match(s[i],p[j-1]))
				else :
					dp[i-1][j-1] (match(s[i],p[j]))
					false (!match(s[i],p[j]))
			}
*/
func isMatchDP(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	isOK := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if s[i-1] == p[j-1] || p[j-1] == '.' {
			return true
		}
		return false
	}
	dp[0][0] = true
	for i := 0; i <= len(s); i++ { //s = "" p = "a***"
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if isOK(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if isOK(i, j) {
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}

//leetcode submit region end(Prohibit modification and deletion)
