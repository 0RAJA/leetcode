//给出一个字符串数组 words 组成的一本英语词典。返回 words 中最长的一个单词，该单词是由 words 词典中其他单词逐步添加一个字母组成。
//
// 若其中有多个可行的答案，则返回答案中字典序最小的单词。若无答案，则返回空字符串。
//
//
//
// 示例 1：
//
//
//输入：words = ["w","wo","wor","worl", "world"]
//输出："world"
//解释： 单词"world"可由"w", "wo", "wor", 和 "worl"逐步添加一个字母组成。
//
//
// 示例 2：
//
//
//输入：words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
//输出："apple"
//解释："apply" 和 "apple" 都能由词典中的单词组成。但是 "apple" 的字典序小于 "apply"
//
//
//
//
// 提示：
//
//
// 1 <= words.length <= 1000
// 1 <= words[i].length <= 30
// 所有输入的字符串 words[i] 都只包含小写字母。
//
// Related Topics 字典树 数组 哈希表 字符串 排序 👍 233 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestWord([]string{"a", "ac", "acd", "acde", "acdef"}))
}

//leetcode submit region begin(Prohibit modification and deletion)
/*
	看懂题意：意思是挨个进行拼凑，第一个单词出第一个字母，第二个单词出第二的字母。。。
	按照长度从小到大（保证可以从小到大开始加字符），长度相同按照字典序从大到小（保证往后遍历到的都是最优解（字典序））。
*/
func longestWord(words []string) (ret string) {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j]) || len(words[i]) == len(words[j]) && words[i] > words[j]
	})
	fmt.Println(words)
	m := make(map[string]struct{})
	m[""] = struct{}{}
	for _, s := range words {
		if _, ok := m[s[:len(s)-1]]; ok {
			m[s] = struct{}{}
			ret = s
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
