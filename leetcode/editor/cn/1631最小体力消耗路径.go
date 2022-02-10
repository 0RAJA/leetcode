//你准备参加一场远足活动。给你一个二维 rows x columns 的地图 heights ，其中 heights[row][col] 表示格子 (row,
// col) 的高度。一开始你在最左上角的格子 (0, 0) ，且你希望去最右下角的格子 (rows-1, columns-1) （注意下标从 0 开始编号）。你
//每次可以往 上，下，左，右 四个方向之一移动，你想要找到耗费 体力 最小的一条路径。
//
// 一条路径耗费的 体力值 是路径上相邻格子之间 高度差绝对值 的 最大值 决定的。
//
// 请你返回从左上角走到右下角的最小 体力消耗值 。
//
//
//
// 示例 1：
//
//
//
//
//输入：heights = [[1,2,2],[3,8,2],[5,3,5]]
//输出：2
//解释：路径 [1,3,5,3,5] 连续格子的差值绝对值最大为 2 。
//这条路径比路径 [1,2,2,2,5] 更优，因为另一条路径差值最大值为 3 。
//
//
// 示例 2：
//
//
//
//
//输入：heights = [[1,2,3],[3,8,4],[5,3,5]]
//输出：1
//解释：路径 [1,2,3,4,5] 的相邻格子差值绝对值最大为 1 ，比路径 [1,3,5,3,5] 更优。
//
//
// 示例 3：
//
//
//输入：heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
//输出：0
//解释：上图所示路径不需要消耗任何体力。
//
//
//
//
// 提示：
//
//
// rows == heights.length
// columns == heights[i].length
// 1 <= rows, columns <= 100
// 1 <= heights[i][j] <= 10⁶
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 二分查找 矩阵 堆（优先队列） 👍 268 👎 0

package main

import (
	"container/heap"
	"fmt"
	"github.com/0RAJA/Rutils/struct/set/unionSet"
	"math"
	"sort"
)

func main() {
	heights := [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}
	fmt.Println(minimumEffortPath(heights))
	heights = [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}
	fmt.Println(minimumEffortPath2(heights))
	fmt.Println(minimumEffortPath2([][]int{{3}}))
}

//leetcode submit region begin(Prohibit modification and deletion)
type Pair struct {
	x, y, maxDiff int
}

type HP []Pair

func (H HP) Len() int { return len(H) }

func (H HP) Less(i, j int) bool { return H[i].maxDiff < H[j].maxDiff }

func (H HP) Swap(i, j int) { H[i], H[j] = H[j], H[i] }

func (H *HP) Push(x interface{}) { *H = append(*H, x.(Pair)) }

func (H *HP) Pop() (t interface{}) {
	t = (*H)[len(*H)-1]
	*H = (*H)[:len(*H)-1]
	return
}

//类似于启发性搜索，从一个点拓展到全部
func minimumEffortPath(heights [][]int) (ret int) {
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	hp := &HP{}
	const INF = math.MaxInt64
	next := [4][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	M, N := len(heights), len(heights[0])
	m := make([][]int, M) //记录每个位置的最小值，用于剪枝。
	for i := range m {
		m[i] = make([]int, N)
		for j := range m[i] {
			m[i][j] = INF
		}
	}
	m[0][0] = 0
	heap.Push(hp, Pair{x: 0, y: 0})
	for true {
		p := heap.Pop(hp).(Pair)
		if p.x == M-1 && p.y == N-1 {
			return p.maxDiff
		}
		if m[p.x][p.y] < p.maxDiff { //剪枝，比之前的记录值大就没必要继续走了
			continue
		}
		for _, v := range next {
			if nx, ny := p.x+v[0], p.y+v[1]; nx >= 0 && ny >= 0 && nx < M && ny < N {
				if diff := max(p.maxDiff, abs(heights[nx][ny]-heights[p.x][p.y])); m[nx][ny] > diff {
					m[nx][ny] = diff
					heap.Push(hp, Pair{x: nx, y: ny, maxDiff: m[nx][ny]})
				}
			}
		}
	}
	return
}

//使用并查集，将所有边按照diff排序，然后加入并查集中，同时判断起点和终点是否重合，重合就返回此刻的diff
func minimumEffortPath2(heights [][]int) (ret int) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	m, n := len(heights), len(heights[0])
	type Edge struct {
		s, e, diff int
	}
	edges := []Edge{}
	for x, row := range heights {
		for y, v := range row {
			idx := x*n + y
			if x > 0 {
				edges = append(edges, Edge{s: idx - n, e: idx, diff: abs(v - heights[x-1][y])})
			}
			if y > 0 {
				edges = append(edges, Edge{s: idx - 1, e: idx, diff: abs(v - heights[x][y-1])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].diff < edges[j].diff })
	myset := unionSet.NewSet(m * n)
	for _, edge := range edges {
		myset.Union(edge.s, edge.e)
		if myset.InSameSet(0, m*n-1) {
			return edge.diff
		}
	}
	return 0 //特殊情况
}

//leetcode submit region end(Prohibit modification and deletion)
