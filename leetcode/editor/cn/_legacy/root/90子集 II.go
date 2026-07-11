//给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
//
//
// 示例 2：
//
//
//输入：nums = [0]
//输出：[[],[0]]
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
//
// Related Topics 位运算 数组 回溯 👍 709 👎 0

package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) (ret [][]int) {
	sort.Ints(nums)
	result := make([]int, 0, len(nums))
	var dfs func(isChoose bool, index int)
	dfs = func(isChoose bool, index int) {
		if index == len(nums) {
			ret = append(ret, append([]int{}, result...))
			return
		}
		dfs(false, index+1)
		if !isChoose && index > 0 && nums[index] == nums[index-1] {
			return
		}
		result = append(result, nums[index])
		dfs(true, index+1)
		result = result[:len(result)-1]
	}
	dfs(false, 0)
	return
}

//leetcode submit region end(Prohibit modification and deletion)
