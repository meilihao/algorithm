package leetcode

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	s := "12"
	s2 := "1345"
	head := GenerateListNodeByChars(s)
	head2 := GenerateListNodeByChars(s2)

	mergedHead := mergeTwoLists(head, head2)
	PrintListNode(mergedHead, true)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummy := &ListNode{} // 虚拟头节点 as 哨兵节点
	head := dummy

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			head.Next = l2
			l2 = l2.Next
		} else {
			head.Next = l1
			l1 = l1.Next
		}

		head = head.Next
	}

	// 将剩余的链表接上, 最多一个链表有剩余即l1或l2为nil
	if l1 != nil {
		head.Next = l1
	} else {
		head.Next = l2
	}

	return dummy.Next
}
