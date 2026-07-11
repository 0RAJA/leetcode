// 统计一个数字在排序数组中出现的次数。
//
//
//
// 示例 1:
//
//
// 输入: nums = [5,7,7,8,8,10], target = 8
// 输出: 2
//
// 示例 2:
//
//
// 输入: nums = [5,7,7,8,8,10], target = 6
// 输出: 0
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
//
//
// 注意：本题与主站 34 题相同（仅返回值不同）：https://leetcode-cn.com/problems/find-first-and-last-
// position-of-element-in-sorted-array/
//
// Related Topics 数组 二分查找 👍 372 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)

func search(nums []int, target int) (ret int) {
	searchL := func() int {
		l, r := 0, len(nums)-1
		for l < r {
			mid := l + (r-l)/2
			if nums[mid] >= target {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return l
	}
	for idx := searchL(); idx < len(nums) && nums[idx] == target; idx++ {
		ret++
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
