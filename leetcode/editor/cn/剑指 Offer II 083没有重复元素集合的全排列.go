// 给定一个不含重复数字的整数数组 nums ，返回其 所有可能的全排列 。可以 按任意顺序 返回答案。
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
//
//
// 注意：本题与主站 46 题相同：https://leetcode-cn.com/problems/permutations/
// Related Topics 数组 回溯 👍 20 👎 0

package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{0, 1}))
}

// leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) (ret [][]int) {
	const USED = -100
	result := make([]int, 0, len(nums))
	var dfs func()
	dfs = func() {
		if len(result) == len(nums) {
			ret = append(ret, append([]int{}, result...))
			return
		}
		for i, v := range nums {
			if v != USED {
				tmp := v
				result = append(result, v)
				nums[i] = USED
				dfs()
				nums[i] = tmp
				result = result[:len(result)-1]
			}
		}
	}
	dfs()
	return
}

// leetcode submit region end(Prohibit modification and deletion)
