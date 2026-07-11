// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
// 
// 
//
// 
//
// 示例 1： 
//
// 
// 输入：nums = [-4,-1,0,3,10]
// 输出：[0,1,9,16,100]
// 解释：平方后，数组变为 [16,1,0,9,100]
// 排序后，数组变为 [0,1,9,16,100]
//
// 示例 2： 
//
// 
// 输入：nums = [-7,-3,2,3,11]
// 输出：[4,9,9,49,121]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10⁴ 
// -10⁴ <= nums[i] <= 10⁴ 
// nums 已按 非递减顺序 排序 
// 
//
// 
//
// 进阶： 
//
// 
// 请你设计时间复杂度为 O(n) 的算法解决本问题 
// 
//
// Related Topics 数组 双指针 排序 👍 1177 👎 0

package main

// leetcode submit region begin(Prohibit modification and deletion)
func sortedSquares(nums []int) []int {
	// 双指针，因为这个平方后实际上是一个V型，所以可以从两边向中间遍历，把较大的放到新数组右边
	// 维护一个新队列，从右往左塞前面比较后的较大值
	newNums := make([]int, len(nums))
	k := len(nums) - 1
	for i, j := 0, len(nums)-1; i <= j; {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			newNums[k] = nums[i] * nums[i]
			i++
		} else {
			newNums[k] = nums[j] * nums[j]
			j--
		}
		k--
	}
	return newNums
}

// leetcode submit region end(Prohibit modification and deletion)
