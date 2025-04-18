package leetcode

import (
	"testing"
)

func TestSwapPairs(t *testing.T) {
	s := "1234"
	head := GenerateListNodeByChars(s)
	PrintListNode(head, true)

	head2 := swapPairs2(head)
	PrintListNode(head2, true)
}

// 需要三个指针: 相邻元素 + 相邻元素前面的一个元素(用于指向后两个相邻元素)
// best
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{ // 哨兵节点
		Next: head,
	}

	pre := dummyHead
	var cur, next *ListNode

	for pre.Next != nil && pre.Next.Next != nil { // 有元素需要交换
		cur = pre.Next
		next = cur.Next

		pre.Next, next.Next, cur.Next = next, cur, next.Next

		pre = cur
	}

	return dummyHead.Next
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode // 相邻元素前面的一个元素
	cur, next := head, head.Next

	head = next
	for cur != nil && next != nil {
		if pre != nil {
			pre.Next = next
		}

		cur.Next, next.Next, pre = next.Next, cur, cur

		if cur.Next == nil { // 偶数节点,即后面还有节点
			return head
		}

		cur = cur.Next
		next = cur.Next
	}

	return head
}
