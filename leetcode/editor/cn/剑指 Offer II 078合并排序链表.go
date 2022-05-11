// 给定一个链表数组，每个链表都已经按升序排列。
//
// 请将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
//
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
//  1->4->5,
//  1->3->4,
//  2->6
// ]
// 将它们合并到一个有序链表中得到。
// 1->1->2->3->4->4->5->6
//
//
// 示例 2：
//
//
// 输入：lists = []
// 输出：[]
//
//
// 示例 3：
//
//
// 输入：lists = [[]]
// 输出：[]
//
//
//
//
// 提示：
//
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
//
//
//
// 注意：本题与主站 23 题相同： https://leetcode-cn.com/problems/merge-k-sorted-lists/
// Related Topics 链表 分治 堆（优先队列） 归并排序 👍 37 👎 0

package main

import (
	"container/heap"
)

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

type PH078 struct {
	list []*ListNode
}

func (p *PH078) Len() int { return len(p.list) }

func (p *PH078) Less(i, j int) bool { return p.list[i].Val < p.list[j].Val }

func (p *PH078) Swap(i, j int) { p.list[i], p.list[j] = p.list[j], p.list[i] }

func (p *PH078) Push(x interface{}) { p.list = append(p.list, x.(*ListNode)) }

func (p *PH078) Pop() interface{} {
	x := p.list[len(p.list)-1]
	p.list = p.list[:len(p.list)-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	ph := &PH078{list: make([]*ListNode, 0, len(lists))}
	for i := range lists {
		if lists[i] != nil {
			heap.Push(ph, lists[i])
		}
	}
	head := new(ListNode)
	tail := head
	for ph.Len() > 0 {
		p := heap.Remove(ph, 0).(*ListNode)
		if p.Next != nil {
			heap.Push(ph, p.Next)
		}
		tail.Next = p
		tail = p
	}
	return head.Next
}

// leetcode submit region end(Prohibit modification and deletion)
