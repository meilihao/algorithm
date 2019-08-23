package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(hasCycle(nil))
}

// map缓存节点, 判重
func hasCycleByHash(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	m := make(map[*ListNode]bool)

	cur := head
	for cur != nil {
		if m[cur] {
			return true
		}

		m[cur] = true
		cur = cur.Next
	}

	return false
}

// 抽象成两名运动员以不同的速度在环形赛道上跑步, 有环总会相遇
// 没环: 时间复杂度= n/2
// 有环:　迭代次数约＝n/1，　ｎ为两者最大差距，　因此时间复杂度为ｎ
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}
