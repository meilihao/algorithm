package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	s := "abcd"
	head := BuildList(s)

	node := middleNode(head)
	if node != nil {
		fmt.Printf("%c\n", node.Val)
	} else {
		fmt.Println("empty list")
	}
}

func BuildList(s string) *ListNode {
	var head, pre, tmp *ListNode

	for _, v := range s {
		tmp = &ListNode{
			Val: int(v),
		}

		if pre != nil {
			pre.Next = tmp
			pre = tmp
		} else {
			pre = tmp
			head = tmp
		}
	}

	return head
}

// slow:fast => fast=2*slow-1
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil { // 确保fast能跳两格
		fast = fast.Next.Next

		slow = slow.Next
	}

	return slow
}
