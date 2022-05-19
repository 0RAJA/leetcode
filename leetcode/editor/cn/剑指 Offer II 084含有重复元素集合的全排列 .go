// 给定一个可包含重复数字的整数集合 nums ，按任意顺序 返回它所有不重复的全排列。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,1,2]
// 输出：
// [[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
//
//
//
// 注意：本题与主站 47 题相同： https://leetcode-cn.com/problems/permutations-ii/
// Related Topics 数组 回溯 👍 20 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 2, 3}))
}

// leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) (ret [][]int) {
	sort.Ints(nums)
	result := make([]int, 0, len(nums))
	isOK := make([]bool, len(nums))
	var dfs func()
	dfs = func() {
		if len(result) == len(nums) {
			ret = append(ret, append([]int{}, result...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if isOK[i] || i > 0 && nums[i] == nums[i-1] && !isOK[i-1] {
				continue
			}
			v := nums[i]
			result = append(result, v)
			isOK[i] = true
			dfs()
			isOK[i] = false
			result = result[:len(result)-1]
		}
	}
	dfs()
	return
}

// leetcode submit region end(Prohibit modification and deletion)
