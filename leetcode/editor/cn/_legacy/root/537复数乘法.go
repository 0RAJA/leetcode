//复数 可以用字符串表示，遵循 "实部+虚部i" 的形式，并满足下述条件：
//
//
// 实部 是一个整数，取值范围是 [-100, 100]
// 虚部 也是一个整数，取值范围是 [-100, 100]
// i² == -1
//
//
// 给你两个字符串表示的复数 num1 和 num2 ，请你遵循复数表示形式，返回表示它们乘积的字符串。
//
//
//
// 示例 1：
//
//
//输入：num1 = "1+1i", num2 = "1+1i"
//输出："0+2i"
//解释：(1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i ，你需要将它转换为 0+2i 的形式。
//
//
// 示例 2：
//
//
//输入：num1 = "1+-1i", num2 = "1+-1i"
//输出："0+-2i"
//解释：(1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i ，你需要将它转换为 0+-2i 的形式。
//
//
//
//
// 提示：
//
//
// num1 和 num2 都是有效的复数表示。
//
// Related Topics 数学 字符串 模拟 👍 76 👎 0

package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := "1+1i"
	num2 := "1+1i"
	fmt.Println(complexNumberMultiply1(num1, num2))
}

// leetcode submit region begin(Prohibit modification and deletion)
func complexNumberMultiply(num1 string, num2 string) string {
	a, _ := strconv.ParseComplex(num1, 10)
	b, _ := strconv.ParseComplex(num2, 10)
	c := a * b
	x, y := real(c), imag(c)
	return strconv.FormatFloat(x, 'f', -1, 64) + "+" + strconv.FormatFloat(y, 'f', -1, 64) + "i"
}
func complexNumberMultiply1(num1 string, num2 string) string {
	var r1, i1, r2, i2 float64
	_, _ = fmt.Sscanf(num1, "%f+%f", &r1, &i1)
	_, _ = fmt.Sscanf(num2, "%f+%f", &r2, &i2)
	return fmt.Sprintf("%.0f+%.0fi", r1*r2+i1*(-i2), r1*i2+i1*r2)
}

//leetcode submit region end(Prohibit modification and deletion)
