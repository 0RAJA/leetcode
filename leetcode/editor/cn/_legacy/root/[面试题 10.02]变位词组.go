// 编写一种方法，对字符串数组进行排序，将所有变位词组合在一起。变位词是指字母相同，但排列不同的字符串。
//
// 注意：本题相对原题稍作修改
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
// 输出:
// [
//
//	["ate","eat","tea"],
//	["nat","tan"],
//	["bat"]
//
// ]
//
// 说明：
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。
//
// Related Topics 哈希表 字符串 排序
// 👍 58 👎 0
package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) (ret [][]string) {
	symbol := make(map[string][]string)
	for x := 0; x < len(strs); x++ {
		str := []byte(strs[x])
		sort.Slice(str, func(i, j int) bool {
			return str[i] < str[j]
		})
		symbol[string(str)] = append(symbol[string(str)], strs[x])
	}
	ret = make([][]string, 0, len(strs))
	for _, v := range symbol {
		ret = append(ret, v)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
