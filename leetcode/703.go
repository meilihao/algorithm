// 703. 数据流中的第K大元素
// 思路: minHeap
// 复杂度: N*log2(k)
package main

import "fmt"

func main() {
	k := 3
	arr := []int{4, 5, 8, 2}

	kth := Constructor(k, arr)
	fmt.Println(kth.Add(3))
	fmt.Println(kth.Add(5))
	fmt.Println(kth.Add(10))
	fmt.Println(kth.Add(9))
	fmt.Println(kth.Add(4))
}

type KthLargest struct {
	k    int
	heap *HeapMin
}

func Constructor(k int, nums []int) KthLargest {
	h := NewHeapMin(k)

	for _, v := range nums {
		h.Push(v)
	}

	return KthLargest{
		k:    k,
		heap: h,
	}
}

func (this *KthLargest) Add(val int) int {
	this.heap.Push(val)

	return this.heap.Peek()
}

type HeapMin struct {
	nums  []int
	max   int
	count int
}

func NewHeapMin(max int) *HeapMin {
	return &HeapMin{
		max:  max,
		nums: make([]int, max+1),
	}
}

func (h *HeapMin) Push(x int) {
	if h.count >= h.max {
		if x > h.Peek() {
			h.nums[1] = x

			h.Heapify(1, h.count)
		}

		return
	}

	h.count++
	h.nums[h.count] = x

	i := h.count
	for i/2 > 0 && h.nums[i] < h.nums[i/2] { // 子节点小于父节点
		h.nums[i], h.nums[i/2] = h.nums[i/2], h.nums[i]

		i /= 2
	}
}

func (h *HeapMin) Pop() int {
	if h.count == 0 {
		return -1
	}

	res := h.nums[1]

	h.nums[1] = h.nums[h.count]
	h.count--

	h.Heapify(1, h.count)

	return res
}

func (h *HeapMin) Peek() int {
	if h.count == 0 {
		return -1
	}

	return h.nums[1]
}

// 从上往下堆化
func (h *HeapMin) Heapify(i, n int) { // i = start, n = end
	var minPos int

	for {
		minPos = i

		if i*2 <= n && h.nums[i*2] < h.nums[i] {
			minPos = i * 2
		}
		if i*2+1 <= n && h.nums[i*2+1] < h.nums[minPos] { // 三个数比较大小, 这里是h.nums[minPos]而不是a[i]
			minPos = i*2 + 1
		}
		if minPos == i {
			break
		}

		h.nums[i], h.nums[minPos] = h.nums[minPos], h.nums[i]

		i = minPos
	}
}
