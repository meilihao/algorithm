package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	s := "abc"
	head := BuildList(s)

	fmt.Println(isPalindrome(head))
}

// func BuildList(s string) *ListNode {
// 	var head, pre, tmp *ListNode

// 	for _, v := range s {
// 		tmp = &ListNode{
// 			Val: int(v),
// 		}

// 		if pre != nil {
// 			pre.Next = tmp
// 			pre = tmp
// 		} else {
// 			pre = tmp
// 			head = tmp
// 		}
// 	}

// 	return head
// }

// slow:fast => fast=2*slow-1
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head

	var pre, tmp *ListNode // 用于翻转链表

	for fast != nil && fast.Next != nil { // 确保fast能跳两格
		fast = fast.Next.Next

		tmp = slow.Next

		slow.Next = pre //翻转链表方向
		pre = slow      // pre已翻转完成

		slow = tmp
	}

	if fast != nil { // 链表长度为奇数, 画图可清晰了解
		slow = slow.Next
	}

	for pre != nil && slow != nil && pre.Val == slow.Val {
		pre = pre.Next
		slow = slow.Next
	}

	return pre == nil && slow == nil
}
