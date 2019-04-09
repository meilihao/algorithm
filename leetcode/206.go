package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	s := "abc"
	head := BuildList(s)

	reversedHead := reverseList(head)
	printList(reversedHead)
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

func printList(head *ListNode) {
	if head == nil {
		fmt.Println("empty list")

		return
	}

	tmp := head
	for tmp != nil {
		fmt.Printf("%c\n", tmp.Val)

		tmp = tmp.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	var pre, tmp *ListNode
	for cur != nil {
		tmp = cur.Next

		cur.Next = pre //翻转链表方向
		pre = cur      // pre已翻转完成

		cur = tmp
	}

	return pre
}
