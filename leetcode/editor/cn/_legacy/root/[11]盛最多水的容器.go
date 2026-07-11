// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。 
//
// 返回容器可以储存的最大水量。 
//
// 说明：你不能倾斜容器。 
//
// 
//
// 示例 1： 
//
// 
//
// 
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
// 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//
// 示例 2： 
//
// 
// 输入：height = [1,1]
// 输出：1
// 
//
// 
//
// 提示： 
//
// 
// n == height.length 
// 2 <= n <= 10⁵ 
// 0 <= height[i] <= 10⁴ 
// 
//
// Related Topics 贪心 数组 双指针 👍 6005 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
// 11 盛最多水的容器
// 双指针，从两边开始遍历，每次移动较短的那根线，因为移动较长的那根线，面积只会更小
// 如果移动较短的那根线，面积可能会更大
// 所以每次移动较短的那根线，直到两根线相遇
func maxArea(height []int) (maxArea int) {
	maxArea = -1
	for i, j := 0, len(height)-1; i < j; {
		maxArea = max(maxArea, min(height[i], height[j])*(j-i))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return maxArea
}

// leetcode submit region end(Prohibit modification and deletion)
