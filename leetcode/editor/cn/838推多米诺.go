//n 张多米诺骨牌排成一行，将每张多米诺骨牌垂直竖立。在开始时，同时把一些多米诺骨牌向左或向右推。
//
// 每过一秒，倒向左边的多米诺骨牌会推动其左侧相邻的多米诺骨牌。同样地，倒向右边的多米诺骨牌也会推动竖立在其右侧的相邻多米诺骨牌。
//
// 如果一张垂直竖立的多米诺骨牌的两侧同时有多米诺骨牌倒下时，由于受力平衡， 该骨牌仍然保持不变。
//
// 就这个问题而言，我们会认为一张正在倒下的多米诺骨牌不会对其它正在倒下或已经倒下的多米诺骨牌施加额外的力。
//
// 给你一个字符串 dominoes 表示这一行多米诺骨牌的初始状态，其中：
//
//
// dominoes[i] = 'L'，表示第 i 张多米诺骨牌被推向左侧，
// dominoes[i] = 'R'，表示第 i 张多米诺骨牌被推向右侧，
// dominoes[i] = '.'，表示没有推动第 i 张多米诺骨牌。
//
//
// 返回表示最终状态的字符串。
//
//
// 示例 1：
//
//
//输入：dominoes = "RR.L"
//输出："RR.L"
//解释：第一张多米诺骨牌没有给第二张施加额外的力。
//
//
// 示例 2：
//
//
//输入：dominoes = ".L.R...LR..L.."
//输出："LL.RR.LLRRLL.."
//
//
//
//
// 提示：
//
//
// n == dominoes.length
// 1 <= n <= 10⁵
// dominoes[i] 为 'L'、'R' 或 '.'
//
// Related Topics 双指针 字符串 动态规划 👍 224 👎 0

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(pushDominoes(".L.R...LR..L.."))
}

//leetcode submit region begin(Prohibit modification and deletion)
func pushDominoes(dominoes string) (ret string) {
	dominoes = "L" + dominoes + "R"
	sb := &strings.Builder{}
	for l, r := 0, 1; r < len(dominoes); r++ {
		if dominoes[r] != '.' {
			if dominoes[l] == dominoes[r] {
				for i := l; i < r; i++ {
					sb.WriteByte(dominoes[l])
				}
			} else {
				if dominoes[l] == 'L' {
					sb.WriteByte(dominoes[l])
					for i := l + 1; i < r; i++ {
						sb.WriteByte('.')
					}
				} else {
					mid := l + (r-l)/2
					var ln, rn, num = mid - l, mid - l - 1, 1
					if (r-l+1)%2 == 0 {
						ln = mid - l + 1
						rn = ln - 1
						num = 0
					}
					for i := 0; i < ln; i++ {
						sb.WriteByte('R')
					}
					if num > 0 {
						sb.WriteByte('.')
					}
					for i := 0; i < rn; i++ {
						sb.WriteByte('L')
					}
				}
			}
			l = r
		}
	}
	return sb.String()[1:]
}

//leetcode submit region end(Prohibit modification and deletion)
