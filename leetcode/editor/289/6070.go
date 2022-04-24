package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(digitSum("11111222223", 3))
	fmt.Println(maxTrailingZeros([][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}))
}

func digitSum(s string, k int) string {
	if len(s) <= k {
		return s
	}
	sb := new(strings.Builder)
	num := 0
	cnt := 0
	for i := range s {
		num += int(s[i] - '0')
		cnt++
		if cnt%k == 0 || i == len(s)-1 {
			sb.WriteString(strconv.Itoa(num))
			num = 0
		}
	}
	return digitSum(sb.String(), k)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func minimumRounds(tasks []int) (ret int) {
	cnts := make(map[int]int)
	for i := range tasks {
		cnts[tasks[i]]++
	}
	dp := make(map[int]int)
	dp[2] = 1
	dp[3] = 1
	var cache func(int) int
	cache = func(n int) (ret int) {
		if n <= 1 {
			return -1
		}
		if dp[n] != -1 && dp[n] > 0 {
			return dp[n]
		}
		defer func() { dp[n] = ret }()
		c2 := cache(n - 2)
		c3 := cache(n - 3)
		if c2 != -1 && c3 != -1 {
			return min(c2, c3) + 1
		} else if c2 == -1 && c3 == -1 {
			return -1
		} else if c2 != -1 {
			return c2 + 1
		} else {
			return c3 + 1
		}
	}
	for _, cnt := range cnts {
		n := cache(cnt)
		if n == -1 {
			return -1
		}
		ret += n
	}
	return
}

func maxTrailingZeros(grid [][]int) int {
	rows := [2][][]int{}
	rows[0] = make([][]int, len(grid[0]))
	rows[1] = make([][]int, len(grid[0]))
	cols := [2][][]int{}
	cols[0] = make([][]int, len(grid))
	cols[1] = make([][]int, len(grid))
	for i, row := range grid {
		for j := range row {
			l := j
			r := len(grid[0]) - j - 1
			if l == 0 {
				rows[0][l] = append(rows[0][l], grid[i][l])
			} else {
				rows[0][l] = append(rows[0][l], grid[i][l]*grid[i][l-1])
			}
			if r == len(grid[0])-1 {
				rows[1][r] = append(rows[1][r], grid[i][r])
			} else {
				rows[1][r] = append(rows[1][r], grid[i][r]*grid[i][r+1])
			}
		}
	}
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			u := j
			d := len(grid) - j - 1
			if u == 0 {
				cols[0][u] = append(cols[0][u], grid[u][i])
			} else {
				cols[0][u] = append(cols[0][u], grid[u][i]*grid[u-1][i])
			}
			if d == len(grid)-1 {
				cols[1][d] = append(cols[1][d], grid[d][i])
			} else {
				cols[1][d] = append(cols[1][d], grid[d][i]*grid[d+1][i])
			}
		}
	}
	fmt.Println(rows)
	fmt.Println(cols)
	return 0
}
