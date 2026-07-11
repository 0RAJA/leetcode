//设计一个使用单词列表进行初始化的数据结构，单词列表中的单词 互不相同 。 如果给出一个单词，请判定能否只将这个单词中一个字母换成另一个字母，使得所形成的新单
//词存在于已构建的神奇字典中。
//
// 实现 MagicDictionary 类：
//
//
// MagicDictionary() 初始化对象
// void buildDict(String[] dictionary) 使用字符串数组 dictionary 设定该数据结构，dictionary 中的字
//符串互不相同
// bool search(String searchWord) 给定一个字符串 searchWord ，判定能否只将字符串中 一个 字母换成另一个字母，使得
//所形成的新字符串能够与字典中的任一字符串匹配。如果可以，返回 true ；否则，返回 false 。
//
//
//
//
//
//
//
// 示例：
//
//
//输入
//inputs = ["MagicDictionary", "buildDict", "search", "search", "search",
//"search"]
//inputs = [[], [["hello", "leetcode"]], ["hello"], ["hhllo"], ["hell"], [
//"leetcoded"]]
//输出
//[null, null, false, true, false, false]
//
//解释
//MagicDictionary magicDictionary = new MagicDictionary();
//magicDictionary.buildDict(["hello", "leetcode"]);
//magicDictionary.search("hello"); // 返回 False
//magicDictionary.search("hhllo"); // 将第二个 'h' 替换为 'e' 可以匹配 "hello" ，所以返回 True
//magicDictionary.search("hell"); // 返回 False
//magicDictionary.search("leetcoded"); // 返回 False
//
//
//
//
// 提示：
//
//
// 1 <= dictionary.length <= 100
// 1 <= dictionary[i].length <= 100
// dictionary[i] 仅由小写英文字母组成
// dictionary 中的所有字符串 互不相同
// 1 <= searchWord.length <= 100
// searchWord 仅由小写英文字母组成
// buildDict 仅在 search 之前调用一次
// 最多调用 100 次 search
//
//
//
//
//
//
//
// 注意：本题与主站 676 题相同： https://leetcode-cn.com/problems/implement-magic-
//dictionary/
// Related Topics 设计 字典树 哈希表 字符串 👍 19 👎 0

package main

import "fmt"

func main() {
	m := Constructor()
	m.BuildDict([]string{"hello", "leetcode"})
	fmt.Println(m.Search("hello"))
}

// leetcode submit region begin(Prohibit modification and deletion)
type MagicDictionary struct {
	next map[byte]*MagicDictionary
	ok   bool
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{next: make(map[byte]*MagicDictionary)}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, s := range dictionary {
		root := this
		for i := range s {
			if root.next[s[i]] == nil {
				root.next[s[i]] = &MagicDictionary{next: make(map[byte]*MagicDictionary)}
			}
			root = root.next[s[i]]
		}
		root.ok = true
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	var dfs func(root *MagicDictionary, idx int, isOK bool) bool
	dfs = func(root *MagicDictionary, idx int, isOK bool) bool {
		if idx == len(searchWord) {
			return root.ok && !isOK
		}
		for c, v := range root.next {
			//fmt.Println(string(c), string(searchWord[idx]))
			if c == searchWord[idx] {
				if dfs(v, idx+1, isOK) {
					return true
				}
			} else if isOK && dfs(v, idx+1, false) {
				return true
			}
		}
		return false
	}
	return dfs(this, 0, true)
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
//leetcode submit region end(Prohibit modification and deletion)
