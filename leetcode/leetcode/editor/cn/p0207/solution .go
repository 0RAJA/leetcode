// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
//
// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表
// 示如果要学习课程 ai 则 必须 先学习课程 bi 。
//
//
// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
//
//
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
// 输入：numCourses = 2, prerequisites = [[1,0]]
// 输出：true
// 解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
//
// 示例 2：
//
//
// 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
// 输出：false
// 解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
//
//
//
// 提示：
//
//
// 1 <= numCourses <= 2000
// 0 <= prerequisites.length <= 5000
// prerequisites[i].length == 2
// 0 <= ai, bi < numCourses
// prerequisites[i] 中的所有课程对 互不相同
//
//
// Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 👍 2394 👎 0

package p0207

// leetcode submit region begin(Prohibit modification and deletion)
// 1. 广搜：统计每个节点的入度，入度为0则表示可以修的课程；
// 把入度为0的节点放到队列，然后遍历队列，每出一个节点把引用的节点的入度-1，为0再入队列
// 实现：1. inDeg 记录节点入度；2. queue 记录每个入度为0的节点；3. edges：记录每个节点引用了哪些节点，用来更新他们的入度；4. result 记录选修结果顺序
func canFinish(numCourses int, prerequisites [][]int) bool {
	inDeg := make(map[int]int, numCourses) // 节点入度
	queue := make([]int, 0, numCourses)    // bfs 队列
	edges := make([][]int, numCourses)     // 每个节点的出边
	result := make([]int, 0, numCourses)   // 可选修课程的顺序结果
	// 1. 构建节点间关系
	for _, edge := range prerequisites {
		a, b := edge[0], edge[1]
		inDeg[a]++
		edges[b] = append(edges[b], a)
		if _, ok := inDeg[b]; !ok {
			inDeg[b] = 0
		}
	}
	// 2. 找到入度为0 的元素放到队列
	for k, v := range inDeg {
		if v == 0 {
			queue = append(queue, k)
		}
	}
	// 3. bfs 遍历队列，把入度为0的元素放到结果，更新出边的元素
	for len(queue) > 0 {
		visitedData := queue[0]
		queue = queue[1:]
		result = append(result, visitedData)
		for _, v := range edges[visitedData] {
			inDeg[v]--
			if inDeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	// 4. 判断结果数量是否和选修课程数一致
	return len(result) == len(inDeg)
}

// leetcode submit region end(Prohibit modification and deletion)
