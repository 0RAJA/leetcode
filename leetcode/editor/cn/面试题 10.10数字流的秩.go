//假设你正在读取一串整数。每隔一段时间，你希望能找出数字 x 的秩(小于或等于 x 的值的个数)。请实现数据结构和算法来支持这些操作，也就是说：
//
// 实现 track(int x) 方法，每读入一个数字都会调用该方法；
//
// 实现 getRankOfNumber(int x) 方法，返回小于或等于 x 的值的个数。
//
// 注意：本题相对原题稍作改动
//
// 示例:
//
// 输入:
//["StreamRank", "getRankOfNumber", "track", "getRankOfNumber"]
//[[], [1], [0], [0]]
//输出:
//[null,0,null,1]
//
//
// 提示：
//
//
// x <= 50000
// track 和 getRankOfNumber 方法的调用次数均不超过 2000 次
//
// Related Topics 设计 树状数组 二分查找 数据流 👍 29 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type StreamRank struct {
	*TreeNum
}

func Constructor() StreamRank {
	max := 50001
	return StreamRank{
		TreeNum: NewTreeNum(max),
	}
}

func (this *StreamRank) Track(x int) {
	this.add(x, 1)
}

func (this *StreamRank) GetRankOfNumber(x int) int {
	return this.ask(x)
}

type TreeNum struct {
	nums []int
}

func NewTreeNum(n int) *TreeNum {
	return &TreeNum{make([]int, n+1)} //下标从1开始
}

//计算x二进制的最低位1及之后的数。101000 -> 001000
func (t *TreeNum) lowbit(x int) int {
	return x & (-x)
}

//x点值+d
func (t *TreeNum) add(x, d int) {
	x++
	for ; x < len(t.nums); x += t.lowbit(x) { //因为动态维护一个前缀和数组，所以依次都要+d
		t.nums[x] += d
	}
}

//查询x下标处的前缀和
func (t *TreeNum) ask(x int) (ret int) {
	x++
	for ; x >= 1 && x < len(t.nums); x -= t.lowbit(x) {
		ret += t.nums[x]
	}
	return
}

/**
 * Your StreamRank object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Track(x);
 * param_2 := obj.GetRankOfNumber(x);
 */
//leetcode submit region end(Prohibit modification and deletion)
