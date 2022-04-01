//给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a ，b ，c ，使得 a + b + c = 0 ？请找出所有和为 0 且
//不重复 的三元组。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
//
//
//
// 注意：本题与主站 15 题相同：https://leetcode-cn.com/problems/3sum/
// Related Topics 数组 双指针 排序 👍 47 👎 0

package main

import "sort"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum007(nums []int) (ret [][]int) {
	target := 0
	n := len(nums)
	sort.Ints(nums)
	for l := 0; l < n-2; l++ {
		//去重
		if l > 0 && nums[l] == nums[l-1] {
			continue
		}
		for a, b := l+1, n-1; a < b; {
			if sum := nums[l] + nums[a] + nums[b]; sum == target {
				ret = append(ret, []int{nums[l], nums[a], nums[b]})
				//去重
				for a++; nums[a] == nums[a-1] && a < b; a++ {
				}
				for b--; nums[b] == nums[b+1] && a < b; b-- {
				}
			} else if sum < target {
				a++
			} else {
				b--
			}
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
