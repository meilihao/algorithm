package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	s := "123"
	head := BuildList(s)
	printList(head)
	fmt.Println("---")

	removedHead := removeNthFromEnd2(head, 3)
	printList(removedHead)
}

func BuildList(s string) *ListNode {
	var head, pre, tmp *ListNode

	for _, v := range s {
		tmp = &ListNode{
			Val: int(v),
		}

		if pre != nil {
			pre.Next = tmp
			pre = tmp
		} else {
			pre = tmp
			head = tmp
		}
	}

	return head
}

func printList(head *ListNode) {
	if head == nil {
		fmt.Println("empty list")

		return
	}

	tmp := head
	for tmp != nil {
		fmt.Printf("%c\n", tmp.Val)

		tmp = tmp.Next
	}
}

// 简化成另一个问题：删除从列表开头数起的第 (L - n + 1) 个结点，其中 L 是列表的长度, 其前一个节点为L-n
// 使用前后指针(均以head为起点,即起点为1),后指针走到L-n需要L-n-1步, 因此同步走时前指针需前进L-n-1步, 此时前指针在最后一个节点上, 那么此前前指针应在L-(L-n-1)=n+1上.
// 具体方法: 前指针先走n步, 此时前指针在n+1上, 还剩L-n-1步到最后一个节点. 此时前指针和后指针同时走，当前指针走到表尾节点时，后指针刚好走到要删除元素的前一个位置上
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}

	first := head
	for i := 1; i <= n && first != nil; i++ {
		first = first.Next
	}

	// first应在n+1上
	if first == nil { // 删除第一个节点(删除不存在的节点时,也是删除第一个节点即可), L - n + 1 => L - L + 1 = 1
		return head.Next
	}

	second := head
	for first.Next != nil { // 到达最后一个节点
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next

	return head
}

// best
// 加哨兵节点
// 倒数第N个节点 -> 正数第(L-N+1)个节点-> 要删除链表的节点就必须知道该节点的前一个节点: L-N
// 假设第(L-N)个节点有指针second,那么此时first应该在nil上, 两者相差N+1步. 因此first应先走N+1步, 之后再同步走.
// 因为是从head开始走, 只需要走N步, 又因为加了哨兵节点, 所以还是要走N+1步.
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
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

	// fmt.Println(first, i)

	if i > 0 { // n>L
		return dummy.Next
	}

	second := dummy
	// fmt.Println(first, second)
	for first != nil { // 到达最后一个节点
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next

	return dummy.Next
}
