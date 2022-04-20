//完全二叉树是每一层（除最后一层外）都是完全填充（即，节点数达到最大，第 n 层有 2ⁿ⁻¹ 个节点）的，并且所有的节点都尽可能地集中在左侧。
//
// 设计一个用完全二叉树初始化的数据结构 CBTInserter，它支持以下几种操作：
//
//
// CBTInserter(TreeNode root) 使用根节点为 root 的给定树初始化该数据结构；
// CBTInserter.insert(int v) 向树中插入一个新节点，节点类型为 TreeNode，值为 v 。使树保持完全二叉树的状态，并返回插入的
//新节点的父节点的值；
// CBTInserter.get_root() 将返回树的根节点。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：inputs = ["CBTInserter","insert","get_root"], inputs = [[[1]],[2],[]]
//输出：[null,1,[1,2]]
//
//
// 示例 2：
//
//
//输入：inputs = ["CBTInserter","insert","insert","get_root"], inputs = [[[1,2,3,4,
//5,6]],[7],[8],[]]
//输出：[null,3,4,[1,2,3,4,5,6,7,8]]
//
//
//
//
// 提示：
//
//
// 最初给定的树是完全二叉树，且包含 1 到 1000 个节点。
// 每个测试用例最多调用 CBTInserter.insert 操作 10000 次。
// 给定节点或插入节点的每个值都在 0 到 5000 之间。
//
//
//
//
// 注意：本题与主站 919 题相同： https://leetcode-cn.com/problems/complete-binary-tree-
//inserter/
// Related Topics 树 广度优先搜索 设计 二叉树 👍 20 👎 0

package main

import "sort"

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

//保存层序遍历的切片，以及第一个满足左右节点有一个为空的节点的下标，之后就只需要维护这个下标然后插入就行。
type CBTInserter struct {
	root *TreeNode
	nums []*TreeNode
	cur  int
}

func Constructor(root *TreeNode) CBTInserter {
	CBT := CBTInserter{
		root: root,
	}
	nums := make([]*TreeNode, 0, 11010)
	nums = append(nums, root)
	for i := 0; i < len(nums); i++ {
		if nums[i].Left != nil {
			nums = append(nums, nums[i].Left)
		}
		if nums[i].Right != nil {
			nums = append(nums, nums[i].Right)
		}
	}
	CBT.nums = nums
	CBT.cur = sort.Search(len(nums), func(i int) bool { return CBT.nums[i].Left == nil || CBT.nums[i].Right == nil })
	return CBT
}

func (this *CBTInserter) Insert(v int) int {
	p := &TreeNode{Val: v}
	now := this.nums[this.cur]
	if now.Left == nil {
		now.Left = p
	} else if now.Right == nil {
		now.Right = p
		this.cur++
	}
	this.nums = append(this.nums, p)
	return now.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */
//leetcode submit region end(Prohibit modification and deletion)
