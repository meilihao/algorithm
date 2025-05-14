/*
82.中 删除排序链表中的重复元素 II

给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。

示例 1：

输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]
示例 2：

输入：head = [1,1,1,2,3]
输出：[2,3]
*/
package leetcode

import (
	"testing"

	. "al/leetcode"
)

// 重复的元素在链表中出现的位置是连续的, 因此重复的元素在链表中出现的位置是连续的
func TestDeleteDuplicates82(t *testing.T) {

}

func deleteDuplicates82(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{0, head} // 哨兵节点, 可能需要删除第一个节点

	cur := dummy
	var tmp int
	for cur.Next != nil && cur.Next.Next != nil { // 处理cur后续的相连两个节点
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
