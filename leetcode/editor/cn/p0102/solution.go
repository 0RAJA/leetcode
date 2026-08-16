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
// Related Topics 树 广度优先搜索 二叉树 👍 2320 👎 0

package p0102

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
func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	treeNodeQueue := []*TreeNode{
		root,
	}
	for len(treeNodeQueue) > 0 {
		length := len(treeNodeQueue)
		data := make([]int, 0, length)
		for i := 0; i < length; i++ {
			node := treeNodeQueue[0]
			data = append(data, node.Val)
			if node.Left != nil {
				treeNodeQueue = append(treeNodeQueue, node.Left)
			}
			if node.Right != nil {
				treeNodeQueue = append(treeNodeQueue, node.Right)
			}
			treeNodeQueue = treeNodeQueue[1:]
		}
		res = append(res, data)
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
