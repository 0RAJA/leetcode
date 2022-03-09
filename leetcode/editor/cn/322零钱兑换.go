//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
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
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//
// 示例 2：
//
//
//输入：coins = [2], amount = 3
//输出：-1
//
// 示例 3：
//
//
//输入：coins = [1], amount = 0
//输出：0
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
// Related Topics 广度优先搜索 数组 动态规划 👍 1766 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(coinChange([]int{474, 83, 404, 3}, 264))
}

//leetcode submit region begin(Prohibit modification and deletion)
/*
dp[i] = min(dp[i],dp[i-j*k]+k)
*/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	sort.Ints(coins)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 1; i <= amount; i++ {
		dp[i] = 1e4 + 1
	}
	//空间
	for _, v := range coins {
		//物品无限
		for i := v; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-v]+1)
		}
	}

	if dp[amount] == 1e4+1 {
		return -1
	}
	return dp[amount]
}

//leetcode submit region end(Prohibit modification and deletion)
