// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
// 。
//
// 返回 滑动窗口中的最大值 。 
//
// 
//
// 示例 1： 
//
// 
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
// 解释：
// 滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
// 
//
// 示例 2： 
//
// 
// 输入：nums = [1], k = 1
// 输出：[1]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10⁵ 
// -10⁴ <= nums[i] <= 10⁴ 
// 1 <= k <= nums.length 
// 
//
// Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 👍 3469 👎 0
package maxSlidingWindow

func main() {
	maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
}

// leetcode submit region begin(Prohibit modification and deletion)
type monotonousQueue struct {
	queue []int
}

func (q *monotonousQueue) Push(data int) {
	for i := len(q.queue) - 1; i >= 0; i-- {
		if data <= q.queue[i] {
			q.queue = append(q.queue, data)
			return
		}
		q.queue = q.queue[:len(q.queue)-1]
	}
	if len(q.queue) == 0 {
		q.queue = append(q.queue, data)
		return
	}
}

func (q *monotonousQueue) Pop() (val int) {
	defer func() {
		q.queue = q.queue[1:]
	}()
	return q.queue[0]
}

func (q *monotonousQueue) Top() int {
	return q.queue[0]
}

func (q *monotonousQueue) Empty() bool {
	return len(q.queue) == 0
}

func maxSlidingWindow(nums []int, k int) (ret []int) {
	// 单调队列 维持入队顺序且元素值递减
	mQ := &monotonousQueue{queue: make([]int, 0, len(nums))}
	// 记录前 k 个元素入队
	for i := 0; i < k; i++ {
		mQ.Push(nums[i])
	}
	// 记录第一个k 区间最大值
	ret = append(ret, mQ.Top())
	for i := k; i < len(nums); i++ {
		// 将滑出窗口的元素出队
		popValue := nums[i-k]
		if popValue == mQ.Top() {
			mQ.Pop()
		}
		// 将滑入窗口的元素入队
		mQ.Push(nums[i])
		// 记录当前窗口最大值
		ret = append(ret, mQ.Top())
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
