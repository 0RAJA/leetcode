//给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,5,11,5]
//输出：true
//解释：数组可以分割成 [1, 5, 5] 和 [11] 。
//
// 示例 2：
//
//
//输入：nums = [1,2,3,5]
//输出：false
//解释：数组不能分割成两个元素和相等的子集。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 100
//
// Related Topics 数组 动态规划 👍 1177 👎 0

package main

import (
	"fmt"
)

func main() {
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}

//leetcode submit region begin(Prohibit modification and deletion)
/*
	通过遍历所有的数，然后尝试将此数在当前范围内进行使用，直到范围最大时确保所有的数都已经尝试使用过。
	dp[i] = dp[i] || dp[i-nums[j]]
	j = [0,n-1]
	i = [sum/2,j] (i反向进行遍历是确保使用的数据都是上一轮更新的而不是此轮更新的。)
*/
func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sum := 0
	max := 0
	for _, v := range nums {
		if max < v {
			max = v
		}
		sum += v
	}
	if sum%2 != 0 || sum/2 < max {
		return false
	}
	dp := make([]bool, sum/2+1)
	dp[0] = true
	for _, v := range nums {
		for j := sum / 2; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	return dp[sum/2]
}

//leetcode submit region end(Prohibit modification and deletion)
