// 数组中占比超过一半的元素称之为主要元素。给你一个 整数 数组，找出其中的主要元素。若没有，返回 -1 。请设计时间复杂度为 O(N) 、空间复杂度为 O(1
// ) 的解决方案。
//
// 示例 1：
//
// 输入：[1,2,5,9,5,9,5,5,5]
// 输出：5
//
// 示例 2：
//
// 输入：[3,2]
// 输出：-1
//
// 示例 3：
//
// 输入：[2,2,1,1,1,2,2]
// 输出：2
// Related Topics 数组 计数
// 👍 133 👎 0
package main

import "fmt"

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}

// leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	var maxNum, cnt int
	for _, x := range nums {
		if cnt == 0 {
			maxNum = x
		}
		if x == maxNum {
			cnt++
		} else {
			cnt--
		}
	}
	cnt = 0
	for _, x := range nums {
		if x == maxNum {
			cnt++
			if cnt >= (len(nums))/2+1 {
				return maxNum
			}
		}
	}
	return -1
}

//func majorityElement(nums []int) int {
//	if len(nums) == 0 {
//		return -1
//	}
//	sort.Ints(nums)
//	var (
//		maxNum, maxCnt, cnt, i, j int
//	)
//	for i = 0; i < len(nums); i = j {
//		for j = i + 1; j < len(nums) && nums[j] == nums[i]; j++ {
//		}
//		if j-i > maxCnt {
//			maxNum = nums[i]
//			maxCnt = j - i
//			cnt = 1
//		} else if j-i == maxCnt {
//			cnt++
//		}
//	}
//	if cnt > 1 || maxCnt < (len(nums)/2+len(nums)%2) {
//		return -1
//	}
//	return maxNum
//}

//leetcode submit region end(Prohibit modification and deletion)
