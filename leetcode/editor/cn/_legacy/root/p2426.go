package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	nums := make([][]int, n+2)
	for i := range nums {
		nums[i] = make([]int, n+2)
		if i >= 1 && i <= n {
			_, _ = fmt.Scan(&nums[i][i])
		}
	}
	fmt.Println(P2426(nums, n))
}

func P2426(dp [][]int, n int) (ret int) {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	cal := func(l, r int) int {
		if l == r {
			return dp[l][r]
		}
		return abs(dp[l][l]-dp[r][r]) * (r - l + 1)
	}
	for i, j := 1, 2; j <= n; j++ {
		s := dp[i][j]
		for k := 2; k <= j; k++ {
			if v := cal(i, k) + dp[k+1][j]; v > s {
				s = v
			}
			if v := cal(k, j) + dp[i][k-1]; v > s {
				s = v
			}
		}
		dp[i][j] = s
	}
	return dp[1][n]
}

/*
6
54 29 196 21 133 118

dp[1][j] = max(cal(1,k)+dp[k+1][j],dp[1][k-1],cal(k,j))
j = [2,n]
k = [2,j]
*/
