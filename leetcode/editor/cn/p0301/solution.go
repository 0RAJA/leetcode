// 给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。
//
// 返回所有可能的结果。答案可以按 任意顺序 返回。
//
//
//
// 示例 1：
//
//
// 输入：s = "()())()"
// 输出：["(())()","()()()"]
//
//
// 示例 2：
//
//
// 输入：s = "(a)())()"
// 输出：["(a())()","(a)()()"]
//
//
// 示例 3：
//
//
// 输入：s = ")("
// 输出：[""]
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 25
// s 由小写英文字母以及括号 '(' 和 ')' 组成
// s 中至多含 20 个括号
//
//
// Related Topics 广度优先搜索 字符串 回溯 👍 1023 👎 0

package p0301

// leetcode submit region begin(Prohibit modification and deletion)
// 删除最小数量的无效括号 考虑 BFS 依次遍历删除第 0-N 个括号下的所有可能，如果在第 i 个下满足了，那就返回该场景下的所有满足的可能即可
//   - 一个校验 s 是否是有效括号的函数：遍历括号，左括号+1，右括号-1，等于 0 时即符合条件
//   - bfs：维护需要访问的字符串的队列，里面表示已经删除了括号需要再次遍历的字符串；每次遍历 队列中所有的字符串，满足则返回当前场景下所有，不满足则删除一个括号加入到队列
//   - 还需要一个去重 map：注意这个需要记录待探索的字符串，避免出现大量重复探索的字符串
func removeInvalidParentheses(s string) (res []string) {
	isValid := func(s string) bool {
		cnt := 0
		for _, c := range s {
			if c == '(' {
				cnt++
			} else if c == ')' {
				cnt--
			}
			// 右括号太多了
			if cnt < 0 {
				return false
			}
		}
		return cnt == 0
	}
	visited := make(map[string]bool)
	queue := []string{s}
	for len(queue) > 0 {
		// 说明上一轮已经找到了符合条件的字符串，就不需要接着找删除更多括号的字符串了
		if len(res) > 0 {
			break
		}
		size := len(queue)
		for i := 0; i < size; i++ {
			headStr := queue[0]
			queue = queue[1:]
			// 当前字符串满足了，添加到结果后直接继续比较其他同类字符串
			if isValid(headStr) {
				res = append(res, headStr)
				continue
			}
			// 不符合则再次裁剪一个括号放到后续扫描中
			for idx, c := range headStr {
				if c != '(' && c != ')' {
					continue
				}
				next := headStr[:idx] + headStr[idx+1:]
				if visited[next] {
					continue
				}
				queue = append(queue, next)
				visited[next] = true
			}
		}
	}
	if len(res) == 0 {
		return []string{""}
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
