// 给定一个整数数组 prices，其中第 prices[i] 表示第 i 天的股票价格 。
//
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
//
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
//
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//
//
// 示例 1:
//
//
// 输入: prices = [1,2,3,0,2]
// 输出: 3
// 解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
//
// 示例 2:
//
//
// 输入: prices = [1]
// 输出: 0
//
//
//
//
// 提示：
//
//
// 1 <= prices.length <= 5000
// 0 <= prices[i] <= 1000
//
//
// Related Topics 数组 动态规划 👍 1935 👎 0

package p0309

import (
	"math"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 状态机-动态规划 持有 -/> 休息 -> 卖出 -> 休息 -/> 休息 -> 持有
// 对于每一天始终存在一种状态：持有、卖出、不持有
// 对于每一种状态的来源：计算天数 i 下当天的最大利润
//   - 持有：昨天持有、昨天休息+今天买入 hold[i] = max(hold[i-1], rest[i-1] - price[i])
//   - 卖出：昨天持有 sale[i] = hold[i-1] + price[i]
//   - 休息：昨天卖出、昨天休息 rest[i] = max(sale[i-1], rest[i-1])
//
// 注意 最开始不可能持有；最后计算 max(sale[n],rest[n])
func maxProfit(prices []int) int {
	hold, sale, rest := make([]int, len(prices)+1), make([]int, len(prices)+1), make([]int, len(prices)+1)
	hold[0] = math.MinInt // 最开始不可能持有
	for idx, price := range prices {
		day := idx + 1
		hold[day] = max(hold[day-1], rest[day-1]-price) // 昨天持有、昨天休息+今天买入
		sale[day] = hold[day-1] + price                 // 昨天持有
		rest[day] = max(sale[day-1], rest[day-1])       // 昨天卖出、昨天休息
	}
	return max(sale[len(prices)], rest[len(prices)])
}

// leetcode submit region end(Prohibit modification and deletion)
