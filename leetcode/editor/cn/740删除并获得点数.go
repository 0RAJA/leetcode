//给你一个整数数组 nums ，你可以对它进行一些操作。
//
// 每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除每个等于 nums[i] - 1 或 nums[i] +
// 1 的元素。
//
// 开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。
//
//
//
// 示例 1：
//
//
//输入：nums = [3,4,2]
//输出：6
//解释：
//删除 4 获得 4 个点数，因此 3 也被删除。
//之后，删除 2 获得 2 个点数。总共获得 6 个点数。
//
//
// 示例 2：
//
//
//输入：nums = [2,2,3,3,3,4]
//输出：9
//解释：
//删除 3 获得 3 个点数，接着要删除两个 2 和 4 。
//之后，再次删除 3 获得 3 个点数，再次删除 3 获得 3 个点数。
//总共获得 9 个点数。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2 * 104
// 1 <= nums[i] <= 104
//
// Related Topics 动态规划
// 👍 279 👎 0
package main

import "fmt"

func main() {
	nums := []int{2, 2, 3, 3, 3, 4}
	fmt.Println(deleteAndEarn(nums))
}

//leetcode submit region begin(Prohibit modification and deletion)
/*
其实每次选择一个数就获取那个数所有个体的值的总和，
转换为打家劫舍问题：
	dp数组计算每个个体的值得总和，则表示选择dp[i] 则可以获取nums[i]的值，
	则进行判断
	dp[i] = max(dp[i-1],dp[i-2]+nums[i])
	i = [2,len(nums)+2]
*/
func deleteAndEarn(nums []int) int {
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	dp := [20002]int{}
	maxIndex, minIndex := -1, 20002
	for _, value := range nums {
		dp[value] += value
		if maxIndex < value {
			maxIndex = value
		}
		if minIndex > value {
			minIndex = value
		}
	}
	if minIndex == maxIndex {
		return dp[minIndex]
	}
	for i := 2; i <= maxIndex+1; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+dp[i])
	}
	return dp[maxIndex+1]
}

//leetcode submit region end(Prohibit modification and deletion)
