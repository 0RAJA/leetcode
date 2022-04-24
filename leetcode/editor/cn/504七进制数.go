//给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。
//
//
//
// 示例 1:
//
//
//输入: num = 100
//输出: "202"
//
//
// 示例 2:
//
//
//输入: num = -7
//输出: "-10"
//
//
//
//
// 提示：
//
//
// -10⁷ <= num <= 10⁷
//
// Related Topics 数学 👍 123 👎 0

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(convertToBase7(-8))
}

//leetcode submit region begin(Prohibit modification and deletion)
func convertToBase7(num int) (ret string) {
	if num == 0 {
		return "0"
	}
	judge := false
	if num < 0 {
		judge = true
		num = -num
	}
	var dfs func(int) string
	dfs = func(num int) string {
		if num == 0 {
			return ""
		}
		return dfs(num/7) + strconv.Itoa(num%7)
	}
	ret = dfs(num)
	if judge {
		return "-" + ret
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
