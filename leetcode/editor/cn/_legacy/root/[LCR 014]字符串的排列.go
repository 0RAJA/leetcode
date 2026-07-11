// 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
//
// 换句话说，第一个字符串的排列之一是第二个字符串的 子串 。 
//
// 
//
// 示例 1： 
//
// 
// 输入: s1 = "ab" s2 = "eidbaooo"
// 输出: True
// 解释: s2 包含 s1 的排列之一 ("ba").
// 
//
// 示例 2： 
//
// 
// 输入: s1= "ab" s2 = "eidboaoo"
// 输出: False
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s1.length, s2.length <= 10⁴ 
// s1 和 s2 仅包含小写字母 
// 
//
// 
//
// 
// 注意：本题与主站 567 题相同： https://leetcode.cn/problems/permutation-in-string/ 
//
// Related Topics 哈希表 双指针 字符串 滑动窗口 👍 114 👎 0
package main

import (
	"fmt"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 排列不包括顺序，只比较字符出现频率，因此统计 s1 的词频，维护 s2 区间的词频，两者长度相同时则比较词频是否相等
func checkInclusion(s1 string, s2 string) bool {
	s1CharCountNums := [26]int{}
	for _,s := range s1 {
		s1CharCountNums[s-'a']++
	}
	s2WinCharCountNums := [26]int{}
	for l, r := 0, 0; r < len(s2); r++ {
		s2WinCharCountNums[s2[r]-'a']++
		for r-l+1 > len(s1) {
			s2WinCharCountNums[s2[l]-'a']--
			l++
		}
		if s2WinCharCountNums == s1CharCountNums {
			return true
		}
	}
	return false
}

// leetcode submit region end(Prohibit modification and deletion)
