// 字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
//
//
//
// 示例 1:
//
// 输入:
// first = "pale"
// second = "ple"
// 输出: True
//
//
//
// 示例 2:
//
// 输入:
// first = "pales"
// second = "pal"
// 输出: False
//
// Related Topics 双指针 字符串 👍 146 👎 0

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(oneEditAway("islander", "slander"))
}

// leetcode submit region begin(Prohibit modification and deletion)
func oneEditAway(first string, second string) bool {
	first = strings.ToLower(first)
	second = strings.ToLower(second)
	if first == second {
		return true
	}
	if len(first) > len(second) {
		first, second = second, first
	}
	if len(second)-len(first) > 1 {
		return false
	}
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			if len(first) == len(second) {
				return first[i+1:] == second[i+1:]
			}
			return first[i:] == second[i+1:]
		}
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)
