// 给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，找出 candidates 中所有可以使数字和为目标数 target 的
// 唯一组合。
//
// candidates 中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是不同的。
//
// 对于给定的输入，保证和为 target 的唯一组合数少于 150 个。
//
//
//
// 示例 1：
//
//
// 输入: candidates = [2,3,6,7], target = 7
// 输出: [[7],[2,2,3]]
//
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
// 示例 4：
//
//
// 输入: candidates = [1], target = 1
// 输出: [[1]]
//
//
// 示例 5：
//
//
// 输入: candidates = [1], target = 2
// 输出: [[1,1]]
//
//
//
//
// 提示：
//
//
// 1 <= candidates.length <= 30
// 1 <= candidates[i] <= 200
// candidate 中的每个元素都是独一无二的。
// 1 <= target <= 500
//
//
//
//
// 注意：本题与主站 39 题相同： https://leetcode-cn.com/problems/combination-sum/
// Related Topics 数组 回溯 👍 23 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) (ret [][]int) {
	sort.Ints(candidates)
	result := make([]int, 0, len(candidates))
	var dfs func(idx, v int)
	dfs = func(idx, v int) {
		if v == 0 { // 得到目标值
			ret = append(ret, append([]int{}, result...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if v-candidates[i] >= 0 { // 尝试选择一次
				result = append(result, candidates[i])
				dfs(i, v-candidates[i])
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
