//给定两个字符串 s 和 t 。返回 s 中包含 t 的所有字符的最短子字符串。如果 s 中不存在符合条件的子字符串，则返回空字符串 "" 。
//
// 如果 s 中存在多个符合条件的子字符串，返回任意一个。
//
//
//
// 注意： 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
//
//
//
// 示例 1：
//
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//解释：最短子字符串 "BANC" 包含了字符串 t 的所有字符 'A'、'B'、'C'
//
// 示例 2：
//
//
//输入：s = "a", t = "a"
//输出："a"
//
//
// 示例 3：
//
//
//输入：s = "a", t = "aa"
//输出：""
//解释：t 中两个字符 'a' 均应包含在 s 的子串中，因此没有符合条件的子字符串，返回空字符串。
//
//
//
// 提示：
//
//
// 1 <= s.length, t.length <= 10⁵
// s 和 t 由英文字母组成
//
//
//
//
// 进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？
//
//
//
// 注意：本题与主站 76 题相似（本题答案不唯一）：https://leetcode-cn.com/problems/minimum-window-
//substring/
// Related Topics 哈希表 字符串 滑动窗口 👍 37 👎 0

package main

import "fmt"

func main() {
	fmt.Println(minWindow("a", "b"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) (ret string) {
	if len(s) < len(t) {
		return ""
	}
	ret = s + t
	cnt := [128]int{}
	lack := 0
	for i := range t {
		cnt[t[i]]++
		if cnt[t[i]] == 1 {
			lack++
		}
	}
	for l, r := 0, 0; r < len(s); r++ {
		cnt[s[r]]--
		if cnt[s[r]] == 0 {
			lack--
		}
		for lack <= 0 {
			cnt[s[l]]++
			if cnt[s[l]] == 1 {
				if lack == 0 {
					if len(ret) > r-l+1 {
						ret = s[l : r+1]
					}
				}
				lack++
			}
			l++
		}
	}
	if ret != s+t {
		return
	}
	return ""
}

//leetcode submit region end(Prohibit modification and deletion)
