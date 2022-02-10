//一个国家有 n 个城市，城市编号为 0 到 n - 1 ，题目保证 所有城市 都由双向道路 连接在一起 。道路由二维整数数组 edges 表示，其中
//edges[i] = [xi, yi, timei] 表示城市 xi 和 yi 之间有一条双向道路，耗费时间为 timei 分钟。两个城市之间可能会有多条耗费时间不同
//的道路，但是不会有道路两头连接着同一座城市。
//
// 每次经过一个城市时，你需要付通行费。通行费用一个长度为 n 且下标从 0 开始的整数数组 passingFees 表示，其中 passingFees[j]
// 是你经过城市 j 需要支付的费用。
//
// 一开始，你在城市 0 ，你想要在 maxTime 分钟以内 （包含 maxTime 分钟）到达城市 n - 1 。旅行的 费用 为你经过的所有城市 通行费
//之和 （包括 起点和终点城市的通行费）。
//
// 给你 maxTime，edges 和 passingFees ，请你返回完成旅行的 最小费用 ，如果无法在 maxTime 分钟以内完成旅行，请你返回 -
//1 。
//
//
//
// 示例 1：
//
//
//
//
//输入：maxTime = 30, edges = [[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]
//], passingFees = [5,1,2,20,20,3]
//输出：11
//解释：最优路径为 0 -> 1 -> 2 -> 5 ，总共需要耗费 30 分钟，需要支付 11 的通行费。
//
//
// 示例 2：
//
//
//
//
//输入：maxTime = 29, edges = [[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]
//], passingFees = [5,1,2,20,20,3]
//输出：48
//解释：最优路径为 0 -> 3 -> 4 -> 5 ，总共需要耗费 26 分钟，需要支付 48 的通行费。
//你不能选择路径 0 -> 1 -> 2 -> 5 ，因为这条路径耗费的时间太长。
//
//
// 示例 3：
//
//
//输入：maxTime = 25, edges = [[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]
//], passingFees = [5,1,2,20,20,3]
//输出：-1
//解释：无法在 25 分钟以内从城市 0 到达城市 5 。
//
//
//
//
// 提示：
//
//
// 1 <= maxTime <= 1000
// n == passingFees.length
// 2 <= n <= 1000
// n - 1 <= edges.length <= 1000
// 0 <= xi, yi <= n - 1
// 1 <= timei <= 1000
// 1 <= passingFees[j] <= 1000
// 图中两个节点之间可能有多条路径。
// 图中不含有自环。
//
// Related Topics 图 动态规划 👍 27 👎 0

package main

import "fmt"

func main() {
	maxTime := 30
	edges := [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}}
	passingFees := []int{5, 1, 2, 20, 20, 3}
	fmt.Println(minCost(maxTime, edges, passingFees))

}

//leetcode submit region begin(Prohibit modification and deletion)
func minCost(maxTime int, edges [][]int, passingFees []int) (ret int) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n := len(passingFees)          //城市数
	dp := make([][]int, maxTime+1) //dp[时间][城市]
	const INF = 1000*1000 + 1      //无穷大
	ret = INF
	for i := range dp { //先将所有初始化为INF表示不可达
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[0][0] = passingFees[0]       //(包括起点和终点城市的通行费)最开始初始化为0了,一直不对...
	for t := 1; t <= maxTime; t++ { //从时间1开始进行松弛
		for _, edge := range edges { //遍历所有边
			x, y, cost := edge[0], edge[1], edge[2]
			if cost <= t { //如果这个边可以走,看看能不能用于松弛
				dp[t][x] = min(dp[t][x], dp[t-cost][y]+passingFees[x])
				dp[t][y] = min(dp[t][y], dp[t-cost][x]+passingFees[y])
				ret = min(ret, dp[t][n-1]) //同时计算途中到达目的地的最短时间(在 maxTime 分钟以内 （包含 maxTime 分钟）到达城市 n - 1)
			}
		}
	}
	if ret == INF {
		return -1
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
