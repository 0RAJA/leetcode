// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
//
//
// 示例 1:
//
//
// 输入: s = "cbaebabacd", p = "abc"
// 输出: [0,6]
// 解释:
// 起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
// 起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
//
//
// 示例 2:
//
//
// 输入: s = "abab", p = "ab"
// 输出: [0,1,2]
// 解释:
// 起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
// 起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
// 起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
//
//
//
//
// 提示:
//
//
// 1 <= s.length, p.length <= 3 * 10⁴
// s 和 p 仅包含小写字母
//
//
// Related Topics 哈希表 字符串 滑动窗口 👍 2004 👎 0

package p0438

// leetcode submit region begin(Prohibit modification and deletion)
// 滑动窗口：维护目标字符串字符分布数组和当前滑动窗口内的字符分布数组
// 每次入右节点，通过出左节点维护窗口数量一致，判断数组是否一致
func findAnagrams(s string, p string) (res []int) {
	if len(p) > len(s) {
		return []int{}
	}
	res = make([]int, 0, len(s))
	targetDistributeList := [26]int{}
	for _, c := range p {
		targetDistributeList[c-'a']++
	}
	nowDistributeList := [26]int{}
	for left, right := 0, 0; right < len(s); right++ {
		nowDistributeList[s[right]-'a']++
		for right-left+1 > len(p) {
			nowDistributeList[s[left]-'a']--
			left++
		}
		if right-left+1 == len(p) && nowDistributeList == targetDistributeList {
			res = append(res, left)
		}
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
