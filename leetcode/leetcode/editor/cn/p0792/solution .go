// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target ，写一个函数搜索 nums 中的 target，如果 target 存在返
// 回下标，否则返回 -1。
//
// 你必须编写一个具有 O(log n) 时间复杂度的算法。
//
// 示例 1:
//
//
// 输入: nums = [-1,0,3,5,9,12], target = 9
// 输出: 4
// 解释: 9 出现在 nums 中并且下标为 4
//
//
// 示例 2:
//
//
// 输入: nums = [-1,0,3,5,9,12], target = 2
// 输出: -1
// 解释: 2 不存在 nums 中因此返回 -1
//
//
//
//
// 提示：
//
//
// 你可以假设 nums 中的所有元素是不重复的。
// n 将在 [1, 10000]之间。
// nums 的每个元素都将在 [-9999, 9999]之间。
//
//
// Related Topics 数组 二分查找 👍 1893 👎 0

package p0792

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	// 方案1 target 在一个 [a,b] 中
	// l, r := 0, len(nums)-1
	// for l <= r {
	// 	m := l + (r-l)/2 // 防止溢出
	// 	v := nums[m]
	// 	if v == target {
	// 		return m
	// 	} else if v > target {
	// 		r = m - 1
	// 	} else {
	// 		l = m + 1
	// 	}
	// }
	// return -1
	// 方案2 target 在 [a,b) 中
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	return -1
}

// leetcode submit region end(Prohibit modification and deletion)
