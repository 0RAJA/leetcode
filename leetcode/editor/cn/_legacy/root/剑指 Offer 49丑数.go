//我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
//
//
//
// 示例:
//
// 输入: n = 10
//输出: 12
//解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
//
// 说明:
//
//
// 1 是丑数。
// n 不超过1690。
//
//
// 注意：本题与主站 264 题相同：https://leetcode-cn.com/problems/ugly-number-ii/
// Related Topics 哈希表 数学 动态规划 堆（优先队列） 👍 259 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	nums := make([]int, 1691)
	nums[1] = 1
	num2, num3, num5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		t := min(nums[num2]*2, min(nums[num3]*3, nums[num5]*5))
		if t == nums[num2]*2 {
			num2++
		}
		if t == nums[num3]*3 {
			num3++
		}
		if t == nums[num5]*5 {
			num5++
		}
		nums[i] = t
	}
	return nums[n]
}

//leetcode submit region end(Prohibit modification and deletion)
