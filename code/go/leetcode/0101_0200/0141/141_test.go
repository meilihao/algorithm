/*
141.简 环形链表

给你一个链表的头节点 head ，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。

如果链表中存在环 ，则返回 true 。 否则，返回 false 。

示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。
示例 2：

输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。
示例 3：

输入：head = [1], pos = -1
输出：false
解释：链表中没有环。

提示：

链表中节点的数目范围是 [0, 104]
-105 <= Node.val <= 105
pos 为 -1 或者链表中的一个 有效索引
*/
package main

import (
	"fmt"
	"testing"

	. "al/leetcode"
)

func TestHasCycle(t *testing.T) {
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
