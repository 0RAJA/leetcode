// 给定一个含有 n 个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其
// 长度。如果不存在符合条件的子数组，返回 0 。
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
// 1 <= nums[i] <= 10⁴ 
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
// Related Topics 数组 二分查找 前缀和 滑动窗口 👍 2672 👎 0

package main

import (
	"math"
)

// leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLen(target int, nums []int) int {
	/*
		滑动窗口：维护一个滑动窗口，左指针和右指针，一个 sum 维护窗口内总和，一个 min_len 维护满足条件的窗口长度最小值
	*/
	// sum := nums[0]
	// totalLen := math.MaxInt32
	// for leftIdx, rightIdx := 0, 0; rightIdx < len(nums) && leftIdx <= rightIdx; {
	// 	if sum >= target {
	// 		totalLen = min(totalLen, rightIdx-leftIdx+1)
	// 		sum -= nums[leftIdx]
	// 		leftIdx++
	// 	} else {
	// 		rightIdx++
	// 		if rightIdx == len(nums) {
	// 			break
	// 		}
	// 		sum += nums[rightIdx]
	// 	}
	// }
	// if totalLen == math.MaxInt32 {
	// 	return 0
	// }
	// return totalLen

	// 版本2: 其实就是达到满足的条件，然后破坏它
	sum := 0
	totalLen := math.MaxInt32
	for leftIdx, rightIdx := 0, 0; rightIdx < len(nums); rightIdx++ {
		sum += nums[rightIdx]
		for sum >= target {
			totalLen = min(totalLen, rightIdx-leftIdx+1)
			sum -= nums[leftIdx]
			leftIdx++
		}
	}
	if totalLen == math.MaxInt32 {
		return 0
	}
	return totalLen
}

// leetcode submit region end(Prohibit modification and deletion)
