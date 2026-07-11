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
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
// ]
// 输出：1
//
//
// 示例 2：
//
//
// 输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
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
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 1972 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
var next = [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func numIslands(grid [][]byte) (ret int) {
	isOk := func(x, y int) bool {
		return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == '1'
	}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		grid[x][y] = '0'
		for i := 0; i < 4; i++ {
			nx, ny := x+next[i][0], y+next[i][1]
			if isOk(nx, ny) {
				dfs(nx, ny)
			}
		}
	}
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == '1' {
				ret++
				dfs(i, j)
			}
		}
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
