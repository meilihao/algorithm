/*
2.中 两数相加 : Add Two Numbers

给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例 1：

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
示例 2：

输入：l1 = [0], l2 = [0]
输出：[0]
示例 3：

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
*/
package leetcode

import (
	"testing"

	. "al/leetcode"
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
