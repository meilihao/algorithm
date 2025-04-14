/*
Add Two Numbers 两数相加

for 单链表
*/
package leetcode

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	TwoNumbers()
	TwoZero()
	TwoFive()
}

func TwoNumbers() {
	num1 := []int{3, 2} // []int{3, 4, 2}
	num2 := []int{4, 6, 5}

	hnum1 := GenerateListNode(num1)
	PrintListNode(hnum1)

	hnum2 := GenerateListNode(num2)
	PrintListNode(hnum2)

	hnum3 := addTwoNumbers(hnum1, hnum2)
	PrintListNode(hnum3)
}

func TwoZero() {
	num1 := []int{0}
	num2 := []int{0}

	hnum1 := GenerateListNode(num1)
	PrintListNode(hnum1)

	hnum2 := GenerateListNode(num2)
	PrintListNode(hnum2)

	hnum3 := addTwoNumbers(hnum1, hnum2)
	PrintListNode(hnum3)
}

func TwoFive() {
	num1 := []int{5}
	num2 := []int{5}

	hnum1 := GenerateListNode(num1)
	PrintListNode(hnum1)

	hnum2 := GenerateListNode(num2)
	PrintListNode(hnum2)

	hnum3 := addTwoNumbers(hnum1, hnum2)
	PrintListNode(hnum3)
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
