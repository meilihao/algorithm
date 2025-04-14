package lcci

import "testing"

func TestKthToLast(t *testing.T) {
	kthToLast(nil, 1)
}

// 链表中倒数第 k 个节点也就是正数第(L-K+1)个节点
// 双指针法:
// 1. 一个节点 fast 先开始跑，指针 fast 跑到 k-1 个节点后
// 1. 另一个节点 slow 开始跑，当 fast 跑到最后时，slow 所指的节点就是倒数第 k 个节点也就是正数第(L-K+1)个节点
// --- 易理解:
// 初始化双指针 slow , fast 都指向头节点 head ；
// 先令 fast 走 k 步，此时 slow , fast 的距离为 k ；
// 令 slow , fast 一起走，直到 fast 走过尾节点时跳出，此时 slow 指向「倒数第 k 个节点」，返回之即可
func kthToLast(head *ListNode, k int) int {
	slow, fast := head, head
	for ; k > 0; k-- {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow.Val
}
