package leetcode

import "testing"

func TestDeleteDuplicates(t *testing.T) {

}

// 给定的链表是排好序的，因此重复的元素在链表中出现的位置是连续的
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head // 已处理好的链表的最后一个节点
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}
