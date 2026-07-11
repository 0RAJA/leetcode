// 给你一个字符串 s 和一个整数 k 。你可以选择字符串中的任一字符，并将其更改为任何其他大写英文字符。该操作最多可执行 k 次。
//
// 在执行上述操作后，返回 包含相同字母的最长子字符串的长度。 
//
// 
//
// 示例 1： 
//
// 
// 输入：s = "ABAB", k = 2
// 输出：4
// 解释：用两个'A'替换为两个'B',反之亦然。
// 
//
// 示例 2： 
//
// 
// 输入：s = "AABABBA", k = 1
// 输出：4
// 解释：
// 将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
// 子串 "BBBB" 有最长重复字母, 答案为 4。
// 可能存在其他的方法来得到同样的结果。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 10⁵ 
// s 仅由大写英文字母组成 
// 0 <= k <= s.length 
// 
//
// Related Topics 哈希表 字符串 滑动窗口 👍 960 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
// 当前区间内想满足要求需要替换的字符数 = 区间总字符数 - 区间最多相同字符总数 <= k 即可满足要求
func characterReplacement(s string, k int) (res int) {
	res = -1
	maxCharCount := 0 // 当前区间最多相同字符数
	charCountMap := make(map[byte]int)
	for l, r := 0, 0; r < len(s); r++ {
		charCountMap[s[r]]++
		maxCharCount = max(maxCharCount, charCountMap[s[r]])
		// 注意 maxCount 是一个历史最大值，用它来控制窗口大小上限即可，这个就算偏大一些也不影响最终答案
		for r-l+1-maxCharCount > k {
			charCountMap[s[l]]--
			l++
		}
		res = max(r-l+1, res)
	}
	if res == -1 {
		return 0
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
