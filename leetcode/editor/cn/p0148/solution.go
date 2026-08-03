// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
//
//
//
//
//
// 示例 1：
//
//
// 输入：head = [4,2,1,3]
// 输出：[1,2,3,4]
//
//
// 示例 2：
//
//
// 输入：head = [-1,5,3,4,0]
// 输出：[-1,0,3,4,5]
//
//
// 示例 3：
//
//
// 输入：head = []
// 输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 5 * 10⁴] 内
// -10⁵ <= Node.val <= 10⁵
//
//
//
//
// 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
//
// Related Topics 链表 双指针 分治 排序 归并排序 👍 2815 👎 0

package p0148

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
// 归并排序：双指针找到两边的 head，然后拆开，然后把两个已经 sorted 的联表 merge 成一个 sorted 的联表
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 找到左半部分的最后一个节点
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 切断链表
	rightHead := slow.Next
	slow.Next = nil

	// 分别排序
	left := sortList(head)
	right := sortList(rightHead)

	// 合并两个有序链表
	return mergeSortedLists(left, right)
}

func mergeSortedLists(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for left != nil && right != nil {
		if left.Val <= right.Val {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}
		current = current.Next
	}

	if left != nil {
		current.Next = left
	} else {
		current.Next = right
	}

	return dummy.Next
}

// leetcode submit region end(Prohibit modification and deletion)
