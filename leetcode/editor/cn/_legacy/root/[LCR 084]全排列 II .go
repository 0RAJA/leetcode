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
// 
// 注意：本题与主站 47 题相同： https://leetcode.cn/problems/permutations-ii/ 
//
// Related Topics 数组 回溯 👍 76 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
// 和常规排列类似，需要从 0 开始，全部排列中每个值用一次（避免用超），每层相同的值只使用一次（避免重复排列）
// 外层 seen 保障每个相同的 num 不会被用超；内层 seen 保障当前层中 每个相同的值只会被排列一次
func permuteUnique(nums []int) (ret [][]int) {
	// 外层 seen 保障每个相同的 num 不会被用超
	seenGlobal := make(map[int]int)
	for _, num := range nums {
		seenGlobal[num]++
	}
	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			data := make([]int, len(path))
			copy(data, path)
			ret = append(ret, data)
			return
		}
		// 内层 seen 保障当前层中 每个相同的值只会被排列一次
		seen := make(map[int]bool, len(nums))
		for i := 0; i < len(nums); i++ {
			if seen[nums[i]] || seenGlobal[nums[i]] == 0 {
				continue
			}
			seen[nums[i]] = true
			seenGlobal[nums[i]]--
			path = append(path, nums[i])
			dfs(path)
			path = path[:len(path)-1]
			seenGlobal[nums[i]]++
		}
	}
	dfs([]int{})
	return
}

// leetcode submit region end(Prohibit modification and deletion)
