// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除了 nums[i] 之外其余各元素的乘积 。
//
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内。
//
// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
//
//
//
// 示例 1:
//
//
// 输入: nums = [1,2,3,4]
// 输出: [24,12,8,6]
//
//
// 示例 2:
//
//
// 输入: nums = [-1,1,0,-3,3]
// 输出: [0,0,9,0,0]
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 10⁵
// -30 <= nums[i] <= 30
// 输入 保证 数组 answer[i] 在 32 位 整数范围内
//
//
//
//
// 进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）
//
// Related Topics 数组 前缀和 👍 2286 👎 0

package p0238

// leetcode submit region begin(Prohibit modification and deletion)
// 前缀乘积 & 后缀乘积
func productExceptSelf(nums []int) (res []int) {
	res = make([]int, len(nums))
	for i := range res {
		res[i] = 1
	}
	// 先算前缀(每个前缀不包含自身)
	prefixMulti := 1
	for i := 1; i < len(nums); i++ {
		prefixMulti *= nums[i-1]
		res[i] *= prefixMulti
	}
	// 再算后缀(每个后缀不包含自身)
	suffixMulti := 1
	for i := len(nums) - 2; i >= 0; i-- {
		suffixMulti *= nums[i+1]
		res[i] *= suffixMulti
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
