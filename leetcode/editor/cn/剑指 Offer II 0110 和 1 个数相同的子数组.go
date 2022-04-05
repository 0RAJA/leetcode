//给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
//
//
//
// 示例 1:
//
//
//输入: nums = [0,1]
//输出: 2
//说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
//
// 示例 2:
//
//
//输入: nums = [0,1,0]
//输出: 2
//说明: [0, 1] (或 [1, 0]) 是具有相同数量 0 和 1 的最长连续子数组。
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// nums[i] 不是 0 就是 1
//
//
//
//
// 注意：本题与主站 525 题相同： https://leetcode-cn.com/problems/contiguous-array/
// Related Topics 数组 哈希表 前缀和 👍 54 👎 0

package main

import "fmt"

func main() {
	fmt.Println(findMaxLength([]int{0, 1, 0, 0, 1, 1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findMaxLength2(nums []int) (ret int) {
	prefix := make([]int, len(nums)+1)
	m := make(map[int]int)
	for i := range nums {
		if nums[i] == 0 {
			nums[i] = -1
		}
		sum := prefix[i] + nums[i]
		if sum == 0 {
			ret = i + 1
		}
		if idx, ok := m[sum]; ok {
			if i-idx+1 > ret {
				ret = i - idx
			}
		} else {
			m[sum] = i
		}
		prefix[i+1] = sum
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
