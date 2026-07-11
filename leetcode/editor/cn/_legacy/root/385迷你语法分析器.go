//给定一个字符串 s 表示一个整数嵌套列表，实现一个解析它的语法分析器并返回解析的结果 NestedInteger 。
//
// 列表中的每个元素只可能是整数或整数嵌套列表
//
//
//
// 示例 1：
//
//
//输入：s = "324",
//输出：324
//解释：你应该返回一个 NestedInteger 对象，其中只包含整数值 324。
//
//
// 示例 2：
//
//
//输入：s = "[123,[456,[789]]]",
//输出：[123,[456,[789]]]
//解释：返回一个 NestedInteger 对象包含一个有两个元素的嵌套列表：
//1. 一个 integer 包含值 123
//2. 一个包含两个元素的嵌套列表：
//    i.  一个 integer 包含值 456
//    ii. 一个包含一个元素的嵌套列表
//         a. 一个 integer 包含值 789
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 5 * 10⁴
// s 由数字、方括号 "[]"、负号 '-' 、逗号 ','组成
// 用例保证 s 是可解析的 NestedInteger
// 输入中的所有值的范围是 [-10⁶, 10⁶]
//
// Related Topics 栈 深度优先搜索 字符串 👍 134 👎 0

package main

import "strconv"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)

// NestedInteger This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// IsInteger Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool { return false }

// GetInteger Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int { return 0 }

// SetInteger Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Add Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {}

// GetList Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger { return nil }

func deserialize(s string) *NestedInteger {
	if s[0] != '[' { //直接给你个整数
		nest := new(NestedInteger)
		num, _ := strconv.Atoi(s)
		nest.SetInteger(num)
		return nest
	}
	var dfs func(idx int) (*NestedInteger, int)
	dfs = func(idx int) (*NestedInteger, int) {
		nest := new(NestedInteger)
		num := 0
		isOK := false
		isN := false
		for idx++; idx < len(s); {
			switch s[idx] {
			case ',', ']': //,和]作为分割都需要把前面的整数提交，只不过]会终止
				if isOK {
					if isN {
						num *= -1
						isN = false
					}
					x := new(NestedInteger)
					x.SetInteger(num)
					nest.Add(*x)
					isOK = false
					num = 0
				}
				idx++
				if s[idx-1] == ']' {
					return nest, idx
				}
			case '[': //下一层嵌套
				result, nx := dfs(idx)
				nest.Add(*result)
				idx = nx
			default: //数字
				isOK = true
				if s[idx] == '-' {
					isN = !isN
				} else {
					num = num*10 + int(s[idx]-'0')
				}
				idx++
			}
		}
		return nest, idx
	}
	result, _ := dfs(0)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
