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
// Related Topics 树 数组 哈希表 分治 二叉树 👍 2752 👎 0
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
// 根据前序遍历（根左右）和中序遍历（左根右），每次获取 根节点 拆分 左右节点构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderIdx := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inorderIdx[v] = i
	}
	var build func(leftIdx, rightIdx int) *TreeNode
	build = func(leftIdx, rightIdx int) *TreeNode {
		if leftIdx > rightIdx {
			return nil
		}
		rootValue := preorder[0]
		preorder = preorder[1:]
		rootIdx := inorderIdx[rootValue]
		return &TreeNode{
			Val:   rootValue,
			Left:  build(leftIdx, rootIdx-1),
			Right: build(rootIdx+1, rightIdx),
		}
	}
	return build(0, len(inorder)-1)
}

// leetcode submit region end(Prohibit modification and deletion)
