package leetcode

import (
	"fmt"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	s := "abcd"
	head := GenerateListNodeByChars(s)

	node := middleNode(head)
	if node != nil {
		fmt.Printf("%c\n", node.Val)
	} else {
		fmt.Println("empty list")
	}
}

// slow:fast => fast=2*slow-1
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	fast := head // 如果 fast 从 head.next 开始，则 slow 指向中点或左中点, 也题目不符合

	for fast != nil && fast.Next != nil { // 确保fast能跳两格
		fast = fast.Next.Next

		slow = slow.Next
	}

	return slow
}
