// Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补全和拼
// 写检查。
//
// 请你实现 Trie 类：
//
//
// Trie() 初始化前缀树对象。
// void insert(String word) 向前缀树中插入字符串 word 。
// boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回
// false 。
// boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否
// 则，返回 false 。
//
//
//
//
// 示例：
//
//
// 输入
// ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
// [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
// 输出
// [null, null, true, false, true, null, true]
//
// 解释
// Trie trie = new Trie();
// trie.insert("apple");
// trie.search("apple");   // 返回 True
// trie.search("app");     // 返回 False
// trie.startsWith("app"); // 返回 True
// trie.insert("app");
// trie.search("app");     // 返回 True
//
//
//
//
// 提示：
//
//
// 1 <= word.length, prefix.length <= 2000
// word 和 prefix 仅由小写英文字母组成
// insert、search 和 startsWith 调用次数 总计 不超过 3 * 10⁴ 次
//
//
// Related Topics 设计 字典树 哈希表 字符串 👍 1971 👎 0

package p0208

// leetcode submit region begin(Prohibit modification and deletion)
// 每层都是一个 map，key 是 字符，value 是下一个 map
type TrieNode struct {
	char    byte              // 当前字符
	nodeMap map[byte]TrieNode // 之后的 map
}

type Trie struct {
	allWord map[string]bool // 记录完整的字符串
	nodeMap TrieNode        // 这里是开头
}

func Constructor() Trie {
	return Trie{
		allWord: make(map[string]bool),
		nodeMap: TrieNode{
			nodeMap: make(map[byte]TrieNode),
		},
	}
}

func (this *Trie) Insert(word string) {
	if _, ok := this.allWord[word]; ok {
		return
	}
	this.allWord[word] = true
	nowNode := this.nodeMap
	for _, char := range word {
		if node, ok := nowNode.nodeMap[byte(char)]; ok {
			nowNode = node
		} else {
			node := TrieNode{
				char:    byte(char),
				nodeMap: make(map[byte]TrieNode),
			}
			nowNode.nodeMap[byte(char)] = node
			nowNode = node
		}
	}
}

func (this *Trie) Search(word string) bool {
	_, ok := this.allWord[word]
	return ok
}

func (this *Trie) StartsWith(prefix string) bool {
	nowNode := this.nodeMap
	for _, char := range prefix {
		if node, ok := nowNode.nodeMap[byte(char)]; ok {
			nowNode = node
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// leetcode submit region end(Prohibit modification and deletion)
