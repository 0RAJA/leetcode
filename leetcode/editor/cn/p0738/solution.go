// 当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。
//
// 给定一个整数 n ，返回 小于或等于 n 的最大数字，且数字呈 单调递增 。
//
// 示例 1:
//
// 输入: n = 10
// 输出: 9
//
// 示例 2:
//
// 输入: n = 1234
// 输出: 1234
//
// 示例 3:
//
// 输入: n = 332
// 输出: 299
//
// 提示:
//
// 0 <= n <= 10⁹
//
// Related Topics 贪心 数学 👍 546 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
// 朴实无华：递减 n，每次判断 n 的每一位得出最后是否满足 -- 肯定超时
// 从低位到高位循环遍历 n ，如果当前位 y < x，那么 x=x-1, 剩余低位 = 9

// 获取数字 n 的每一位(反向) 1224 -> 4221
func num2List(n int) (res []int) {
	if n == 0 {
		return []int{0}
	}
	for ; n != 0; n /= 10 {
		res = append(res, n%10)
	}
	return res
}

// 将 list(反向) 转为 num 4221 -> 1224
func list2Num(digits []int) (res int) {
	for i := len(digits) - 1; i >= 0; i-- {
		res = res*10 + digits[i]
	}
	return
}

func monotoneIncreasingDigits(n int) int {
	// 将 n -> list 利于遍历每一位
	numDigitList := num2List(n)
	for {
		isOK := true
		// 遍历每一位判断是否符合预期（反向遍历，贪心，只要不满足就把高位-1，剩余低位整体调整为 9）
		for i := 0; i < len(numDigitList)-1; i++ {
			// 发现高位比低位大，不满足单调递增
			if numDigitList[i] < numDigitList[i+1] {
				// 高位 - 1
				numDigitList[i+1]--
				// 剩余低位调整为 9
				for j := 0; j <= i; j++ {
					numDigitList[j] = 9
				}
				isOK = false
				break
			}
		}
		// 所有位都满足
		if isOK {
			return list2Num(numDigitList)
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)
