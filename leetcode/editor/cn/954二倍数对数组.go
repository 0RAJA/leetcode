//给定一个长度为偶数的整数数组 arr，只有对 arr 进行重组后可以满足 “对于每个 0 <= i < len(arr) / 2，都有 arr[2 * i
//+ 1] = 2 * arr[2 * i]” 时，返回 true；否则，返回 false。
//
//
//
// 示例 1：
//
//
//输入：arr = [3,1,3,6]
//输出：false
//
//
// 示例 2：
//
//
//输入：arr = [2,1,2,6]
//输出：false
//
//
// 示例 3：
//
//
//输入：arr = [4,-2,2,-4]
//输出：true
//解释：可以用 [-2,-4] 和 [2,4] 这两组组成 [-2,-4,2,4] 或是 [2,4,-2,-4]
//
//
//
//
// 提示：
//
//
// 0 <= arr.length <= 3 * 10⁴
// arr.length 是偶数
// -10⁵ <= arr[i] <= 10⁵
//
// Related Topics 贪心 数组 哈希表 排序 👍 171 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canReorderDoubled([]int{4, -2, 2, -4}))
}

//leetcode submit region begin(Prohibit modification and deletion)
//其实就是分成一对一对，每一对一个是另一个的两倍,所以小的的次数一定要小于大的的次数
func canReorderDoubled(arr []int) bool {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	cnt := make(map[int]int, len(arr)) //k的次数
	for _, x := range arr {
		cnt[x]++
	}
	if cnt[0]%2 == 1 { //0只能匹配0，所以必须是偶数
		return false
	}

	vals := make([]int, 0, len(cnt))
	for x := range cnt {
		vals = append(vals, x)
	}
	sort.Slice(vals, func(i, j int) bool { return abs(vals[i]) < abs(vals[j]) }) //按数的绝对值从小到大排序,cnt[k1] == cnt[2*k2] 所以k1要小

	for _, x := range vals {
		if cnt[2*x] < cnt[x] { // 无法找到足够的 2x 与 x 配对
			return false
		}
		cnt[2*x] -= cnt[x]
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
