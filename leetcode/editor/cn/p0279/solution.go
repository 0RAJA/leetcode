// 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
//
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
//
//
//
//
// 示例 1：
//
//
// 输入：n = 12
// 输出：3
// 解释：12 = 4 + 4 + 4
//
// 示例 2：
//
//
// 输入：n = 13
// 输出：2
// 解释：13 = 4 + 9
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁴
//
//
// Related Topics 广度优先搜索 数学 动态规划 背包问题 完全背包 👍 2394 👎 0

package p0279

// leetcode submit region begin(Prohibit modification and deletion)
// 动态规划：完全背包求最少数量
// dp[j] = min(dp[j],dp[j-w]+1)
// 完全背包，不限制物品数：容量正序遍历
func numSquares(n int) (res int) {
	// 构造物品
	nums := make([]int, 0, n)
	for i := 1; i*i <= n; i++ {
		nums = append(nums, i*i)
	}
	dp := make([]int, n+1)
	// 可复用物品，容量从小到大
	for capacity := 1; capacity <= n; capacity++ {
		for idx := 0; idx < len(nums) && nums[idx] <= capacity; idx++ {
			// 避免被默认值误导
			if dp[capacity] == 0 {
				dp[capacity] = dp[capacity-nums[idx]] + 1
			} else {
				dp[capacity] = min(dp[capacity], dp[capacity-nums[idx]]+1)
			}
		}
	}
	return dp[n]
}

// leetcode submit region end(Prohibit modification and deletion)
