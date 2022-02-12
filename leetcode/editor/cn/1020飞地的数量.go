//给你一个大小为 m x n 的二进制矩阵 grid ，其中 0 表示一个海洋单元格、1 表示一个陆地单元格。
//
// 一次 移动 是指从一个陆地单元格走到另一个相邻（上、下、左、右）的陆地单元格或跨过 grid 的边界。
//
// 返回网格中 无法 在任意次数的移动中离开网格边界的陆地单元格的数量。
//
//
//
// 示例 1：
//
//
//输入：grid = [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
//输出：3
//解释：有三个 1 被 0 包围。一个 1 没有被包围，因为它在边界上。
//
//
// 示例 2：
//
//
//输入：grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
//输出：0
//解释：所有 1 都在边界上或可以到达边界。
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 500
// grid[i][j] 的值为 0 或 1
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 72 👎 0

package main

import (
	"container/list"
	"fmt"
)

func main() {
	grid := [][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}
	fmt.Println(numEnclaves(grid))
}

//leetcode submit region begin(Prohibit modification and deletion)
func numEnclaves(grid [][]int) (ret int) {
	type Point struct{ x, y int }
	next := [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	bfs := func(x, y int) {
		queue := list.New()
		isOK := false
		num := 1
		queue.PushBack(Point{x, y})
		grid[x][y] = 2
		for queue.Len() > 0 {
			p := queue.Front().Value.(Point)
			queue.Remove(queue.Front())
			for _, v := range next {
				nx, ny := p.x+v[0], p.y+v[1]
				if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[nx]) {
					isOK = true
				} else if grid[nx][ny] == 1 {
					grid[nx][ny] = 2
					num++
					queue.PushBack(Point{x: nx, y: ny})
				}
			}
		}
		if !isOK {
			ret += num
		}
	}
	for x, row := range grid {
		for y, v := range row {
			if v == 1 {
				bfs(x, y)
			}
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
