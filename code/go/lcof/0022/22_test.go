/*
22：链表中倒数第k个结点

输入一个链表，输出该链表中倒数第k个结点。为了符合大多数人的习惯，
本题从1开始计数，即链表的尾结点是倒数第1个结点。例如一个链表有6个结点，
从头结点开始它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个结点是
值为4的结点。
*/
package demo

import (
	"fmt"
	"testing"

	. "al/leetcode"
)

func TestGetNthFromEnd(t *testing.T) {
	s := "123456"
	head := GenerateListNodeByChars(s)
	PrintListNode(head, true)
	fmt.Println("---")

	node := getNthFromEnd(head, 7)
	fmt.Println(node, node == nil)

	node = getNthFromEnd(head, 6)
	fmt.Println(node, fmt.Sprintf("%c", node.Val))

	node = getNthFromEnd(head, 3)
	fmt.Println(node, fmt.Sprintf("%c", node.Val))
}

// best
// 加哨兵节点
// 倒数第N个节点 -> 正数第(L-N+1)个节点-> 要删除链表的节点就必须知道该节点的前一个节点: L-N
// 假设第(L-N)个节点有指针second,那么此时first应该在nil上, 两者相差N+1步(k+最后的nil). 因此first应先走N+1步, 之后再同步走.
func getNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}

	dummy := &ListNode{ // 哨兵节点
		Next: head,
	}

	first := dummy

	i := n + 1
	for ; i > 0 && first != nil; i-- {
		first = first.Next
	}

	if i > 0 { // n>L
		return nil
	}

	fmt.Println("first", fmt.Sprintf("%#v", first))

	second := dummy
	for first != nil { // 到达最后一个节点
		first = first.Next
		second = second.Next
	}

	return second.Next
}
