// 给你一个非负整数数组 nums 和一个整数 target 。
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
// 输入：nums = [1,1,1,1,1], target = 3
// 输出：5
// 解释：一共有 5 种方法让最终目标和为 3 。
// -1 + 1 + 1 + 1 + 1 = 3
// +1 - 1 + 1 + 1 + 1 = 3
// +1 + 1 - 1 + 1 + 1 = 3
// +1 + 1 + 1 - 1 + 1 = 3
// +1 + 1 + 1 + 1 - 1 = 3
//
//
// 示例 2：
//
//
// 输入：nums = [1], target = 1
// 输出：1
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
//
// Related Topics 数组 动态规划 回溯 背包问题 0-1 背包 👍 2321 👎 0

package p0494

// leetcode submit region begin(Prohibit modification and deletion)
// 假设选为正数的 num 和为 X，选为负数的 num 和为 Y
// X+Y = sum；X - Y = target
// 2X = sum + target ; X = (sum+target)/2
// 即：从 nums 中选择 num 使得总和为 (sum+target)/2 的方法个数 -> 0-1背包计数问题
// dp[i] += dp[i-num] + 1 ; dp[i] 表示总数为 i 的使用 num 的方案数
// 注意每个 num 只能使用一次
// 0-1背包组合计数问题
// 0-1背包：从大到小遍历 bag
// 组合：先遍历 nums 再遍历 bag
// 计数：dp[i] += dp[i-num] + 1
func findTargetSumWays(nums []int, target int) (res int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	bag := (sum + target) / 2
	if (sum+target)%2 != 0 || bag < 0 {
		return 0
	}
	dp := make([]int, bag+1)
	dp[0] = 1
	for _, num := range nums {
		for i := bag; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[bag]
}

// leetcode submit region end(Prohibit modification and deletion)
