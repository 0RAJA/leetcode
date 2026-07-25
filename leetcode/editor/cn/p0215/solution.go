// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//
//
// 示例 1:
//
//
// 输入: [3,2,1,5,6,4], k = 2
// 输出: 5
//
//
// 示例 2:
//
//
// 输入: [3,2,3,1,2,4,5,5,6], k = 4
// 输出: 4
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 3041 👎 0

package p0215

// leetcode submit region begin(Prohibit modification and deletion)
type MaxHeap []int

// 小根堆 python
/*
import heapq
from typing import List


class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        min_heap: list[int] = []

        for num in nums:
            heapq.heappush(min_heap, num)

            if len(min_heap) > k:
                heapq.heappop(min_heap)

        return min_heap[0]
*/
func findKthLargest(nums []int, k int) int {
	bucket := make([]int, 2*10000+1)
	for _, v := range nums {
		bucket[v+10000]++
	}
	for i := len(bucket) - 1; i >= 0; i-- {
		k -= bucket[i]
		if k <= 0 {
			return i - 10000
		}
	}
	return -1
}

// leetcode submit region end(Prohibit modification and deletion)
