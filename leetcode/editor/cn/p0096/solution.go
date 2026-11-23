// 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
//
//
//
// 示例 1：
//
//
// 输入：n = 3
// 输出：5
//
//
// 示例 2：
//
//
// 输入：n = 1
// 输出：1
//
//
//
//
// 提示：
//
//
// 1 <= n <= 19
//
//
// Related Topics 树 二叉搜索树 数学 动态规划 二叉树 👍 2755 👎 0

package p0096

// leetcode submit region begin(Prohibit modification and deletion)
// dp[n] = n 个不同节点能组成多少种 BST
// 其实就是以每个节点为root节点能够早的BST类型总和
// 例如，以当前以 root 为根，那当前的BST类型总和为：dp[root] = dp[root-1] * dp[n-root] 种
// 全部总和为：dp[1] + dp[2] +...+ dp[n]
// 这道题就是得先从小往大遍历节点数，再以每个节点数下计算在该节点数下，以每一个节点为根节点的 BST 数量，最后把它们总体再加起来
func numTrees(n int) (res int) {
	dp := make([]int, n+1) // dp[n] = n 个不同节点能组成多少种 BST
	dp[0] = 1
	dp[1] = 1
	for x := 2; x <= n; x++ {
		// 计算以当前节点为根时的 BST 数量
		for root := 1; root <= x; root++ {
			dp[x] += dp[root-1] * dp[x-root]
		}
	}
	return dp[n]
}

// leetcode submit region end(Prohibit modification and deletion)
