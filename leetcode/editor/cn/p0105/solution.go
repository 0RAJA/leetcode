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
// Related Topics 树 数组 哈希表 分治 二叉树 👍 2823 👎 0

package p0105

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
/*
前序遍历： 确定当前子树的根节点是谁。
中序遍历： 确定左右子树的边界，以及各自有多少个节点。
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 中序值 -> 下标
	index := make(map[int]int)
	for i, v := range inorder {
		index[v] = i
	}

	// 所有区间采用左闭右开
	var build func(preL, preR, inL, inR int) *TreeNode

	build = func(preL, preR, inL, inR int) *TreeNode {
		if preL >= preR {
			return nil
		}

		// 1. 前序确定根节点
		rootVal := preorder[preL]
		root := &TreeNode{Val: rootVal}

		// 2. 中序确定左子树节点数量
		mid := index[rootVal]
		leftSize := mid - inL

		// 3. 用 leftSize 划分前序
		root.Left = build(
			preL+1,
			preL+1+leftSize,
			inL,
			mid,
		)

		root.Right = build(
			preL+1+leftSize,
			preR,
			mid+1,
			inR,
		)

		return root
	}

	n := len(preorder)
	return build(0, n, 0, n)
}

// leetcode submit region end(Prohibit modification and deletion)
