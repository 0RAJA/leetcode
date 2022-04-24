package main

import "fmt"

func main() {
	fmt.Println(P1077())
}

/*
3 5
3 4 4
*/
func P1077() int {
	const MOD = 1e6 + 7
	var n, m int
	fmt.Scan(&n, &m)
	nums := make([]int, n)
	for i := range nums {
		fmt.Scan(&nums[i])
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		v := nums[i-1]
		for j := 0; j <= m; j++ {
			for k := 0; k <= v && k <= j; k++ {
				dp[i][j] = (dp[i][j] + dp[i-1][j-k]) % MOD
			}
		}
	}
	return dp[n][m]
}
