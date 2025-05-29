package demo

import (
	"container/heap"
	"sort"
)

type KthLargest2 struct {
	sort.IntSlice
	k int
}

func Constructor2(k int, nums []int) KthLargest2 {
	kl := KthLargest2{k: k}
	for _, val := range nums {
		kl.Add(val)
	}
	return kl
}

func (kl *KthLargest2) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

func (kl *KthLargest2) Pop() interface{} {
	a := kl.IntSlice
	v := a[len(a)-1]
	kl.IntSlice = a[:len(a)-1]
	return v
}

func (kl *KthLargest2) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.IntSlice[0]
}
