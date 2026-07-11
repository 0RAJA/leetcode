//给你一个由 无重复 正整数组成的集合 nums ，请你找出并返回其中最大的整除子集 answer ，子集中每一元素对 (answer[i], answer[
//j]) 都应当满足：
//
// answer[i] % answer[j] == 0 ，或
// answer[j] % answer[i] == 0
//
//
// 如果存在多个有效解子集，返回其中任何一个均可。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[1,2]
//解释：[1,3] 也会被视为正确答案。
//
//
// 示例 2：
//
//
//输入：nums = [1,2,4,8]
//输出：[1,2,4,8]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 1000
// 1 <= nums[i] <= 2 * 10⁹
// nums 中的所有整数 互不相同
//
// Related Topics 数组 数学 动态规划 排序 👍 429 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestDivisibleSubset([]int{576, 288, 144, 72, 64, 32, 16, 3}))
}

//leetcode submit region begin(Prohibit modification and deletion)
/*
	// a % b == 0 && b % c == 0  得: a % c == 0
	sort.Ints()
	for i range [1,n-1] :
		for j range [0,i-1] :
			if nums[i] % nums[j] == 0 :
				dp[i] = max(dp[i],dp[j]+1) // 经典计算前i个中最多多少个能被nums[i]整除的。
*/
func largestDivisibleSubset(nums []int) (res []int) {
	sort.Ints(nums)

	// 第 1 步：动态规划找出最大子集的个数、最大子集中的最大整数
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	maxSize, maxVal := 1, 1
	for i := 1; i < n; i++ {
		for j, v := range nums[:i] {
			if nums[i]%v == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxSize {
			maxSize, maxVal = dp[i], nums[i]
		}
	}

	if maxSize == 1 {
		return []int{nums[0]}
	}

	// 第 2 步：倒推获得最大子集
	for i := n - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxVal%nums[i] == 0 {
			res = append(res, nums[i])
			maxVal = nums[i]
			maxSize--
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
