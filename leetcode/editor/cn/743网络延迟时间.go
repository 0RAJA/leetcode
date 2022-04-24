//有 n 个网络节点，标记为 1 到 n。
//
// 给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点，
//wi 是一个信号从源节点传递到目标节点的时间。
//
// 现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。
//
//
//
// 示例 1：
//
//
//
//
//输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
//输出：2
//
//
// 示例 2：
//
//
//输入：times = [[1,2,1]], n = 2, k = 1
//输出：1
//
//
// 示例 3：
//
//
//输入：times = [[1,2,1]], n = 2, k = 2
//输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= k <= n <= 100
// 1 <= times.length <= 6000
// times[i].length == 3
// 1 <= ui, vi <= n
// ui != vi
// 0 <= wi <= 100
// 所有 (ui, vi) 对都 互不相同（即，不含重复边）
//
// Related Topics 深度优先搜索 广度优先搜索 图 最短路 堆（优先队列） 👍 493 👎 0

package main

import "math"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func networkDelayTime(times [][]int, n int, k int) (ret int) {
	const INF = math.MaxInt32
	g := make([][]int, n+1)
	for i := range g {
		g[i] = make([]int, n+1)
		for j := range g {
			g[i][j] = INF
		}
	}
	for _, v := range times {
		x, y, w := v[0], v[1], v[2]
		g[x][y] = w
	}
	book := make(map[int]bool, n) // 记录是否已经通过此顶点松弛过
	dis := make([]int, n+1)       //距离
	for i := 1; i <= n; i++ {     //初始化距离
		dis[i] = g[k][i]
	}
	dis[k] = 0
	book[k] = true //把自己先标记
	for i := 1; i <= n; i++ {
		min := INF
		idx := 0
		for j := 1; j <= n; j++ { //找到最近的一个点
			if !book[j] && dis[j] < min {
				min, idx = dis[j], j
			}
		}
		for j := 1; j <= n; j++ { //通过最近的这个点松弛周围的点
			if g[idx][j] != INF && dis[j] > g[idx][j]+dis[idx] {
				dis[j] = g[idx][j] + dis[idx]
			}
		}
		book[idx] = true
	}
	for i := range dis {
		if dis[i] > ret {
			ret = dis[i]
		}
	}
	if ret == INF {
		return -1
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
