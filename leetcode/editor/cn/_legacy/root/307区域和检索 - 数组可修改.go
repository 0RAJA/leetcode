// 给你一个数组 nums ，请你完成两类查询。
//
//
// 其中一类查询要求 更新 数组 nums 下标对应的值
// 另一类查询要求返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 ，其中 left <= right
//
//
// 实现 NumArray 类：
//
//
// NumArray(int[] nums) 用整数数组 nums 初始化对象
// void update(int index, int val) 将 nums[index] 的值 更新 为 val
// int sumRange(int left, int right) 返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元
// 素的 和 （即，nums[left] + nums[left + 1], ..., nums[right]）
//
//
//
//
// 示例 1：
//
//
// 输入：
// ["NumArray", "sumRange", "update", "sumRange"]
// [[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
// 输出：
// [null, 9, null, 8]
//
// 解释：
// NumArray numArray = new NumArray([1, 3, 5]);
// numArray.sumRange(0, 2); // 返回 1 + 3 + 5 = 9
// numArray.update(1, 2);   // nums = [1,2,5]
// numArray.sumRange(0, 2); // 返回 1 + 2 + 5 = 8
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 10⁴
// -100 <= nums[i] <= 100
// 0 <= index < nums.length
// -100 <= val <= 100
// 0 <= left <= right < nums.length
// 调用 pdate 和 sumRange 方法次数不大于 3 * 10⁴
//
// Related Topics 设计 树状数组 线段树 数组 👍 379 👎 0

package main

import "fmt"

func main() {
	numArray := Constructor([]int{7, 2, 7, 2, 0})
	fmt.Println(numArray.SumRange(0, 4))
	numArray.Update(4, 6) // 7, 2, 7, 2, 6
	numArray.Update(0, 2) // 2, 2, 7, 2, 6
	numArray.Update(0, 9) // 9, 2, 7, 2, 6
	fmt.Println(numArray.SumRange(4, 4))
	numArray.Update(3, 8) // 9, 2, 7, 8, 6
	fmt.Println(numArray.SumRange(0, 4))
	numArray.Update(4, 1)
	fmt.Println(numArray.SumRange(0, 3))
	fmt.Println(numArray.SumRange(0, 4))
}

// leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	*TreeNum
	nums []int
}

func Constructor(nums []int) NumArray {
	treeNum := NewTreeNum(len(nums))
	for i := range nums {
		treeNum.add(i, nums[i])
	}
	return NumArray{
		TreeNum: treeNum,
		nums:    nums,
	}
}

func (this *NumArray) Update(index int, val int) {
	a := this.nums[index]
	this.add(index, val-a)
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.ask(right) - this.ask(left-1)
}

type TreeNum struct {
	nums []int
}

func NewTreeNum(n int) *TreeNum {
	return &TreeNum{make([]int, n+1)} // 下标从1开始
}

// 计算x二进制的最低位1及之后的数。101000 -> 001000
func (t *TreeNum) lowbit(x int) int {
	return x & (-x)
}

// x点值+d
func (t *TreeNum) add(x, d int) {
	x++
	for ; x < len(t.nums); x += t.lowbit(x) { // 因为动态维护一个前缀和数组，所以依次都要+d
		t.nums[x] += d
	}
}

// 查询x下标处的前缀和
func (t *TreeNum) ask(x int) (ret int) {
	x++
	for ; x >= 1 && x < len(t.nums); x -= t.lowbit(x) {
		ret += t.nums[x]
	}
	return
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
// leetcode submit region end(Prohibit modification and deletion)
