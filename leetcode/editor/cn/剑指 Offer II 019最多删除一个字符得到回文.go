//给定一个非空字符串 s，请判断如果 最多 从字符串中删除一个字符能否得到一个回文字符串。
//
//
//
// 示例 1:
//
//
//输入: s = "aba"
//输出: true
//
//
// 示例 2:
//
//
//输入: s = "abca"
//输出: true
//解释: 可以删除 "c" 字符 或者 "b" 字符
//
//
// 示例 3:
//
//
//输入: s = "abc"
//输出: false
//
//
//
// 提示:
//
//
// 1 <= s.length <= 10⁵
// s 由小写英文字母组成
//
//
//
//
// 注意：本题与主站 680 题相同： https://leetcode-cn.com/problems/valid-palindrome-ii/
// Related Topics 贪心 双指针 字符串 👍 29 👎 0

package main

import "fmt"

func main() {
	fmt.Println(validPalindrome("abc"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func validPalindrome(s string) bool {
	isPalindrome := func(i, j int) bool {
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return isPalindrome(i+1, j) || isPalindrome(i, j-1)
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
