// 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
//
//
// 示例 1：
//
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[20,9],[15,7]]
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
// -100 <= Node.val <= 100
//
//
// Related Topics 树 广度优先搜索 二叉树 👍 708 👎 0

package main

func main() {

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
func zigzagLevelOrder(root *TreeNode) (ret [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	push := func(node *TreeNode) {
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	cnt := 1
	for len(q) > 0 {
		n := len(q)
		var res []int
		if cnt%2 != 0 {
			for i := 0; i < n; i++ {
				res = append(res, q[i].Val)
			}
		} else {
			for i := n - 1; i >= 0; i-- {
				res = append(res, q[i].Val)
			}
		}
		ret = append(ret, res)
		for _, v := range q {
			push(v)
		}
		q = q[n:]
		cnt++
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
