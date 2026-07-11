//如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数
//值排序之后中间两个数的平均值。
//
// 例如，
//
// [2,3,4] 的中位数是 3
//
// [2,3] 的中位数是 (2 + 3) / 2 = 2.5
//
// 设计一个支持以下两种操作的数据结构：
//
//
// void addNum(int num) - 从数据流中添加一个整数到数据结构中。
// double findMedian() - 返回目前所有元素的中位数。
//
//
// 示例 1：
//
// 输入：
//["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
//[[],[1],[2],[],[3],[]]
//输出：[null,null,null,1.50000,null,2.00000]
//
//
// 示例 2：
//
// 输入：
//["MedianFinder","addNum","findMedian","addNum","findMedian"]
//[[],[2],[],[3],[]]
//输出：[null,null,2.00000,null,2.50000]
//
//
//
// 限制：
//
//
// 最多会对 addNum、findMedian 进行 50000 次调用。
//
//
// 注意：本题与主站 295 题相同：https://leetcode-cn.com/problems/find-median-from-data-
//stream/
// Related Topics 设计 双指针 数据流 排序 堆（优先队列） 👍 223 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	f := FConstructor()
	f.AddNum(6)
	fmt.Println(f.FindMedian())
	f.AddNum(10)
	fmt.Println(f.FindMedian())
	f.AddNum(2)
	fmt.Println(f.FindMedian())
	f.AddNum(6)
	fmt.Println(f.FindMedian())
	f.AddNum(5)
	fmt.Println(f.FindMedian())
}

// leetcode submit region begin(Prohibit modification and deletion)
type MedianFinder struct {
	nums sort.IntSlice
}

/** initialize your data structure here. */
func FConstructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	idx := this.nums.Search(num)
	if idx == 0 {
		this.nums = append([]int{num}, this.nums...)
	} else if idx == len(this.nums) {
		this.nums = append(this.nums, num)
	} else {
		p := this.nums[idx:]
		this.nums = append(this.nums[:idx:idx], num)
		this.nums = append(this.nums, p...)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	n := len(this.nums)
	if n%2 == 0 {
		return (float64(this.nums[n/2]) + float64(this.nums[n/2-1])) / 2
	}
	return float64(this.nums[n/2])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
//leetcode submit region end(Prohibit modification and deletion)
