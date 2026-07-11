// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子
// 序列。
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
// Related Topics 数组 二分查找 动态规划 👍 2856 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)

func lengthOfLIS(nums []int) (cnt int) {
	dp := make([]int, len(nums)) // dp[i] 为考虑前 i 个元素，以第 i 个数字结尾的最长上升子序列的长度，注意 nums[i] 必须被选取。
	// 每次选前面最长的
	for i := 0; i < len(nums); i++ {
		max := 0
		for j := 0; j < i; j++ {
			if max < dp[j] && nums[i] > nums[j] {
				max = dp[j]
			}
		}
		dp[i] = max + 1
		if cnt < dp[i] {
			cnt = dp[i]
		}
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
