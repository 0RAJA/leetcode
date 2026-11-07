// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
// 示例 1：
//
//
// 输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word =
// "ABCCED"
// 输出：true
//
//
// 示例 2：
//
//
// 输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word =
// "SEE"
// 输出：true
//
//
// 示例 3：
//
//
// 输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word =
// "ABCB"
// 输出：false
//
//
//
//
// 提示：
//
//
// m == board.length
// n = board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board 和 word 仅由大小写英文字母组成
//
//
//
//
// 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？
//
// Related Topics 深度优先搜索 数组 字符串 回溯 矩阵 👍 2251 👎 0

package p0079

// leetcode submit region begin(Prohibit modification and deletion)
// 深搜 DFS，每轮依次遍历下一个可选位置，然后记录当前访问了哪个位置（回退后可以取消标记）;注意一个元素的情况
func exist(board [][]byte, word string) bool {
	seen := make(map[int]map[int]bool)
	maxX := len(board)
	maxY := len(board[0])
	for x := range maxX {
		seen[x] = make(map[int]bool)
	}
	var dfs func(x, y, nowIdx int) bool
	dfs = func(x, y, nowIdx int) bool {
		if nowIdx == len(word) {
			return true
		}
		if board[x][y] == word[nowIdx] && !seen[x][y] {
			seen[x][y] = true
			if nowIdx+1 == len(word) {
				return true
			}
			nextStep := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
			for _, step := range nextStep {
				nextX := x + step[0]
				nextY := y + step[1]
				if nextX < 0 || nextX >= maxX || nextY < 0 || nextY >= maxY {
					continue
				}
				if dfs(nextX, nextY, nowIdx+1) {
					return true
				}
			}
			seen[x][y] = false
		}
		return false
	}
	for x, val := range board {
		for y, _ := range val {
			if dfs(x, y, 0) {
				return true
			}
		}
	}
	return false
}

// leetcode submit region end(Prohibit modification and deletion)
