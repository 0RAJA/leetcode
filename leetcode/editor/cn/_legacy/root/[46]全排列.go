// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
// 
//
// 示例 1： 
//
// 
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 
//
// 示例 2： 
//
// 
// 输入：nums = [0,1]
// 输出：[[0,1],[1,0]]
// 
//
// 示例 3： 
//
// 
// 输入：nums = [1]
// 输出：[[1]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 6 
// -10 <= nums[i] <= 10 
// nums 中的所有整数 互不相同 
// 
//
// Related Topics 数组 回溯 👍 3344 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
// 排列问题，主要点是有序的，顺序改变也会导致结果不同（每层循环从0开始利用每个数字），
// 其次是每个数字不可重复使用（需要把当前排列中使用过的数字记录下来，避免重复使用）
func permute(nums []int) (ret [][]int) {
	seen := make(map[int]bool)
	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			data := make([]int, len(path), len(path))
			copy(data, path)
			ret = append(ret, data)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !seen[nums[i]] {
				seen[nums[i]] = true
				path = append(path, nums[i])
				dfs(path)
				path = path[:len(path)-1]
				seen[nums[i]] = false
			}
		}
	}
	dfs([]int{})
	return
}

// leetcode submit region end(Prohibit modification and deletion)
