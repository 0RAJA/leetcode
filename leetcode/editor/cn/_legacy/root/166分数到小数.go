//给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以 字符串形式返回小数 。
//
// 如果小数部分为循环小数，则将循环的部分括在括号内。
//
// 如果存在多个答案，只需返回 任意一个 。
//
// 对于所有给定的输入，保证 答案字符串的长度小于 10⁴ 。
//
//
//
// 示例 1：
//
//
//输入：numerator = 1, denominator = 2
//输出："0.5"
//
//
// 示例 2：
//
//
//输入：numerator = 2, denominator = 1
//输出："2"
//
//
// 示例 3：
//
//
//输入：numerator = 2, denominator = 3
//输出："0.(6)"
//
//
// 示例 4：
//
//
//输入：numerator = 4, denominator = 333
//输出："0.(012)"
//
//
// 示例 5：
//
//
//输入：numerator = 1, denominator = 5
//输出："0.2"
//
//
//
//
// 提示：
//
//
// -2³¹ <= numerator, denominator <= 2³¹ - 1
// denominator != 0
//
// Related Topics 哈希表 数学 字符串 👍 306 👎 0

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(fractionToDecimal(1, -6))
}

// leetcode submit region begin(Prohibit modification and deletion)
func fractionToDecimal(numerator int, denominator int) (s string) {
	type Num struct {
		num     int
		isStart bool
	}
	if numerator*denominator < 0 { //加小数点
		s += "-"
	}
	numerator = int(math.Abs(float64(numerator))) //取绝对值
	denominator = int(math.Abs(float64(denominator)))
	mapStr := make(map[string]*Num)
	integer := make([]int, 0)
	decimal := make([]*Num, 0)
	if numerator < denominator {
		s += "0"
	}
	for numerator > denominator {
		integer = append(integer, numerator/denominator)
		numerator %= denominator
	}
	for numerator *= 10; numerator != 0; numerator *= 10 {
		k := strconv.Itoa(numerator) + "/" + strconv.Itoa(denominator)
		if _, ok := mapStr[k]; ok {
			mapStr[k].isStart = true
			break
		}
		num := &Num{
			num:     numerator / denominator,
			isStart: false,
		}
		mapStr[k] = num
		decimal = append(decimal, num)
		numerator %= denominator
	}
	for i := 0; i < len(integer); i++ {
		s += strconv.Itoa(integer[i])
	}
	if len(decimal) > 0 {
		s += "."
	}
	isOk := false
	for i := 0; i < len(decimal); i++ {
		if decimal[i].isStart {
			isOk = true
			s += "("
		}
		s += strconv.Itoa(decimal[i].num)
	}
	if isOk {
		s += ")"
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
