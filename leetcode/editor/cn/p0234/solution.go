// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
// 输入：head = [1,2,2,1]
// 输出：true
//
//
// 示例 2：
//
//
// 输入：head = [1,2]
// 输出：false
//
//
//
//
// 提示：
//
//
// 链表中节点数目在范围[1, 10⁵] 内
// 0 <= Node.val <= 9
//
//
//
//
// 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//
// Related Topics 栈 递归 链表 双指针 👍 2284 👎 0

package p0234

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

// 双指针+递归：先用双指针算出中间的节点，然后递归到中间节点后回退，维护递归节点和中间节点，两者比较;
// 注意需要跳过奇数时的中间节点
func isPalindrome(head *ListNode) (res bool) {
	res = true
	slow, fast := head, head
	// slow 就是中点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 跳过奇数的中间节点
	midNode := slow
	if fast != nil {
		slow = slow.Next
	}
	var dfs func(root *ListNode)
	dfs = func(root *ListNode) {
		if root == midNode {
			return
		}
		dfs(root.Next)
		if !res {
			return
		}
		res = root.Val == slow.Val
		slow = slow.Next
	}
	dfs(head)
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
