package leetcode

import "fmt"

// --- array
func RevertInts(a []int) []int {
	for i, j := 0, len(a)-1; i < j; {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}

	return a
}

// --- tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- list
type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateListNode(nums []int) *ListNode {
	var head *ListNode = nil

	for i := 0; i < len(nums); i++ {
		cur := &ListNode{
			Val: nums[i],
		}

		if head == nil {
			head = cur

			continue
		}

		cur.Next = head
		head = cur
	}

	return head
}

func GenerateListNodeByChars(s string) *ListNode {
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

func PrintListNode(head *ListNode, isChars ...bool) {
	if head == nil {
		fmt.Println("empty list")

		return
	}

	var isChar bool
	if len(isChars) > 0 {
		isChar = isChars[0]
	}

	for head != nil {
		if isChar {
			fmt.Printf("%c", head.Val)
		} else {
			fmt.Print(head.Val)
		}

		head = head.Next
		if head != nil {
			fmt.Print("->")
		}
	}
	fmt.Println()
}
