/*
142.中 环形链表 II

给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

不允许修改 链表。

示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。
示例 2：

输入：head = [1,2], pos = 0
输出：返回索引为 0 的链表节点
解释：链表中有一个环，其尾部连接到第一个节点。
示例 3：

输入：head = [1], pos = -1
输出：返回 null
解释：链表中没有环。

提示：

链表中节点的数目范围在范围 [0, 104] 内
-105 <= Node.val <= 105
pos 的值为 -1 或者链表中的一个有效索引
*/
package main

import (
	"fmt"
	"testing"

	. "al/leetcode"
)

func TestDetectCycle(t *testing.T) {
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
			// X----<a>----Y---<b>-----Z
			//             |----<c>----|
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
