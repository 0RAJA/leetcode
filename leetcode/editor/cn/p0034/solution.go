// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
//
// 如果数组中不存在目标值 target，返回 [-1, -1]。
//
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
//
//
//
// 示例 1：
//
//
// 输入：nums = [5,7,7,8,8,10], target = 8
// 输出：[3,4]
//
// 示例 2：
//
//
// 输入：nums = [5,7,7,8,8,10], target = 6
// 输出：[-1,-1]
//
// 示例 3：
//
//
// 输入：nums = [], target = 0
// 输出：[-1,-1]
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// nums 是一个非递减数组
// -10⁹ <= target <= 10⁹
//
//
// Related Topics 数组 二分查找 👍 3323 👎 0

package p0034

// leetcode submit region begin(Prohibit modification and deletion)
// 二分查找两次，分别找开头和结尾的 target
/*
```
left, right := 0, len(nums)-1
for left <= right {
	mid := left + (right-left)/2
	if nums[mid] >= target {
		// 往左找
		if nums[mid] == target {
			leftIdx = min(leftIdx, mid)
		}
		right = mid - 1
	} else {
		// 往右找
		left = mid + 1
	}
}
```
*/
func searchRange(nums []int, target int) []int {
	leftIdx, rightIdx := -1, -1
	searchLeft := func() {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] >= target {
				if nums[mid] == target {
					if leftIdx == -1 {
						leftIdx = mid
					} else {
						leftIdx = min(leftIdx, mid)
					}
				}
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	searchRight := func() {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] > target {
				right = mid - 1
			} else {
				if nums[mid] == target {
					if rightIdx == -1 {
						rightIdx = mid
					} else {
						rightIdx = max(rightIdx, mid)
					}
				}
				left = mid + 1
			}
		}
	}
	searchLeft()
	searchRight()
	return []int{leftIdx, rightIdx}
}

// leetcode submit region end(Prohibit modification and deletion)
