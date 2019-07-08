package main

import "fmt"

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

	l:= LRUCache{
		M:    make(map[int]*ListNode, capacity),
		C:    capacity,
		Head: &ListNode{},
		Last: &ListNode{},
	}

	l.Head.Next=l.Last
	l.Last.Pre=l.Head

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
		n = this.Last.Pre
		delete(this.M, n.Key)

		n.Key=key
		n.Val=value

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

func main() {
	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	fmt.Println(l.Get(1))
	l.Put(3, 3)
	fmt.Println("---",l)
	fmt.Println(l.Get(2))
	l.Put(4, 4)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(3))
	fmt.Println(l.Get(4))
}
