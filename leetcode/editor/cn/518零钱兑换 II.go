// 给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。
//
//
//
//
//
//
// 示例 1:
//
// 输入: amount = 5, coins = [1, 2, 5]
// 输出: 4
// 解释: 有四种方式可以凑成总金额:
// 5=5
// 5=2+2+1
// 5=2+1+1+1
// 5=1+1+1+1+1
//
//
// 示例 2:
//
// 输入: amount = 3, coins = [2]
// 输出: 0
// 解释: 只用面额2的硬币不能凑成总金额3。
//
//
// 示例 3:
//
// 输入: amount = 10, coins = [10]
// 输出: 1
//
//
//
//
// 注意:
//
// 你可以假设：
//
//
// 0 <= amount (总金额) <= 5000
// 1 <= coin (硬币面额) <= 5000
// 硬币种类不超过 500 种
// 结果符合 32 位符号整数
//
// 👍 426 👎 0
package main

import (
	"fmt"
)

func main() {
	amount := 5
	coins := []int{1, 2, 5}
	fmt.Println(change(amount, coins))
}

// leetcode submit region begin(Prohibit modification and deletion)
/*
完全背包，和279类似
/*
01和完全背包真的很像，主要就是遍历的方向
排列和组合的区别影响的是物品和空间的遍历顺序（两个for的先后），排列是固定的，组合是可以替换的
*/
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	// 物品 无限
	for _, v := range coins {
		// 空间
		for i := v; i <= amount; i++ {
			dp[i] += dp[i-v]
		}
	}
	return dp[amount]
}

// leetcode submit region end(Prohibit modification and deletion)
