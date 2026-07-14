// 请你设计并实现一个满足
// LRU (最近最少使用) 缓存 约束的数据结构。
//
// 实现
// LRUCache 类：
//
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
// key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
//
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
// 示例：
//
// 输入
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// 输出
// [null, null, null, 1, null, -1, null, -1, 3, 4]
//
// 解释
// LRUCache lRUCache = new LRUCache(2);
// lRUCache.put(1, 1); // 缓存是 {1=1}
// lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
// lRUCache.get(1);    // 返回 1
// lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
// lRUCache.get(2);    // 返回 -1 (未找到)
// lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
// lRUCache.get(1);    // 返回 -1 (未找到)
// lRUCache.get(3);    // 返回 3
// lRUCache.get(4);    // 返回 4
//
// 提示：
//
// 1 <= capacity <= 3000
// 0 <= key <= 10000
// 0 <= value <= 10⁵
// 最多调用 2 * 10⁵ 次 get 和 put
//
// Related Topics 设计 哈希表 链表 双向链表 👍 3848 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
// LRU 最近最少使用
// map+双向链表，map 用于 get O1，双向列表用于 put 时 O1，以及删除数据时 O1
// 1. get：map 找到节点，更新 节点链表位置 到开头
// 2. put
//  1. 未触发上限：map 找到节点，更新值，更新节点位置到开头
//  2. 触发上限：map 未找到节点，新增节点到开头，map 更新引用，删除最后一个节点

type Node struct {
	preNode  *Node
	nextNode *Node
	key      int
	value    int
}

type LRUCache struct {
	capacity  int
	m         map[int]*Node
	startNode *Node // 双向链表开始节点
	endNode   *Node // 双向链表结束节点
}

func Constructor(capacity int) LRUCache {
	startNode := new(Node)
	endNode := new(Node)
	startNode.nextNode = endNode
	endNode.preNode = startNode
	return LRUCache{
		m:         make(map[int]*Node),
		startNode: startNode,
		endNode:   endNode,
		capacity:  capacity,
	}
}

// 移动当前节点到开头
func (this *LRUCache) moveNode2Head(node *Node) {
	// 删除当前节点
	this.removeNode(node)
	// 更新当前节点引用
	node.preNode = this.startNode
	node.nextNode = this.startNode.nextNode
	// 更新开头节点引用
	this.startNode.nextNode.preNode = node
	this.startNode.nextNode = node
}

// 删除当前节点引用
func (this *LRUCache) removeNode(node *Node) {
	if node.preNode != nil {
		node.preNode.nextNode = node.nextNode
	}
	if node.nextNode != nil {
		node.nextNode.preNode = node.preNode
	}
}

// Get 找到则移动到开头，否则返回-1
func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.moveNode2Head(node)
		return node.value
	}
	return -1
}

// Put 尝试找到并更新节点，否则新增节点，节点移动到开头，最后判断是否超限，超过则删除最后一个节点
func (this *LRUCache) Put(key int, value int) {
	// 尝试找到更新节点，否则插入新节点
	if node, ok := this.m[key]; ok {
		node.value = value
		this.moveNode2Head(node)
	} else {
		newNode := &Node{
			preNode:  nil,
			nextNode: nil,
			key:      key,
			value:    value,
		}
		// 双向关联
		this.moveNode2Head(newNode)
		this.m[key] = newNode
	}
	// 判断是否超限，超限则进行驱逐
	if len(this.m) > this.capacity {
		removeNode := this.endNode.preNode
		this.removeNode(removeNode)
		delete(this.m, removeNode.key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// leetcode submit region end(Prohibit modification and deletion)
