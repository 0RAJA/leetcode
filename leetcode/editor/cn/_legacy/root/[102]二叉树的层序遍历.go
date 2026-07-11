// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
// 
//
// 示例 1： 
// 
// 
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]
// 
//
// 示例 2： 
//
// 
// 输入：root = [1]
// 输出：[[1]]
// 
//
// 示例 3： 
//
// 
// 输入：root = []
// 输出：[]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目在范围 [0, 2000] 内 
// -1000 <= Node.val <= 1000 
// 
//
// Related Topics 树 广度优先搜索 二叉树 👍 2294 👎 0
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

type QueueData struct {
	node  *TreeNode // 节点
	depth int       // 深度
}

func levelOrder(root *TreeNode) (ret [][]int) {
	if root == nil {
		return
	}
	q := []QueueData{
		{
			node:  root,
			depth: 0,
		},
	}
	for len(q) > 0 {
		leftData := q[0]
		q = q[1:]
		if len(ret) <= leftData.depth {
			ret = append(ret, []int{})
		}
		ret[leftData.depth] = append(ret[leftData.depth], leftData.node.Val)
		if leftData.node.Left != nil {
			q = append(q, QueueData{
				node:  leftData.node.Left,
				depth: leftData.depth + 1,
			})
		}
		if leftData.node.Right != nil {
			q = append(q, QueueData{
				node:  leftData.node.Right,
				depth: leftData.depth + 1,
			})
		}
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
