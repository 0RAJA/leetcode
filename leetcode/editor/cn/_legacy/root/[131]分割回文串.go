// 给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
// 
//
// 示例 1： 
//
// 
// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
// 
//
// 示例 2： 
//
// 
// 输入：s = "a"
// 输出：[["a"]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 16 
// s 仅由小写英文字母组成 
// 
//
// Related Topics 字符串 动态规划 回溯 👍 2222 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// 切割问题，每次切割出一个回文子串，然后拿剩下的再去切割
func partition(s string) (ret [][]string) {
	var dfs func(path []string, idx int)
	dfs = func(path []string, idx int) {
		if idx >= len(s) {
			ret = append(ret, path)
			return
		}
		cur := ""
		for i := idx; i < len(s); i++ {
			cur += string(s[i])
			if isPalindrome(cur) {
				dfs(append(append(make([]string, 0, len(path)+1), path...), cur), i+1)
			}
		}
	}
	dfs([]string{}, 0)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
