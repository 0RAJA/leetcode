// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//
//
// 示例 1：
//
//
// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
//
//
// 示例 2：
//
//
// 输入：head = [1], n = 1
// 输出：[]
//
//
// 示例 3：
//
//
// 输入：head = [1,2], n = 1
// 输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
//
//
//
// 进阶：你能尝试使用一趟扫描实现吗？
//
// Related Topics 链表 双指针 👍 3342 👎 0

package p0019

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
// 维护两个节点（前后节点），间隔 n 个节点，当其中一个节点跑到头的时候删掉 Next 节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	preNode, curNode := dummyHead, head
	for i := 0; i < n; i++ {
		curNode = curNode.Next
	}
	for curNode != nil {
		preNode = preNode.Next
		curNode = curNode.Next
	}
	preNode.Next = preNode.Next.Next
	return dummyHead.Next
}

// leetcode submit region end(Prohibit modification and deletion)
