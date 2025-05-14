/*
61.中 旋转链表

给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

示例 1：

输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]
示例 2：

输入：head = [0,1,2], k = 4
输出：[2,0,1]
*/
package leetcode

import (
	"testing"

	. "al/leetcode"
)

func TestRotateRight(t *testing.T) {

}

/*
思路: 闭合为环

记给定链表的长度为 n，注意到当向右移动的次数 k≥n 时，实际仅需要向右移动 kmodn 次即可.
因为每 n 次移动都会让链表变为原状。这样可以知道，新链表的最后一个节点为原链表的第 (n−1)−(k mod n) 个节点（从 0 开始计数）= 第 n−(k mod n) 个节点（从 1 开始计数）
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	n := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}

	add := n - k%n // 移动次数
	if add == n {  // k%n==0, 原样
		return head
	}
	iter.Next = head // 闭合为环. iter当前是尾节点, 即从尾开始移动
	for add > 0 {
		iter = iter.Next
		add--
	}
	// 断开
	ret := iter.Next
	iter.Next = nil
	return ret
}
