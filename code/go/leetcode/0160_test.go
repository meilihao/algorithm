package leetcode

import "testing"

func TestGetIntersectionNode(t *testing.T) {

}

/*
本质: 跑步相遇问题

假设相交点为node, 两链表的公共尾部的节点数量为 c, len(headA)=a, len(headB)=b那么
1. 如何让等式成立: `X + (a - c) = Y + (b - c)`
1. 指针 A 先遍历完链表 headA ，再开始遍历链表 headB ，当走到 node 时，共走步数为：a+(b−c)
1. 指针 B 先遍历完链表 headB ，再开始遍历链表 headA ，当走到 node 时，共走步数为：b+(a−c)
1. a+(b−c) = b+(a−c): 此时所处位置相同

其实就是大家都跑a+b的长度, 如果相交, 那么此时所处位置相同

时间复杂度 O(a+b) ： 最差情况下（即 ∣a−b∣=1 , c=0 ），此时需遍历 a+b 个节点。
空间复杂度 O(1) ： 节点指针 A , B 使用常数大小的额外空间
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}

	return pa
}
