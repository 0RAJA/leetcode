// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
//
//
// 示例 1:
//
//
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//
//
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// 解释：
//
//
// 在 strs 中没有字符串可以通过重新排列来形成 "bat"。
// 字符串 "nat" 和 "tan" 是字母异位词，因为它们可以重新排列以形成彼此。
// 字符串 "ate" ，"eat" 和 "tea" 是字母异位词，因为它们可以重新排列以形成彼此。
//
//
//
// 示例 2:
//
//
// 输入: strs = [""]
//
//
// 输出: [[""]]
//
//
// 示例 3:
//
//
// 输入: strs = ["a"]
//
//
// 输出: [["a"]]
//
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 10⁴
// 0 <= strs[i].length <= 100
// strs[i] 仅包含小写字母
//
//
// Related Topics 数组 哈希表 字符串 排序 👍 2851 👎 0

package p0049

// leetcode submit region begin(Prohibit modification and deletion)
// 使用 m := make(map[[26]int][]string) 来记录即可
func groupAnagrams(strs []string) (res [][]string) {
	m := make(map[[26]int][]string)
	for _, str := range strs {
		cnt := [26]int{}
		for _, c := range str {
			cnt[c-'a']++
		}
		m[cnt] = append(m[cnt], str)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
