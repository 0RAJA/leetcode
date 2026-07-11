// 给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
//
// 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。 
//
// 
//
// 示例 1： 
//
// 
// 输入：nums = [4,6,7,7]
// 输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
// 
//
// 示例 2： 
//
// 
// 输入：nums = [4,4,3,2,1]
// 输出：[[4,4]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 15 
// -100 <= nums[i] <= 100 
// 
//
// Related Topics 位运算 数组 哈希表 回溯 👍 905 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
// 回溯算法
func findSubsequences(nums []int) (ret [][]int) {
	var dfs func(path []int, idx int)
	dfs = func(path []int, idx int) {
		if idx == len(nums) {
			return
		}
		seen := make(map[int]bool)
		for i := idx; i < len(nums); i++ {
			if len(path) == 0 || path[len(path)-1] <= nums[i] {
				if seen[nums[i]] {
					continue
				}
				next := append(append(make([]int, 0, len(path)), path...), nums[i])
				if len(next) >= 2 {
					ret = append(ret, next)
				}
				seen[nums[i]] = true
				dfs(next, i+1)
			}
		}
	}
	dfs([]int{}, 0)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
