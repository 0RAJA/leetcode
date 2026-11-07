// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
// 示例 1：
//
//
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
//
//
// 示例 2：
//
//
// 输入：l1 = [0], l2 = [0]
// 输出：[0]
//
//
// 示例 3：
//
//
// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]
//
//
//
//
// 提示：
//
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
//
//
// Related Topics 递归 链表 数学 👍 12160 👎 0

package p0002

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 依次合并 l1 和 l2 节点以及前一个节点的进位；终止条件为：node1 == nil && node2 == nil && 前一个进位 == 0
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 合并 node1、node2 以及进位
	var mergeNode func(node1 *ListNode, node2 *ListNode, num int) *ListNode
	mergeNode = func(node1 *ListNode, node2 *ListNode, num int) *ListNode {
		node := &ListNode{
			Val:  0,
			Next: nil,
		}
		if node1 != nil {
			node.Val += node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			node.Val += node2.Val
			node2 = node2.Next
		}
		node.Val += num
		num = node.Val / 10
		node.Val = node.Val % 10
		if node1 == nil && node2 == nil && num == 0 {
			return node
		}
		node.Next = mergeNode(node1, node2, num)
		return node
	}
	return mergeNode(l1, l2, 0)
}

// leetcode submit region end(Prohibit modification and deletion)
