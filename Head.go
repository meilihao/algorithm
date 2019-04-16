package main

import "fmt"

type Heap struct {
	a     []int
	max   int // 堆可以存储的最大个数
	count int // 堆中已存储的个数
}

func NewHeap(max int) *Heap {
	return &Heap{
		a:     make([]int, max+1),
		max:   max,
		count: 0,
	}
}

// 插入法建最大堆
func (h *Heap) Insert(data int) {
	if h.count >= h.max {
		fmt.Println("heap is max")
		return
	}

	h.count++
	h.a[h.count] = data

	i := h.count
	for i/2 > 0 && h.a[i] > h.a[i/2] {
		h.a[i], h.a[i/2] = h.a[i/2], h.a[i]

		i /= 2
	}
}

func (h *Heap) RemoveMax() {
	if h.count == 0 {
		fmt.Println("heap is empty")
		return
	}

	h.a[1] = h.a[h.count]
	//h.a[h.count] = 0
	h.count--

	heapify(h.a, h.count, 1)
}

func main() {
	nums := []int{0, 7, 5, 19, 8, 4, 1, 20, 13, 16}
	fmt.Println(nums)

	// h := NewHeap(10)
	// for _, v := range nums[1:] {
	// 	h.Insert(v)
	// }
	// fmt.Println(h.a[1 : h.count+1])
	// h.RemoveMax()
	// fmt.Println(h.a[1 : h.count+1])

	buildHead(nums, len(nums)-1)
	fmt.Println(nums)
}

// n为堆中元素的个数
func buildHead(a []int, n int) {
	if n <= 2 {
		return
	}

	for i := n / 2; i >= 1; i-- {
		heapify(a, n, i)
	}
}

// 从上往下堆化
func heapify(a []int, n, i int) {
	var maxPos int

	for {
		maxPos = i

		if i*2 <= n && a[i] < a[i*2] {
			maxPos = i * 2
		}
		if i*2+1 <= n && a[i] < a[i*2+1] {
			maxPos = i*2 + 1
		}
		if maxPos == i {
			break
		}

		a[i], a[maxPos] = a[maxPos], a[i]

		fmt.Println("--", i, maxPos, a)

		i = maxPos
	}
}
