/*
24.中 两两交换链表中的节点

给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

示例 1：

输入：head = [1,2,3,4]
输出：[2,1,4,3]
示例 2：

输入：head = []
输出：[]
示例 3：

输入：head = [1]
输出：[1]
*/
package leetcode

import (
	"testing"

	. "al/leetcode"
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

// 两个临时变量cur,next
func swapPairs_1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{ // 哨兵节点
		Next: head,
	}

	pre := dummyHead
	var cur, next *ListNode

	// 处理两边再中间
	// pre->1->2->3 => (pre->[2)->(1]->3)
	for pre.Next != nil && pre.Next.Next != nil { // 有元素需要交换
		cur = pre.Next
		next = cur.Next

		// 这里处理先首尾, 或先尾首都可以
		pre.Next = next      // pre->2
		cur.Next = next.Next // 1->3
		next.Next = cur      // 2->1
		pre = cur
	}

	return dummyHead.Next
}

// pre->1->2->3 => pre->2->1->3 : 逆推: 因为要修改1.next, 而原1.next指向2, 所以要暂存2
// 一个临时变量tmp
func swapPairs3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{ // 哨兵节点
		Next: head,
	}

	pre := dummyHead
	var tmp *ListNode

	for pre.Next != nil && pre.Next.Next != nil { // 有元素需要交换
		tmp = pre.Next.Next // 2

		pre.Next.Next = tmp.Next // 1->3
		tmp.Next = pre.Next      // 2 -> 1
		pre.Next = tmp           // pre ->2

		pre = pre.Next.Next
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
