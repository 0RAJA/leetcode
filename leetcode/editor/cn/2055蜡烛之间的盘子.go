//给你一个长桌子，桌子上盘子和蜡烛排成一列。给你一个下标从 0 开始的字符串 s ，它只包含字符 '*' 和 '|' ，其中 '*' 表示一个 盘子 ，'|
//' 表示一支 蜡烛 。
//
// 同时给你一个下标从 0 开始的二维整数数组 queries ，其中 queries[i] = [lefti, righti] 表示 子字符串 s[
//lefti...righti] （包含左右端点的字符）。对于每个查询，你需要找到 子字符串中 在 两支蜡烛之间 的盘子的 数目 。如果一个盘子在 子字符串中 左边和右边
// 都 至少有一支蜡烛，那么这个盘子满足在 两支蜡烛之间 。
//
//
// 比方说，s = "||**||**|*" ，查询 [3, 8] ，表示的是子字符串 "*||**|" 。子字符串中在两支蜡烛之间的盘子数目为 2 ，子字符
//串中右边两个盘子在它们左边和右边 都 至少有一支蜡烛。
//
//
// 请你返回一个整数数组 answer ，其中 answer[i] 是第 i 个查询的答案。
//
//
//
// 示例 1:
//
//
//
// 输入：s = "**|**|***|", queries = [[2,5],[5,9]]
//输出：[2,3]
//解释：
//- queries[0] 有两个盘子在蜡烛之间。
//- queries[1] 有三个盘子在蜡烛之间。
//
//
// 示例 2:
//
//
//
// 输入：s = "***|**|*****|**||**|*", queries = [[1,17],[4,5],[14,17],[5,11],[15,16
//]]
//输出：[9,0,0,0,0]
//解释：
//- queries[0] 有 9 个盘子在蜡烛之间。
//- 另一个查询没有盘子在蜡烛之间。
//
//
//
//
// 提示：
//
//
// 3 <= s.length <= 10⁵
// s 只包含字符 '*' 和 '|' 。
// 1 <= queries.length <= 10⁵
// queries[i].length == 2
// 0 <= lefti <= righti < s.length
//
// Related Topics 数组 字符串 二分查找 前缀和 👍 30 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "**|**|***|"
	queries := [][]int{{2, 5}, {5, 9}}
	fmt.Println(platesBetweenCandles(s, queries))
}

//leetcode submit region begin(Prohibit modification and deletion)
func platesBetweenCandles(s string, queries [][]int) (ret []int) {
	ret = make([]int, 0, len(queries))
	nums := make(map[int]int)         //记录每个蜡烛前面的有效的盘子数（只计算在两个蜡烛之间的）
	candles := make([]int, 0, len(s)) //记录每个蜡烛的下标，二分查找用
	var candle, plate, sum int        //前面的蜡烛数(0,1) 与前面蜡烛之间的盘子数 蜡烛前面的有效的盘子数
	for i, c := range s {
		if c == '|' {
			if candle > 0 {
				sum += plate
				plate = 0
			}
			candle = 1
			if len(candles) != 0 { //第一个蜡烛前面肯定是0个有效的盘子
				nums[i] = sum
			}
			candles = append(candles, i)
		} else if candle > 0 {
			plate++
		}
	}
	for _, qs := range queries {
		var sum int
		a, b := qs[0], qs[1]                                               //左右区间
		li, ri := sort.SearchInts(candles, a), sort.SearchInts(candles, b) //二分找到最近的两个蜡烛的下标
		if li != ri {                                                      //左右下标相同直接返回0
			if ri == len(candles) { //没找到且右边没蜡烛了
				ri--
			}
			if candles[ri] > b { //超过右边界了 左边界不用考虑
				ri--
			}
			l, r := candles[li], candles[ri]
			sum = nums[r] - nums[l]
		} else {
			sum = 0
		}
		ret = append(ret, sum)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
