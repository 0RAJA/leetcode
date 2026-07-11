//列表 arr 由在范围 [1, n] 中的所有整数组成，并按严格递增排序。请你对 arr 应用下述算法：
//
//
//
//
// 从左到右，删除第一个数字，然后每隔一个数字删除一个，直到到达列表末尾。
// 重复上面的步骤，但这次是从右到左。也就是，删除最右侧的数字，然后剩下的数字每隔一个删除一个。
// 不断重复这两步，从左到右和从右到左交替进行，直到只剩下一个数字。
//
//
// 给你整数 n ，返回 arr 最后剩下的数字。
//
//
//
// 示例 1：
//
//
//输入：n = 9
//输出：6
//解释：
//arr = [1, 2, 3, 4, 5, 6, 7, 8, 9]
//arr = [2, 4, 6, 8]
//arr = [2, 6]
//arr = [6]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁹
//
//
//
// Related Topics 数学 👍 194 👎 0

package main

import "fmt"

func main() {
	fmt.Println(lastRemaining(9))
}

//leetcode submit region begin(Prohibit modification and deletion)
/*
	每次都会减少一半的数，从左往右每次减去第一个数，从右往左只有是奇数个数才会减少第一个数
	开始每个数间距为1,然后每删除一轮间距*2
	每当要删除第一个数就要加上当前第二个数前面删除的数的个数(即和第一个数的间距)
*/
func lastRemaining(n int) (ret int) {
	step := 1
	ret = 1
	direct := true
	for n > 1 {
		if direct || n%2 != 0 {
			ret += step
		}
		direct = !direct
		step *= 2
		n /= 2
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
