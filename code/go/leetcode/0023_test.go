package leetcode

import "testing"

func TestMergeKLists(t *testing.T) {
	MergeKLists(nil)
}

func MergeKLists(lists []*ListNode) *ListNode {
	//return mergeKLists(lists)
	return mergeKLists2(lists)
}

// k个链表的合并，可以看做是k-1次，每两个链表之间的合并
func mergeKLists(lists []*ListNode) *ListNode {
	var pre, cur *ListNode
	n := len(lists)
	for i := 0; i < n; i++ {
		if i == 0 {
			pre = lists[i]
			continue
		}
		cur = lists[i]
		pre = mergeTwoLists(pre, cur)
	}
	return pre
}

// 整体是分治，归并排序思想

// 时间 nlogn
// 空间 logn
// best
func mergeKLists2(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	l, r := mergeKLists2(lists[:n/2]), mergeKLists2(lists[n/2:])
	return mergeTwoLists(l, r)
}
