// 两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。
//
// 给你两个整数 x 和 y，计算并返回它们之间的汉明距离。
//
//
//
// 示例 1：
//
//
// 输入：x = 1, y = 4
// 输出：2
// 解释：
// 1   (0 0 0 1)
// 4   (0 1 0 0)
//       ↑   ↑
// 上面的箭头指出了对应二进制位不同的位置。
//
//
// 示例 2：
//
//
// 输入：x = 3, y = 1
// 输出：1
//
//
//
//
// 提示：
//
//
// 0 <= x, y <= 2³¹ - 1
//
//
//
//
// 注意：本题与 2220. 转换数字的最少位翻转次数 相同。
//
// Related Topics 位运算 👍 787 👎 0

package p0461

// leetcode submit region begin(Prohibit modification and deletion)
// 异或（相同为0，不同为1），然后算1的个数(与1取and)
func hammingDistance(x int, y int) (res int) {
	for target := x ^ y; target != 0; target >>= 1 {
		if target&1 == 1 {
			res++
		}
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
