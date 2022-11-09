// 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并
// 返回其根节点。
//
//
//
// 示例 1:
//
//
// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]
//
//
// 示例 2:
//
//
// 输入: preorder = [-1], inorder = [-1]
// 输出: [-1]
//
//
//
//
// 提示:
//
//
// 1 <= preorder.length <= 3000
// inorder.length == preorder.length
// -3000 <= preorder[i], inorder[i] <= 3000
// preorder 和 inorder 均 无重复 元素
// inorder 均出现在 preorder
// preorder 保证 为二叉树的前序遍历序列
// inorder 保证 为二叉树的中序遍历序列
//
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 1788 👎 0

package main

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

func buildTree(preorder []int, inorder []int) *TreeNode {
	var dfs func(preorder, inorder []int) *TreeNode
	dfs = func(preorder, inorder []int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		}
		k := preorder[0]
		node := &TreeNode{Val: k}
		idx := 0 // 找到分界点
		for inorder[idx] != k {
			idx++
		}
		if idx != 0 {
			node.Left = dfs(preorder[1:idx+1], inorder[:idx])
		}
		if idx != len(inorder)-1 {
			node.Right = dfs(preorder[idx+1:], inorder[idx+1:])
		}
		return node
	}
	return dfs(preorder, inorder)
}

// leetcode submit region end(Prohibit modification and deletion)
