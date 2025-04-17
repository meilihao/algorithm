package leetcode

import "testing"

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
