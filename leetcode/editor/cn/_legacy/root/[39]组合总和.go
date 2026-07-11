// 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的
// 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。 
//
// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。 
//
// 对于给定的输入，保证和为 target 的不同组合数少于 150 个。 
//
// 
//
// 示例 1： 
//
// 
// 输入：candidates = [2,3,6,7], target = 7
// 输出：[[2,2,3],[7]]
// 解释：
// 2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
// 7 也是一个候选， 7 = 7 。
// 仅有这两种组合。
//
// 示例 2： 
//
// 
// 输入: candidates = [2,3,5], target = 8
// 输出: [[2,2,2,2],[2,3,3],[3,5]]
//
// 示例 3： 
//
// 
// 输入: candidates = [2], target = 1
// 输出: []
// 
//
// 
//
// 提示： 
//
// 
// 1 <= candidates.length <= 30 
// 2 <= candidates[i] <= 40 
// candidates 的所有元素 互不相同 
// 1 <= target <= 40 
// 
//
// Related Topics 数组 回溯 👍 3212 👎 0
package main

import (
	"sort"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 思路：递归穷举法来无限组合，注意剪枝，先排序从小到大，然后每轮从头开始遍历组合，如果当前值 == target 则选取，否则如果小则继续，大则 return;
// 还有个点需要注意，不能有重复的组合，那每次把当前值遍历完，然后从下一个开始枚举组合，避免出现 223 232 这种情况，直接 223 233 ...
func combinationSum(candidates []int, target int) (ret [][]int) {
	var dfs func(path []int, idx, sum int)
	dfs = func(path []int, idx, sum int) {
		if sum >= target {
			if sum == target {
				ret = append(ret, path)
			}
			return
		}
		for i := idx; i < len(candidates); i++ {
			v := candidates[i]
			if v+sum > target {
				break
			}
			dfs(append(append(make([]int, 0, len(path)+1), path...), v), i, sum+v)
		}
	}
	sort.Ints(candidates)
	dfs([]int{}, 0, 0)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
