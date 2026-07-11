//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//
//
// 示例 1:
//
//
//
//
//输入：heights = [2,1,5,6,2,3]
//输出：10
//解释：最大的矩形为图中红色区域，面积为 10
//
//
// 示例 2：
//
//
//
//
//输入： heights = [2,4]
//输出： 4
//
//
//
// 提示：
//
//
// 1 <= heights.length <=10⁵
// 0 <= heights[i] <= 10⁴
//
// Related Topics 栈 数组 单调栈 👍 1882 👎 0

package main

import "fmt"

func main() {
	fmt.Println(largestRectangleArea([]int{1, 1}))
}

// leetcode submit region begin(Prohibit modification and deletion)
func largestRectangleArea(heights []int) (ret int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	heights = append(heights, 0)
	stack := make([]int, 0, len(heights))
	stack = append(stack, 0)
	for i := range heights {
		ret = max(ret, heights[i])
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			last := len(stack) - 1
			ret = max(ret, (i-last+1)*heights[stack[last-1]])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
