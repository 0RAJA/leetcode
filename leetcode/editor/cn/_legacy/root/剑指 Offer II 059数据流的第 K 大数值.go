//设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。
//
// 请实现 KthLargest 类：
//
//
// KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
// int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。
//
//
//
//
// 示例：
//
//
//输入：
//["KthLargest", "add", "add", "add", "add", "add"]
//[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
//输出：
//[null, 4, 5, 5, 8, 8]
//
//解释：
//KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
//kthLargest.add(3);   // return 4
//kthLargest.add(5);   // return 5
//kthLargest.add(10);  // return 5
//kthLargest.add(9);   // return 8
//kthLargest.add(4);   // return 8
//
//
//
//
// 提示：
//
//
// 1 <= k <= 10⁴
// 0 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// -10⁴ <= val <= 10⁴
// 最多调用 add 方法 10⁴ 次
// 题目数据保证，在查找第 k 大元素时，数组中至少有 k 个元素
//
//
//
//
// 注意：本题与主站 703 题相同： https://leetcode-cn.com/problems/kth-largest-element-in-a-
//stream/
// Related Topics 树 设计 二叉搜索树 二叉树 数据流 堆（优先队列） 👍 26 👎 0

package main

import "container/heap"

func main() {
	Constructor(3, []int{4, 5, 8, 2})
}

// leetcode submit region begin(Prohibit modification and deletion)
type KthLargest struct {
	IntSlice []int
	k        int
}

func (this *KthLargest) Len() int           { return len(this.IntSlice) }
func (this *KthLargest) Less(i, j int) bool { return this.IntSlice[i] < this.IntSlice[j] }
func (this *KthLargest) Swap(i, j int) {
	this.IntSlice[i], this.IntSlice[j] = this.IntSlice[j], this.IntSlice[i]
}

func (this *KthLargest) Push(x interface{}) { this.IntSlice = append(this.IntSlice, x.(int)) }

func (this *KthLargest) Pop() interface{} {
	x := this.IntSlice[len(this.IntSlice)-1]
	this.IntSlice = this.IntSlice[:len(this.IntSlice)-1]
	return x
}

func Constructor(k int, nums []int) KthLargest {
	ret := &KthLargest{IntSlice: make([]int, 0, k+1), k: k}
	for i := range nums {
		heap.Push(ret, nums[i])
		if ret.Len() > k {
			heap.Pop(ret)
		}
	}
	return *ret
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.k {
		heap.Pop(this)
	}
	return this.IntSlice[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
//leetcode submit region end(Prohibit modification and deletion)
