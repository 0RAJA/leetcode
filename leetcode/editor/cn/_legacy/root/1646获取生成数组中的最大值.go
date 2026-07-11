// 给你一个整数 n 。按下述规则生成一个长度为 n + 1 的数组 nums ：
//
// nums[0] = 0
// nums[1] = 1
// 当 2 <= 2 * i <= n 时，nums[2 * i] = nums[i]
// 当 2 <= 2 * i + 1 <= n 时，nums[2 * i + 1] = nums[i] + nums[i + 1]
//
// 返回生成数组 nums 中的 最大 值。
//
// 示例 1：
//
// 输入：n = 7
// 输出：3
// 解释：根据规则：
//
//	nums[0] = 0
//	nums[1] = 1
//	nums[(1 * 2) = 2] = nums[1] = 1
//	nums[(1 * 2) + 1 = 3] = nums[1] + nums[2] = 1 + 1 = 2
//	nums[(2 * 2) = 4] = nums[2] = 1
//	nums[(2 * 2) + 1 = 5] = nums[2] + nums[3] = 1 + 2 = 3
//	nums[(3 * 2) = 6] = nums[3] = 2
//	nums[(3 * 2) + 1 = 7] = nums[3] + nums[4] = 2 + 1 = 3
//
// 因此，nums = [0,1,1,2,1,3,2,3]，最大值 3
//
// 示例 2：
//
// 输入：n = 2
// 输出：1
// 解释：根据规则，nums[0]、nums[1] 和 nums[2] 之中的最大值是 1
//
// 示例 3：
//
// 输入：n = 3
// 输出：2
// 解释：根据规则，nums[0]、nums[1]、nums[2] 和 nums[3] 之中的最大值是 2
//
// 提示：
//
// 0 <= n <= 100
//
// Related Topics 数组 动态规划 模拟 👍 40 👎 0
package main

import "fmt"

func main() {
	fmt.Println(getMaximumGenerated(100))
}

// leetcode submit region begin(Prohibit modification and deletion)
func getMaximumGenerated(n int) int {
	num := make([]int, 101)
	num[0], num[1] = 0, 1
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	max := 0
	for i := 1; ; i++ {
		if 2*i >= 2 && 2*i <= n {
			num[2*i] = num[i]
			if max < num[2*i] {
				max = num[2*i]
			}
		}
		if 2*i+1 >= 2 && 2*i+1 <= n {
			num[2*i+1] = num[i] + num[i+1]
			if max < num[2*i+1] {
				max = num[2*i+1]
			}
		}
		if 2*i+1 > n {
			break
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
