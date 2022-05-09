//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返
//回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
//
//
//
// 示例 1：
//
//
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
//
// 示例 2：
//
//
//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
//
//
// 提示：
//
//
// 1 <= intervals.length <= 10⁴
// intervals[i].length == 2
// 0 <= starti <= endi <= 10⁴
//
//
//
//
// 注意：本题与主站 56 题相同： https://leetcode-cn.com/problems/merge-intervals/
// Related Topics 数组 排序 👍 21 👎 0

package main

import "sort"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) (ret [][]int) {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	var l, r = intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		nL, nR := intervals[i][0], intervals[i][1]
		switch {
		case r < nL:
			ret = append(ret, []int{l, r})
			l, r = nL, nR
		case r >= nL && r <= nR:
			r = nR
		}
	}
	ret = append(ret, []int{l, r})
	return
}

//leetcode submit region end(Prohibit modification and deletion)
