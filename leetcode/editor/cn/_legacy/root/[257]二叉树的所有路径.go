// 给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
//
// 叶子节点 是指没有子节点的节点。 
//
// 示例 1： 
// 
// 
// 输入：root = [1,2,3,null,5]
// 输出：["1->2->5","1->3"]
// 
//
// 示例 2： 
//
// 
// 输入：root = [1]
// 输出：["1"]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点的数目在范围 [1, 100] 内 
// -100 <= Node.val <= 100 
// 
//
// Related Topics 树 深度优先搜索 字符串 回溯 二叉树 👍 1300 👎 0
package main

import (
	"strconv"
)

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
// dfs 前序遍历即可
func binaryTreePaths(root *TreeNode) (ret []string) {
	var dfs func(root *TreeNode, path string)
	dfs = func(root *TreeNode, path string) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			ret = append(ret, path+strconv.Itoa(root.Val))
			return
		}
		if root.Left != nil {
			dfs(root.Left, path+strconv.Itoa(root.Val)+"->")
		}
		if root.Right != nil {
			dfs(root.Right, path+strconv.Itoa(root.Val)+"->")
		}
	}
	dfs(root, "")
	return
}

// leetcode submit region end(Prohibit modification and deletion)
