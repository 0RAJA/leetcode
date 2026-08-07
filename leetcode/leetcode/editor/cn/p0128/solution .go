// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//
//
// 示例 1：
//
//
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
// 示例 2：
//
//
// 输入：nums = [0,3,7,2,5,8,4,6,0,1]
// 输出：9
//
//
// 示例 3：
//
//
// 输入：nums = [1,0,1,2]
// 输出：3
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
//
//
// Related Topics 并查集 数组 哈希表 👍 2945 👎 0

package p0128

// leetcode submit region begin(Prohibit modification and deletion)
// 使用 num_set 来记录去重后的num，然后遍历 num，如果 num-1 存在则跳过，
// 如果num-1不存在则该num为数组第一个，从该num开始向后遍历来计算最长长度
func longestConsecutive(nums []int) (res int) {
	numSet := make(map[int]bool, len(nums))
	for _, v := range nums {
		numSet[v] = true
	}
	for k := range numSet {
		if numSet[k-1] {
			continue
		}
		cnt := 1
		for k++; numSet[k]; k++ {
			cnt++
		}
		res = max(res, cnt)
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
