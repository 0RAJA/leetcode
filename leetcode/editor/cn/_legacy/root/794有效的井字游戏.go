//给你一个字符串数组 board 表示井字游戏的棋盘。当且仅当在井字游戏过程中，棋盘有可能达到 board 所显示的状态时，才返回 true 。
//
// 井字游戏的棋盘是一个 3 x 3 数组，由字符 ' '，'X' 和 'O' 组成。字符 ' ' 代表一个空位。
//
// 以下是井字游戏的规则：
//
//
// 玩家轮流将字符放入空位（' '）中。
// 玩家 1 总是放字符 'X' ，而玩家 2 总是放字符 'O' 。
// 'X' 和 'O' 只允许放置在空位中，不允许对已放有字符的位置进行填充。
// 当有 3 个相同（且非空）的字符填充任何行、列或对角线时，游戏结束。
// 当所有位置非空时，也算为游戏结束。
// 如果游戏结束，玩家不允许再放置字符。
//
//
//
//
// 示例 1：
//
//
//输入：board = ["O  ","   ","   "]
//输出：false
//解释：玩家 1 总是放字符 "X" 。
//
//
// 示例 2：
//
//
//输入：board = ["XOX"," X ","   "]
//输出：false
//解释：玩家应该轮流放字符。
//
// 示例 3：
//
//
//输入：board = ["XXX","   ","OOO"]
//输出：false
//
//
// Example 4:
//
//
//输入：board = ["XOX","O O","XOX"]
//输出：true
//
//
//
//
// 提示：
//
//
// board.length == 3
// board[i].length == 3
// board[i][j] 为 'X'、'O' 或 ' '
//
// Related Topics 数组 字符串 👍 86 👎 0

package main

import "fmt"

func main() {
	board := []string{"XXX", "XOO", "OO "}
	fmt.Println(validTicTacToe(board))
}

// leetcode submit region begin(Prohibit modification and deletion)
func validTicTacToe(board []string) bool {
	numsA := 0 //数量 ->AB只差小于等于1
	numsB := 0
	rowsA := [3]int{} //行中数的个数
	rowsB := [3]int{}
	colsA := [3]int{} //列中数的个数
	colsB := [3]int{}
	diagonalA := [2]int{} //对角中数的个数
	diagonalB := [2]int{}
	winA := 0 //A赢的个数
	winB := 0 //B赢的个数
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' {
				numsA++
				rowsA[i]++
				colsA[j]++
				if i == j {
					diagonalA[0]++
				}
				if i+j == 2 {
					diagonalA[1]++
				}
			} else if board[i][j] == 'O' {
				numsB++
				rowsB[i]++
				colsB[j]++
				if i == j {
					diagonalB[0]++
				}
				if i+j == 2 {
					diagonalB[1]++
				}
			}
		}
	}
	if numsA-numsB > 1 || numsA < numsB || numsA == 0 {
		return false
	}
	for i := 0; i < 3; i++ {
		if rowsA[i] == 3 {
			winA++
		}
		if colsA[i] == 3 {
			winA++
		}
		if rowsB[i] == 3 {
			winB++
		}
		if colsB[i] == 3 {
			winB++
		}
	}
	for i := 0; i < 2; i++ {
		if diagonalA[i] == 3 {
			winA++
		}
		if diagonalB[i] == 3 {
			winB++
		}
	}
	/*
		前提: A数量>=B数量 且 A和B不能同时赢
		1. A 和 B 都不赢
		2. A赢且A比B数量多1
		3. B赢且A与B数量相等
	*/
	return !(winA > 0 && winB > 0) && (winA == 0 && winB == 0 || winA > 0 && numsA-numsB == 1 || winB > 0 && numsB == numsA)
}

//leetcode submit region end(Prohibit modification and deletion)
