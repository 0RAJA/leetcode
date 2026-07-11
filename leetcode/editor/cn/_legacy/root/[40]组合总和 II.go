// 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用 一次 。 
//
// 注意：解集不能包含重复的组合。 
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
// Related Topics 数组 回溯 👍 1751 👎 0
package main

import (
	"sort"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 注意几个点：
// 1. 组合（枚举）；
// 2. 每个数字只能出现一次（每层需要用下一个数，不能复用）；
// 3. 结果的组合不能重复（需要考虑下 重复的原因 是 当前值可能匹配到了，下一个值可能和当前值相同，也会被匹配到导致重复，只需要每层相同的值枚举使用一次即可）
// 满足三者条件即可
func combinationSum2(candidates []int, target int) (ret [][]int) {
	var dfs func(path []int, idx, sum int)
	dfs = func(path []int, idx, sum int) {
		if sum >= target {
			if sum == target {
				ret = append(ret, path)
			}
			return
		}
		for i := idx; i < len(candidates); i++ {
			// 每层直接跳过相同的人来避免结果重复
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			v := candidates[i]
			// 剪枝
			if sum+v > target {
				break
			}
			dfs(append(append(make([]int, 0, len(path)+1), path...), v), i+1, sum+v)
		}
	}
	// 从小到大 便于跳过相同值 && 剪枝
	sort.Ints(candidates)
	dfs([]int{}, 0, 0)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
