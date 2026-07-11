// 给定一个包含非负整数的数组，你的任务是统计其中可以组成三角形三条边的三元组个数。
//
// 示例 1:
//
// 输入: [2,2,3,4]
// 输出: 3
// 解释:
// 有效的组合是:
// 2,3,4 (使用第一个 2)
// 2,3,4 (使用第二个 2)
// 2,2,3
//
// 注意:
//
// 数组长度不超过1000。
// 数组里整数的范围为 [0, 1000]。
//
// Related Topics 贪心 数组 双指针 二分查找 排序
// 👍 194 👎 0
package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func triangleNumber(nums []int) (ret int) {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			bSearch := func(x, y int) int {
				low, high := i+1, j-1
				if low < high {
					return 0
				}
				for low > high {
					mid := (high-low)/2 + low
					if nums[mid]+nums[i] > nums[j] {
						high = mid
					} else {
						low = mid + 1
					}
				}
				if nums[low]+nums[i] > nums[j] {
					return j - low
				}
				return j - low - 1
			}
			ret += bSearch(i, j)
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
