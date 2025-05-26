package leetcode

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
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

// from 题206
func ReverseList(head *ListNode) *ListNode {
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

// from 题2
func AddTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, pre *ListNode // pre: 上一个节点
	var carry int           // carry是否进位

	for {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}

		// must have
		tmp := &ListNode{}

		if l1 != nil {
			tmp.Val += l1.Val

			l1 = l1.Next
		}

		if l2 != nil {
			tmp.Val += l2.Val

			l2 = l2.Next
		}

		tmp.Val += carry

		carry = tmp.Val / 10
		tmp.Val %= 10

		if pre == nil {
			pre = tmp
			head = tmp

			continue
		}

		pre.Next = tmp
		pre = tmp
	}

	return head
}
