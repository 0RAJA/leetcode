// 有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
//
// 现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i -
// 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。
//
// 求所能获得硬币的最大数量。
//
//
// 示例 1：
//
//
// 输入：nums = [3,1,5,8]
// 输出：167
// 解释：
// nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
// coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
//
// 示例 2：
//
//
// 输入：nums = [1,5]
// 输出：10
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 300
// 0 <= nums[i] <= 100
//
//
// Related Topics 数组 动态规划 👍 1518 👎 0

package p0312

// leetcode submit region begin(Prohibit modification and deletion)
// 区间 dp
// dp[left][right] 表示区间内所有气球能获得的最大硬币
// dp[left][right] = max(dp[left][right], dp[left][k] + nums[left]*nums[k]*nums[right] + dp[k][right])
// 区间长度 [2,N+1];遍历每个区间下的 k 从 [left+1,right-1]
func maxCoins(nums []int) int {
	// 为两侧填充 1 便于计算
	nums = append([]int{1}, append(nums, 1)...)
	// dp[left][right]包含从left到right所能获取的最大硬币数
	dp := make(map[int]map[int]int)
	for i := 0; i < len(nums)+2; i++ {
		dp[i] = make(map[int]int)
	}
	for interval := 2; interval <= len(nums); interval++ {
		for left, right := 0, interval; right < len(nums); {
			for k := left + 1; k < right; k++ {
				dp[left][right] = max(dp[left][right], dp[left][k]+nums[left]*nums[k]*nums[right]+dp[k][right])
			}
			left++
			right++
		}
	}
	return dp[0][len(nums)-1]
}

// leetcode submit region end(Prohibit modification and deletion)
