// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//
//
// 示例 1：
//
//
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
// 示例 2：
//
//
// 输入：nums = [0,3,7,2,5,8,4,6,0,1]
// 输出：9
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
//
//
// Related Topics 并查集 数组 哈希表 👍 1466 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)

/*
外层循环需要 O(n) 的时间复杂度，
只有当一个数是连续序列的第一个数的情况下才会进入内层循环，
然后在内层循环中匹配连续序列中的数，
因此数组中的每个数只会进入内层循环一次
*/
func longestConsecutive(nums []int) (ret int) {
	numSet := make(map[int]bool) // 去重
	for _, num := range nums {
		numSet[num] = true
	}
	for num := range numSet {
		if !numSet[num-1] { // 前一个数不存在，说明这个是第一个数
			cN := num
			nowLength := 1
			for numSet[cN+1] { // 向后判断是否存在
				cN++
				nowLength++
			}
			if ret < nowLength {
				ret = nowLength
			}
		}
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
