/*
146.中 LRU 缓存

请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4

提示：

1 <= capacity <= 3000
0 <= key <= 10000
0 <= value <= 105
最多调用 2 * 105 次 get 和 put
*/
package main

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Key  int
	Pre  *ListNode
	Next *ListNode
}

type LRUCache struct {
	M map[int]*ListNode
	C int

	// Head,Last为哨兵节点, 方便节点挪移
	Head *ListNode
	Last *ListNode
}

func Constructor(capacity int) LRUCache {
	if capacity <= 0 {
		panic("invalid capacity")
	}

	l := LRUCache{
		M:    make(map[int]*ListNode, capacity),
		C:    capacity,
		Head: &ListNode{},
		Last: &ListNode{},
	}

	l.Head.Next = l.Last
	l.Last.Pre = l.Head

	return l
}

func (this *LRUCache) Get(key int) int {
	n := this.M[key]
	if n == nil { // no node
		return -1
	}

	this.Remove(n)
	this.Insert2Head(n)

	return n.Val
}

// 分离节点
// 将前一个的Next指向n.Next
func (this *LRUCache) Remove(n *ListNode) {
	n.Next.Pre = n.Pre
	n.Pre.Next = n.Next
}

func (this *LRUCache) Insert2Head(n *ListNode) {
	n.Pre = this.Head
	n.Next = this.Head.Next

	this.Head.Next = n
	n.Next.Pre = n
}

func (this *LRUCache) Put(key int, value int) {
	if n := this.M[key]; n != nil { // node exist
		n.Val = value

		this.Remove(n)
		this.Insert2Head(n)

		return
	}

	// node not exist

	var n *ListNode
	if len(this.M) == this.C { // is full
		n = this.Last.Pre     // 将结尾哨兵的前一个元素即最后一个元素替换成新value
		delete(this.M, n.Key) // n将被挪用

		n.Key = key
		n.Val = value

		this.Remove(n)
	} else {
		n = &ListNode{
			Key: key,
			Val: value,
		}
	}

	this.Insert2Head(n)

	this.M[key] = n
}

func TestConstructor(t *testing.T) {
	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	fmt.Println(l.Get(1))
	l.Put(3, 3)
	fmt.Println("---", l)
	fmt.Println(l.Get(2))
	l.Put(4, 4)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(3))
	fmt.Println(l.Get(4))
}
