// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
//
//
// 示例 1:
//
//
// 输入: [3,2,1,5,6,4] 和 k = 2
// 输出: 5
//
//
// 示例 2:
//
//
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
// 输出: 4
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
// 注意：本题与主站 215 题相同： https://leetcode-cn.com/problems/kth-largest-element-in-an-
// array/
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 32 👎 0

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

// leetcode submit region begin(Prohibit modification and deletion)
type Hp076 struct {
	sort.IntSlice
}

func (h *Hp076) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Hp076) Pop() interface{} {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

func findKthLargest(nums []int, k int) int {
	h := &Hp076{IntSlice: make([]int, 0, k+1)}
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}
	for i := k; i < len(nums); i++ {
		heap.Push(h, nums[i])
		heap.Remove(h, 0)
	}
	return heap.Remove(h, 0).(int)
}

// leetcode submit region end(Prohibit modification and deletion)
