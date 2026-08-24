// 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
//
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//
//
//
// 示例 1：
//
//
//
//
// 输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
// 输出：3
// 解释：和等于 8 的路径有 3 条，如图所示。
//
//
// 示例 2：
//
//
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出：3
//
//
//
//
// 提示:
//
//
// 二叉树的节点个数的范围是 [0,1000]
// -10⁹ <= Node.val <= 10⁹
// -1000 <= targetSum <= 1000
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 2347 👎 0

package p0437

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 以 root 为根，和为 target 的总数
func _dfs(root *TreeNode, target int64) (res int) {
	if root == nil {
		return 0
	}
	if int64(root.Val) == target {
		res++
	}
	res += _dfs(root.Left, target-int64(root.Val))
	res += _dfs(root.Right, target-int64(root.Val))
	return res
}

// 两层深搜：一层搜所有节点，一层以当前节点为根搜路径数
func _pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return _dfs(root, int64(targetSum)) + _pathSum(root.Left, targetSum) + _pathSum(root.Right, targetSum)
}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 使用前缀和 prefixMap 记录每个前缀区间数，求次数时用 res += prefixMap[sum-targetSum]
func pathSum(root *TreeNode, targetSum int) (res int) {
	prefixMap := make(map[int]int)
	prefixMap[0] = 1 // 表示刚好 sum == targetSum 的次数初始 1
	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum += root.Val
		res += prefixMap[sum-targetSum]
		prefixMap[sum]++
		dfs(root.Left, sum)
		dfs(root.Right, sum)
		prefixMap[sum]--
	}
	dfs(root, 0)
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
