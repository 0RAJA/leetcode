// 给定一个含有 n 个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长
// 度。如果不存在符合条件的子数组，返回 0 。
//
// 
//
// 示例 1： 
//
// 
// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
// 
//
// 示例 2： 
//
// 
// 输入：target = 4, nums = [1,4,4]
// 输出：1
// 
//
// 示例 3： 
//
// 
// 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
// 输出：0
// 
//
// 
//
// 提示： 
//
// 
// 1 <= target <= 10⁹ 
// 1 <= nums.length <= 10⁵ 
// 1 <= nums[i] <= 10⁵ 
// 
//
// 
//
// 进阶： 
//
// 
// 如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。 
// 
//
// 
//
// 
// 注意：本题与主站 209 题相同：https://leetcode.cn/problems/minimum-size-subarray-sum/ 
//
// Related Topics 数组 二分查找 前缀和 滑动窗口 👍 172 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
// 正数说明总和一定是递增的
// 维护 l 和 r 组成的子数组，符合条件了，然后计算总长度，然后弹出元素直到不符合预期
func minSubArrayLen(target int, nums []int) (res int) {
	res = len(nums) + 1 // 一个大值
	sum := 0
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		// 维护 l 和 r 组成的子数组，符合条件了，然后计算总长度，然后弹出元素直到不符合预期
		for sum >= target {
			res = min(res, r-l+1)
			sum -= nums[l]
			l++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
