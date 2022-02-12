//使用下面描述的算法可以扰乱字符串 s 得到字符串 t ：
//
// 如果字符串的长度为 1 ，算法停止
// 如果字符串的长度 > 1 ，执行下述步骤：
//
// 在一个随机下标处将字符串分割成两个非空的子字符串。即，如果已知字符串 s ，则可以将其分成两个子字符串 x 和 y ，且满足 s = x + y 。
// 随机 决定是要「交换两个子字符串」还是要「保持这两个子字符串的顺序不变」。即，在执行这一步骤之后，s 可能是 s = x + y 或者 s = y +
//x 。
// 在 x 和 y 这两个子字符串上继续从步骤 1 开始递归执行此算法。
//
//
//
//
// 给你两个 长度相等 的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。如果是，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：s1 = "great", s2 = "rgeat"
//输出：true
//解释：s1 上可能发生的一种情形是：
//"great" --> "gr/eat" // 在一个随机下标处分割得到两个子字符串
//"gr/eat" --> "gr/eat" // 随机决定：「保持这两个子字符串的顺序不变」
//"gr/eat" --> "g/r / e/at" // 在子字符串上递归执行此算法。两个子字符串分别在随机下标处进行一轮分割
//"g/r / e/at" --> "r/g / e/at" // 随机决定：第一组「交换两个子字符串」，第二组「保持这两个子字符串的顺序不变」
//"r/g / e/at" --> "r/g / e/ a/t" // 继续递归执行此算法，将 "at" 分割得到 "a/t"
//"r/g / e/ a/t" --> "r/g / e/ a/t" // 随机决定：「保持这两个子字符串的顺序不变」
//算法终止，结果字符串和 s2 相同，都是 "rgeat"
//这是一种能够扰乱 s1 得到 s2 的情形，可以认为 s2 是 s1 的扰乱字符串，返回 true
//
//
// 示例 2：
//
//
//输入：s1 = "abcde", s2 = "caebd"
//输出：false
//
//
// 示例 3：
//
//
//输入：s1 = "a", s2 = "a"
//输出：true
//
//
//
//
// 提示：
//
//
// s1.length == s2.length
// 1 <= s1.length <= 30
// s1 和 s2 由小写英文字母组成
//
// Related Topics 字符串 动态规划 👍 436 👎 0

package main

import (
	"fmt"
)

func main() {
	s1 := "abcdbdacbdac"
	s2 := "bdacabcdbdac"
	//s1 := "abcde"
	//s2 := "caebd"
	fmt.Println(isScramble3(s1, s2))
}

//leetcode submit region begin(Prohibit modification and deletion)
//朴素搜索
func isScramble1(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if !isScrambleCheck(s1, s2) {
		return false
	}
	for i := 1; i < len(s1); i++ {
		a, b := s1[:i], s1[i:]
		c, d := s2[:i], s2[i:]
		if isScramble1(a, c) && isScramble1(b, d) {
			return true
		}
		e, f := s2[:len(s2)-i], s2[len(s2)-i:]
		if isScramble1(b, e) && isScramble1(a, f) {
			return true
		}
	}
	return false
}

//检查s1,s2的词频和长度
func isScrambleCheck(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	cnt1 := [26]byte{}
	cnt2 := [26]byte{}
	for i := range s1 {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}
	for i := range cnt1 {
		if cnt1[i] != cnt2[i] {
			return false
		}
	}
	return true
}

//记忆化搜索
func isScramble2(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	const (
		EMPTY = iota
		Y
		N
	)
	length := len(s1)
	cache := [30][30][31]int{} //cache[i][j][l] 记录从i开始,长度为l的字符串s1和从j开始,长度为l的字符串s2是否满足题目要求。
	var dfs func(i, j, l int) bool
	dfs = func(i, j, l int) bool {
		if cache[i][j][l] != EMPTY { //先看缓存
			return cache[i][j][l] == Y
		}
		a, b := s1[i:i+l], s2[j:j+l] //两个全长字符串
		if a == b {
			cache[i][j][l] = Y
			return true
		}
		if !isScrambleCheck(a, b) { //检查长度和字频
			cache[i][j][l] = N
			return false
		}
		for k := 1; k < l; k++ { //内部分割递归检查
			if dfs(i, j, k) && dfs(i+k, j+k, l-k) {
				cache[i][j][l] = Y
				return true
			}
			if dfs(i+k, j, l-k) && dfs(i, j+l-k, k) {
				cache[i][j][l] = Y
				return true
			}
		}
		cache[i][j][l] = N
		return false
	}
	return dfs(0, 0, length)
}

//区间DP
func isScramble3(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	length := len(s1)
	dp := [30][30][31]bool{} //同方法2的cache
	//先处理长度为1的情况
	for i := range s1 {
		for j := range s2 {
			dp[i][j][1] = s1[i] == s2[j]
		}
	}
	//再处理其他情况
	for l := 2; l <= length; l++ {
		for i := 0; i <= length-l; i++ {
			for j := 0; j <= length-l; j++ {
				for k := 1; k < l; k++ {
					a := dp[i][j][k] && dp[i+k][j+k][l-k]
					b := dp[i+k][j][l-k] && dp[i][j+l-k][k]
					if a || b {
						dp[i][j][l] = true
					}
				}
			}
		}
	}
	return dp[0][0][length]
}

//leetcode submit region end(Prohibit modification and deletion)
