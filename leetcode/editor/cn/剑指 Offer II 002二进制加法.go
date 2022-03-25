//给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。
//
// 输入为 非空 字符串且只包含数字 1 和 0。
//
//
//
// 示例 1:
//
//
//输入: a = "11", b = "10"
//输出: "101"
//
// 示例 2:
//
//
//输入: a = "1010", b = "1011"
//输出: "10101"
//
//
//
// 提示：
//
//
// 每个字符串仅由字符 '0' 或 '1' 组成。
// 1 <= a.length, b.length <= 10^4
// 字符串如果不是 "0" ，就都不含前导零。
//
//
//
//
// 注意：本题与主站 67 题相同：https://leetcode-cn.com/problems/add-binary/
// Related Topics 位运算 数学 字符串 模拟 👍 24 👎 0

package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("11", "10"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) (ret string) {
	count := func(s string) (ret int) {
		for i := range s {
			ret = ret*2 + int(s[i]-'0')
		}
		return
	}
	sum := count(a) + count(b)
	if sum == 0 {
		return "0"
	}
	for sum > 0 {
		ret = string((sum%2)+'0') + ret
		sum /= 2
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
