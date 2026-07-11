// 给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
//
// 假设二叉树中至少有一个节点。 
//
// 
//
// 示例 1： 
//
// 
//
// 
// 输入: root = [2,1,3]
// 输出: 1
// 
//
// 示例 2： 
//
// 
//
// 
// 输入: [1,2,3,4,null,5,6,null,null,7]
// 输出: 7
// 
//
// 
//
// 提示： 
//
// 
// 二叉树的节点个数的范围是 [1,10⁴] 
// 
// -2³¹ <= Node.val <= 2³¹ - 1 
// 
//
// 
//
// 
// 注意：本题与主站 513 题相同： https://leetcode.cn/problems/find-bottom-left-tree-value/ 
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 58 👎 0
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
// 前序遍历，中节点判断当前深度，如果是世上最大则更新结果
func findBottomLeftValue(root *TreeNode) (ret int) {
	maxDepth := -1
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth > maxDepth {
			ret = root.Val
			maxDepth = depth
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
