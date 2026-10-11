// 整数数组 nums 按升序排列，数组中的值 互不相同 。
//
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 向左旋转，使数组变为 [nums[k],
// nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1
// ,2,4,5,6,7] 下标 3 上向左旋转后可能变为 [4,5,6,7,0,1,2] 。
//
// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
//
// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
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
// -10⁴ <= nums[i] <= 10⁴
// nums 中的每个值都 独一无二
// 题目数据保证 nums 在预先未知的某个下标上进行了旋转
// -10⁴ <= target <= 10⁴
//
//
// Related Topics 数组 二分查找 👍 3465 👎 0

package p0033

// 个人理解，先通过首尾两个元素找到转置点，然后根据转置点来定位到有序的部分进行二分
// 找到转置点需要用 不断夹逼 来找到边界点
// 第二步查找则是使用常规二分
func search0(nums []int, target int) int {
	first, last := nums[0], nums[len(nums)-1]
	// 前提：数组已旋转，元素互不相同
	// left 始终在左侧较大值区间
	// right 始终在右侧较小值区间
	// 不断夹逼，直到两个指针相邻
	searchTransPoint := func() (int, int) {
		left, right := 0, len(nums)-1
		for right-left > 1 {
			mid := left + (right-left)/2
			if nums[mid] > nums[0] {
				left = mid
			} else {
				right = mid
			}
		}
		return left, right
	}
	// 在一个有序排列中通过二分找指定元素
	searchPoint := func(left, right int) int {
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] == target {
				return mid
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return -1
	}
	// 默认有序和单个节点直接找即可
	if first <= last {
		return searchPoint(0, len(nums)-1)
	}
	// 转置过了先找转置点
	leftMaxPoint, rightMinPoint := searchTransPoint()
	// 找 target 对应的区间
	left, right := 0, 0
	if target <= last {
		left, right = rightMinPoint, len(nums)-1
	} else {
		left, right = 0, leftMaxPoint
	}
	return searchPoint(left, right)
}

// leetcode submit region begin(Prohibit modification and deletion)

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		// 左半部分有序
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 右半部分有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// leetcode submit region end(Prohibit modification and deletion)
