// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子
// 序列。
//
//
// 示例 1：
//
//
// 输入：nums = [10,9,2,5,3,7,101,18]
// 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
//
//
// 示例 2：
//
//
// 输入：nums = [0,1,0,3,2,3]
// 输出：4
//
//
// 示例 3：
//
//
// 输入：nums = [7,7,7,7,7,7,7]
// 输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2500
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
// 进阶：
//
//
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
//
//
// Related Topics 数组 二分查找 动态规划 最长上升子序列 👍 4309 👎 0

package p0300

// leetcode submit region begin(Prohibit modification and deletion)
// dp[i] 表示以当前字符结尾的最长递增子序列数，对于 j < i && dp[j] > dp[i] 那么 dp[i] = dp[j] + 1
// 注意最大长度可能在其中的任意一个位置哦
func lengthOfLIS(nums []int) (res int) {
	dp := make([]int, len(nums))

	for i, num := range nums {
		for j := 0; j < i; j++ {
			if nums[j] >= num {
				continue
			}
			dp[i] = max(dp[i], dp[j]+1)
		}
		res = max(res, dp[i])
	}
	return res + 1
}

// leetcode submit region end(Prohibit modification and deletion)
