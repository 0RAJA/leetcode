//给定一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [3,10,5,25,2,8]
//输出：28
//解释：最大运算结果是 5 XOR 25 = 28.
//
// 示例 2：
//
//
//输入：nums = [0]
//输出：0
//
//
// 示例 3：
//
//
//输入：nums = [2,4]
//输出：6
//
//
// 示例 4：
//
//
//输入：nums = [8,10,2]
//输出：10
//
//
// 示例 5：
//
//
//输入：nums = [14,70,53,83,49,91,36,80,92,51,66,70]
//输出：127
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2 * 10⁴
// 0 <= nums[i] <= 2³¹ - 1
//
//
//
//
//
//
// 进阶：你可以在 O(n) 的时间解决这个问题吗？
//
//
//
// 注意：本题与主站 421 题相同： https://leetcode-cn.com/problems/maximum-xor-of-two-
//numbers-in-an-array/
// Related Topics 位运算 字典树 数组 哈希表 👍 30 👎 0

package main

import "fmt"

func main() {
	fmt.Println(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
}

//leetcode submit region begin(Prohibit modification and deletion)
/*
	ret ^ a = b
	其实就是从高位往低位走，高位优先为1,每次先让ret<<1+1再遍历a，如果存在ret ^ a = b则此时这个ret此位为1,否则为0
*/
func findMaximumXOR(nums []int) (ret int) {
	maxbits := 30
	for i := maxbits; i >= 0; i-- {
		seen := make(map[int]bool)
		for _, num := range nums {
			seen[num>>i] = true
		}
		ret = ret<<1 + 1
		for _, num := range nums {
			if seen[(num>>i)^ret] {
				goto over
			}
		}
		ret--
	over:
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
