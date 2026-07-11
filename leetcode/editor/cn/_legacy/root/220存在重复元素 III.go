//给你一个整数数组 nums 和两个整数 K 和 t 。请你判断是否存在 两个不同下标 i 和 j，使得 abs(nums[i] - nums[j]) <=
//t ，同时又满足 abs(i - j) <= K 。
//
// 如果存在则返回 true，不存在返回 false。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3,1], K = 3, t = 0
//输出：true
//
// 示例 2：
//
//
//输入：nums = [1,0,1,1], K = 1, t = 2
//输出：true
//
// 示例 3：
//
//
//输入：nums = [1,5,9,1,5,9], K = 2, t = 3
//输出：false
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 2 * 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
// 0 <= K <= 10⁴
// 0 <= t <= 2³¹ - 1
//
// Related Topics 数组 桶排序 有序集合 排序 滑动窗口 👍 517 👎 0

package main

import "fmt"

func main() {
	nums := []int{1, 2, 5, 6, 7, 2, 4}
	k := 4
	t := 0
	fmt.Println(containsNearbyAlmostDuplicate(nums, k, t))
}

// leetcode submit region begin(Prohibit modification and deletion)
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	for i, j := 0, 0; j < len(nums); j++ {
		for i < j && j-i > k {
			i++
		}
		for l := i; l < j; l++ {
			if abs(nums[l]-nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
