// 给定链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
//
//
//
//
//
// 示例 1：
//
//
//
//
// 输入：head = [4,2,1,3]
// 输出：[1,2,3,4]
//
//
// 示例 2：
//
//
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
//
//
// 注意：本题与主站 148 题相同：https://leetcode-cn.com/problems/sort-list/
// Related Topics 链表 双指针 分治 排序 归并排序 👍 58 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// SortList 给出链表的头节点和最后一个有效节点的下一个节点
func SortList(head, tail *ListNode) *ListNode {
	// 0
	if head == nil {
		return head
	}
	// 1
	if head.Next == tail {
		head.Next = nil
		return head
	}
	// >=2 快慢指针找中点
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	// merge
	return mergeList(SortList(head, mid), SortList(mid, tail))
}

func mergeList(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	t, t1, t2 := dummyHead, head1, head2
	for t1 != nil && t2 != nil {
		if t1.Val <= t2.Val {
			t.Next = t1
			t1 = t1.Next
		} else {
			t.Next = t2
			t2 = t2.Next
		}
		t = t.Next
	}
	if t1 != nil {
		t.Next = t1
	} else {
		t.Next = t2
	}
	return dummyHead.Next
}

func sortList(head *ListNode) *ListNode {
	return SortList(head, nil)
}

// leetcode submit region end(Prohibit modification and deletion)
