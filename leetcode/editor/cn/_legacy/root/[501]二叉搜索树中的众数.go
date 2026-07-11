// 给你一个含重复值的二叉搜索树（BST）的根节点 root ，找出并返回 BST 中的所有 众数（即，出现频率最高的元素）。
//
// 如果树中有不止一个众数，可以按 任意顺序 返回。 
//
// 假定 BST 满足如下定义： 
//
// 
// 结点左子树中所含节点的值 小于等于 当前节点的值 
// 结点右子树中所含节点的值 大于等于 当前节点的值 
// 左子树和右子树都是二叉搜索树 
// 
//
// 
//
// 示例 1： 
// 
// 
// 输入：root = [1,null,2,2]
// 输出：[2]
// 
//
// 示例 2： 
//
// 
// 输入：root = [0]
// 输出：[0]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点的数目在范围 [1, 10⁴] 内 
// -10⁵ <= Node.val <= 10⁵ 
// 
//
// 
//
// 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内） 
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 834 👎 0
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
// 中序遍历 元素值由小到大，此时维护 最大次数、当前值、当前值对应次数 以此来维护 众数列表即可
func findMode(root *TreeNode) (ret []int) {
	maxTimes := 0
	currentValue := math.MaxInt
	currentTimes := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 左
		dfs(root.Left)
		// 中
		if currentValue != root.Val {
			currentTimes = 1
		} else {
			currentTimes++
		}
		currentValue = root.Val
		if currentTimes == maxTimes {
			ret = append(ret, root.Val)
		} else if currentTimes > maxTimes {
			maxTimes = currentTimes
			ret = []int{root.Val}
		}
		// 右
		dfs(root.Right)
	}
	dfs(root)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
