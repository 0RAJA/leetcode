//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
// 如果数组中不存在目标值 target，返回 [-1, -1]。
//
// 进阶：
//
//
// 你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
//
//
//
//
// 示例 1：
//
//
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
//
// 示例 2：
//
//
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]
//
// 示例 3：
//
//
//输入：nums = [], target = 0
//输出：[-1,-1]
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
// Related Topics 数组 二分查找 👍 1308 👎 0

package main

import "fmt"

func main() {
	nums := []int{1}
	fmt.Println(searchRange(nums, 1))
}

// leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	searchMin := func() int {
		l, r := 0, len(nums)-1
		for l < r {
			mid := l + (r-l)/2
			if nums[l] == target {
				return l
			}
			if nums[mid] < target {
				l = mid + 1
			} else {
				r = mid
			}
		}
		if nums[l] == target {
			return l
		}
		return -1
	}
	searchMax := func() int {
		l, r := 0, len(nums)-1
		for l <= r {
			mid := l + (r-l)/2
			if nums[r] == target {
				return r
			}
			if nums[mid] <= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		if l > 0 && nums[l-1] == target {
			return l - 1
		}
		return -1
	}
	x, y := searchMin(), searchMax()
	if x == -1 || y == -1 {
		return []int{-1, -1}
	}
	return []int{x, y}
}

//leetcode submit region end(Prohibit modification and deletion)
