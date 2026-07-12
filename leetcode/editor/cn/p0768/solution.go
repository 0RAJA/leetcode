// 给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。例如，字符串 "ababcc" 能够被分为 ["abab",
// "cc"]，但类似 ["aba", "bcc"] 或 ["ab", "ab", "cc"] 的划分是非法的。
//
// 注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。
//
// 返回一个表示每个字符串片段的长度的列表。
//
//
// 示例 1：
//
//
// 输入：s = "ababcbacadefegdehijhklij"
// 输出：[9,7,8]
// 解释：
// 划分结果为 "ababcbaca"、"defegde"、"hijhklij" 。
// 每个字母最多出现在一个片段中。
// 像 "ababcbacadefegde", "hijhklij" 这样的划分是错误的，因为划分的片段数较少。
//
// 示例 2：
//
//
// 输入：s = "eccbbbbdec"
// 输出：[10]
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 500
// s 仅由小写英文字母组成
//
//
// Related Topics 贪心 哈希表 双指针 字符串 👍 1431 👎 0

package p0768

import (
	"sort"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 本质还是合并区间，先遍历 s 统计每个字符出现的 idx 起止区间，然后统计每个字符的 start_idx 和 end_idx，然后合并所有字符的 起止区间，不能有重叠
func partitionLabels(s string) (res []int) {
	charPartition := make([][]int, 26)
	// 统计每个字符出现的 idx 起止列表,注意可能出现一些空位
	for idx, c := range s {
		charPartitionIdx := c - 'a'
		if len(charPartition[charPartitionIdx]) == 0 {
			charPartition[charPartitionIdx] = []int{idx, idx}
		} else {
			charPartition[charPartitionIdx][1] = idx
		}
	}
	// 针对起止区间按 start_idx 排序，注意有一些字符空位的放到最后后续直接跳过
	sort.Slice(charPartition, func(i, j int) bool {
		if len(charPartition[i]) > len(charPartition[j]) {
			return true
		} else if len(charPartition[i]) < len(charPartition[j]) {
			return false
		} else if len(charPartition[i]) != 2 || len(charPartition[j]) != 2 {
			return true
		}
		return charPartition[i][0] < charPartition[j][0]
	})
	// 保存每轮的区间边界
	startIdx := charPartition[0][0]
	endIdx := charPartition[0][1]
	// 合并区间,注意跳过空值
	for i := 1; i < len(charPartition) && len(charPartition[i]) > 0; i++ {
		// 区间冲突需要合并,更新右边界
		if charPartition[i][0] < endIdx {
			endIdx = max(charPartition[i][1], endIdx)
		} else {
			// 没有区间冲突则追加上轮结果，更新本轮区间边界
			res = append(res, endIdx-startIdx+1)
			startIdx = charPartition[i][0]
			endIdx = charPartition[i][1]
		}
	}
	// 注意追加最后一个边界
	res = append(res, endIdx-startIdx+1)
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
