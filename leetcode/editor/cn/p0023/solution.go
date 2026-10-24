// 给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
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
// 输入：lists = []
// 输出：[]
//
//
// 示例 3：
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
// Related Topics 链表 分治 堆（优先队列） 归并排序 锦标赛排序 👍 3302 👎 0

package p0023

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

type NodeHeap struct {
	nl []*ListNode
}

// Push 维护一个完全二叉树数组，Push 就是追加到末尾，然后向上冒泡
func (n *NodeHeap) Push(node *ListNode) {
	n.nl = append(n.nl, node)
	for i := len(n.nl) - 1; i > 0; {
		nextIdx := (i - 1) / 2
		if n.nl[nextIdx].Val <= node.Val {
			break
		}
		n.nl[nextIdx], n.nl[i] = n.nl[i], n.nl[nextIdx]
		i = nextIdx
	}
}

// Pop 剔除首个元素（最小值），然后把最后一个元素放到开头，然后下沉，注意，左右节点中找最小的比较交换
func (n *NodeHeap) Pop() (res *ListNode) {
	if len(n.nl) == 0 {
		return nil
	}
	res = n.nl[0]
	n.nl[0], n.nl[len(n.nl)-1] = n.nl[len(n.nl)-1], n.nl[0]
	n.nl = n.nl[:len(n.nl)-1]
	for i := 0; i < len(n.nl); {
		nextIdx := i*2 + 1
		if nextIdx >= len(n.nl) {
			break
		}
		rightIdx := i*2 + 2
		if rightIdx < len(n.nl) && n.nl[nextIdx].Val > n.nl[rightIdx].Val {
			nextIdx = rightIdx
		}
		if n.nl[i].Val <= n.nl[nextIdx].Val {
			break
		}
		n.nl[i], n.nl[nextIdx] = n.nl[nextIdx], n.nl[i]
		i = nextIdx
	}
	return res
}

// 1. 两两合并
// 2. 维护一个 k 个元素的最小堆，每次 pop 出最小元素追加到新的联表中，然后把该元素的 next push 到堆中
// 维护一个完全二叉树数组: x 的左节点: x*2+1；右节点: x*2+2；父节点: (x-1)/2
// 注意堆维护两个方法：
// Push：追加到末尾，然后向上找比自己大的父节点冒泡
// Pop：剔除首个元素（最小值），然后把最后一个元素放到开头，左右节点中找最小的比较下沉
func mergeKLists(lists []*ListNode) *ListNode {
	heap := &NodeHeap{nl: make([]*ListNode, 0, len(lists))}
	dummy := &ListNode{Next: nil}
	next := dummy
	for _, n := range lists {
		if n != nil {
			heap.Push(n)
		}
	}
	for node := heap.Pop(); node != nil; {
		next.Next = node
		next = node
		if node.Next != nil {
			heap.Push(node.Next)
		}
		node = heap.Pop()
	}
	return dummy.Next
}

// leetcode submit region end(Prohibit modification and deletion)
