// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。 
//
// 
//
// 示例 1： 
//
// 
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
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
// nums 中的所有元素 互不相同 
// 
//
// Related Topics 位运算 数组 回溯 👍 2650 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
// 组合问题，直接枚举，只不过 子集 这个比较特殊，就是任何元素的组合都是属于子集的范畴，所以只要能组合出来就加到结果里
func subsets(nums []int) (ret [][]int) {
	ret = append(ret, []int{})
	var dfs func(path []int, idx int)
	dfs = func(path []int, idx int) {
		if idx == len(nums) {
			return
		}
		for i := idx; i < len(nums); i++ {
			next := append(append(make([]int, 0, len(path)+1), path...), nums[i])
			ret = append(ret, next)
			dfs(next, i+1)
		}
	}
	dfs([]int{}, 0)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
