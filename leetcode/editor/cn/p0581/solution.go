// 给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
//
// 请你找出符合题意的 最短 子数组，并输出它的长度。
//
//
//
//
//
// 示例 1：
//
//
//
// 输入：nums = [2,6,4,8,10,9,15]
// 输出：5
// 解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
//
//
// 示例 2：
//
//
// 输入：nums = [1,2,3,4]
// 输出：0
//
//
// 示例 3：
//
//
// 输入：nums = [1]
// 输出：0
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁴
// -10⁵ <= nums[i] <= 10⁵
//
//
//
//
// 进阶：你可以设计一个时间复杂度为 O(n) 的解决方案吗？
//
//
//
// Related Topics 栈 贪心 数组 双指针 排序 单调栈 👍 1262 👎 0

package p0581

// leetcode submit region begin(Prohibit modification and deletion)
// 就是找到最小无序的子数组，其实就是拿有序和无序对比最短差异的数组。
// 最简单的就是把它排序，然后再和原先的去做一个差异对比，但这个时间复杂度就是 O (N log N)。
// 另一种就是我假设从已排序的角度出发，从左到右就是从小到大，我维护一个最大值，如果发现后面的位置比前面维护的最大值还小，拿这个位置就是需要排序的，
// 找到最靠右的。同理反过来找到最靠左的位置，这个区间就是待排序的。
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	left, right := n, -1
	maxVal := nums[0]
	minVal := nums[n-1]
	// 从左往右，寻找右边界
	for i := range n {
		if nums[i] < maxVal {
			right = i
		} else {
			maxVal = nums[i]
		}
	}
	// 从右往左，寻找左边界
	for i := n - 1; i >= 0; i-- {
		if nums[i] > minVal {
			left = i
		} else {
			minVal = nums[i]
		}
	}
	// 原数组已经有序
	if right == -1 {
		return 0
	}
	return right - left + 1
}

// leetcode submit region end(Prohibit modification and deletion)
