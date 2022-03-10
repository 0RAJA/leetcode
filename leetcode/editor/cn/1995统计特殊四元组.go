//给你一个 下标从 0 开始 的整数数组 nums ，返回满足下述条件的 不同 四元组 (a, b, c, d) 的 数目 ：
//
//
// nums[a] + nums[b] + nums[c] == nums[d] ，且
// a < b < c < d
//
//
//
//
// 示例 1：
//
// 输入：nums = [1,2,3,6]
//输出：1
//解释：满足要求的唯一一个四元组是 (0, 1, 2, 3) 因为 1 + 2 + 3 == 6 。
//
//
// 示例 2：
//
// 输入：nums = [3,3,6,4,5]
//输出：0
//解释：[3,3,6,4,5] 中不存在满足要求的四元组。
//
//
// 示例 3：
//
// 输入：nums = [1,1,1,3,5]
//输出：4
//解释：满足要求的 4 个四元组如下：
//- (0, 1, 2, 3): 1 + 1 + 1 == 3
//- (0, 1, 3, 4): 1 + 1 + 3 == 5
//- (0, 2, 3, 4): 1 + 1 + 3 == 5
//- (1, 2, 3, 4): 1 + 1 + 3 == 5
//
//
//
//
// 提示：
//
//
// 4 <= nums.length <= 50
// 1 <= nums[i] <= 100
//
// Related Topics 数组 枚举 👍 79 👎 0

package main

import (
	"fmt"
)

func main() {
	fmt.Println(countQuadruplets([]int{1, 1, 1, 3, 5}))
}

//leetcode submit region begin(Prohibit modification and deletion)
/*
	dp[target][cnt] = 数值为target位数cnt
*/
func countQuadruplets(nums []int) (ret int) {
	if len(nums) < 4 {
		return 0
	}
	dp := make([][]int, 110)
	for i := range dp {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = 1
	for _, v := range nums {
		for j := 109; j >= v; j-- {
			for k := 3; k >= 1; k-- {
				dp[j][k] += dp[j-v][k-1]
			}
		}
		ret += dp[v][3]
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
