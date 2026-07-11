// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足： 
//
// 
// 左括号必须用相同类型的右括号闭合。 
// 左括号必须以正确的顺序闭合。 
// 每个右括号都有一个对应的相同类型的左括号。 
// 
//
// 
//
// 示例 1： 
//
// 
// 输入：s = "()" 
// 
//
// 输出：true 
//
// 示例 2： 
//
// 
// 输入：s = "()[]{}" 
// 
//
// 输出：true 
//
// 示例 3： 
//
// 
// 输入：s = "(]" 
// 
//
// 输出：false 
//
// 示例 4： 
//
// 
// 输入：s = "([])" 
// 
//
// 输出：true 
//
// 示例 5： 
//
// 
// 输入：s = "([)]" 
// 
//
// 输出：false 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 10⁴ 
// s 仅由括号 '()[]{}' 组成 
// 
//
// Related Topics 栈 字符串 👍 4959 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	leftMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	rightMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]rune, 0, len(s)/2+1)
	// 遍历 s
	for _, c := range s {
		// 左边入栈
		if _, ok := leftMap[c]; ok {
			stack = append(stack, c)
		}
		// 右边出栈进行匹配
		if matchChar, ok := rightMap[c]; ok {
			if len(stack) == 0 {
				return false
			}
			leftChar := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if leftChar != matchChar {
				return false
			}
		}
	}
	// 栈被清空才说明都匹配到了
	return len(stack) == 0
}

// leetcode submit region end(Prohibit modification and deletion)
