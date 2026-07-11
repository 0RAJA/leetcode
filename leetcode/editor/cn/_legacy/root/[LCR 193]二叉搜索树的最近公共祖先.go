// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（
// 一个节点也可以是它自己的祖先）。”
//
// 例如，给定如下二叉搜索树: root = [6,2,8,0,4,7,9,null,null,3,5] 
//
// 
//
// 
//
// 示例 1： 
//
// 
// 输入：root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
// 输出：6
// 解释：节点 2 和节点 8 的最近公共祖先是 6。
// 
//
// 示例 2： 
//
// 
// 输入：root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
// 输出：2
// 解释：节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
//
// 
//
// 说明： 
//
// 
// 所有节点的值都是唯一的。 
// p、q 为不同节点且均存在于给定的二叉搜索树中。 
// 
//
// 注意：本题与主站 235 题相同：https://leetcode.cn/problems/lowest-common-ancestor-of-a-
// binary-search-tree/
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 353 👎 0
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
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

// 前序遍历：若当前节点值落在 p、q 之间或等于其中一个，它就是最近公共祖先；否则 p、q 都在同一侧，按大小递归去左/右子树继续找。
func lowestCommonAncestor(root, p, q *TreeNode) (ret *TreeNode) {
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if root.Val >= p.Val && root.Val <= q.Val || root.Val <= p.Val && root.Val >= q.Val {
			return root
		}
		if root.Val < p.Val {
			return dfs(root.Right)
		} else {
			return dfs(root.Left)
		}
	}
	return dfs(root)
}

// leetcode submit region end(Prohibit modification and deletion)
