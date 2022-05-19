// 给你一个长度为 n 的整数数组 nums ，返回使所有数组元素相等需要的最少移动数。
//
// 在一步操作中，你可以使数组中的一个元素加 1 或者减 1 。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,2,3]
// 输出：2
// 解释：
// 只需要两步操作（每步操作指南使一个元素加 1 或减 1）：
// [1,2,3]  =>  [2,2,3]  =>  [2,2,2]
//
//
// 示例 2：
//
//
// 输入：nums = [1,10,2,9]
// 输出：16
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
//
// Related Topics 数组 数学 排序 👍 170 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minMoves2([]int{1, 0, 0, 8, 6}))
}

// leetcode submit region begin(Prohibit modification and deletion)
/*
	最终形态是中位数
	可以理解为左右指针移动，左右指针需要相同则更改的数量为两者的差值
	{
		for l < r {
			ret += r-l
		}
	}
*/
func minMoves2(nums []int) (ret int) {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	sort.Ints(nums)
	x := nums[len(nums)/2]
	for _, v := range nums {
		ret += abs(v - x)
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
