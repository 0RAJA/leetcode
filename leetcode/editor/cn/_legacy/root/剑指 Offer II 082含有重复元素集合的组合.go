// 给定一个可能有重复数字的整数数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合
// 。
//
// candidates 中的每个数字在每个组合中只能使用一次，解集不能包含重复的组合。
//
//
//
// 示例 1:
//
//
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 输出:
// [
// [1,1,6],
// [1,2,5],
// [1,7],
// [2,6]
// ]
//
// 示例 2:
//
//
// 输入: candidates = [2,5,2,1,2], target = 5,
// 输出:
// [
// [1,2,2],
// [5]
// ]
//
//
//
// 提示:
//
//
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30
//
//
//
//
// 注意：本题与主站 40 题相同： https://leetcode-cn.com/problems/combination-sum-ii/
// Related Topics 数组 回溯 👍 21 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) (ret [][]int) {
	sort.Ints(candidates)
	result := make([]int, 0, len(candidates))
	var dfs func(idx, target int)
	dfs = func(idx, target int) {
		if target == 0 {
			ret = append(ret, append([]int{}, result...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] { // 保证一个数只用一次
				continue
			}
			if target-candidates[i] >= 0 {
				result = append(result, candidates[i])
				dfs(i+1, target-candidates[i])
				result = result[:len(result)-1]
			} else {
				break
			}
		}
	}
	dfs(0, target)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
