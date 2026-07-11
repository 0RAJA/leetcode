// 编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
//
// 示例 1：
//
// 输入："hello"
// 输出："holle"
//
// 示例 2：
//
// 输入："leetcode"
// 输出："leotcede"
//
// 提示：
//
// 元音字母不包含字母 "y" 。
//
// Related Topics 双指针 字符串 👍 177 👎 0
package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func reverseVowels(s string) string {
	v := []byte(s)
	for i, j := 0, len(v)-1; i < j; {
		switch v[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		Loop:
			for {
				switch v[j] {
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					v[i], v[j] = v[j], v[i]
					break Loop
				default:
					j--
				}
			}
			i++
			j--
		default:
			i++
		}
	}
	return string(v)
}

//leetcode submit region end(Prohibit modification and deletion)
