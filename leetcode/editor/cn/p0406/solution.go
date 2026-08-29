// 假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。每个 people[i] = [hi, ki] 表示第 i
// 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。
//
// 请你重新构造并返回输入数组 people 所表示的队列。返回的队列应该格式化为数组 queue ，其中 queue[j] = [hj, kj] 是队列中第
// j 个人的属性（queue[0] 是排在队列前面的人）。
//
//
//
//
//
//
// 示例 1：
//
//
// 输入：people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
// 输出：[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
// 解释：
// 编号为 0 的人身高为 5 ，没有身高更高或者相同的人排在他前面。
// 编号为 1 的人身高为 7 ，没有身高更高或者相同的人排在他前面。
// 编号为 2 的人身高为 5 ，有 2 个身高更高或者相同的人排在他前面，即编号为 0 和 1 的人。
// 编号为 3 的人身高为 6 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
// 编号为 4 的人身高为 4 ，有 4 个身高更高或者相同的人排在他前面，即编号为 0、1、2、3 的人。
// 编号为 5 的人身高为 7 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
// 因此 [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]] 是重新构造后的队列。
//
//
// 示例 2：
//
//
// 输入：people = [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]
// 输出：[[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
//
//
//
//
// 提示：
//
//
// 1 <= people.length <= 2000
// 0 <= hi <= 10⁶
// 0 <= ki < people.length
// 题目数据确保队列可以被重建
//
//
// Related Topics 树状数组 线段树 数组 排序 👍 1984 👎 0

package p0406

import (
	"sort"
)

// leetcode submit region begin(Prohibit modification and deletion)
/*
核心思路：
[h, k] 表示：
身高为 h，前面恰好有 k 个身高 >= h 的人。
关键是：
先处理高个子，再处理矮个子。
因为矮个子后面插入队列时，不会影响高个子的 k，当前队列里已经存在的人，身高一定都 >= 当前人，因此都是矮个子需要的人
*/
// 先按 h 从大到小，k 从小到大排序
// 然后从 h 由大到小遍历，且由于是按 h 从大到小，直接按 k 插入则说明需要前面有 k 个 >= 自己的元素
func reconstructQueue(people [][]int) (res [][]int) {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] {
			return true
		} else if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return false
	})
	for _, person := range people {
		k := person[1]
		// 长度先 +1
		res = append(res, nil)
		// [k:] 整体向后移动一个位置
		copy(res[k+1:], res[k:])
		// 插入 person
		res[k] = person
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
