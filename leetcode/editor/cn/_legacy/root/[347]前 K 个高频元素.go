// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//
// 
//
// 示例 1： 
//
// 
// 输入：nums = [1,1,1,2,2,3], k = 2 
// 
//
// 输出：[1,2] 
//
// 示例 2： 
//
// 
// 输入：nums = [1], k = 1 
// 
//
// 输出：[1] 
//
// 示例 3： 
//
// 
// 输入：nums = [1,2,1,2,1,2,3,1,3,2], k = 2 
// 
//
// 输出：[1,2] 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10⁵ 
// -10⁴ <= nums[i] <= 10⁴ 
// k 的取值范围是 [1, 数组中不相同的元素的个数] 
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的 
// 
//
// 
//
// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。 
//
// Related Topics 数组 哈希表 分治 桶排序 计数 快速选择 排序 堆（优先队列） 👍 2179 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func topKFrequent(nums []int, k int) (ret []int) {
	// 统计每个元素出现的次数
	countMap := make(map[int]int, len(nums))
	for _, num := range nums {
		countMap[num]++
	}
	// 使用桶排序，桶下标为元素出现次数，值为出现该次数的元素列表
	bucket := make([][]int, len(nums)+1)
	for num, count := range countMap {
		bucket[count] = append(bucket[count], num)
	}
	// 由大到小遍历桶，把出现过的元素加入结果集
	for i := len(bucket) - 1; i >= 0; i-- {
		if bucket[i] != nil {
			ret = append(ret, bucket[i]...)
		}
	}
	return ret[:k]
}

// leetcode submit region end(Prohibit modification and deletion)
