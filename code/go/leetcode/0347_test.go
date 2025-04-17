/*
前 K 个高频元素 = 使用排序算法对元素按照频率由高到低进行排序，然后再取前 k 个元素

方法:
- 堆:
  - 最大堆: 前k个
  - 最小堆: k个, 取出时从最小到最大
*/
package leetcode

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1, 2, 2, 3}

	fmt.Println(topKFrequent(nums, 1))
}

/*
维护一个元素数目为 k 的最小堆
每次都将新的元素与堆顶元素（堆中频率最小的元素）进行比较
如果新的元素的频率比堆顶端的元素大，则弹出堆顶端的元素，将新的元素添加进堆中
*/
func topKFrequent(nums []int, k int) []int {
	countM := make(map[int]int)
	for i := range nums {
		countM[nums[i]]++
	}

	fmt.Println(countM)

	h := &IHeap{}
	heap.Init(h)

	for key, c := range countM {
		heap.Push(h, [2]int{key, c})

		if h.Len() > k { //只要前k个
			heap.Pop(h)
		}
	}

	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}

	return ret
}

type IHeap [][2]int // [key,count]

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
