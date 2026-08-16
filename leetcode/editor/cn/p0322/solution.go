// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//
// 你可以认为每种硬币的数量是无限的。
//
//
//
// 示例 1：
//
//
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1
//
// 示例 2：
//
//
// 输入：coins = [2], amount = 3
// 输出：-1
//
// 示例 3：
//
//
// 输入：coins = [1], amount = 0
// 输出：0
//
//
//
//
// 提示：
//
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 2³¹ - 1
// 0 <= amount <= 10⁴
//
//
// Related Topics 广度优先搜索 数组 动态规划 背包问题 完全背包 👍 3295 👎 0

package p0322

// leetcode submit region begin(Prohibit modification and deletion)
// 完全背包组合最小值问题
// dp[i] 表达凑成金额 i 的最小硬币数；dp[i] = min(dp[i],dp[i-num]+1)
// 完全背包：从小到大遍历 bag
// 组合：由于是最小值问题，所以内外层遍历方式不影响答案
// 最小值：dp[i] = min(dp[i],dp[i-num]+1)
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0 // 组成 0 的最小硬币数为0
	// 从小到大，因为硬币是可以被反复使用的，那么后者可以利用前者形成的结果
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin > i {
				continue
			}
			if i != coin && dp[i-coin] == 0 {
				continue
			}
			if dp[i] == 0 {
				dp[i] = dp[i-coin] + 1
			} else {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == 0 {
		return -1
	}
	return dp[amount]
}

// leetcode submit region end(Prohibit modification and deletion)
