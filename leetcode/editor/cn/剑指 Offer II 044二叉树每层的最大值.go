//给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
//
//
//
// 示例1：
//
//
//输入: root = [1,3,2,5,3,null,9]
//输出: [1,3,9]
//解释:
//          1
//         / \
//        3   2
//       / \   \
//      5   3   9
//
//
// 示例2：
//
//
//输入: root = [1,2,3]
//输出: [1,3]
//解释:
//          1
//         / \
//        2   3
//
//
// 示例3：
//
//
//输入: root = [1]
//输出: [1]
//
//
// 示例4：
//
//
//输入: root = [1,null,2]
//输出: [1,2]
//解释:
//           1
//            \
//             2
//
//
// 示例5：
//
//
//输入: root = []
//输出: []
//
//
//
//
// 提示：
//
//
// 二叉树的节点个数的范围是 [0,10⁴]
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
//
//
// 注意：本题与主站 515 题相同： https://leetcode-cn.com/problems/find-largest-value-in-
//each-tree-row/
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 21 👎 0

package main

import (
	"container/list"
	"math"
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
func largestValues(root *TreeNode) (ret []int) {
	aQueue := list.New()
	bQueue := list.New()
	if root == nil {
		return
	}
	aQueue.PushBack(root)
	for aQueue.Len() > 0 || bQueue.Len() > 0 {
		max := math.MinInt32
		for aQueue.Len() > 0 {
			p := aQueue.Front().Value.(*TreeNode)
			aQueue.Remove(aQueue.Front())
			if p.Val > max {
				max = p.Val
			}
			if p.Left != nil {
				bQueue.PushBack(p.Left)
			}
			if p.Right != nil {
				bQueue.PushBack(p.Right)
			}
		}
		ret = append(ret, max)
		aQueue.PushBackList(bQueue)
		bQueue = list.New()
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
