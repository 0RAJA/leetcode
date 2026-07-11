//定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
//
//
//
// 示例:
//
// MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.min();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.min();   --> 返回 -2.
//
//
//
//
// 提示：
//
//
// 各函数的调用总次数不超过 20000 次
//
//
//
//
// 注意：本题与主站 155 题相同：https://leetcode-cn.com/problems/min-stack/
// Related Topics 栈 设计 👍 242 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type MinNode struct {
	val, minNum int
}
type MinStack struct {
	nodes []MinNode
}

/** initialize your data structure here. */
func Constructor30() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	minNum := x
	if len(this.nodes) != 0 && this.nodes[len(this.nodes)-1].minNum < minNum {
		minNum = this.nodes[len(this.nodes)-1].minNum
	}
	this.nodes = append(this.nodes, MinNode{
		val:    x,
		minNum: minNum,
	})
}

func (this *MinStack) Pop() {
	this.nodes = this.nodes[:len(this.nodes)-1]
}

func (this *MinStack) Top() int {
	return this.nodes[len(this.nodes)-1].val
}

func (this *MinStack) Min() int {
	return this.nodes[len(this.nodes)-1].minNum
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
//leetcode submit region end(Prohibit modification and deletion)
