package intervalsum

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://programmercarl.com/kamacoder/0058.%E5%8C%BA%E9%97%B4%E5%92%8C.html

/*
给定一个整数数组 Array，请计算该数组在每个指定区间内元素的总和。

输入描述

第一行输入为整数数组 Array 的长度 n，接下来 n 行，每行一个整数，表示数组的元素。随后的输入为需要计算总和的区间，直至文件结束。

输出描述

输出每个指定区间内元素的总和。

输入示例

5
1
2
3
4
5
0 1
1 3

输出示例
3
9

数据范围：

0 < n <= 100000

*/

// intervalSum 计算区间和
// params: nums 数组, intervals 区间
// return: 区间和
func intervalSum(nums []int, intervals [][2]int) (result []int) {
	// 维持一个数组，记录前缀和即可，需要求哪个范围则用后一个减去前一个
	result = make([]int, len(intervals))
	for idx := 1; idx < len(nums); idx++ {
		nums[idx] += nums[idx-1]
	}
	for idx, interval := range intervals {
		st, et := interval[0], interval[1]
		if st == 0 {
			result[idx] = nums[et]
		} else {
			result[idx] = nums[et] - nums[st-1]
		}
	}
	return result
}

func main() {
	// bufio中读取数据的接口，因为数据卡的比较严，导致使用fmt.Scan会超时
	scanner := bufio.NewScanner(os.Stdin)
	// 获取数组大小
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		nums[i], _ = strconv.Atoi(scanner.Text())
	}
	var intervals [][2]int
	for {
		var st, et int
		scanner.Scan()
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &st, &et)
		if err != nil {
			break
		}
		intervals = append(intervals, [2]int{st, et})
	}
	result := intervalSum(nums, intervals)
	for _, sum := range result {
		fmt.Println(sum)
	}
}
