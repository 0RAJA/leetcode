// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//
// 有效 二叉搜索树定义如下： 
//
// 
// 节点的左子树只包含 严格小于 当前节点的数。 
// 节点的右子树只包含 严格大于 当前节点的数。 
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
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 2771 👎 0
package main

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
// 判断核心原则：当前 root 大于左边的最大值 && 小于右边的最小值 并且 返回当前节点下的最小值和最大值 来用于上游判断
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(root *TreeNode) (int, int, bool)
	dfs = func(root *TreeNode) (int, int, bool) {
		if root.Left == nil && root.Right == nil {
			return root.Val, root.Val, true
		}
		minValue, maxValue := root.Val, root.Val
		if root.Left != nil {
			leftMin, leftMax, leftOk := dfs(root.Left)
			if !leftOk || leftMax >= root.Val {
				return 0, 0, false
			}
			minValue = leftMin
		}
		if root.Right != nil {
			rightMin, rightMax, rightOk := dfs(root.Right)
			if !rightOk || rightMin <= root.Val {
				return 0, 0, false
			}
			maxValue = rightMax
		}
		return minValue, maxValue, true
	}
	_, _, isOk := dfs(root)
	return isOk
}

// leetcode submit region end(Prohibit modification and deletion)
