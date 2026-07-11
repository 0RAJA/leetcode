// 给你一个正整数 n ，生成一个包含 1 到 n² 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
//
// 
//
// 示例 1： 
// 
// 
// 输入：n = 3
// 输出：[[1,2,3],[8,9,4],[7,6,5]]
// 
//
// 示例 2： 
//
// 
// 输入：n = 1
// 输出：[[1]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 20 
// 
//
// Related Topics 数组 矩阵 模拟 👍 1550 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	// 模拟题，把控好规律和边界条件即可
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int, n)
	}
	num := 1
	// 定义好上下左边边界，这个边界会变小
	minX, minY, maxX, maxY := 0, 0, n-1, n-1
	for minX <= maxX && minY <= maxY {
		// 从左到右, x = minX; y 从 minY -> maxY
		for y := minY; y <= maxY; y++ {
			nums[minX][y] = num
			num++
		}
		// minX 上边缩小
		minX++
		// 从上到下，y = maxY； x 从 minX -> maxX
		for x := minX; x <= maxX; x++ {
			nums[x][maxY] = num
			num++
		}
		// maxY 右边缩小
		maxY--
		// 从右往左；x = maxX；y 从 maxY -> minY
		for y := maxY; y >= minY; y-- {
			nums[maxX][y] = num
			num++
		}
		// maxX 右边缩小
		maxX--
		// 从下到上；y = minY;x 从 maxX -> minX
		for x := maxX; x >= minX; x-- {
			nums[x][minY] = num
			num++
		}
		// minY 下边缩小
		minY++
	}
	return nums
}

// leetcode submit region end(Prohibit modification and deletion)
