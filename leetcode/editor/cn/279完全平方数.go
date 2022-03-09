//给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
//
// 给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
//
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
//
//
//
//
// 示例 1：
//
//
//输入：n = 12
//输出：3
//解释：12 = 4 + 4 + 4
//
// 示例 2：
//
//
//输入：n = 13
//输出：2
//解释：13 = 4 + 9
//
//
// 提示：
//
//
// 1 <= n <= 104
//
// Related Topics 广度优先搜索 数学 动态规划
// 👍 922 👎 0
package main

import "math"

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
/*
dp[i] 表示 表示最少需要多少个数的平方来表示整数i
	这些数必然落在[1,sqrt(i)],通过j来枚举这些数则选择此数的总个数为 dp[i-j*j]+1 而dp[i-j*j]属于更小的问题规模，于是构成
动态规划。
	转移方程: dp[i] = min(dp[i],dp[i-j*j]+1)
	i = [1,n]
	j = [1,sqrt(i)]
*/
func numSquares(n int) int {
	dp := make([]int, n+1)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := range dp {
		dp[i] = i
	}
	//相当于一个数组[1,n] 每个数量不限
	for j := 1; j <= int(math.Sqrt(float64(n))); j++ {
		v := j * j
		for i := v; i <= n; i++ {
			//可以重复用，所以需要正着遍历
			dp[i] = min(dp[i], dp[i-v]+1)
		}
	}
	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
