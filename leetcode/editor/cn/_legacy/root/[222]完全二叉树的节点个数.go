// 给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
//
// 完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层
// 为第 h 层（从第 0 层开始），则该层包含 1~ 2ʰ 个节点。
//
// 
//
// 示例 1： 
// 
// 
// 输入：root = [1,2,3,4,5,6]
// 输出：6
// 
//
// 示例 2： 
//
// 
// 输入：root = []
// 输出：0
// 
//
// 示例 3： 
//
// 
// 输入：root = [1]
// 输出：1
// 
//
// 
//
// 提示： 
//
// 
// 树中节点的数目范围是[0, 5 * 10⁴] 
// 0 <= Node.val <= 5 * 10⁴ 
// 题目数据保证输入的树是 完全二叉树 
// 
//
// 
//
// 进阶：遍历树来统计节点是一种时间复杂度为 O(n) 的简单解决方案。你可以设计一个更快的算法吗？ 
//
// Related Topics 位运算 树 二分查找 二叉树 👍 1314 👎 0
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
// 通过完全二叉树的定义来计算：一个完全二叉树可以拆解为 N 个满二叉树(2**N-1个节点)+最后叶子节点没有满(拆解为满二叉树+当前节点)；如果是满二叉树（左右高度相同）则返回2**N-1个节点；如果不是，则返回递归统计左右节点数之和+1
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth, rightDepth := 1, 1
	for leftNode := root.Left; leftNode != nil; leftDepth++ {
		leftNode = leftNode.Left
	}
	for rightNode := root.Right; rightNode != nil; rightDepth++ {
		rightNode = rightNode.Right
	}
	if leftDepth == rightDepth {
		return (1 << leftDepth) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

// leetcode submit region end(Prohibit modification and deletion)
