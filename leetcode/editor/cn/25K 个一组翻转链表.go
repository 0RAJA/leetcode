// 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
//
//
// 示例 1：
//
//
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[2,1,4,3,5]
//
//
// 示例 2：
//
//
//
//
// 输入：head = [1,2,3,4,5], k = 3
// 输出：[3,2,1,4,5]
//
//
//
// 提示：
//
//
// 链表中的节点数目为 n
// 1 <= k <= n <= 5000
// 0 <= Node.val <= 1000
//
//
//
//
// 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
//
//
//
//
// Related Topics 递归 链表 👍 1842 👎 0

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
func reverseKGroup(head *ListNode, k int) (ret *ListNode) {
	nh := &ListNode{Next: head} // 作为每轮反转节点的头节点
	ret = nh                    // 保存第一个头节点
	for lh, rh := head, head; lh != nil; {
		for i := 0; i < k; i++ { // rh找到下一轮反转的第一个节点，找不到说明凑不够,直接返回
			if rh == nil {
				return ret.Next
			}
			rh = rh.Next
		}
		t := lh // 保存此轮反转后的最后一个节点
		for lh != rh {
			q := lh.Next
			lh.Next = nh.Next
			nh.Next = lh
			lh = q
		}
		t.Next = rh // 连接下一轮的首节点
		nh = t      // 更新每轮反转节点的首节点
	}
	return ret.Next
}

// leetcode submit region end(Prohibit modification and deletion)
