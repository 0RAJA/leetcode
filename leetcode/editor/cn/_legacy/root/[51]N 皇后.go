// 按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
//
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。 
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。 
//
// 
// 
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。 
// 
// 
//
// 
//
// 示例 1： 
// 
// 
// 输入：n = 4
// 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// 解释：如上图所示，4 皇后问题存在两个不同的解法。
// 
//
// 示例 2： 
//
// 
// 输入：n = 1
// 输出：[["Q"]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 9 
// 
//
// Related Topics 数组 回溯 👍 2493 👎 0
package main

import (
	"strings"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 回溯算法，穷举法；如何定义维度：每层表示第几行，每层循环表示第几列，
// 判定 ok 的条件，最终层数超过 n；
// 满足条件，x，y 表示当前占用的行、列，其他人不可占用；对角线也不可占用
func solveNQueens(n int) (res [][]string) {
	// m 记录元素
	m := make(map[int]map[int]bool)
	for i := 0; i < n; i++ {
		m[i] = make(map[int]bool)
	}
	// x 表示行，y 表示列，diag1 表示对角线1，diag2 表示对角线2
	x := make(map[int]bool)     // 表示第几行被占用
	y := make(map[int]bool)     // 表示第几列被占用
	diag1 := make(map[int]bool) // 表示第几条对角线1被占用
	diag2 := make(map[int]bool) // 表示第几条对角线2被占用
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == n {
			data := make([]string, 0, n)
			for i := 0; i < n; i++ {
				tmp := make([]string, n)
				for j := 0; j < n; j++ {
					if m[i][j] {
						tmp[j] = "Q"
					} else {
						tmp[j] = "."
					}
				}
				data = append(data, strings.Join(tmp, ""))
			}
			res = append(res, data)
			return
		}
		for j := 0; j < n; j++ {
			if !y[j] && !diag1[idx-j] && !diag2[idx+j] {
				// 占用位置
				x[idx] = true
				y[j] = true
				diag1[idx-j] = true
				diag2[idx+j] = true
				m[idx][j] = true
				dfs(idx + 1)
				// 解除位置
				m[idx][j] = false
				y[j] = false
				diag2[idx+j] = false
				diag1[idx-j] = false
				x[idx] = false
			}
		}
	}
	dfs(0)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
