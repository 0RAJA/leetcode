// 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数 。
//
// 你可以对一个单词进行如下三种操作：
//
//
// 插入一个字符
// 删除一个字符
// 替换一个字符
//
//
//
//
// 示例 1：
//
//
// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
//
//
// 示例 2：
//
//
// 输入：word1 = "intention", word2 = "execution"
// 输出：5
// 解释：
// intention -> inention (删除 't')
// inention -> enention (将 'i' 替换为 'e')
// enention -> exention (将 'n' 替换为 'x')
// exention -> exection (将 'n' 替换为 'c')
// exection -> execution (插入 'u')
//
//
//
//
// 提示：
//
//
// 0 <= word1.length, word2.length <= 500
// word1 和 word2 由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 3991 👎 0

package p0072

// leetcode submit region begin(Prohibit modification and deletion)
/*
其实这题可以这样考虑：比较两个字符串时，从后往前进行，每次只比较最后一位。
如果最后一位相同，就继续比较前一位；如果两位不同，再考虑做操作（插入、删除或替换），目的是尽可能让当前比较的这一位相同
针对最后一个不同的情况，有三种处理方式：
	1. 插入字符：在第一个字符串后加一个字符，使其与第二个字符串最后一位相同，然后再比较。
	2. 删除字符：直接删除第一个字符串的最后一个字符，再与第二个字符串对比。
	3. 替换字符：将第一个字符串的最后一个字符替换为与第二个字符串相同。

dp[i][j] = 第一个字符串的第 i 个位置和第二个字符串的这个位置对比，所需要的最少操作数。
dp[i][0] = i; dp[0][j] = j // 因为当一个字符串是 0 的时候，另一个字符串为了跟它相同，最少操作次数是全加/全减对应的字符数

从小到大遍历 i,j
if s1[i] == s2[j]:
	dp[i][j] = dp[i-1][j-1]
else:
	dp[i][j] = min(dp[i][j-1],dp[i-1][j],dp[i-1][j-1]) +1
*/
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

// leetcode submit region end(Prohibit modification and deletion)
