//给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
//
//
// 请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。
//
// 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。
//
//
//
//
// 示例 1：
//
//
//输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
//输出：4
//解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
//其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于
//n 的值 3 。
//
//
// 示例 2：
//
//
//输入：strs = ["10", "0", "1"], m = 1, n = 1
//输出：2
//解释：最大的子集是 {"0", "1"} ，所以答案是 2 。
//
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 600
// 1 <= strs[i].length <= 100
// strs[i] 仅由 '0' 和 '1' 组成
// 1 <= m, n <= 100
//
// Related Topics 动态规划
// 👍 440 👎 0
package main

import (
	"fmt"
)

//不会
func main() {
	strs := []string{"10", "0", "1"}
	fmt.Println(findMaxForm(strs, 1, 1))
}

//leetcode submit region begin(Prohibit modification and deletion)
/*
	dp[x][y] = max(dp[x][y],dp[x-a][y-b]+1)
	a,b = nums[i][0],nums[i][1]
	i = [0,len(strs)-1]
	x = [m,1]
	y = [n,1]
*/
func findMaxForm(strs []string, m int, n int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	nums := make([][2]int, 0, len(strs)) //用来记录每个字符串的0,1个数
	for _, s := range strs {
		var num0, num1 int
		for _, v := range s {
			if v == '0' {
				num0++
			} else {
				num1++
			}
		}
		nums = append(nums, [2]int{num0, num1})
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, v := range nums {
		a, b := v[0], v[1]
		for x := m; x >= 0; x-- { //注意反向是为了使用没有更新的数据来进行更新
			for y := n; y >= 0; y-- {
				if x >= a && y >= b {
					dp[x][y] = max(dp[x][y], dp[x-a][y-b]+1)
				}
			}
		}
	}
	return dp[m][n]
}

//leetcode submit region end(Prohibit modification and deletion)
