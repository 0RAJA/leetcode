// 给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并
// 返回这颗 二叉树 。
//
// 
//
// 示例 1: 
// 
// 
// 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// 输出：[3,9,20,null,null,15,7]
// 
//
// 示例 2: 
//
// 
// 输入：inorder = [-1], postorder = [-1]
// 输出：[-1]
// 
//
// 
//
// 提示: 
//
// 
// 1 <= inorder.length <= 3000 
// postorder.length == inorder.length 
// -3000 <= inorder[i], postorder[i] <= 3000 
// inorder 和 postorder 都由 不同 的值组成 
// postorder 中每一个值都在 inorder 中 
// inorder 保证是树的中序遍历 
// postorder 保证是树的后序遍历 
// 
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 1406 👎 0
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
// 利用中序（左根右）和后序（左右根）的特性；每次 pop 后序数组根节点，然后拆分中序左右节点来递归构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 构造 v -> idx 的索引，用来快速根据后序找到的 root 找对应的 idx
	inorderPos := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inorderPos[v] = i
	}
	var build func(leftIdx, rightIdx int) *TreeNode
	build = func(leftIdx, rightIdx int) *TreeNode {
		// 结束条件，当前分区没有元素则表明是一个 nil 节点
		if leftIdx > rightIdx {
			return nil
		}
		// 弹出 root 节点
		rootValue := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{Val: rootValue}
		rootIdx := inorderPos[rootValue]
		// 根据 root 节点拆分左右子树 ，注意 后序遍历 是 左右根，反过来就是 根右左
		root.Right = build(rootIdx+1, rightIdx)
		root.Left = build(leftIdx, rootIdx-1)
		return root
	}
	return build(0, len(inorder)-1)
}

// leetcode submit region end(Prohibit modification and deletion)
