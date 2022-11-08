// 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
//
// 示例:
//
//
// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
// 输出: [3,3,5,5,6,7]
// 解释:
//
//  滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
//
// 提示：
//
// 你可以假设 k 总是有效的，在输入数组 不为空 的情况下，1 ≤ k ≤ nums.length。
//
// 注意：本题与主站 239 题相同：https://leetcode-cn.com/problems/sliding-window-maximum/
//
// Related Topics 队列 滑动窗口 单调队列 堆（优先队列） 👍 509 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) (ret []int) {
	var q []int
	in := func(idx int) {
		for len(q) > 0 && nums[idx] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, idx)
	}
	max := func() int {
		return q[0]
	}
	for i := 0; i < k; i++ {
		in(i)
	}
	ret = append(ret, nums[max()]) // 第一个窗口取最大值
	for l, r := 0, k; r < len(nums); r++ {
		if max() == l { // t淘汰
			q = q[1:]
		}
		l++                            // 左窗口右移
		in(r)                          // 右窗口右移
		ret = append(ret, nums[max()]) // 取当前窗口最大值
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
