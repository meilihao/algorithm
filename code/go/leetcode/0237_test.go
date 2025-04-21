package leetcode

import "testing"

func TestDeleteNode(t *testing.T) {

}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
