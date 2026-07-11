// 给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
//
// 
//
// 示例 1： 
//
// 
// 输入：root = [1,null,2,3] 
// 
//
// 输出：[3,2,1] 
//
// 解释： 
//
// 
//
// 示例 2： 
//
// 
// 输入：root = [1,2,3,4,5,null,8,null,null,6,7,9] 
// 
//
// 输出：[4,6,7,5,2,9,8,3,1] 
//
// 解释： 
//
// 
//
// 示例 3： 
//
// 
// 输入：root = [] 
// 
//
// 输出：[] 
//
// 示例 4： 
//
// 
// 输入：root = [1] 
// 
//
// 输出：[1] 
//
// 
//
// 提示： 
//
// 
// 树中节点的数目在范围 [0, 100] 内 
// -100 <= Node.val <= 100 
// 
//
// 
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？ 
//
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1277 👎 0
package main

import (
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归实现后序遍历
func postorderTraversalRecursion(root *TreeNode) (ret []int) {
	if root == nil {
		return
	}
	ret = append(ret, postorderTraversalRecursion(root.Left)...)
	ret = append(ret, postorderTraversalRecursion(root.Right)...)
	ret = append(ret, root.Val)
	return
}

// 栈实现后序遍历 -- 方案1 获取栈顶元素，如果没有孩子就说明是 mid 元素，弹出并返回值；有右孩子则入右孩子，有左孩子则入左孩子（为了避免重复，需要入孩子后把对应指向设为 nil）
func postorderTraversal01(root *TreeNode) (ret []int) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		if cur.Right == nil && cur.Left == nil {
			ret = append(ret, cur.Val)
			stack = stack[:len(stack)-1]
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
			cur.Right = nil
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
			cur.Left = nil
		}
	}
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

// 栈实现后序遍历 -- 方案2: 中左右 -> 中右左 -> 结果反转 -> 左右中
func postorderTraversal(root *TreeNode) (ret []int) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		ret = append(ret, cur.Val)
		stack = stack[:len(stack)-1]
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}
	slices.Reverse(ret)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
