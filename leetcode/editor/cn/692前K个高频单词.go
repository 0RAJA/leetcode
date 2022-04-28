//给一非空的单词列表，返回前 K 个出现次数最多的单词。
//
// 返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。
//
// 示例 1：
//
//
//输入: ["i", "love", "leetcode", "i", "love", "coding"], K = 2
//输出: ["i", "love"]
//解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
//    注意，按字母顺序 "i" 在 "love" 之前。
//
//
//
//
// 示例 2：
//
//
//输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], K
// = 4
//输出: ["the", "is", "sunny", "day"]
//解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
//    出现次数依次为 4, 3, 2 和 1 次。
//
//
//
//
// 注意：
//
//
// 假定 K 总为有效值， 1 ≤ K ≤ 集合元素数。
// 输入的单词均由小写字母组成。
//
//
//
//
// 扩展练习：
//
//
// 尝试以 O(n log K) 时间复杂度和 O(n) 空间复杂度解决。
//
// Related Topics 堆 字典树 哈希表
// 👍 291 👎 0

package main

import (
	"container/heap"
	"fmt"
	"strings"
)

func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	fmt.Println(topKFrequent(words, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
type myType struct {
	strs []string
	cnts map[string]int
}

func topKFrequent(words []string, k int) (ret []string) {
	m := &myType{cnts: make(map[string]int), strs: make([]string, 0, len(words))}
	for i := range words {
		m.cnts[words[i]]++
	}
	for k := range m.cnts {
		heap.Push(m, k)
	}
	ret = make([]string, 0, k)
	for i := 0; i < k; i++ {
		ret = append(ret, m.strs[0])
		heap.Remove(m, 0)
	}
	return
}

func (m myType) Len() int {
	return len(m.strs)
}

func (m myType) Less(i, j int) bool {
	if m.cnts[m.strs[i]] != m.cnts[m.strs[j]] {
		return m.cnts[m.strs[i]] > m.cnts[m.strs[j]]
	} else {
		return strings.Compare(m.strs[i], m.strs[j]) < 0
	}
}

func (m myType) Swap(i, j int) {
	m.strs[i], m.strs[j] = m.strs[j], m.strs[i]
}

func (m *myType) Push(v interface{}) {
	m.strs = append(m.strs, v.(string))
}

func (m *myType) Pop() interface{} {
	x := m.strs[len(m.strs)-1]
	m.strs = m.strs[:len(m.strs)-1]
	return x
}

//leetcode submit region end(Prohibit modification and deletion)
