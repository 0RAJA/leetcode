//n 对情侣坐在连续排列的 2n 个座位上，想要牵到对方的手。
//
// 人和座位由一个整数数组 row 表示，其中 row[i] 是坐在第 i 个座位上的人的 ID。情侣们按顺序编号，第一对是 (0, 1)，第二对是 (2,
//3)，以此类推，最后一对是 (2n-2, 2n-1)。
//
// 返回 最少交换座位的次数，以便每对情侣可以并肩坐在一起。 每次交换可选择任意两人，让他们站起来交换座位。
//
//
//
// 示例 1:
//
//
//输入: row = [0,2,1,3]
//输出: 1
//解释: 只需要交换row[1]和row[2]的位置即可。
//
//
// 示例 2:
//
//
//输入: row = [3,2,0,1]
//输出: 0
//解释: 无需交换座位，所有的情侣都已经可以手牵手了。
//
//
//
//
// 提示:
//
//
// 2n == row.length
// 2 <= n <= 30
// n 是偶数
// 0 <= row[i] < 2n
// row 中所有元素均无重复
//
// Related Topics 贪心 深度优先搜索 广度优先搜索 并查集 图 👍 319 👎 0

package main

import "fmt"

func main() {
	fmt.Println(minSwapsCouples([]int{0, 2, 1, 3}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minSwapsCouples(row []int) (ret int) {
	indexs := make([]int, len(row)) //记录每个情侣所在下标
	for i := range row {
		indexs[row[i]] = i
	}
	for i := 0; i < len(row); i += 2 { //根据题意,由第一个人决定座位,所以通过每组的第一个人找另一个人的位置
		lover := row[i] ^ 1            // 找到对应的情侣,+1/-1
		if x := row[i+1]; x != lover { //如果不是情侣,就把自己的情侣和右边的人进行交换,然后更新数据即可.
			indexs[x], indexs[lover] = indexs[lover], indexs[x]
			row[indexs[x]], row[indexs[lover]] = x, lover
			ret++
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
