//给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//
//
//
// 示例 1:
//
//
//
//
//输入: [1,2,3,null,5,null,4]
//输出: [1,3,4]
//
//
// 示例 2:
//
//
//输入: [1,null,3]
//输出: [1,3]
//
//
// 示例 3:
//
//
//输入: []
//输出: []
//
//
//
//
// 提示:
//
//
// 二叉树的节点个数的范围是 [0,100]
// -100 <= Node.val <= 100
//
//
//
//
// 注意：本题与主站 199 题相同：https://leetcode-cn.com/problems/binary-tree-right-side-
//view/
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 25 👎 0

package main

import "container/list"

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
func rightSideView(root *TreeNode) (ret []int) {
	aQueue := list.New()
	bQueue := list.New()
	aQueue.PushBack(root)
	if root == nil {
		return
	}
	for aQueue.Len() > 0 || bQueue.Len() > 0 {
		var t int
		for aQueue.Len() > 0 {
			p := aQueue.Front().Value.(*TreeNode)
			t = p.Val
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
		ret = append(ret, t)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
