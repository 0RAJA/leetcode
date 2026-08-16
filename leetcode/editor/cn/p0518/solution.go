// 给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
//
// 请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
//
// 假设每一种面额的硬币有无限个。
//
// 题目数据 保证 最终 结果符合 32 位 带符号整数。
//
//
//
//
//
//
// 示例 1：
//
//
// 输入：amount = 5, coins = [1, 2, 5]
// 输出：4
// 解释：有四种方式可以凑成总金额：
// 5=5
// 5=2+2+1
// 5=2+1+1+1
// 5=1+1+1+1+1
//
//
// 示例 2：
//
//
// 输入：amount = 3, coins = [2]
// 输出：0
// 解释：只用面额 2 的硬币不能凑成总金额 3 。
//
//
// 示例 3：
//
//
// 输入：amount = 10, coins = [10]
// 输出：1
//
//
//
//
// 提示：
//
//
// 1 <= coins.length <= 300
// 1 <= coins[i] <= 5000
// coins 中的所有值 互不相同
// 0 <= amount <= 5000
//
//
// Related Topics 数组 动态规划 背包问题 完全背包 👍 1504 👎 0

package p0518

// leetcode submit region begin(Prohibit modification and deletion)
// 完全背包组合计数
// 完全背包：从小到大遍历 bag
// 组合：外层 coin 内层 bag
// 计数：dp[i] 表示总数为 i 下的方案数; dp[i] += dp[i-coin] + 1;
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1 // 总数为 0 的方案只有 1 种
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if dp[i-coin] == 0 {
				continue
			}
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

// leetcode submit region end(Prohibit modification and deletion)
