//给定一个二进制数组， 计算其中最大连续 1 的个数。
//
//
//
// 示例：
//
//
//输入：[1,1,0,1,1,1]
//输出：3
//解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.
//
//
//
//
// 提示：
//
//
// 输入的数组只包含 0 和 1 。
// 输入数组的长度是正整数，且不超过 10,000。
//
// Related Topics 数组 👍 276 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findMaxConsecutiveOnes(nums []int) (ret int) {
	cnt := 0
	for i := range nums {
		if nums[i] == 1 {
			cnt++
		} else {
			if ret < cnt {
				ret = cnt
			}
			cnt = 0
		}
	}
	if ret < cnt {
		ret = cnt
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
