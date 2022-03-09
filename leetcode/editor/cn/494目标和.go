//给你一个整数数组 nums 和一个整数 target 。
//
// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//
//
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
//
//
// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,1,1,1], target = 3
//输出：5
//解释：一共有 5 种方法让最终目标和为 3 。
//-1 + 1 + 1 + 1 + 1 = 3
//+1 - 1 + 1 + 1 + 1 = 3
//+1 + 1 - 1 + 1 + 1 = 3
//+1 + 1 + 1 - 1 + 1 = 3
//+1 + 1 + 1 + 1 - 1 = 3
//
//
// 示例 2：
//
//
//输入：nums = [1], target = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 20
// 0 <= nums[i] <= 1000
// 0 <= sum(nums[i]) <= 1000
// -1000 <= target <= 1000
//
// Related Topics 深度优先搜索 动态规划
// 👍 804 👎 0
package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 107, 109, 113, 127, 131, 137, 3, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 47, 53}
	target := 1000
	fmt.Println(findTargetSumWays3(nums, target))
}

//leetcode submit region begin(Prohibit modification and deletion)

//DFS
func findTargetSumWays(nums []int, target int) int {
	var dfs func(int, int) int
	dfs = func(sum, cnt int) (all int) {
		if cnt == len(nums) {
			if sum == target {
				return 1
			} else {
				return 0
			}
		}
		return dfs(sum+nums[cnt], cnt+1) + dfs(sum-nums[cnt], cnt+1)
	}
	return dfs(0, 0)
}

//DFS记忆化搜索
func findTargetSumWays2(nums []int, target int) int {
	cache := make([]map[int]int, len(nums))
	for i := range cache {
		cache[i] = make(map[int]int)
	}
	var dfs func(sum, idx int) int
	dfs = func(sum, idx int) (ret int) {
		if idx == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}
		if v, ok := cache[idx][sum]; ok {
			return v
		}
		defer func() {
			cache[idx][sum] = ret
		}()
		return dfs(sum+nums[idx], idx+1) + dfs(sum-nums[idx], idx+1)
	}
	return dfs(0, 0)
}

//DP
/*
	假设总和为sum,标记为负数的元素和为x,则标记为正数的元素和为(sum-x)
	则: (sum-x)-x = target
	即: x = (sum-target)/2 为被标记为负数的元素的和，则此题就转换为0-1背包问题和416很像，只需要确定每个数选或不选的情况
	dp[i] += dp[i-v] //dp[i]表示和为i的情况数，因为要求全部的方案数，有些像走方格，只需要加起来自己的上一步的方案数即可。
*/
func findTargetSumWays3(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum-target < 0 || (sum-target)%2 != 0 {
		return 0
	}
	num := (sum - target) / 2
	dp := make([]int, num+1)
	dp[0] = 1
	//物品
	for _, v := range nums {
		//每个只有一个，所以需要反着来，防止使用之前更新的。空间
		for i := num; i >= v; i-- {
			dp[i] += dp[i-v]
		}
	}
	return dp[num]
}

//leetcode submit region end(Prohibit modification and deletion)
