// 25. K 个一组翻转链表
//
// 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	s := "12345"
	head := BuildList(s)

	head2 := reverseKGroup(head, 3)
	printList(head2)
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

// 链表分区为已翻转部分+待翻转部分+未翻转部分
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 || head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{ // 哨兵节点
		Next: head,
	}

	pre, end := dummy, dummy // pre 代表待翻转链表的前驱，end 代表待翻转链表的末尾
	var i int

	var start, next *ListNode // start 待翻转链表的head, next保存未处理链表的head
	for end.Next != nil {
		for i = 0; i < k && end != nil; i++ {
			end = end.Next
		}
		// 不够k个一组
		if end == nil {
			break
		}

		next = end.Next

		// 构造一个需要翻转的链表
		start = pre.Next
		end.Next = nil

		pre.Next = reverse(start) // pre指向翻转好的链表

		// 翻转后变成: pre -> end -> ... -> start
		start.Next = next // 指向待处理的链表
		pre = start
		end = start
	}

	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}

	return pre
}

// todo use stack
