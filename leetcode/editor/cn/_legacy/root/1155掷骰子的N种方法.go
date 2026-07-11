//这里有 n 个一样的骰子，每个骰子上都有 k 个面，分别标号为 1 到 k 。
//
// 给定三个整数 n , k 和 target ，返回可能的方式(从总共 kⁿ 种方式中)滚动骰子的数量，使正面朝上的数字之和等于 target 。
//
// 答案可能很大，你需要对 10⁹ + 7 取模 。
//
//
//
// 示例 1：
//
//
//输入：n = 1, k = 6, target = 3
//输出：1
//解释：你扔一个有6张脸的骰子。
//得到3的和只有一种方法。
//
//
// 示例 2：
//
//
//输入：n = 2, k = 6, target = 7
//输出：6
//解释：你扔两个骰子，每个骰子有6个面。
//得到7的和有6种方法1+6 2+5 3+4 4+3 5+2 6+1。
//
//
// 示例 3：
//
//
//输入：n = 30, k = 30, target = 500
//输出：222616187
//解释：返回的结果必须是对 10⁹ + 7 取模。
//
//
//
// 提示：
//
//
// 1 <= n, k <= 30
// 1 <= target <= 1000
//
// Related Topics 动态规划 👍 123 👎 0

package main

import "fmt"

func main() {
	fmt.Println(numRollsToTarget(30, 30, 500))
}

//leetcode submit region begin(Prohibit modification and deletion)
/*
	dp[n][target] 选n个骰子和为target的方案数
*/
func numRollsToTarget(n int, d int, target int) int {
	const MOD = 1e9 + 7
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	if d > target {
		d = target
	}
	for i := 1; i <= d; i++ {
		dp[1][i] = 1
	}
	//遍历物品
	for i := 1; i <= n; i++ {
		//遍历容量
		for j := 2; j <= target; j++ {
			//因为求总方案数，所以要把所有取值的可能都加起来。
			for k := 1; k <= j && k <= d; k++ {
				dp[i][j] = (dp[i][j] + dp[i-1][j-k]) % MOD
			}
		}
	}
	return dp[n][target]
}

//leetcode submit region end(Prohibit modification and deletion)
