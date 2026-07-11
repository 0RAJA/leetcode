// 整数数组 nums 按升序排列，数组中的值 互不相同 。
//
// 在传递给函数之前，nums 在预先未知的某个下标 K（0 <= K < nums.length）上进行了 旋转，使数组变为 [nums[K], nums[
// K+1], ..., nums[n-1], nums[0], nums[1], ..., nums[K-1]]（下标 从 0 开始 计数）。例如， [0,1,2
// ,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
//
// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
//
//
//
// 示例 1：
//
//
// 输入：nums = [4,5,6,7,0,1,2], target = 0
// 输出：4
//
//
// 示例 2：
//
//
// 输入：nums = [4,5,6,7,0,1,2], target = 3
// 输出：-1
//
// 示例 3：
//
//
// 输入：nums = [1], target = 0
// 输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5000
// -10^4 <= nums[i] <= 10^4
// nums 中的每个值都 独一无二
// 题目数据保证 nums 在预先未知的某个下标上进行了旋转
// -10^4 <= target <= 10^4
//
//
//
//
// 进阶：你可以设计一个时间复杂度为 O(log n) 的解决方案吗？
// Related Topics 数组 二分查找 👍 1681 👎 0

package main

import (
	"fmt"
)

func main() {
	nums := []int{9, 10, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(search(nums, 1))
}

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) (ret int) {
	// 咋能有递减的呢？
	// 找到排序数组中最小值
	searchMin := func() int {
		l, r := 0, len(nums)-1
		for l < r {
			mid := l + (r-l)/2
			// 左边界小于右边界，说明左边界就是最小值
			if nums[l] < nums[r] {
				return l
			} else if nums[mid] > nums[r] { // 左边界右移
				l = mid + 1
			} else { // 缩小右边界
				r = mid
			}
		}
		return l
	}
	searchX := func(l, r int) (ret int) {
		mid := 0
		for l < r {
			mid = l + (r-l)/2
			switch {
			case nums[mid] == target:
				return mid
			case nums[mid] < target:
				l = mid + 1
			default:
				r = mid
			}
		}
		if ret = l; nums[l] != target {
			ret = -1
		}
		return
	}
	index := searchMin()
	if ret = searchX(0, index); ret == -1 {
		ret = searchX(index, len(nums)-1)
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
