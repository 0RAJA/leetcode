//给你一个按递增顺序排序的数组 arr 和一个整数 k 。数组 arr 由 1 和若干 素数 组成，且其中所有整数互不相同。
//
// 对于每对满足 0 <= i < j < arr.length 的 i 和 j ，可以得到分数 arr[i] / arr[j] 。
//
// 那么第 k 个最小的分数是多少呢? 以长度为 2 的整数数组返回你的答案, 这里 answer[0] == arr[i] 且 answer[1] ==
//arr[j] 。
//
//
// 示例 1：
//
//
//输入：arr = [1,2,3,5], k = 3
//输出：[2,5]
//解释：已构造好的分数,排序后如下所示:
//1/5, 1/3, 2/5, 1/2, 3/5, 2/3
//很明显第三个最小的分数是 2/5
//
//
// 示例 2：
//
//
//输入：arr = [1,7], k = 1
//输出：[1,7]
//
//
//
//
// 提示：
//
//
// 2 <= arr.length <= 1000
// 1 <= arr[i] <= 3 * 10⁴
// arr[0] == 1
// arr[i] 是一个 素数 ，i > 0
// arr 中的所有数字 互不相同 ，且按 严格递增 排序
// 1 <= k <= arr.length * (arr.length - 1) / 2
//
// Related Topics 数组 二分查找 堆（优先队列） 👍 202 👎 0

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	/*
		[1,29,47]
		1
	*/
	fmt.Println(kthSmallestPrimeFraction([]int{1, 29, 47}, 1))
}

// leetcode submit region begin(Prohibit modification and deletion)
type Data struct {
	sum float64
	a   int
	b   int
}

type HpK []Data

func (H HpK) Len() int {
	return len(H)
}

func (H HpK) Less(i, j int) bool {
	return H[i].sum < H[j].sum
}

func (H HpK) Swap(i, j int) {
	H[i], H[j] = H[j], H[i]
}

func (H *HpK) Push(x interface{}) {
	*H = append(*H, x.(Data))
}

func (H *HpK) Pop() interface{} {
	p := (*H)[len(*H)-1]
	*H = (*H)[:len(*H)-1]
	return p
}

func kthSmallestPrimeFraction(arr []int, k int) (ret []int) {
	n := len(arr) - 1
	hp := new(HpK)
	for i := 0; i <= n; i++ {
		for j := i + 1; j < n+1; j++ {
			heap.Push(hp, Data{sum: float64(arr[i]) / float64(arr[j]), a: i, b: j})
		}
	}
	for hp.Len() > 0 {
		p := heap.Pop(hp).(Data)
		a := p.a
		b := p.b
		if k--; k == 0 {
			ret = []int{arr[a], arr[b]}
			break
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
