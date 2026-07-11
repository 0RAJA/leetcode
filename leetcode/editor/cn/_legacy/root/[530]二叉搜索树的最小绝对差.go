// 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
//
// 差值是一个正数，其数值等于两值之差的绝对值。 
//
// 
//
// 示例 1： 
// 
// 
// 输入：root = [4,2,6,1,3]
// 输出：1
// 
//
// 示例 2： 
// 
// 
// 输入：root = [1,0,48,null,null,12,49]
// 输出：1
// 
//
// 
//
// 提示： 
//
// 
// 树中节点的数目范围是 [2, 10⁴] 
// 0 <= Node.val <= 10⁵ 
// 
//
// 
//
// 注意：本题与 783 https://leetcode.cn/problems/minimum-distance-between-bst-nodes/ 相
// 同
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉搜索树 二叉树 👍 664 👎 0
package main

import (
	"math"
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
// 计算 root 和 左边最大值 和 右边最小值 差值，然后返回上游 最大值和最小值
func getMinimumDifference(root *TreeNode) (ret int) {
	ret = math.MaxInt
	var dfs func(root *TreeNode) (min, max int)
	dfs = func(root *TreeNode) (min, max int) {
		min, max = root.Val, root.Val
		if root.Left == nil && root.Right == nil {
			return
		}
		if root.Left != nil {
			leftMin, leftMax := dfs(root.Left)
			ret = int(math.Min(float64(ret), float64(root.Val-leftMax)))
			min = leftMin
		}
		if root.Right != nil {
			rightMin, rightMax := dfs(root.Right)
			ret = int(math.Min(float64(ret), float64(rightMin-root.Val)))
			max = rightMax
		}
		return
	}
	dfs(root)
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
