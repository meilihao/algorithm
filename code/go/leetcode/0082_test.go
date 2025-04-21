package leetcode

import "testing"

func TestDeleteDuplicates82(t *testing.T) {

}

func deleteDuplicates82(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{0, head} // 哨兵节点, 可能需要删除第一个节点

	cur := dummy
	var tmp int
	for cur.Next != nil && cur.Next.Next != nil { // 处理cur后续的两个节点
		if cur.Next.Val == cur.Next.Next.Val {
			tmp = cur.Next.Val
			for cur.Next != nil && cur.Next.Val == tmp { // 删除重复节点
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
