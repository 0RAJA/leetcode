package main

import (
	"sort"
)

// 和子集的处理很像，但是关键点是避免重复的子集，先想下为啥会重复 1222 按之前的解法，会出现两个 (2,2)，其实就是在于本层相同的处理的值处理一次需要跳过，那么先排序
func subsetsWithDup(nums []int) (res [][]int) {
	sort.Ints(nums)
	var dfs func(path []int, idx int)
	dfs = func(path []int, idx int) {
		if idx == len(nums) {
			return
		}
		for i := idx; i < len(nums); i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			next := append(append(make([]int, 0, len(path)+1), path...), nums[i])
			res = append(res, next)
			dfs(next, i+1)
		}
	}
	res = append(res, []int{})
	dfs([]int{}, 0)
	return
}
