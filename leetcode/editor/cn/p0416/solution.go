// 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,5,11,5]
// 输出：true
// 解释：数组可以分割成 [1, 5, 5] 和 [11] 。
//
// 示例 2：
//
//
// 输入：nums = [1,2,3,5]
// 输出：false
// 解释：数组不能分割成两个元素和相等的子集。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 100
//
//
// Related Topics 数组 动态规划 背包问题 0-1 背包 👍 2592 👎 0

package p0416

// leetcode submit region begin(Prohibit modification and deletion)
// 动态规划 0-1 背包组合盘存在问题：从 N 个元素中找出总价值能否达到 target
// dp[i] 表示能否凑出 i
// dp[i] = dp[i-1] || dp[target-i] target = sum/2
// 0-1 背包：从大到小遍历 bag
// 组合：外层 遍历 num，内层遍历 bag
// 存在：dp[i] = dp[i] || dp[i-num]
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[target]
}

// leetcode submit region end(Prohibit modification and deletion)
