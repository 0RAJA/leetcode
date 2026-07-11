// 设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。
//
// 示例：
//
// 输入： arr = [1,3,5,7,2,4,6,8], K = 4
// 输出： [1,2,3,4]
//
// 提示：
//
// 0 <= len(arr) <= 100000
// 0 <= K <= min(100000, len(arr))
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 111 👎 0
package main

import (
	"container/heap"
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func smallestK(arr []int, k int) (ret []int) {
	if k == 0 {
		return nil
	}
	h := &hp{arr[:k]}
	heap.Init(h)
	for _, v := range arr[k:] {
		if h.IntSlice[0] > v {
			h.IntSlice[0] = v
			heap.Fix(h, 0)
		}
	}
	return h.IntSlice
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }

func (h hp) Push(interface{})   {}
func (hp) Pop() (_ interface{}) { return }

//leetcode submit region end(Prohibit modification and deletion)
