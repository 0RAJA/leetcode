package main

import (
	"fmt"
)

/*
从起点开始接下来有 100 个方块，相邻方块间的距离都为 1，每个方块上有增加体力的食用蘑菇或减少体力的毒蘑菇，蘑菇带来的体力改变是已知的。
一个人初始体力为 m，每次可以往前跳任意个方块，体力耗尽就会死掉。

每跳一次消耗的体力与跳的距离成正比，比例为1。问这个人能否跳到终点，如果能，求可能剩余的最大体力。
每跳一次消耗的体力是跳的距离的平方。问这个人能否跳到终点，如果能，求可能剩余的最大体力。
每跳一次消耗的体力是跳的距离的平方，每一次跳跃得一分，求到达终点的最大分数。

链接：https://leetcode.cn/circle/discuss/bMmxI9/
*/

func bMmxI9(nums []int, m int) (int, bool) {
	end := m
	for now := 0; now < len(nums) && now < end; now++ {
		if nums[now] > 0 {
			end += nums[now]
		}
	}
	return end - len(nums), end-len(nums) >= 0
}

func main() {
	fmt.Println(bMmxI9([]int{1, 2, 1, 2, 1, 2, 1, 1, 2}, 3))
}
