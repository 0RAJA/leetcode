// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
//
// 
//
// 示例 1： 
// 
// 
// 输入：root = [1,null,2,3]
// 输出：[1,3,2]
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
// 输入：root = [1]
// 输出：[1]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目在范围 [0, 100] 内 
// -100 <= Node.val <= 100 
// 
//
// 
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？ 
//
// Related Topics 栈 树 深度优先搜索 二叉树 👍 2424 👎 0

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归实现中序遍历
func inorderTraversalRecursion(root *TreeNode) (ret []int) {
	if root == nil {
		return
	}
	ret = append(ret, inorderTraversalRecursion(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, inorderTraversalRecursion(root.Right)...)
	return
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
// 栈实现中序遍历: cur 始终指向待入栈的元素，cur 一直往左走，有元素就入栈，直到nil，pop 回到 mid 节点，输出值，然后 cur 到右节点 循环
func inorderTraversal(root *TreeNode) (ret []int) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	cur := root.Left
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			MidNode := stack[len(stack)-1]
			ret = append(ret, MidNode.Val)
			stack = stack[:len(stack)-1]
			cur = MidNode.Right
		}
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
