// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上
// 被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
//
//
//
// 示例 1：
//
//
// 输入：[1,2,3,1]
// 输出：4
// 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//      偷窃到的最高金额 = 1 + 3 = 4 。
//
// 示例 2：
//
//
// 输入：[2,7,9,3,1]
// 输出：12
// 解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//      偷窃到的最高金额 = 2 + 9 + 1 = 12 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 400
//
//
// Related Topics 数组 动态规划 👍 3557 👎 0

package p0198

// leetcode submit region begin(Prohibit modification and deletion)
// 动态规划：dp[i] = max(dp[i-2] + nums[i], dp[i-1])
// 选择 max(当前房子+不选择相邻房子,不选当前房子+选择相邻房子)
func rob(nums []int) (res int) {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		num1, num2 := nums[i], 0
		if i >= 2 {
			num1 += dp[i-2]
		}
		if i >= 1 {
			num2 += dp[i-1]
		}
		dp[i] = max(num1, num2)
	}
	return dp[len(dp)-1]
}

// leetcode submit region end(Prohibit modification and deletion)
