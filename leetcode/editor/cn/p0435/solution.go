// 给定一个区间的集合 intervals ，其中 intervals[i] = [starti, endi] 。返回 需要移除区间的最小数量，使剩余区间互不重
// 叠 。
//
// 注意 只在一点上接触的区间是 不重叠的。例如 [1, 2] 和 [2, 3] 是不重叠的。
//
//
//
// 示例 1:
//
//
// 输入: intervals = [[1,2],[2,3],[3,4],[1,3]]
// 输出: 1
// 解释: 移除 [1,3] 后，剩下的区间没有重叠。
//
//
// 示例 2:
//
//
// 输入: intervals = [ [1,2], [1,2], [1,2] ]
// 输出: 2
// 解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
//
//
// 示例 3:
//
//
// 输入: intervals = [ [1,2], [2,3] ]
// 输出: 0
// 解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
//
//
//
//
// 提示:
//
//
// 1 <= intervals.length <= 10⁵
// intervals[i].length == 2
// -5 * 10⁴ <= starti < endi <= 5 * 10⁴
//
//
// Related Topics 贪心 数组 动态规划 排序 👍 1290 👎 0

package p0435

import (
	"sort"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 贪心：按 start_idx 从小到大排序，判断当前节点的 start_idx 和 上一个节点的 end_idx 如果重叠就去掉两个节点中 end_idx 最小的那个
func eraseOverlapIntervals(intervals [][]int) (res int) {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] <= intervals[j][1]
		}
		return false
	})
	lastEndIdx := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < lastEndIdx {
			// 不一定去掉当前，而是选择去掉两者中 end_idx 较小的，给后者留更多空间
			lastEndIdx = min(lastEndIdx, intervals[i][1])
			res++
		} else {
			lastEndIdx = intervals[i][1]
		}
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
