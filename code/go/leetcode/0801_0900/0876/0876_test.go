/*
876.简 链表的中间结点

给你单链表的头结点 head ，请你找出并返回链表的中间结点。

如果有两个中间结点，则返回第二个中间结点。

示例 1：

输入：head = [1,2,3,4,5]
输出：[3,4,5]
解释：链表只有一个中间结点，值为 3 。
示例 2：

输入：head = [1,2,3,4,5,6]
输出：[4,5,6]
解释：该链表有两个中间结点，值分别为 3 和 4 ，返回第二个结点。

提示：

链表的结点数范围是 [1, 100]
1 <= Node.val <= 100
*/
package leetcode

import (
	"fmt"
	"testing"

	. "al/leetcode"
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
