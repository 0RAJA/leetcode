// 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a ，b ，c ，使得 a + b + c = 0 ？请找出所有和为 0 且
// 不重复 的三元组。
//
// 
//
// 示例 1： 
//
// 
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 
//
// 示例 2： 
//
// 
// 输入：nums = []
// 输出：[]
// 
//
// 示例 3： 
//
// 
// 输入：nums = [0]
// 输出：[]
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
// 
// 注意：本题与主站 15 题相同：https://leetcode.cn/problems/3sum/ 
//
// Related Topics 数组 双指针 排序 👍 172 👎 0
package main

import (
	"sort"
)

// leetcode submit region begin(Prohibit modification and deletion)
// func threeSum(nums []int) (ret [][]int) {
// 	// 构造一个有序数组
// 	sort.Ints(nums)
// 	for left := 0; left < len(nums)-1; left++ {
// 		// 避免 left 重复
// 		if left != 0 && nums[left] == nums[left-1] {
// 			continue
// 		}
// 		// 固定 left，引入 mid 和 right 双指针来探索边界，寻找另外两个数
// 		for mid, right := left+1, len(nums)-1; mid < right; {
// 			sum := nums[left] + nums[mid] + nums[right]
// 			if sum == 0 {
// 				ret = append(ret, []int{nums[left], nums[mid], nums[right]})
// 				// 避免 mid 重复
// 				for mid++; mid < right && nums[mid] == nums[mid-1]; mid++ {
// 				}
// 				// 避免 right 重复
// 				for right--; mid < right && nums[right] == nums[right+1]; right-- {
// 				}
// 			} else if sum > 0 {
// 				right--
// 			} else {
// 				mid++
// 			}
// 		}
// 	}
// 	return
// }

// 排序，每次固定一个 i，然后遍历 j，k
// 关键是去重，跳过 i == 前一个的
func threeSum(nums []int) (res [][]int) {
	// 从小到大
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		// 跳过重复的 i
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		// i 固定了，然后右侧找 j 和 k
		for j, k := i+1, len(nums)-1; j < k; {
			target := nums[i] + nums[j] + nums[k]
			// 找到了
			if target == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				// 然后跳过重复值
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
				for k--; j < k && nums[k] == nums[k+1]; k-- {
				}
				// 否则根据大小再遍历
			} else if target > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
