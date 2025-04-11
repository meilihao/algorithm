/*
	Add Two Numbers 两数相加

	for 单链表
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	Test()
	TestTwoZero()
	TestTwoFive()
}

func Test() {
	num1 := []int{3, 2} // []int{3, 4, 2}
	num2 := []int{4, 6, 5}

	hnum1 := generateNumber(num1)
	printNum(hnum1)

	hnum2 := generateNumber(num2)
	printNum(hnum2)

	hnum3 := addTwoNumbers(hnum1, hnum2)
	printNum(hnum3)
}

func TestTwoZero() {
	num1 := []int{0}
	num2 := []int{0}

	hnum1 := generateNumber(num1)
	printNum(hnum1)

	hnum2 := generateNumber(num2)
	printNum(hnum2)

	hnum3 := addTwoNumbers(hnum1, hnum2)
	printNum(hnum3)
}

func TestTwoFive() {
	num1 := []int{5}
	num2 := []int{5}

	hnum1 := generateNumber(num1)
	printNum(hnum1)

	hnum2 := generateNumber(num2)
	printNum(hnum2)

	hnum3 := addTwoNumbers(hnum1, hnum2)
	printNum(hnum3)
}

func generateNumber(num []int) *ListNode {
	var head *ListNode

	for _, v := range num {
		tmp := &ListNode{
			Val: v,
		}

		if head == nil {
			head = tmp

			continue
		}

		tmp.Next = head
		head = tmp
	}

	return head
}

func printNum(head *ListNode) {
	for head != nil {
		if head.Next != nil {
			fmt.Printf("%d -> ", head.Val)

			head = head.Next
		} else {
			fmt.Print(head.Val)

			break
		}
	}

	fmt.Println()
}

// 两个链表长度不一致时，要处理较长链表剩余的高位和进位计算的值
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, pre *ListNode // pre: 上一个节点
	var carry int           // carry是否进位

	for {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}

		// must have
		tmp := &ListNode{}

		if l1 != nil {
			tmp.Val += l1.Val

			l1 = l1.Next
		}

		if l2 != nil {
			tmp.Val += l2.Val

			l2 = l2.Next
		}

		tmp.Val += carry

		carry = tmp.Val / 10
		tmp.Val %= 10

		if pre == nil {
			pre = tmp
			head = tmp

			continue
		}

		pre.Next = tmp
		pre = tmp
	}

	return head
}
