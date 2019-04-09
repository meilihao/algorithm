package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	s := "12"
	s2 := "1345"
	head := BuildList(s)
	head2 := BuildList(s2)

	mergedHead := mergeTwoLists(head, head2)
	printList(mergedHead)
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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	new := new(ListNode) // 哨兵节点
	pre := new

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			pre.Next = l2
			l2 = l2.Next
		} else {
			pre.Next = l1
			l1 = l1.Next
		}

		pre = pre.Next
	}

	// l1或l2为nil
	if l1 != nil {
		pre.Next = l1
	} else {
		pre.Next = l2
	}

	return new.Next
}
