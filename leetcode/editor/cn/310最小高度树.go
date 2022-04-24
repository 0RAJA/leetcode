//树是一个无向图，其中任何两个顶点只通过一条路径连接。 换句话说，一个任何没有简单环路的连通图都是一棵树。
//
// 给你一棵包含 n 个节点的树，标记为 0 到 n - 1 。给定数字 n 和一个有 n - 1 条无向边的 edges 列表（每一个边都是一对标签），其中
// edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间存在一条无向边。
//
// 可选择树中任何一个节点作为根。当选择节点 x 作为根节点时，设结果树的高度为 h 。在所有可能的树中，具有最小高度的树（即，min(h)）被称为 最小高度
//树 。
//
// 请你找到所有的 最小高度树 并按 任意顺序 返回它们的根节点标签列表。
//树的 高度 是指根节点和叶子节点之间最长向下路径上边的数量。
//
//
//
// 示例 1：
//
//
//输入：n = 4, edges = [[1,0],[1,2],[1,3]]
//输出：[1]
//解释：如图所示，当根是标签为 1 的节点时，树的高度是 1 ，这是唯一的最小高度树。
//
// 示例 2：
//
//
//输入：n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
//输出：[3,4]
//
//
//
//
//
//
//
// 提示：
//
//
// 1 <= n <= 2 * 10⁴
// edges.length == n - 1
// 0 <= ai, bi < n
// ai != bi
// 所有 (ai, bi) 互不相同
// 给定的输入 保证 是一棵树，并且 不会有重复的边
//
// Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 👍 539 👎 0

package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(findMinHeightTrees(2, [][]int{{0, 1}}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findMinHeightTrees(n int, edges [][]int) (ret []int) {
	if n == 1 {
		return []int{0}
	}
	type (
		VexNode struct {
			cnt    int
			childs []int
		}
		AdjList []VexNode
	)
	adj := make(AdjList, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a].cnt++
		adj[a].childs = append(adj[a].childs, b)
		adj[b].cnt++
		adj[b].childs = append(adj[b].childs, a)
	}
	cnt := 0
	queue := list.New()
	queue2 := list.New()
	for i := range adj {
		if adj[i].cnt == 1 {
			queue.PushBack(i)
		}
	}
	cnt += queue.Len()
	if cnt == n {
		for queue.Len() > 0 {
			idx := queue.Front().Value.(int)
			queue.Remove(queue.Front())
			ret = append(ret, idx)
		}
		return
	}
	for queue.Len() > 0 {
		idx := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		for _, v := range adj[idx].childs {
			adj[v].cnt--
			if adj[v].cnt == 1 {
				queue2.PushBack(v)
				cnt++
			}
		}
		if queue.Len() == 0 {
			if cnt == n {
				for queue2.Len() > 0 {
					idx := queue2.Front().Value.(int)
					queue2.Remove(queue2.Front())
					ret = append(ret, idx)
				}
				return
			}
			queue.PushBackList(queue2)
			queue2 = list.New()
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
