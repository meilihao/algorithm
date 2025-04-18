package leetcode

import "testing"

func TestRemoveElements(t *testing.T) {

}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	for pre := dummyHead; pre.Next != nil; { // 找pre
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return dummyHead.Next
}
