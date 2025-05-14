/*
92.中 反转链表 II

给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

示例 1：

输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]
示例 2：

输入：head = [5], left = 1, right = 1
输出：[5]
*/
package leetcode

import (
	"fmt"
	"testing"

	. "al/leetcode"
)

func TestReverseBetween(t *testing.T) {
	head := GenerateListNode(RevertInts([]int{9, 7, 2, 5, 4, 3, 6}))
	PrintListNode(head)

	rh := reverseBetween(head, 2, 3)
	//rh := reverseBetweenWithValue(head, 2, 3)
	PrintListNode(rh)
}

// left, right是链表节点的索引, 从1开始
/*
curr：永远指向待反转区域的第一个节点 left；
next：永远指向 curr 的下一个节点，循环过程中，curr 变化以后 next 会变化；
pre：永远指向待反转区域的第一个节点 left 的前一个节点，在循环过程中不变
*/
func reverseBetween(head *ListNode, left, right int) *ListNode {
	// 设置 dummyNode 是这一类问题的一般做法
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	var next *ListNode
	for i := 0; i < right-left; i++ {
		next = cur.Next // 暂存

		// `pre->cur->next->next.next`=>`pre->next->cur->next.next`
		cur.Next = next.Next // cur指向next.Next
		next.Next = pre.Next // 将next插到cur前
		pre.Next = next      // pre指向next
	}
	return dummyNode.Next
}

// left, right是ListNode.Val, 且不重复, 扩展, 非本题解法
/*
curr：永远指向待反转区域的第一个节点 left；
next：永远指向 curr 的下一个节点，循环过程中，curr 变化以后 next 会变化；
pre：永远指向待反转区域的第一个节点 left 的前一个节点，在循环过程中不变
*/
func reverseBetweenWithValue(head *ListNode, left, right int) *ListNode {
	// 设置 dummyNode 是反转这一类问题的一般做法, 解决left是第一个节点的问题
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for pre != nil && pre.Next != nil { //  pre.Next is match node for left
		if pre.Next.Val == left {
			break
		}

		pre = pre.Next
	}

	cur := pre.Next
	if cur == nil { // no found left
		return dummyNode.Next
	}

	next := cur.Next

	fmt.Println(pre.Val, cur.Val)

	var stop bool

	for next != nil { // 找不到就一直反转
		if next.Val == right {
			stop = true
		}

		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next

		next = cur.Next

		if stop { // 需变更完right所在节点再停止
			break
		}
	}

	return dummyNode.Next
}
