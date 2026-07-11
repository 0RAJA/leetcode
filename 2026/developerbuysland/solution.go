package developerbuysland

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
【题目描述】

在一个城市区域内，被划分成了n * m个连续的区块，每个区块都拥有不同的权值，代表着其土地价值。目前，有两家开发公司，A 公司和 B 公司，希望购买这个城市区域的土地。

现在，需要将这个城市区域的所有区块分配给 A 公司和 B 公司。

然而，由于城市规划的限制，只允许将区域按横向或纵向划分成两个子区域，而且每个子区域都必须包含一个或多个区块。

为了确保公平竞争，你需要找到一种分配方式，使得 A 公司和 B 公司各自的子区域内的土地总价值之差最小。

注意：区块不可再分。

【输入描述】

第一行输入两个正整数，代表 n 和 m。

接下来的 n 行，每行输出 m 个正整数。

输出描述

请输出一个整数，代表两个子区域内土地总价值之间的最小差距。

【输入示例】

3 3
1 2 3
2 1 3
1 2 3

【输出示例】

0

【提示信息】

如果将区域按照如下方式划分：

1 2 | 3
2 1 | 3
1 2 | 3
两个子区域内土地总价值之间的最小差距可以达到 0。

【数据范围】：

1 <= n, m <= 100；
n 和 m 不同时为 1。
*/

// developerBuysLand 开发商购买土地
// params: data 土地数据
// return: 最小差距
func developerBuysLand(data [][]int) int {
	// 二维前缀和，data[x][y] = data[x-1][y] + data[x][y-1] - data[x-1][y-1]
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			if i == 0 && j == 0 {
				data[i][j] = data[i][j]
			} else if i == 0 {
				data[i][j] += data[i][j-1]
			} else if j == 0 {
				data[i][j] += data[i-1][j]
			} else {
				data[i][j] += data[i-1][j] + data[i][j-1] - data[i-1][j-1]
			}
		}
	}
	minDiff := math.MaxFloat64
	total := data[len(data)-1][len(data[0])-1]
	// 从左往右走，竖着切
	for i := 0; i < len(data)-1; i++ {
		dataA := data[i][len(data[0])-1]
		minDiff = math.Min(minDiff, math.Abs(float64(2*dataA-total)))
	}
	// 从上往下走，横着切
	for j := 0; j < len(data[0])-1; j++ {
		dataB := data[len(data)-1][j]
		minDiff = math.Min(minDiff, math.Abs(float64(2*dataB-total)))
	}
	return int(minDiff)
}

func main() {
	var n, m int

	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	params := strings.Split(line, " ")

	n, _ = strconv.Atoi(params[0])
	m, _ = strconv.Atoi(params[1]) // n和m读取完成

	land := make([][]int, n) // 土地矩阵初始化

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		values := strings.Split(line, " ")
		land[i] = make([]int, m)
		for j := 0; j < m; j++ {
			value, _ := strconv.Atoi(values[j])
			land[i][j] = value
		}
	} // 所有读取完成
	fmt.Println(developerBuysLand(land))
}
