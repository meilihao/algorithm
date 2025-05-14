/*
328.中 奇偶链表

给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。

第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。

请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。

你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。

示例 1:

输入: head = [1,2,3,4,5]
输出: [1,3,5,2,4]
示例 2:

输入: head = [2,1,3,5,6,4,7]
输出: [2,3,6,7,1,5,4]

提示:

n ==  链表中的节点数
0 <= n <= 104
-106 <= Node.val <= 106
*/
package leetcode

import (
	"testing"

	. "al/leetcode"
)

// 将奇数节点和偶数节点分离成奇数链表和偶数链表，然后将偶数链表连接在奇数链表之后，合并后的链表即为结果链表
func TestOddEvenList(t *testing.T) {
	//nums := []int{2, 1, 3, 5, 6, 4, 7}
	nums := []int{1, 2, 3, 4, 5}
	head := GenerateListNode(RevertInts(nums))

	nHead := oddEvenList2(head)
	PrintListNode(nHead)
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // len(head)=0/1
		return head
	}

	evenHead := head.Next                 // 偶数head, 不动
	odd := head                           // 奇数head, 不动
	even := evenHead                      // 偶数
	for even != nil && even.Next != nil { //??? `even != nil && even.Next != nil`好处: 总节点奇数时, even=nil; 偶数时,  even=最后一个偶数节点, odd.Next = evenHead后不会产生循环
		odd.Next = even.Next // 将奇数节点的Next指向下一个奇数节点（even.Next）
		odd = odd.Next       // 移动odd指针到新的奇数节点
		even.Next = odd.Next // 将偶数节点的Next指向下一个偶数节点（odd.Next）
		even = even.Next     // 移动even指针到新的偶数节点
	}

	odd.Next = evenHead // 最后一个奇数节点要指向偶数链表
	return head
}

// 好理解
func oddEvenList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// oH,eH: 分别是奇数节点和偶数节点的表头, oH可用充当
	odd := head
	even, eH := head.Next, head.Next

	for {
		if even.Next == nil { // len(nodes)=偶数
			break
		}

		odd.Next = even.Next
		odd = odd.Next

		if odd.Next == nil { // len(nodes)=奇数
			break
		}

		even.Next = odd.Next
		even = even.Next
	}

	even.Next = nil // 断开循环, 因为len(nodes)=奇数时, even最后指向最后一个奇数节点
	odd.Next = eH
	return head
}
