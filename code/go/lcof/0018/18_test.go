/*
在O(1)时间删除链表结点

给定单向链表的头指针和一个结点指针，定义一个函数在O(1)时间删除该
结点。
*/
package demo

import (
	. "al/leetcode"
)

// 其实是假定了p一定在链表head中. 因为需要O(n)判断链表中是否包含p, 与题目冲突
// 时间复杂度: [(n-1)*O(1)+O(n)]/n=O(1)
func DeleteNode(head *ListNode, p *ListNode) *ListNode {
	if head == nil || p == nil {
		return head
	}

	// p不是尾节点
	if p.Next != nil {
		// 实际是用p.Next的内容替换到p, 再删除p.Next
		tmp := p.Next
		p.Val = tmp.Val
		p.Next = tmp.Next
	} else if head == p { // 链表只有一个节点
		head = nil
	} else { // 多个节点中, 删除尾节点
		tmp := head
		for tmp.Next != p {
			tmp = tmp.Next
		}

		tmp.Next = nil
	}

	return head
}
