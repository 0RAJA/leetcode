// 给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数
// 字，并以数组的形式返回结果。
//
//
//
// 示例 1：
//
//
// 输入：nums = [4,3,2,7,8,2,3,1]
// 输出：[5,6]
//
//
// 示例 2：
//
//
// 输入：nums = [1,1]
// 输出：[2]
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 10⁵
// 1 <= nums[i] <= n
//
//
// 进阶：你能在不使用额外空间且时间复杂度为 O(n) 的情况下解决这个问题吗? 你可以假定返回的数组不算在额外空间内。
//
// Related Topics 数组 哈希表 👍 1432 👎 0

package p0448

// leetcode submit region begin(Prohibit modification and deletion)
// 遍历数组，把所有遍历过的位置置为 -1，把所有出现过的位置置为 -2，最后再遍历一遍，找出 -1 的点
func findDisappearedNumbers(nums []int) (res []int) {
	res = make([]int, 0, len(nums))
	var dfs func(i int)
	dfs = func(i int) {
		if nums[i] == -2 {
			return
		}
		if nums[i] == -1 {
			nums[i] = -2
			return
		}
		nextIdx := nums[i]
		nums[i] = -2
		dfs(nextIdx - 1)
	}
	for i, nextIdx := range nums {
		if nextIdx == -1 || nextIdx == -2 {
			continue
		} else {
			nums[i] = -1
			dfs(nextIdx - 1)
		}
	}
	for i, nextIdx := range nums {
		if nextIdx == -1 {
			res = append(res, i+1)
		}
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
