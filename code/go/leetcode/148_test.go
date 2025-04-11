/*
思路:
1. 遍历list, 获取元素后重新构建list
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

	head1 := sortList1(head)
	PrintListNode(head1)
}

func sortList1(head *ListNode) *ListNode {
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
