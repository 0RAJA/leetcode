// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的 子集（幂集）。
//
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。 
//
// 
// 
// 
// 
// 
//
// 示例 1： 
//
// 
// 输入：nums = [1,2,2]
// 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
// 
//
// 示例 2： 
//
// 
// 输入：nums = [0]
// 输出：[[],[0]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10 
// -10 <= nums[i] <= 10 
// 
//
// Related Topics 位运算 数组 回溯 👍 1361 👎 0
package main

import (
	"sort"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 和子集的处理很像，但是关键点是避免重复的子集，先想下为啥会重复 1222 按之前的解法，会出现两个 (2,2)，其实就是在于本层相同的处理的值处理一次需要跳过，那么先排序
func subsetsWithDup(nums []int) (res [][]int) {
	sort.Ints(nums)
	var dfs func(path []int, idx int)
	dfs = func(path []int, idx int) {
		if idx == len(nums) {
			return
		}
		for i := idx; i < len(nums); i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			next := append(append(make([]int, 0, len(path)+1), path...), nums[i])
			res = append(res, next)
			dfs(next, i+1)
		}
	}
	res = append(res, []int{})
	dfs([]int{}, 0)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
