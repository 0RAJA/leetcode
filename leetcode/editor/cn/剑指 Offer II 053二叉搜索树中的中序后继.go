//给定一棵二叉搜索树和其中的一个节点 p ，找到该节点在树中的中序后继。如果节点没有中序后继，请返回 null 。
//
// 节点 p 的后继是值比 p.val 大的节点中键值最小的节点，即按中序遍历的顺序节点 p 的下一个节点。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [2,1,3], p = 1
//输出：2
//解释：这里 1 的中序后继是 2。请注意 p 和返回值都应是 TreeNode 类型。
//
//
// 示例 2：
//
//
//
//
//输入：root = [5,3,6,2,4,null,null,1], p = 6
//输出：null
//解释：因为给出的节点没有中序后继，所以答案就返回 null 了。
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [1, 10⁴] 内。
// -10⁵ <= Node.val <= 10⁵
// 树中各节点的值均保证唯一。
//
//
//
//
// 注意：本题与主站 285 题相同： https://leetcode-cn.com/problems/inorder-successor-in-bst/
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 28 👎 0

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
func inorderSuccessor(root *TreeNode, p *TreeNode) (ret *TreeNode) {
	var dfs func(root *TreeNode)
	var pre *TreeNode
	dfs = func(root *TreeNode) {
		if ret == nil && root != nil {
			dfs(root.Left)              //一直往左边探索直到探索完
			if ret == nil && pre == p { //如果发现前驱节点是目标值则直接返回自己
				ret = root
				return
			}
			pre = root //左边走完了，更新前驱
			dfs(root.Right)
		}
	}
	dfs(root)
	return
}

//利用二叉搜索树的条件
/*
充分利用二叉搜索树的特点，这个题设条件没用上的话，面试肯定被怼的！
	如果p节点有右子树，那么返回其右子树最左边的节点
	如果p无右子树，需要返回p所在的左子树的根节点（应该尽量接近p）
	如果p不在任意一棵左子树上，那么返回null
*/

func inorderSuccessor2(root *TreeNode, p *TreeNode) (ret *TreeNode) {
	var findLeft func(root *TreeNode) *TreeNode
	findLeft = func(root *TreeNode) *TreeNode {
		if root.Left == nil {
			return root
		}
		return findLeft(root.Left)
	}
	if p.Right != nil {
		return findLeft(p.Right)
	}
	var prv *TreeNode
	for root != p { //注意往右走是不更新前驱的，不会再走了
		if root.Val < p.Val {
			root = root.Right
		} else {
			prv = root
			root = root.Left
		}
	}
	return prv
}

//leetcode submit region end(Prohibit modification and deletion)
