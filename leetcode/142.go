// 142. 环形链表 II : 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(detectCycle(nil))
}

// 抽象成两名运动员以不同的速度在环形赛道上跑步, 有环总会相遇, ｎ为节点个数
// 没环: 时间复杂度= n/2
// 有环:　迭代次数约＝n/1(1为速度差), 即慢指针移动n次必定相遇, 因此时间复杂度为ｎ
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			// https://github.com/TomorrowWu/golang-algorithms/tree/master/leetcode/0142.linked-list-cycle-ii
			// X---------Y---------Z
			// x起点，y环的起点，z是相遇点(fast比slow多走一圈)
			// 链表头是X，环的第一个节点是Y，slow和fast第一次的交点是Z. 各段的长度分别是a,b,c
			// 相遇时，慢节点走了：a+b; 由于快指针速度是慢指针的2倍，快节点走了：2(a+b); 同时快节点比慢节点刚好多走了一圈环形节点, 快节点走了：(a+b)+(c+b)
			// fast走过的距离为 2(a+b)=a+b+c+b => a=c
			// 则slow从head开始走，fast从Z点开始走，速度都是一次走一步，slow和fast同时到达Y点，即环的入口
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}

	return nil
}
