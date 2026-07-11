// 在一个 2 x 3 的板上（board）有 5 块砖瓦，用数字 1~5 来表示, 以及一块空缺用 0 来表示.
//
// 一次移动定义为选择 0 与一个相邻的数字（上下左右）进行交换.
//
// 最终当板 board 的结果是 [[1,2,3],[4,5,0]] 谜板被解开。
//
// 给出一个谜板的初始状态，返回最少可以通过多少次移动解开谜板，如果不能解开谜板，则返回 -1 。
//
// 示例：
//
// 输入：board = [[1,2,3],[4,0,5]]
// 输出：1
// 解释：交换 0 和 5 ，1 步完成
//
// 输入：board = [[1,2,3],[5,4,0]]
// 输出：-1
// 解释：没有办法完成谜板
//
// 输入：board = [[4,1,2],[5,0,3]]
// 输出：5
// 解释：
// 最少完成谜板的最少移动次数是 5 ，
// 一种移动路径:
// 尚未移动: [[4,1,2],[5,0,3]]
// 移动 1 次: [[4,1,2],[0,5,3]]
// 移动 2 次: [[0,1,2],[4,5,3]]
// 移动 3 次: [[1,0,2],[4,5,3]]
// 移动 4 次: [[1,2,0],[4,5,3]]
// 移动 5 次: [[1,2,3],[4,5,0]]
//
// 输入：board = [[3,2,4],[1,5,0]]
// 输出：14
//
// 提示：
//
// board 是一个如上所述的 2 x 3 的数组.
// board[i][j] 是一个 [0, 1, 2, 3, 4, 5] 的排列.
//
// Related Topics 广度优先搜索 数组 矩阵
// 👍 148 👎 0
package main

import (
	"fmt"
	"strconv"
)

func main() {
	board := [][]int{{4, 1, 2}, {5, 0, 3}}
	fmt.Println(slidingPuzzle(board))
}

// leetcode submit region begin(Prohibit modification and deletion)
type pail struct {
	key  string
	step int
}

func slidingPuzzle(board [][]int) int {
	var (
		end  = [][]int{{1, 2, 3}, {4, 5, 0}}
		rows = 2
		cols = 3
		next = [4][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	)
	getKey := func(num [][]int) (str string) { //确保6个数位置对应key唯一
		for _, v := range num {
			for _, x := range v {
				str += strconv.Itoa(x)
			}
		}
		return str
	}
	endKey := getKey(end)
	startKey := getKey(board)
	if startKey == endKey {
		return 0
	}
	viewMap := make(map[string]bool)         //是否访问过的列表
	get := func(key string) (ret []string) { //获取此状态所有一次变换的结果
		var sx, sy int
		num := make([][]int, 2)
		for i := 0; i < 2; i++ {
			num[i] = make([]int, 3)
		}
		for i := 0; i < 3; i++ {
			num[0][i] = int(key[i] - '0')
		}
		for i := 0; i < 3; i++ {
			num[1][i] = int(key[i+3] - '0')
		}
		for i, v := range num {
			for j, x := range v {
				if x == 0 {
					sx = i
					sy = j
				}
			}
		}
		for i := 0; i < 4; i++ {
			nx := sx + next[i][0]
			ny := sy + next[i][1]
			if nx < 0 || ny < 0 || nx >= rows || ny >= cols {
				continue
			}
			num[nx][ny], num[sx][sy] = num[sx][sy], num[nx][ny]
			ret = append(ret, getKey(num))
			num[nx][ny], num[sx][sy] = num[sx][sy], num[nx][ny]
		}
		return ret
	}
	//BFS
	queue := []pail{{
		key:  startKey,
		step: 0,
	}}
	viewMap[startKey] = true
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for _, key := range get(p.key) {
			if !viewMap[key] {
				if key == endKey {
					return p.step + 1
				}
				queue = append(queue, pail{
					key:  key,
					step: p.step + 1,
				})
				viewMap[key] = true
			}
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
