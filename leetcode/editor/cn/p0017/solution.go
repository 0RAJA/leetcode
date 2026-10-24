// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//
//
// 示例 1：
//
//
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
//
// 示例 2：
//
//
// 输入：digits = "2"
// 输出：["a","b","c"]
//
//
//
//
// 提示：
//
//
// 1 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。
//
//
// Related Topics 哈希表 字符串 回溯 👍 3314 👎 0

package p0017

// leetcode submit region begin(Prohibit modification and deletion)
// 递归遍历，记录每个层级，每个层级进行对应数字的尝试
func letterCombinations(digits string) (res []string) {
	m := map[uint8]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var dfs func(idx int, path string)
	dfs = func(idx int, path string) {
		if idx == len(digits) {
			res = append(res, path)
			return
		}
		for _, c := range m[digits[idx]] {
			path += string(c)
			dfs(idx+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, "")
	return
}

// leetcode submit region end(Prohibit modification and deletion)
