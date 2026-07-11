//输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
//
//
//
// 示例 1：
//
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//
//
// 示例 2：
//
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//
//
// 限制：
//
//
// 0 <= matrix.length <= 100
// 0 <= matrix[i].length <= 100
//
//
// 注意：本题与主站 54 题相同：https://leetcode-cn.com/problems/spiral-matrix/
// Related Topics 数组 矩阵 模拟 👍 334 👎 0

package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

// leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) (ret []int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	cnt := len(matrix) * len(matrix[0])
	ret = make([]int, 0, cnt)
	u, d, l, r := 0, len(matrix)-1, 0, len(matrix[0])-1
	for len(ret) < cnt {
		for i := l; i <= r; i++ {
			ret = append(ret, matrix[u][i])
		}
		for i := u + 1; i <= d; i++ {
			ret = append(ret, matrix[i][r])
		}
		for i := r - 1; i >= l; i-- {
			ret = append(ret, matrix[d][i])
		}
		for i := d - 1; i >= u+1; i-- {
			ret = append(ret, matrix[i][l])
		}
		l++
		r--
		u++
		d--
	}
	return ret[:cnt]
}

//leetcode submit region end(Prohibit modification and deletion)
