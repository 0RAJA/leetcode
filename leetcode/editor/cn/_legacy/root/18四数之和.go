//给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[
//b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
//
//
// 0 <= a, b, c, d < n
// a、b、c 和 d 互不相同
// nums[a] + nums[b] + nums[c] + nums[d] == target
//
//
// 你可以按 任意顺序 返回答案 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = [2,2,2,2,2], target = 8
//输出：[[2,2,2,2]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
//
// Related Topics 数组 双指针 排序 👍 998 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSum(nums, target))
}

// leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) (ret [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for a := 0; a < n; a++ { //a从左向右
		if a > 0 && nums[a] == nums[a-1] { //a不重复
			continue
		}
		for d := n - 1; a+2 < d; d-- { //d从右向左
			if d < n-1 && nums[d] == nums[d+1] { //d不重复
				continue
			}
			for b, c := a+1, d-1; c > b; { //b,c作为两个指针
				if sum := nums[a] + nums[b] + nums[c] + nums[d]; sum == target { //通过求和决定指针的移动
					ret = append(ret, []int{nums[a], nums[b], nums[c], nums[d]})
					for b++; nums[b] == nums[b-1] && b < c; b++ {
					}
					for c--; nums[c] == nums[c+1] && b < c; c-- {
					}
				} else if sum < target {
					b++
				} else {
					c--
				}
			}
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
