/*
206.简 反转链表

给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

示例 1：

输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
示例 2：

输入：head = [1,2]
输出：[2,1]
示例 3：

输入：head = []
输出：[]

提示：

链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000
*/
package leetcode

import (
	"testing"

	. "al/leetcode"
)

func TestReverseList(t *testing.T) {
	s := "abc"

	head := GenerateListNodeByChars(s)
	PrintListNode(head, true)

	reversedHead := reverseList(head)
	PrintListNode(reversedHead, true)
}

// best
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	var pre, next *ListNode
	for cur != nil {
		next = cur.Next // 保存要反转到头的那个节点

		cur.Next = pre //翻转链表方向, 第一次反转的时候会指向null

		pre = cur  // 上一个已经反转到头部的节点
		cur = next // 一直向链表尾走
	}

	return pre
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	var pre *ListNode

	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}

	return pre
}

// 不推荐
func reverseList3(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// if head.Next==nil{
	// 	return head
	// }

	pre, cur := head, head.Next
	var tmp *ListNode

	for cur != nil {
		tmp = cur.Next

		cur.Next = pre
		pre = cur

		cur = tmp
	}

	head.Next = nil // head变为了末尾节点, 因为它仍指向原先的第二个节点. 反转后该节点已指向末尾节点head, 导致循环了, 需要打破
	head = pre      // 设置新head

	return head
}
