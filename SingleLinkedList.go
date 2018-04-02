package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main() {
	//head := GenerateList1(10)
	head := GenerateList2(10)
	PrintList(head)

	head = ReverseList(head)
	PrintList(head)
}

func GenerateList1(num int) *Node {
	var head *Node = nil

	for i := 0; i < num; i++ {
		cur := &Node{
			Value: i,
		}

		if head == nil {
			head = cur

			continue
		}

		cur.Next = head
		head = cur
	}

	return head
}

func GenerateList2(num int) *Node {
	var head *Node = nil

	// retrurn *Node: fn's head is a copy of GenerateList2's head
	fn := func(i int, head *Node) *Node {
		n := &Node{
			Value: i,
		}

		if head == nil {
			head = n

			return head
		}

		n.Next = head
		head = n

		return head
	}

	for i := 0; i < num; i++ {
		head = fn(i, head)
	}

	return head
}

func PrintList(head *Node) {
	for head != nil {
		fmt.Println(head.Value)

		head = head.Next
	}
}

func ReverseList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}

	pre, cur := head, head.Next
	var tmp *Node

	for cur != nil {
		tmp = cur.Next

		cur.Next = pre
		pre = cur

		cur = tmp
	}

	head.Next = nil
	head = pre

	return head
}
