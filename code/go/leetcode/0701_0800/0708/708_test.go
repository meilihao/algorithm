/*
143.中 重排链表

给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1：

输入：head = [1,2,3,4]
输出：[1,4,2,3]
示例 2：

输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]

提示：

链表的长度范围为 [1, 5 * 104]
1 <= node.val <= 1000
*/
package leetcode

import (
	"fmt"
	"testing"

	. "al/leetcode"
)

func TestReverseList(t *testing.T) {
	head := &Node{Val: 3}
	n1 := &Node{Val: 3}
	n2 := &Node{Val: 3}
	n3 := &Node{Val: 0}

	head.Next = n1
	n1.Next = n2
	n2.Next = n3

	fmt.Println(insert(head, 3))
}

func insert(head *Node, insertVal int) *Node {
	node := &Node{Val: insertVal}
	if head == nil { // 空链表
		node.Next = node
		return node
	}
	if head.Next == head { // 只有一个节点
		head.Next = node
		node.Next = head
		return head
	}
	curr, next := head, head.Next
	for next != head { // 遍历一圈即可
		if insertVal >= curr.Val && insertVal <= next.Val { // curr.Val <=next.Val // 新节点的值介于循环链表中的两个节点值之间，在 curr 和 next 之间插入新节点
			break
		}
		if curr.Val > next.Val { // curr 和 next 分别是循环链表中的值最大的节点和值最小的节点
			if insertVal > curr.Val || insertVal < next.Val {
				break
			}
		}
		curr = curr.Next
		next = next.Next
	}
	curr.Next = node
	node.Next = next
	return head
}

func insert2(head *Node, insertVal int) *Node {
	node := &Node{Val: insertVal}
	if head == nil { // 空链表
		node.Next = node
		return node
	}
	if head.Next == head { // 只有一个节点
		head.Next = node
		node.Next = head
		return head
	}
	curr, next := head, head.Next
	biggest := head
	for next != head && !(insertVal >= curr.Val && insertVal <= next.Val) {
		curr = next
		next = next.Next
		if curr.Val >= biggest.Val {
			biggest = curr
		}
	}

	if insertVal >= curr.Val && insertVal <= next.Val {
		curr.Next = node
		node.Next = next
	} else { // 查到biggest后面
		node.Next = biggest.Next
		biggest.Next = node
	}

	return head
}
