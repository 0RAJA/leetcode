// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
//
//
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。
//
//
//
//
// 示例 1：
//
//
// 输入：root = [1,2,5,3,4,null,6]
// 输出：[1,null,2,null,3,null,4,null,5,null,6]
//
//
// 示例 2：
//
//
// 输入：root = []
// 输出：[]
//
//
// 示例 3：
//
//
// 输入：root = [0]
// 输出：[0]
//
//
//
//
// 提示：
//
//
// 树中结点数在范围 [0, 2000] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？
//
// Related Topics 栈 树 深度优先搜索 链表 二叉树 👍 2063 👎 0

package p0114

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
// 前序遍历进行展开，并返回当前节点前序遍历后的最后一个节点；也就是先把左节点接到右节点，再把当前节点右节点换成左节点
// if theLastOneNode(root.Left):
//  theLastOneNode(root.Left).Right = root.Right
//  root.Right = root.Left
// root.Left = nil
// return theLastOneNode(root.Right) or theLastOneNode(root.Left) or root
func flatten(root *TreeNode) {
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		leftReturnNode := dfs(root.Left)
		rightReturnNode := dfs(root.Right)
		if leftReturnNode != nil {
			leftReturnNode.Right = root.Right
			root.Right = root.Left
		}
		root.Left = nil
		if rightReturnNode != nil {
			return rightReturnNode
		}
		if leftReturnNode != nil {
			return leftReturnNode
		}
		return root
	}
	root = dfs(root)
}

// leetcode submit region end(Prohibit modification and deletion)
