//可以用字符串表示一个学生的出勤记录，其中的每个字符用来标记当天的出勤情况（缺勤、迟到、到场）。记录中只含下面三种字符：
//
// 'A'：Absent，缺勤
// 'L'：Late，迟到
// 'P'：Present，到场
//
//
// 如果学生能够 同时 满足下面两个条件，则可以获得出勤奖励：
//
//
// 按 总出勤 计，学生缺勤（'A'）严格 少于两天。
// 学生 不会 存在 连续 3 天或 连续 3 天以上的迟到（'L'）记录。
//
//
// 给你一个整数 n ，表示出勤记录的长度（次数）。请你返回记录长度为 n 时，可能获得出勤奖励的记录情况 数量 。答案可能很大，所以返回对 10⁹ + 7
//取余 的结果。
//
//
//
// 示例 1：
//
//
//输入：n = 2
//输出：8
//解释：
//有 8 种长度为 2 的记录将被视为可奖励：
//"PP" , "AP", "PA", "LP", "PL", "AL", "LA", "LL"
//只有"AA"不会被视为可奖励，因为缺勤次数为 2 次（需要少于 2 次）。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：3
//
//
// 示例 3：
//
//
//输入：n = 10101
//输出：183236316
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁵
//
// Related Topics 动态规划 👍 258 👎 0

package main

import "fmt"

func main() {
	fmt.Println(checkRecord3(10101))
}

//leetcode submit region begin(Prohibit modification and deletion)
//状态机
func checkRecord1(n int) int {
	const MOD = 1e9 + 7
	a, b, c, d, e, f := 1, 0, 0, 0, 0, 0
	for i := 1; i <= n; i++ {
		a, b, c, d, e, f = a+b+c, a, b, a+b+c+d+e+f, d, e
		a %= MOD
		b %= MOD
		c %= MOD
		d %= MOD
		e %= MOD
		f %= MOD
	}
	return (a + b + c + d + e + f) % MOD
}

//记忆化搜索
func checkRecord2(n int) int {
	const MOD = 1e9 + 7
	cache := make([][2][3]int, n+1) //day a l
	var dfs func(day, A, L int) int
	dfs = func(day, A, L int) (ret int) {
		if day >= n {
			return 1
		}
		if cache[day][A][L] > 0 {
			return cache[day][A][L]
		}
		defer func() { cache[day][A][L] = ret }()
		ret = (ret + dfs(day+1, A, 0)) % MOD
		if A < 1 {
			ret = (ret + dfs(day+1, A+1, 0)) % MOD
		}
		if L < 2 {
			ret = (ret + dfs(day+1, A, L+1)) % MOD
		}
		return
	}
	return dfs(0, 0, 0)
}
func checkRecord3(n int) int {
	const MOD = 1e9 + 7
	cache := [2][3]int{} //day a l
	cache[0][0] = 1
	cache[1][0] = 1
	cache[0][1] = 1
	for i := 1; i < n; i++ {
		/*P*/
		cache[0][0], cache[1][0], cache[0][1], cache[0][2], cache[1][1], cache[1][2] = (cache[0][0]+cache[0][1]+cache[0][2])%MOD, (cache[1][0]+cache[1][1]+cache[1][2]+cache[0][0]+cache[0][1]+cache[0][2])%MOD, cache[0][0], cache[0][1], cache[1][0], cache[1][1]
	}
	return (cache[0][0] + cache[0][1] + cache[0][2] + cache[1][0] + cache[1][1] + cache[1][2]) % MOD
}

//leetcode submit region end(Prohibit modification and deletion)
