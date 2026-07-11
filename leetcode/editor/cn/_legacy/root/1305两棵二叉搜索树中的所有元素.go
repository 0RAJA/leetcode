//给你 root1 和 root2 这两棵二叉搜索树。请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.
//
//
//
// 示例 1：
//
//
//
//
//输入：root1 = [2,1,4], root2 = [1,0,3]
//输出：[0,1,1,2,3,4]
//
//
// 示例 2：
//
//
//
//
//输入：root1 = [1,null,8], root2 = [8,1]
//输出：[1,1,8,8]
//
//
//
//
// 提示：
//
//
// 每棵树的节点数在 [0, 5000] 范围内
// -10⁵ <= Node.val <= 10⁵
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 排序 👍 113 👎 0

package main

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
func getAllElements(root1 *TreeNode, root2 *TreeNode) (ret []int) {
	stack1, stack2 := make([]*TreeNode, 0, 5010), make([]*TreeNode, 0, 5010)
	ret = make([]int, 0, 10010)
	addRoo1 := func(root *TreeNode) {
		for root != nil {
			stack1 = append(stack1, root)
			root = root.Left
		}
	}
	addRoo2 := func(root *TreeNode) {
		for root != nil {
			stack2 = append(stack2, root)
			root = root.Left
		}
	}
	addRoo1(root1)
	addRoo2(root2)
	for len(stack1) > 0 || len(stack2) > 0 {
		var p1, p2 *TreeNode
		if len(stack1) > 0 {
			p1 = stack1[len(stack1)-1]
		}
		if len(stack2) > 0 {
			p2 = stack2[len(stack2)-1]
		}
		if p1 != nil && (p2 == nil || p1.Val <= p2.Val) {
			ret = append(ret, p1.Val)
			stack1 = stack1[:len(stack1)-1]
			addRoo1(p1.Right)
		} else {
			ret = append(ret, p2.Val)
			stack2 = stack2[:len(stack2)-1]
			addRoo2(p2.Right)
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
