// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 示例 1：
//
//
//
//
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
//
// 示例 2：
//
//
// 输入：height = [4,2,0,3,2,5]
// 输出：9
//
//
//
//
// 提示：
//
//
// n == height.length
// 1 <= n <= 2 * 10⁴
// 0 <= height[i] <= 10⁵
//
//
// Related Topics 栈 数组 双指针 动态规划 单调栈 👍 6357 👎 0

package p0042

// leetcode submit region begin(Prohibit modification and deletion)
// 双指针：左右分别往里面遍历计算，两个重点
// 1. 什么时候算雨水：只有当高度相比于最高高度度降低的时候计算，如果高度升高是不会有雨水的
// 2. 用于计算的最高度怎么算：左右两边维护各自的当前最高高度，取其中的较小的高度来算高度
func trap(height []int) (water int) {
	maxL, maxR := 0, 0 // 左右两边当前最高高度
	for l, r := 0, len(height)-1; l < r; {
		maxL = max(maxL, height[l])
		maxR = max(maxR, height[r])
		if maxL < maxR {
			// 降低了就要算雨水了
			water += maxL - height[l]
			l++
		} else {
			water += maxR - height[r]
			r--
		}
	}
	return water
}

// leetcode submit region end(Prohibit modification and deletion)
