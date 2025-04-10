// 239. 滑动窗口最大值
// 思路: 1. MaxHeap(N*log2(k)); 2. Deque(N)
package main

import "fmt"

func main() {
	k := 3                         // 1               // 3
	nums := []int{1, 3, -1, 0, -2} // []int{1, -1} // []int{1, 3, -1, -3, 5, 3, 6, 7}

	fmt.Println(maxSlidingWindow(nums, k))
}

// why : 不适合使用NewDeque, 因为需要将小于等于当前值的队尾元素依次出队
// func maxSlidingWindow2(nums []int, k int) []int {
// 	if !(len(nums) > 0 && k >= 1 && k <= len(nums)) {
// 		return nil
// 	}

// 	res := []int{}
// 	q := NewDeque() // 保存索引下标, 因此保持数据时, 无法弹出超出窗口的过期数据

// 	var tmp int
// 	var isExist bool

// 	for i, x := range nums {
// 		tmp, isExist = q.PeekFront() // 最大值在queue的最前面, 同时它的索引值也是最小的, 因为我们保存的索引值是随i递增的
// 		if isExist && tmp <= i-k {   // 弹出过期索引
// 			q.PopFront()
// 		}

// 		for !q.IsEmpty() && true {
// 			q.PopBack()
// 		}

// 		q.PushBack(x)

// 		if i >= k-1 { // 输出q中最左边的值
// 			tmp, _ = q.PeekFront()
// 			res = append(res, nums[tmp])
// 		}
// 	}

// 	return res
// }

func maxSlidingWindow(nums []int, k int) []int {
	if !(len(nums) > 0 && k >= 1 && k <= len(nums)) {
		return nil
	}

	// 再优化, k == 1 return nums

	window := make([]int, 0, k) // 下标对应的值是有序的
	result := []int{}

	for i, x := range nums {
		if len(window) > 0 {
			if i-k >= window[0] { // 仅 i>=k时需要清理过期索引, 同时也仅用清理一个即可, 因为i的步长是1
				window = window[1:]
			}

			for len(window) != 0 && nums[window[len(window)-1]] <= x { // 将小于等于当前值的队尾元素依次出队(包含等于是要相应值的索引能更久地保留在window中, 即不过期), 可保证window下标对应的值是有序的
				window = window[:len(window)-1]
			}
		}

		window = append(window, i)

		if i >= k-1 {
			result = append(result, nums[window[0]])
		}
	}
	return result
}

type Deque struct {
	nums []int
}

func NewDeque() *Deque {
	return &Deque{}
}

func (q *Deque) PushBack(x int) {
	q.nums = append(q.nums, x)
}

func (q *Deque) PushFront(x int) {
	tmp := []int{x}

	if len(q.nums) > 0 {
		tmp = append(tmp, q.nums...)
	}

	q.nums = tmp
}

func (q *Deque) PopBack() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	res := q.nums[len(q.nums)-1]
	q.nums = q.nums[:len(q.nums)-1]

	return res, true
}

func (q *Deque) PopFront() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	res := q.nums[0]
	q.nums = q.nums[1:]

	return res, true
}

func (q *Deque) PeekBack() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	return q.nums[len(q.nums)-1], true
}

func (q *Deque) PeekFront() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	return q.nums[0], true
}

func (q *Deque) IsEmpty() bool {
	return len(q.nums) == 0
}

func (q *Deque) Len() int {
	return len(q.nums)
}
