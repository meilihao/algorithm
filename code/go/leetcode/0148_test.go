/*
思路:
1. 遍历list, 获取元素后重新构建list
2. 归并排序
*/
package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortList(t *testing.T) {
	a := []int{4, 2, 1, 3}
	head := GenerateListNode(RevertInts(a))
	PrintListNode(head)

	head1 := sortList(head)
	PrintListNode(head1)
}

func sortList(head *ListNode) *ListNode {
	//return sortList2(head)
	//return sortList1_2(head)
	return sortList2(head)
}

func sortList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	vals := []int{}
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	sort.Ints(vals)

	fmt.Println(vals)

	var nhead *ListNode
	if len(vals) > 0 {
		nhead = &ListNode{Val: vals[0]}
	}
	cur := nhead
	for i := 1; i < len(vals); i++ {
		cur.Next = &ListNode{Val: vals[i]}
		cur = cur.Next
	}

	return nhead
}

// best
func sortList1_2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	vals := []int{}
	cur := head
	for cur != nil {
		vals = append(vals, cur.Val)
		cur = cur.Next
	}
	sort.Ints(vals)

	// 延用原有的链表空间
	cur = head
	idx := 0
	for cur != nil {
		cur.Val = vals[idx]
		cur = cur.Next
		idx++
	}

	return head
}

// 归并排序
func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := findMiddle(head)
	tail := mid.Next // 右半部分的头节点
	mid.Next = nil   // 断开链表

	left := sortList2(head)
	right := sortList2(tail)

	return mergeTwoLists(left, right)
}

// 快慢指针找链表的中点
func findMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head.Next // fast 从 head.next 开始，确保 slow 指向中点或左中点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// todo: 官方方法二：自底向上归并排序
func sortList3(head *ListNode) *ListNode {
	return nil
}
