package leetcode

import (
	"testing"
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
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
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
