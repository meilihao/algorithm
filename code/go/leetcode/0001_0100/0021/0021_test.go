/*
21. 合并两个有序链表

将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：

输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
示例 2：

输入：l1 = [], l2 = []
输出：[]
示例 3：

输入：l1 = [], l2 = [0]
输出：[0]
*/
package leetcode

import (
	"testing"

	. "al/leetcode"
)

func TestMergeTwoLists(t *testing.T) {
	s := "12"
	s2 := "1345"
	head := GenerateListNodeByChars(s)
	head2 := GenerateListNodeByChars(s2)

	mergedHead := mergeTwoLists(head, head2)
	PrintListNode(mergedHead, true)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return MergeTwoLists(l1, l2)
}
