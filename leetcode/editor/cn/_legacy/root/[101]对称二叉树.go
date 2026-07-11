// 给你一个二叉树的根节点 root ， 检查它是否轴对称。
//
// 
//
// 示例 1： 
// 
// 
// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
// 
//
// 示例 2： 
// 
// 
// 输入：root = [1,2,2,null,3,null,3]
// 输出：false
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目在范围 [1, 1000] 内 
// -100 <= Node.val <= 100 
// 
//
// 
//
// 进阶：你可以运用递归和迭代两种方法解决这个问题吗？ 
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 3139 👎 0
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetricHelper(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	} else if node1 != nil && node2 != nil {
		return node1.Val == node2.Val && isSymmetricHelper(node1.Left, node2.Right) && isSymmetricHelper(node1.Right, node2.Left)
	} else {
		return false
	}
}

// 拿左节点的左和右节点的右比较，拿左节点的右和右节点的左比较
func isSymmetric01(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricHelper(root.Left, root.Right)
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

// 层序遍历法：横向遍历，每次拿两个节点比较，然后把左节点的右孩子、右节点的做孩子、左节点的左孩子、右节点的右孩子放入进行比较
func isSymmetric(root *TreeNode) bool {
	var q []*TreeNode
	if root != nil {
		q = append(q, root.Left, root.Right)
	}
	for len(q) > 0 {
		node0 := q[0]
		node1 := q[1]
		q = q[2:]
		if node0 == nil && node1 == nil {
			continue
		}
		if node0 != nil && node1 != nil {
			if node0.Val != node1.Val {
				return false
			}
			q = append(q, node0.Left, node1.Right, node0.Right, node1.Left)
		} else {
			return false
		}
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)
