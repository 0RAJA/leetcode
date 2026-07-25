// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
//
//
//
// 示例 1：
//
//
// 输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"]
// ,["1","0","0","1","0"]]
// 输出：4
//
//
// 示例 2：
//
//
// 输入：matrix = [["0","1"],["1","0"]]
// 输出：1
//
//
// 示例 3：
//
//
// 输入：matrix = [["0"]]
// 输出：0
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 300
// matrix[i][j] 为 '0' 或 '1'
//
//
// Related Topics 数组 动态规划 矩阵 👍 1925 👎 0

package p0221

// 1. 暴力求解：时间复杂度：O(mn × min(m,n))  空间复杂度：O(mn)
// 以每个为 1 的元素为正方形左上角向外拓展计算可能的最大正方形（其中每个元素都是1，也就是总和为 (j-i+1) ** 2）
//
//	计算前缀和，用于后续计算范围内的元素总和，来判断是否全为 1
//	    prefixSum[i][j] = prefixSum[i][j] + prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1]
//	以当前 i，j 为左上角,x,y 为右下角，计算最大的全为1的正方形面积
//	    sum = prefixSum[x][y] - prefixSum[i-1][y] - prefixSum[x][j-1] + prefixSum[i-1][j-1]
//	    if sum == (x-i+1) ** 2 则全是 1 用来更新 maxValue
//	    else 直接结束判断
func maximalSquare1(matrix [][]byte) (res int) {
	maxX := len(matrix)
	maxY := len(matrix[0])
	// 计算前缀和，用于后续计算范围内的元素总和，用来判断是否全为 1
	// prefixSum[i][j] = prefixSum[i][j] + prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1]
	prefixSum := make([][]int, len(matrix))
	for i := 0; i < maxX; i++ {
		prefixSum[i] = make([]int, maxY)
	}
	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			if matrix[i][j] == '1' {
				prefixSum[i][j] += 1
			}
			if i != 0 {
				prefixSum[i][j] += prefixSum[i-1][j]
			}
			if j != 0 {
				prefixSum[i][j] += prefixSum[i][j-1]
			}
			if i != 0 && j != 0 {
				prefixSum[i][j] -= prefixSum[i-1][j-1]
			}
		}
	}
	// 以当前 i，j 为左上角，计算最大的正方形面积
	countMaximalSquare := func(i, j int) {
		// sum = prefixSum[x][y] - prefixSum[i-1][y] - prefixSum[x][j-1] + prefixSum[i-1][j-1]
		// if sum == (x-i+1) ** 2 则全是 1 用来更新 maxValue
		// else 直接结束判断
		for x, y := i, j; x < maxX && y < maxY; {
			sum := prefixSum[x][y]
			if i != 0 {
				sum -= prefixSum[i-1][y]
			}
			if j != 0 {
				sum -= prefixSum[x][j-1]
			}
			if i != 0 && j != 0 {
				sum += prefixSum[i-1][j-1]
			}
			if sum == (x-i+1)*(x-i+1) {
				res = max(res, sum)
			} else {
				return
			}
			x++
			y++
		}
	}
	// 遍历matrix中每个1来挨个计算可能的正方形
	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			countMaximalSquare(i, j)
		}
	}
	return
}

// leetcode submit region begin(Prohibit modification and deletion)
// 2. 动态规划：每个为1的元素 i,j 的最大为1的边数 = min(dp[i-1,j],dp[i][j-1],dp[i-1][j-1]) + 1
// dp[i][j] = 以 matrix[i][j] 为右下角的最大正方形边长
// 当 matrix[i][j] == '1' 时：
// dp[i][j] = min(
//
//	dp[i-1][j],
//	dp[i][j-1],
//	dp[i-1][j-1],
//
// ) + 1
// 想让当前正方形边长从 k 扩展到 k+1，需要保证：
//
// 左边至少能提供 k 行；
// 上边至少能提供 k 列；
// 左上区域本身至少是 k × k 的全 1 正方形。
//
// 只要其中一个方向较短，就会限制正方形扩张，因此取最小值。
func maximalSquare(matrix [][]byte) (res int) {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				continue
			}
			// 每个为1的元素 i,j 的最大为1的边数 = min(dp[i-1,j],dp[i][j-1],dp[i-1][j-1]) + 1
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
			}
			res = max(res, dp[i][j]*dp[i][j])
		}
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
