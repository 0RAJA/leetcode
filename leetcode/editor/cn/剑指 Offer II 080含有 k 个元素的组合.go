// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//
//
//
// 示例 1:
//
//
// 输入:n = 4, k = 2
// 输出:
// [
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
// ]
//
// 示例 2:
//
//
// 输入: n = 1, k = 1
// 输出: [[1]]
//
//
//
// 提示:
//
//
// 1 <= n <= 20
// 1 <= k <= n
//
//
//
//
// 注意：本题与主站 77 题相同： https://leetcode-cn.com/problems/combinations/
// Related Topics 数组 回溯 👍 19 👎 0

package main

import (
	"fmt"
)

func main() {
	fmt.Println(combine(4, 2))
}

// leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) (ret [][]int) {
	result := make([]int, 0, k)
	var dfs func(idx int)
	dfs = func(idx int) {
		if len(result)+(n-idx+1) < k {
			return
		}
		if len(result) == k {
			ret = append(ret, append([]int{}, result...))
			return
		}
		result = append(result, idx)
		dfs(idx + 1)
		result = result[:len(result)-1]
		dfs(idx + 1)
	}
	dfs(1)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
