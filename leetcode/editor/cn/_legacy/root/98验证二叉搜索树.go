// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//
// 有效 二叉搜索树定义如下：
//
//
// 节点的左子树只包含 小于 当前节点的数。
// 节点的右子树只包含 大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
//
//
//
// 示例 1：
//
//
// 输入：root = [2,1,3]
// 输出：true
//
//
// 示例 2：
//
//
// 输入：root = [5,1,4,null,null,3,6]
// 输出：false
// 解释：根节点的值是 5 ，但是右子节点的值是 4 。
//
//
//
//
// 提示：
//
//
// 树中节点数目范围在[1, 10⁴] 内
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 1808 👎 0

package main

import (
	"math"
)

func main() {

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
// 递归，每次告知下一层最大值和最小值，让其在每一层进行判断
func isValidBST(root *TreeNode) bool {
	var dfs func(root *TreeNode, low, up int) bool
	dfs = func(root *TreeNode, low, up int) bool {
		if root == nil {
			return true
		}
		if root.Val <= low || root.Val >= up {
			return false
		}
		return dfs(root.Left, low, root.Val) && dfs(root.Right, root.Val, up)
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}

// 中序遍历，维护最小值依次进行比较
func isValidBST2(root *TreeNode) bool {
	var stack []*TreeNode
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)
