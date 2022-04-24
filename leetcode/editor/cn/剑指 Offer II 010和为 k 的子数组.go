//给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数。
//
//
//
// 示例 1 :
//
//
//输入:nums = [1,1,1], k = 2
//输出: 2
//解释: 此题 [1,1] 与 [1,1] 为两种不同的情况
//
//
// 示例 2 :
//
//
//输入:nums = [1,2,3], k = 3
//输出: 2
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 2 * 10⁴
// -1000 <= nums[i] <= 1000
//
// -10⁷ <= k <= 10⁷
//
//
//
//
//
// 注意：本题与主站 560 题相同： https://leetcode-cn.com/problems/subarray-sum-equals-k/
// Related Topics 数组 哈希表 前缀和 👍 52 👎 0

package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}

//leetcode submit region begin(Prohibit modification and deletion)
func subarraySum(nums []int, k int) (ret int) {
	m := make(map[int]int) //记录之前的前缀和出现的次数
	m[0] = 1               //初始化0为1
	suffix := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		suffix[i] = suffix[i-1] + nums[i-1]
		ret += m[suffix[i]-k] //看看之前有没有出现过suffix[i]-k，有就说明中间这段的和为k
		m[suffix[i]]++
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
