/*
83.简 删除排序链表中的重复元素

题目描述
给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。

示例 1：

输入：head = [1,1,2]
输出：[1,2]
示例 2：

输入：head = [1,1,2,3,3]
输出：[1,2,3]
*/
package leetcode

import (
	"testing"

	. "al/leetcode"
)

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
