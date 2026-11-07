// 给你一个字符串 s，找到 s 中最长的 回文 子串。
//
//
//
// 示例 1：
//
//
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
// 输入：s = "cbbd"
// 输出："bb"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
//
//
// Related Topics 双指针 字符串 动态规划 Manacher 算法 👍 8205 👎 0

package p0005

// leetcode submit region begin(Prohibit modification and deletion)
// 中心拓展法：每次尝试1个元素 or 2个元素为中心向外拓展 O(N2)
func longestPalindrome(s string) (res string) {
	tryMatch := func(l, r int) {
		tmpString := ""
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if l == r {
				tmpString += string(s[l])
			} else {
				tmpString = string(s[l]) + tmpString + string(s[r])
			}
			l--
			r++
		}
		if len(tmpString) > len(res) {
			res = tmpString
		}
	}
	for i := 0; i < len(s); i++ {
		tryMatch(i, i)
		tryMatch(i, i+1)
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
