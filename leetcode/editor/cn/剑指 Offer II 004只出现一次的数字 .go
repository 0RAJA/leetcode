//给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,2,3,2]
//输出：3
//
//
// 示例 2：
//
//
//输入：nums = [0,1,0,1,0,1,100]
//输出：100
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
// nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
//
//
//
//
// 进阶：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
//
//
// 注意：本题与主站 137 题相同：https://leetcode-cn.com/problems/single-number-ii/
// Related Topics 位运算 数组 👍 54 👎 0

package main

import "fmt"

func main() {
	fmt.Println(singleNumber2([]int{-1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func singleNumber2(nums []int) int {
	sums := [32]int{}
	for _, v := range nums {
		bit := 0
		for v>>bit != 0 && bit < 32 {
			sums[bit] += (v >> bit) & 1
			bit++
		}
	}
	var key int32
	for bit, cnt := range sums {
		key += int32(cnt%3) << bit
	}
	return int(key)
}

//leetcode submit region end(Prohibit modification and deletion)
