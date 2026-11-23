// 给定一个二叉树 root ，返回其最大深度。
//
// 二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。
//
//
//
// 示例 1：
//
//
//
//
//
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：3
//
//
// 示例 2：
//
//
// 输入：root = [1,null,2]
// 输出：2
//
//
//
//
// 提示：
//
//
// 树中节点的数量在 [0, 10⁴] 区间内。
// -100 <= Node.val <= 100
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 2169 👎 0

package p0104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
func maxDepth(root *TreeNode) (ret int) {
	var dfs func(root *TreeNode, height int)
	dfs = func(root *TreeNode, height int) {
		if root == nil {
			return
		}
		ret = max(ret, height)    // 中，记录最大高度
		dfs(root.Left, height+1)  // 左，递归左子树，高度+1
		dfs(root.Right, height+1) // 右，递归右子树，高度+1
	}
	dfs(root, 1)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
