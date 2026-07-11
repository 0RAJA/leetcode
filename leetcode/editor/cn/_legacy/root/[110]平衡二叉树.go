// 给定一个二叉树，判断它是否是 平衡二叉树
//
// 
//
// 示例 1： 
// 
// 
// 输入：root = [3,9,20,null,null,15,7]
// 输出：true
// 
//
// 示例 2： 
// 
// 
// 输入：root = [1,2,2,3,3,null,null,4,4]
// 输出：false
// 
//
// 示例 3： 
//
// 
// 输入：root = []
// 输出：true
// 
//
// 
//
// 提示： 
//
// 
// 树中的节点数在范围 [0, 5000] 内 
// -10⁴ <= Node.val <= 10⁴ 
// 
//
// Related Topics 树 深度优先搜索 二叉树 👍 1713 👎 0
package main

import (
	"math"
)

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

// 其实就是分求每个节点左右子树的最大高度，然后算高度差的绝对值是否 > 1
func isBalanced(root *TreeNode) (ok bool) {
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) (maxHeight int) {
		if root == nil {
			return 0
		}
		leftHeight := dfs(root.Left)
		if leftHeight == -1 {
			return -1
		}
		rightHeight := dfs(root.Right)
		if rightHeight == -1 {
			return -1
		}
		if math.Abs(float64(leftHeight-rightHeight)) > 1 {
			return -1
		}
		return max(leftHeight, rightHeight) + 1
	}
	return dfs(root) != -1
}

// leetcode submit region end(Prohibit modification and deletion)
