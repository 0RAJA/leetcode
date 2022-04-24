//给你一个字符串 text，你需要使用 text 中的字母来拼凑尽可能多的单词 "balloon"（气球）。
//
// 字符串 text 中的每个字母最多只能被使用一次。请你返回最多可以拼凑出多少个单词 "balloon"。
//
//
//
// 示例 1：
//
//
//
// 输入：text = "nlaebolko"
//输出：1
//
//
// 示例 2：
//
//
//
// 输入：text = "loonbalxballpoon"
//输出：2
//
//
// 示例 3：
//
// 输入：text = "leetcode"
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= text.length <= 10^4
// text 全部由小写英文字母组成
//
// Related Topics 哈希表 字符串 计数 👍 101 👎 0

package main

import "fmt"

func main() {
	text := "balon"
	fmt.Println(maxNumberOfBalloons(text))
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxNumberOfBalloons(text string) (ret int) {
	ret = 1e4
	alphabet := "balloon"
	n := [26]int{}
	for i := range alphabet {
		n[alphabet[i]-'a']++
	}
	m := [26]int{}
	for i := range text {
		m[text[i]-'a']++
	}
	for i := range alphabet {
		if p := m[alphabet[i]-'a'] / n[alphabet[i]-'a']; ret > p {
			ret = p
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
