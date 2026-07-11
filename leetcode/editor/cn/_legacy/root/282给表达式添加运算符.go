//给定一个仅包含数字 0-9 的字符串 num 和一个目标值整数 target ，在 num 的数字之间添加 二元 运算符（不是一元）+、- 或 * ，返回所
//有能够得到目标值的表达式。
//
//
//
// 示例 1:
//
//
//输入: num = "123", target = 6
//输出: ["1+2+3", "1*2*3"]
//
//
// 示例 2:
//
//
//输入: num = "232", target = 8
//输出: ["2*3+2", "2+3*2"]
//
// 示例 3:
//
//
//输入: num = "105", target = 5
//输出: ["1*0+5","10-5"]
//
// 示例 4:
//
//
//输入: num = "00", target = 0
//输出: ["0+0", "0-0", "0*0"]
//
//
// 示例 5:
//
//
//输入: num = "3456237490", target = 9191
//输出: []
//
//
//
// 提示：
//
//
// 1 <= num.length <= 10
// num 仅含数字
// -2³¹ <= target <= 2³¹ - 1
//
// Related Topics 数学 字符串 回溯 👍 316 👎 0

package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	num := "105"
	fmt.Println(addOperators(num, 5))
}

// leetcode submit region begin(Prohibit modification and deletion)
var Priority = map[int]int{
	'+': 0,
	'-': 0,
	'*': 1,
	'/': 1,
}

type (
	Stack struct {
		stack []Num
	}

	Num struct {
		val   int
		isNum bool
	}
)

func addOperators(num string, target int) (ret []string) {
	next := []string{"+", "-", "*"}
	var dfs func(str string, index int)
	dfs = func(str string, index int) {
		if v, err := CalculateTheSuffix(InfixToSuffix(ToNum(str))); err == nil && v == target {
			ret = append(ret, str)
			return
		}
		if index >= len(str) {
			return
		}
		for i := 0; i < 3; i++ {
			dfs(str[:index]+next[i]+str[index:], index+2)
		}
	}
	dfs(num, 1)
	return
}

func ToNum(str string) (num []Num) {
	tempIndex := 0
	for i := 1; i < len(str); i++ {
		if !IsNum(str[i]) {
			temp, err := strconv.Atoi(str[tempIndex:i])
			if err != nil {
				return
			}
			p := Num{
				val:   temp,
				isNum: true,
			}
			num = append(num, p, Num{
				val:   int(str[i]),
				isNum: false,
			})
			tempIndex = i + 1
		}
	}
	if tempIndex == 0 {
		temp, err := strconv.Atoi(str)
		if err != nil {
			return
		}
		p := Num{
			val:   temp,
			isNum: true,
		}
		num = append(num, p)
	}
	return
}

func IsNum(c byte) bool {
	return c >= '0' && c <= '9'
}

// InfixToSuffix 中缀转后缀
func InfixToSuffix(s []Num) (ret []Num) {
	charStack := CreatStack() //字符栈
	for i := 0; i < len(s); i++ {
		if s[i].IsNum() {
			ret = append(ret, s[i])
		} else {
			switch s[i].Val() {
			case '(':
				charStack.Push(s[i])
			case ')':
				for !charStack.IsEmpty() {
					ch := charStack.Pop()
					if ch.Val() == '(' {
						break
					}
					ret = append(ret, ch)
				}
			default:
				for !charStack.IsEmpty() {
					ch := charStack.Pop()
					if ch.Val() != '(' && Priority[s[i].Val()] <= Priority[ch.Val()] {
						ret = append(ret, ch)
					} else {
						charStack.Push(ch)
						break
					}
				}
				charStack.Push(s[i])
			}
		}
	}
	for !charStack.IsEmpty() {
		ret = append(ret, charStack.Pop())
	}
	return
}

func (n Num) IsNum() bool {
	return n.isNum
}

func (n Num) Val() int {
	return n.val
}

func CalculateTheSuffix(s []Num) (int, error) {
	numStack := CreatStack()
	for i := 0; i < len(s); i++ {
		if s[i].IsNum() {
			numStack.Push(s[i])
		} else {
			if numStack.Len() < 2 {
				return 0, errors.New("err")
			}
			right := numStack.Pop()
			left := numStack.Pop()
			numStack.Push(Calculate(left.Val(), right.Val(), s[i].Val()))
		}
	}
	if numStack.Len() > 1 {
		return 0, errors.New("err")
	}
	return numStack.Pop().Val(), nil
}

func Calculate(left, right, operate int) Num {
	num := Num{
		isNum: true,
	}
	switch operate {
	case '+':
		num.val = left + right
	case '-':
		num.val = left - right
	case '*':
		num.val = left * right
	case '/':
		num.val = left / right
	}
	return num
}

func CreatStack() *Stack {
	return &Stack{stack: make([]Num, 0)}
}

func (s *Stack) Push(ch Num) {
	s.stack = append(s.stack, ch)
}

func (s *Stack) Pop() (ch Num) {
	ch = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack) Len() int {
	return len(s.stack)
}

//leetcode submit region end(Prohibit modification and deletion)
