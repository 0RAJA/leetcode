// 给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和
// values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。
//
// 另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj =
// ? 的结果作为答案。
//
// 返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替
// 代这个答案。
//
// 注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。
//
// 注意：未在等式列表中出现的变量是未定义的，因此无法确定它们的答案。
//
//
//
// 示例 1：
//
//
// 输入：equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"]
// ,["b","a"],["a","e"],["a","a"],["x","x"]]
// 输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
// 解释：
// 条件：a / b = 2.0, b / c = 3.0
// 问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
// 结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]
// 注意：x 是未定义的 => -1.0
//
// 示例 2：
//
//
// 输入：equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0],
// queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
// 输出：[3.75000,0.40000,5.00000,0.20000]
//
//
// 示例 3：
//
//
// 输入：equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],[
// "a","c"],["x","y"]]
// 输出：[0.50000,2.00000,-1.00000,-1.00000]
//
//
//
//
// 提示：
//
//
// 1 <= equations.length <= 20
// equations[i].length == 2
// 1 <= Ai.length, Bi.length <= 5
// values.length == equations.length
// 0.0 < values[i] <= 20.0
// 1 <= queries.length <= 20
// queries[i].length == 2
// 1 <= Cj.length, Dj.length <= 5
// Ai, Bi, Cj, Dj 由小写英文字母与数字组成
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 图 数组 字符串 最短路 Floyd 算法 Bellman–Ford 算法 👍 128
// 5 👎 0

package p0399

// leetcode submit region begin(Prohibit modification and deletion)
type Edge struct {
	to  string
	val float64
}

// 将 a / b = x 转换成带权边 a -> b = x 和 b -> a = 1/x，每个查询就是在图中寻找从起点到终点的路径，并将路径上的边权相乘。
// 先根据 a/b = x 构造 graph(带权邻接表)：a -> b = x; b -> a = 1/x
// 然后对于每个 query dfs 遍历 graph，维护 累计乘法结果、当前元素、目标元素、visited 访问记录
// 1. 根据 equations 构造带权邻接表。
// 2. 对每个 query 从起点执行 DFS。
// 3. DFS 过程中维护当前累计乘积。
// 4. 找到目标节点时，累计乘积就是答案。
// 5. 使用 visited 防止图中环导致重复搜索。
func calcEquation(equations [][]string, values []float64, queries [][]string) (res []float64) {
	// 1. 构造 graph（带权邻接表）
	graph := make(map[string][]*Edge)
	for idx, equation := range equations {
		from, to := equation[0], equation[1]
		val := values[idx]
		graph[from] = append(graph[from], &Edge{
			to:  to,
			val: val,
		})
		graph[to] = append(graph[to], &Edge{
			to:  from,
			val: 1 / val,
		})
	}
	// 2. 构造 dfs 深搜函数
	visited := make(map[string]bool)
	var dfs func(currentStr, targetStr string, sum float64) (float64, bool)
	dfs = func(currentStr, targetStr string, sum float64) (float64, bool) {
		if currentStr == targetStr {
			return sum, true
		}
		for _, toNode := range graph[currentStr] {
			if visited[toNode.to] {
				continue
			}
			// 标记访问过，避免 a->b,b->a 陷入死循环
			visited[toNode.to] = true
			result, ok := dfs(toNode.to, targetStr, sum*toNode.val)
			visited[toNode.to] = false
			if ok {
				return result, ok
			}
		}
		return -1, false
	}
	// 3. 遍历 query 数组构造答案
	for _, query := range queries {
		from, to := query[0], query[1]
		// 对于从未出现在graph的元素认为是未定义的，因此无法确定它们的答案
		_, ok := graph[to]
		result := float64(-1)
		if ok {
			result, _ = dfs(from, to, 1)
		}
		res = append(res, result)
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
