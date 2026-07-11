// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置在下标 0。
//
// 每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在索引 i 处，你可以跳转到任意 (i + j) 处： 
//
// 
// 0 <= j <= nums[i] 且 
// i + j < n 
// 
//
// 返回到达 n - 1 的最小跳跃次数。测试用例保证可以到达 n - 1。 
//
// 
//
// 示例 1: 
//
// 
// 输入: nums = [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
//      从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
// 
//
// 示例 2: 
//
// 
// 输入: nums = [2,3,0,1,4]
// 输出: 2
// 
//
// 
//
// 提示: 
//
// 
// 1 <= nums.length <= 10⁴ 
// 0 <= nums[i] <= 1000 
// 题目保证可以到达 n - 1 
// 
//
// Related Topics 贪心 数组 动态规划 👍 3072 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
// 贪心：每次算当前能到的最远的位置，以及路程中算下一次能到达的最远位置，如果当前到了本次最远位置还没到终点，那就新跳一此，更新下能到的最远位置
func jump(nums []int) (res int) {
	// 当前能到的最大范围
	nowMaxDest := 0
	nextMaxDest := 0
	for i := 0; i < len(nums); i++ {
		// 每次更新下次能到的最远边界，为下次跳做好准备
		nextMaxDest = max(nextMaxDest, i+nums[i])
		// 当前直接能到终点，那就不用再跳了，直接结束
		if nowMaxDest >= len(nums)-1 {
			break
		}
		// 如果已经到到本轮的最大边界，那就新跳一此，更新下能到的最远位置
		if i == nowMaxDest {
			res++
			nowMaxDest = nextMaxDest
		}
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
