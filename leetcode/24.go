package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	s := "1234"
	head := BuildList(s)

	head2 := swapPairs2(head)
	printList(head2)
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

// 需要三个指针: 相邻元素 + 相邻元素前面的一个元素
// best
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{ // 哨兵节点
		Next: head,
	}

	pre := dummyHead
	var cur, next *ListNode

	for pre.Next != nil && pre.Next.Next != nil {
		cur = pre.Next
		next = cur.Next

		pre.Next, next.Next, cur.Next = next, cur, next.Next

		pre = cur
	}

	return dummyHead.Next
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode // 相邻元素前面的一个元素
	cur, next := head, head.Next

	head = next
	for cur != nil && next != nil {
		if pre != nil {
			pre.Next = next
		}

		cur.Next, next.Next, pre = next.Next, cur, cur

		if cur.Next == nil { // 偶数节点,即后面还有节点
			return head
		}

		cur = cur.Next
		next = cur.Next
	}

	return head
}
