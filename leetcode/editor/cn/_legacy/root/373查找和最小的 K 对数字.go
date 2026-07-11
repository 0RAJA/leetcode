//给定两个以 升序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。
//
// 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
//
// 请找到和最小的 k 个数对 (u1,v1), (u2,v2) ... (uk,vk) 。
//
//
//
// 示例 1:
//
//
//输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
//输出: [1,2],[1,4],[1,6]
//解释: 返回序列中的前 3 对数：
//     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
//
//
// 示例 2:
//
//
//输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
//输出: [1,1],[1,1]
//解释: 返回序列中的前 2 对数：
//     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
//
//
// 示例 3:
//
//
//输入: nums1 = [1,2], nums2 = [3], k = 3
//输出: [1,3],[2,3]
//解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]
//
//
//
//
// 提示:
//
//
// 1 <= nums1.length, nums2.length <= 10⁵
// -10⁹ <= nums1[i], nums2[i] <= 10⁹
// nums1 和 nums2 均为升序排列
// 1 <= k <= 1000
//
// Related Topics 数组 堆（优先队列） 👍 323 👎 0

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	/*
		[-10,-4,0,0,6]
		[3,5,6,7,8,100]
		10
	*/
	nums1 := []int{-10, -4, 0, 0, 6}
	nums2 := []int{3, 5, 6, 7, 8, 100}
	k := 1
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}

// leetcode submit region begin(Prohibit modification and deletion)
type HP [][]int

func (hp *HP) Push(x interface{}) {
	*hp = append(*hp, x.([]int))
}
func (hp *HP) Pop() interface{} {
	x := (*hp)[len(*hp)-1]
	*hp = (*hp)[:len(*hp)-1]
	return x
}
func (hp *HP) Len() int {
	return len(*hp)
}
func (hp *HP) Less(i, j int) bool {
	return (*hp)[i][0] < (*hp)[j][0]
}
func (hp *HP) Swap(i, j int) {
	(*hp)[i], (*hp)[j] = (*hp)[j], (*hp)[i]
}
func kSmallestPairs(nums1 []int, nums2 []int, k int) (ret [][]int) {
	hp := new(HP)
	swap := false
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
		swap = true
	}
	for i := 0; i < k && i < len(nums1); i++ {
		heap.Push(hp, []int{nums1[i] + nums2[0], i, 0})
	}
	ret = make([][]int, 0, k)
	for i := 0; i < k && hp.Len() > 0; i++ {
		p := heap.Pop(hp).([]int)
		a, b := p[1], p[2]
		if swap {
			ret = append(ret, []int{nums2[b], nums1[a]})
		} else {
			ret = append(ret, []int{nums1[a], nums2[b]})
		}
		if b+1 < len(nums2) {
			heap.Push(hp, []int{nums1[a] + nums2[b+1], a, b + 1})
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
