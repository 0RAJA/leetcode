// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//
//
// 示例 1：
//
//
// 输入：nums = [3,2,3]
// 输出：3
//
// 示例 2：
//
//
// 输入：nums = [2,2,1,1,1,2,2]
// 输出：2
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 5 * 10⁴
// -10⁹ <= nums[i] <= 10⁹
// 输入保证数组中一定有一个多数元素。
//
//
//
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
//
// Related Topics 数组 哈希表 分治 计数 排序 👍 2715 👎 0

package p0169

// leetcode submit region begin(Prohibit modification and deletion)
// 我们只需要找到多数元素即可：其出现数量一定大于 n/2；那么就维护一个多数元素，出现其他元素就cnt--，如果cnt==0则更换元素
func majorityElement(nums []int) int {
	maj := 0
	cnt := 0
	for _, v := range nums {
		if v == maj {
			cnt++
		} else if cnt == 0 {
			maj = v
			cnt = 1
		} else {
			cnt--
		}
	}
	return maj
}

// leetcode submit region end(Prohibit modification and deletion)
