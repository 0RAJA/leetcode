//给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
//
// 假设二叉树中至少有一个节点。
//
//
//
// 示例 1:
//
//
//
//
//输入: root = [2,1,3]
//输出: 1
//
//
// 示例 2:
//
//
//
//
//输入: [1,2,3,4,null,5,6,null,null,7]
//输出: 7
//
//
//
//
// 提示:
//
//
// 二叉树的节点个数的范围是 [1,10⁴]
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
//
//
// 注意：本题与主站 513 题相同： https://leetcode-cn.com/problems/find-bottom-left-tree-
//value/
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 20 👎 0

package main

import (
	"container/list"
)

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) (ret int) {
	aQueue := list.New()
	bQueue := list.New()
	aQueue.PushBack(root)
	for aQueue.Len() > 0 || bQueue.Len() > 0 {
		ret = aQueue.Front().Value.(*TreeNode).Val
		for aQueue.Len() > 0 {
			p := aQueue.Front().Value.(*TreeNode)
			aQueue.Remove(aQueue.Front())
			if p.Left != nil {
				bQueue.PushBack(p.Left)
			}
			if p.Right != nil {
				bQueue.PushBack(p.Right)
			}
		}
		aQueue.PushBackList(bQueue)
		bQueue = list.New()
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
