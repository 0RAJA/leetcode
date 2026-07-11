// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。
// 如果 needle 不是 haystack 的一部分，则返回 -1 。
//
// 
//
// 示例 1： 
//
// 
// 输入：haystack = "sadbutsad", needle = "sad"
// 输出：0
// 解释："sad" 在下标 0 和 6 处匹配。
// 第一个匹配项的下标是 0 ，所以返回 0 。
// 
//
// 示例 2： 
//
// 
// 输入：haystack = "leetcode", needle = "leeto"
// 输出：-1
// 解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= haystack.length, needle.length <= 10⁴ 
// haystack 和 needle 仅由小写英文字符组成 
// 
//
// Related Topics 双指针 字符串 字符串匹配 👍 2565 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
// KMP 算法求最长相等前后缀，原理：拿 needle 与自身匹配，求出每个位置的最大相等前后缀（每次当前位置不同则尝试到前一个最长相同前缀的位置，然后拿下一个数再和右侧比较）
// 算出 next 数组后，拿 haystack 与 needle 匹配，每次不匹配则尝试回退 needle 的位置到前一个最长相同前缀处，然后拿下一个数再和右侧比较
func getNext(needle string) (next []int) {
	next = make([]int, len(needle))
	next[0] = 0
	for i, j := 1, 0; i < len(needle); i++ {
		// 如果不匹配则回退 j，找到最长相同前缀处，然后拿下一个数再和右侧比较
		for ; j > 0 && needle[i] != needle[j]; j = next[j-1] {
		}
		// 如果匹配则 j++，稍后 i++
		if needle[i] == needle[j] {
			j++
		}
		// 记录next[i]的最大相等前后缀数(恰好=j)，然后 i++
		next[i] = j
	}
	return next
}
func strStr(haystack string, needle string) int {
	next := getNext(needle)
	for i, j := 0, 0; i < len(haystack); i++ {
		// 如果不匹配则回退 j，找到最长相同前缀处，然后拿下一个数再和右侧比较
		for ; j > 0 && haystack[i] != needle[j]; j = next[j-1] {
		}
		// 如果匹配则 j++，稍后 i++
		if haystack[i] == needle[j] {
			j++
		}
		// 如果 j 到达 needle 的长度，则说明匹配成功
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}

// leetcode submit region end(Prohibit modification and deletion)
