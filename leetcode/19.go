package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	s := "12"
	head := BuildList(s)

	removedHead := removeNthFromEnd(head, 2)
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
// 使用前后指针(均以head为起点): 后指针需前进L-n-1步, 此时后指针的下一个节点即为要删除的节点, 此时前指针应在L-(L-n-1)=n+1上.
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
