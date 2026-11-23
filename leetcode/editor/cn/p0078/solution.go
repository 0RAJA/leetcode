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
// Related Topics 位运算 数组 回溯 👍 2704 👎 0

package p0078

// leetcode submit region begin(Prohibit modification and deletion)
// 这里其实就是去做回溯，遍历出所有可能性。但是注意两点：
// 1. 子集的顺序是无关的：1和2、2和1其实是一个答案。因此，这里每次每层往下搜索的时候，不能回到之前的，必须要从下一个元素开始
// 2. 任何一个情况都要记录到结果中。
func subsets(nums []int) (res [][]int) {
	res = make([][]int, 0)
	var path []int
	var dfs func(idx int)
	dfs = func(idx int) {
		// 任意一个情况都要算到结果中
		res = append(res, append(make([]int, 0, len(path)), path...))
		// 针对每一个当前的 number，去往下一层进行深搜，然后注意搜索完一层之后，要把当前这个元素去掉
		// 而且由于是集合子集，顺序是无关的，所以我们只能往后搜，不能回到前面的元素去使用
		for i := idx; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
