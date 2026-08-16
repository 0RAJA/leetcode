// 给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
//
// 回文字符串 是正着读和倒过来读一样的字符串。
//
// 子字符串 是字符串中的由连续字符组成的一个序列。
//
//
//
// 示例 1：
//
//
// 输入：s = "abc"
// 输出：3
// 解释：三个回文子串: "a", "b", "c"
//
//
// 示例 2：
//
//
// 输入：s = "aaa"
// 输出：6
// 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 由小写英文字母组成
//
//
// Related Topics 双指针 字符串 动态规划 👍 1525 👎 0

package p0647

// leetcode submit region begin(Prohibit modification and deletion)
// 每个回文串有两种算法，奇数(i,i)、偶数(i,i+1)，每个从中间往两边拓展计数
func countSubstrings(s string) (res int) {
	countCycleStr := func(i, j int) {
		for i >= 0 && j < len(s) {
			if s[i] != s[j] {
				break
			}
			i--
			j++
			res++
		}
	}
	for i := 0; i < len(s); i++ {
		countCycleStr(i, i)
		if i+1 != len(s) {
			countCycleStr(i, i+1)
		}
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
