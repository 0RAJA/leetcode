// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
//
//
// 示例 1：
//
//
// 输入：grid = [
//   ['1','1','1','1','0'],
//   ['1','1','0','1','0'],
//   ['1','1','0','0','0'],
//   ['0','0','0','0','0']
// ]
// 输出：1
//
//
// 示例 2：
//
//
// 输入：grid = [
//   ['1','1','0','0','0'],
//   ['1','1','0','0','0'],
//   ['0','0','1','0','0'],
//   ['0','0','0','1','1']
// ]
// 输出：3
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 3013 👎 0

package p0200

// leetcode submit region begin(Prohibit modification and deletion)
// 染色问题，我个人理解是遍历每个未染色节点，bfs 进行染色
func numIslands(grid [][]byte) (res int) {
	m := len(grid)
	n := len(grid[0])
	visited := make(map[[2]int]bool) // key 记录 x，y，value 记录是否访问过
	q := make([][2]int, 0, 300)
	bfs := func(x, y int) {
		q = q[:0]
		q = append(q, [2]int{x, y})
		for len(q) > 0 {
			now := q[0]
			q = q[1:]
			nextCur := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
			for _, cur := range nextCur {
				nextX := now[0] + cur[0]
				nextY := now[1] + cur[1]
				if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n || grid[nextX][nextY] == '0' || visited[[2]int{nextX, nextY}] {
					continue
				}
				visited[[2]int{nextX, nextY}] = true
				q = append(q, [2]int{nextX, nextY})
			}
		}
	}
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '0' || visited[[2]int{x, y}] {
				continue
			} else {
				visited[[2]int{x, y}] = true
				res++
				bfs(x, y)
			}
		}
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
