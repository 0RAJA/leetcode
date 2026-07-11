// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
// 你可以按 任何顺序 返回答案。 
//
// 
//
// 示例 1： 
//
// 
// 输入：n = 4, k = 2
// 输出：
// [
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
// ]
//
// 示例 2： 
//
// 
// 输入：n = 1, k = 1
// 输出：[[1]]
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 20 
// 1 <= k <= n 
// 
//
// Related Topics 回溯 👍 1865 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
// 组合：递归遍历，选择一个数，递归下一层，直到选够，实际上是用递归层级来动态模拟 for 循环
// 剪枝：判断我还剩几个需要找，然后拿当前 N - 剩余个数 == 我前面可以枚举的个数，超过的及时遍历也无用（因为组合会被其他遍历到）
func combine(n int, k int) (ret [][]int) {
	var dfs func(path []int, start int)
	dfs = func(path []int, start int) {
		if len(path) == k {
			ret = append(ret, path)
			return
		}
		// 剪枝：判断我还剩几个需要找，然后拿当前 N - 剩余个数 == 我前面可以枚举的个数，超过的及时遍历也无用（因为组合会被其他遍历到）
		for i := start; i <= n-(k-len(path))+1; i++ {
			next := make([]int, len(path), len(path)+1)
			copy(next, path)
			dfs(append(next, i), i+1)
		}
	}
	dfs([]int{}, 1)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
