//给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
//
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
//
//
//
//
//
// 示例 1:
//
//
//输入: rowIndex = 3
//输出: [1,3,3,1]
//
//
// 示例 2:
//
//
//输入: rowIndex = 0
//输出: [1]
//
//
// 示例 3:
//
//
//输入: rowIndex = 1
//输出: [1,1]
//
//
//
//
// 提示:
//
//
// 0 <= rowIndex <= 33
//
//
//
//
// 进阶：
//
// 你可以优化你的算法到 O(rowIndex) 空间复杂度吗？
// Related Topics 数组 动态规划 👍 366 👎 0

package main

import "fmt"

func main() {
	fmt.Println(getRow(0))
}

//leetcode submit region begin(Prohibit modification and deletion)
func getRow(rowIndex int) []int {
	dp := make([][]int, rowIndex+2)
	for i := range dp {
		dp[i] = make([]int, rowIndex+2)
	}
	dp[0][0] = 1
	for i := 1; i <= rowIndex+1; i++ {
		for j := 1; j <= i; j++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
		}
	}
	return dp[rowIndex+1][1:]
}

//leetcode submit region end(Prohibit modification and deletion)
