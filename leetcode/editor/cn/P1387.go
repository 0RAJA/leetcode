package main

import "fmt"

func main() {
	fmt.Println(P1387())
}

func P1387() (ret int) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var m, n int
	fmt.Scan(&m, &n)
	/*
		dp[i][j] 表示以【i,j】为右下角的正方形的最大的边长，选取四个方向的最小值
	*/
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			fmt.Scan(&dp[i][j])
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i][j] == 1 {
				dp[i][j] += min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
				if dp[i][j] > ret {
					ret = dp[i][j]
				}
			}
		}
	}
	return
}
