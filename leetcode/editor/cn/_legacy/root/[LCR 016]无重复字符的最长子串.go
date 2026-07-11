// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长连续子字符串 的长度。
//
// 
//
// 示例 1： 
//
// 
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子字符串是 "abc"，所以其长度为 3。
// 
//
// 示例 2： 
//
// 
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子字符串是 "b"，所以其长度为 1。
// 
//
// 示例 3： 
//
// 
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
// 
//
// 示例 4： 
//
// 
// 输入: s = ""
// 输出: 0
// 
//
// 
//
// 提示： 
//
// 
// 0 <= s.length <= 5 * 10⁴ 
// s 由英文字母、数字、符号和空格组成 
// 
//
// 
//
// 
// 注意：本题与主站 3 题相同： https://leetcode.cn/problems/longest-substring-without-
// repeating-characters/
//
// Related Topics 哈希表 字符串 滑动窗口 👍 130 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
// 找最长的不含重复字符的子串长度。
// right 扩张
// 如果某个字符重复了，就移动 left，直到不重复
// 每次窗口合法后更新最大长度
func lengthOfLongestSubstring(s string) (res int) {
	res = -1
	charCountMap := make(map[byte]int) // 记录字符出现次数
	left := 0
	for right := 0; right < len(s); right++ {
		// 把新字符加进去
		charCountMap[s[right]]++
		// 剔除左边的重复字符
		for charCountMap[s[right]] > 1 {
			charCountMap[s[left]]--
			left++
		}
		// 计算当前符合要求的子串长度
		res = max(res, right-left+1)
	}
	if res == -1 {
		return 0
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
