// 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
//
// 题目数据保证答案符合 32 位整数范围。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,2,3], target = 4
// 输出：7
// 解释：
// 所有可能的组合为：
// (1, 1, 1, 1)
// (1, 1, 2)
// (1, 2, 1)
// (1, 3)
// (2, 1, 1)
// (2, 2)
// (3, 1)
// 请注意，顺序不同的序列被视作不同的组合。
//
//
// 示例 2：
//
//
// 输入：nums = [9], target = 3
// 输出：0
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 1000
// nums 中的所有元素 互不相同
// 1 <= target <= 1000
//
//
//
//
// 进阶：如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？
//
// Related Topics 数组 动态规划 👍 1233 👎 0

package p0377

// leetcode submit region begin(Prohibit modification and deletion)
// 完全背包排列计数问题
// 完全背包：从小到大遍历 bag
// 排列：外层遍历 bag，内层遍历 num
// 计数：dp[i] += dp[i-num]; dp[i] 表示 bag 为 i 时的方案数
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1 // bag 为0时的方案只有一种（什么都不选）
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i < num {
				continue
			}
			dp[i] += dp[i-num]
		}
	}
	return dp[target]
}

// leetcode submit region end(Prohibit modification and deletion)
