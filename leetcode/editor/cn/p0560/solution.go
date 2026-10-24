// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
//
// 子数组是数组中元素的连续非空序列。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,1,1], k = 2
// 输出：2
//
//
// 示例 2：
//
//
// 输入：nums = [1,2,3], k = 3
// 输出：2
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2 * 10⁴
// -1000 <= nums[i] <= 1000
// -10⁷ <= k <= 10⁷
//
//
// Related Topics 数组 哈希表 前缀和 👍 3288 👎 0

package p0560

// leetcode submit region begin(Prohibit modification and deletion)
// 解法：数组前缀和 + map
// map[k] = 这个前缀和 k 此前出现的次数
// map[0] = 1 // 统计第一个元素
// for i,num range nums:
//
//	prefixSum += num 累加和;
//	res += map[prefixSum - k] -- 注意先查历史匹配次数后再记录, 避免 k = 0 时会匹配到自己
//	map[prefixSum]++
func subarraySum(nums []int, k int) (res int) {
	m := make(map[int]int)
	m[0] = 1 // 数组开始前有一个前缀和 0，用于统计从数组第一个元素开始的合法子数组
	prefixSum := 0
	for _, num := range nums {
		prefixSum += num
		res += m[prefixSum-k]
		m[prefixSum]++
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
