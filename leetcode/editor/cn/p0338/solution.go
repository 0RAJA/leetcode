// 给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。
//
//
// 不要使用内置函数来解决（例如，C++ 中的 __builtin_popcount）。
//
//
//
//
// 示例 1：
//
//
//
// 输入：n = 2
// 输出：[0,1,1]
// 解释：
// 0 --> 0
// 1 --> 1
// 2 --> 10
//
//
// 示例 2：
//
//
// 输入：n = 5
// 输出：[0,1,1,2,1,2]
// 解释：
// 0 --> 0
// 1 --> 1
// 2 --> 10
// 3 --> 11
// 4 --> 100
// 5 --> 101
//
//
//
//
// 提示：
//
//
// 0 <= n <= 10⁵
//
//
//
//
// 进阶：
//
//
// 很容易就能实现时间复杂度为 O(n log n) 的解决方案，你可以在线性时间复杂度 O(n) 内用一趟扫描解决此问题吗？
//
//
//
// Related Topics 位运算 动态规划 👍 1442 👎 0

package p0338

// leetcode submit region begin(Prohibit modification and deletion)
// f(k) = f(k >> 1) + k & 1
// 数字 k 中 1 的个数 = k 右移一位的数字中 1 的个数 + k 最后一位是否为 1
// 6 = 110
// 5 = 101
// 4 = 100
// 3 = 011
// 2 = 010
// 1 = 001
func countBits(n int) (res []int) {
	res = make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i>>1] + i&1
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
