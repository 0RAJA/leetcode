// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
// 实现 MinStack 类:
//
//
// MinStack() 初始化堆栈对象。
// void push(int value) 将元素 value 推入堆栈。
// void pop() 删除堆栈顶部的元素。
// int top() 获取堆栈顶部的元素。
// int getMin() 获取堆栈中的最小元素。
//
//
//
//
// 示例 1:
//
//
// 输入：
// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
//
// 输出：
// [null,null,null,null,-3,null,0,-2]
//
// 解释：
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 0.
// minStack.getMin();   --> 返回 -2.
//
//
//
//
// 提示：
//
//
// -2³¹ <= val <= 2³¹ - 1
// pop、top 和 getMin 操作总是在 非空栈 上调用
// push, pop, top, and getMin最多被调用 3 * 10⁴ 次
//
//
// Related Topics 栈 设计 👍 2098 👎 0

package p0155

// leetcode submit region begin(Prohibit modification and deletion)

// MinNode
// 每个节点里存储当前 val 和当前最小值
type MinNode struct {
	val, minNum int
}

type MinStack struct {
	nodes []MinNode
}

/** initialize your data structure here. */
func Constructor() MinStack {
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

func (this *MinStack) GetMin() int {
	return this.nodes[len(this.nodes)-1].minNum
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// leetcode submit region end(Prohibit modification and deletion)
