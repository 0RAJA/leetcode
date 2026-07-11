//给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
//
// 示例 1:
//
//
//输入: "abab"
//
//输出: True
//
//解释: 可由子字符串 "ab" 重复两次构成。
//
//
// 示例 2:
//
//
//输入: "aba"
//
//输出: False
//
//
// 示例 3:
//
//
//输入: "abcabcabcabc"
//
//输出: True
//
//解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
//
// Related Topics 字符串 字符串匹配 👍 539 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func repeatedSubstringPattern(s string) bool {
	next := GetNext(s)
	if next[len(s)] == 0 {
		return false
	}
	return len(s)%(len(s)-next[len(s)]) == 0
}

func GetNext(s string) []int {
	next := make([]int, len(s)+1)
	next[0] = -1
	for i, j := 0, -1; i <= len(s)-1; {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

//leetcode submit region end(Prohibit modification and deletion)
