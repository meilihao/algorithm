/*
143.中 重排链表

给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1：

输入：head = [1,2,3,4]
输出：[1,4,2,3]
示例 2：

输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]

提示：

链表的长度范围为 [1, 5 * 104]
1 <= node.val <= 1000
*/
package leetcode

import (
	"fmt"
	"testing"

	. "al/leetcode"
)

func TestReverseList(t *testing.T) {
	s := "12345"
	//s := "1234"

	head := GenerateListNodeByChars(s)
	PrintListNode(head, true)

	reorderList(head)
	PrintListNode(head, true)
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	mid := findMidNode(head)
	l1 := head
	l2 := mid.Next
	fmt.Printf("%c\n", mid.Val)
	mid.Next = nil // 断开链表
	l2 = ReverseList(l2)
	mergeList(l1, l2)
}

// len(head)=奇数, return mid; 偶数, return 后半部的前一个节点
func findMidNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 两链表长度相差不超过 1，因此直接合并即可
func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}
