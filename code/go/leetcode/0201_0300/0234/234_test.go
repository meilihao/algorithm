/*
234.简单 回文链表

给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

示例 1：

输入：head = [1,2,2,1]
输出：true
示例 2：

输入：head = [1,2]
输出：false

提示：

链表中节点数目在范围[1, 105] 内
0 <= Node.val <= 9
*/
package main

import (
	"fmt"
	"testing"

	. "al/leetcode"
)

func TestIsPalindrome(t *testing.T) {
	s := "abca"
	head := GenerateListNodeByChars(s)

	//fmt.Println(isPalindrome(head))
	fmt.Println(isPalindrome2(head))
}

// slow:fast => fast=2*slow-1
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head

	var pre, next *ListNode // 用于翻转链表

	// 在找中点的过程中同步反转前半部分
	// for结束时：pre 指向反转后的前半部分链表头; slow 指向后半部分链表头（或中间节点）
	for fast != nil && fast.Next != nil { // 确保fast能跳两格
		fast = fast.Next.Next

		//翻转链表方向
		next = slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}

	if fast != nil { // 链表长度为奇数, slow 需要跳过中间节点, 画图可清晰了解
		slow = slow.Next
	}

	for pre != nil && slow != nil && pre.Val == slow.Val {
		pre = pre.Next
		slow = slow.Next
	}

	return pre == nil && slow == nil
}

// 反转后半部分
func isPalindrome2(head *ListNode) bool {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil { // 当fast==nil, slow刚好在链表中间或在前半部分的最后一个节点
		slow, fast = slow.Next, fast.Next.Next
	}
	fmt.Printf("%c\n", slow.Val)

	var pre *ListNode // 最终会指向反转后的后半部分链表的头节点
	cur := slow.Next
	for cur != nil {
		t := cur.Next
		cur.Next = pre
		pre = cur
		cur = t
	}
	for pre != nil {
		if pre.Val != head.Val {
			return false
		}
		pre, head = pre.Next, head.Next
	}
	return true
}
