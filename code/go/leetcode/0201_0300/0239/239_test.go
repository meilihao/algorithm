/*
239.困 滑动窗口最大值

给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。

示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3

	1 [3  -1  -3] 5  3  6  7       3
	1  3 [-1  -3  5] 3  6  7       5
	1  3  -1 [-3  5  3] 6  7       5
	1  3  -1  -3 [5  3  6] 7       6
	1  3  -1  -3  5 [3  6  7]      7

示例 2：

输入：nums = [1], k = 1
输出：[1]

提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104
1 <= k <= nums.length
*/
package main

import (
	"fmt"
	"testing"
)

// 思路: 1. MaxHeap(N*log2(k)); 2. Deque(N)
func TestMaxSlidingWindow(t *testing.T) {
	k := 3 // 1               // 3
	//nums := []int{1, 3, -1, 0, -2}
	//nums :=  []int{1, -1}
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}

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

// 解决滑动窗口最大值问题最有效的方法是使用单调队列.
// 单调队列是一个双端队列 (deque)，它维护着窗口内可能成为最大值的元素的索引，并且这些索引对应的元素值是单调递减的
func maxSlidingWindow(nums []int, k int) []int {
	if !(len(nums) > 0 && k >= 1 && k <= len(nums)) {
		return nil
	}

	// 再优化, k == 1 return nums

	// window是单调队列: 左, 大 -> 右, 小
	window := make([]int, 0, k) // 下标对应的值是有序的
	result := []int{}

	for i, x := range nums {
		if len(window) > 0 {
			//fmt.Println("b", x, window)
			if i-k >= window[0] { // 仅 i>=k时需要清理过期索引, 同时也仅用清理一个即可, 因为i的步长是1
				window = window[1:] // 从队首移除过期元素
			}

			// nums[window[len(window)-1]] <= x： 如果队尾的元素值小于或等于新来的 x，那么队尾的这个元素永远不可能成为最大值, 弹出
			for len(window) != 0 && nums[window[len(window)-1]] <= x { // 将小于等于当前值的队尾元素依次出队(包含等于是要相应值的索引能更久地保留在window中, 即不过期), 可保证window下标对应的值是有序的
				window = window[:len(window)-1]
			}

			//fmt.Println("a", x, window)
		}

		window = append(window, i)

		// 当 i 达到 k-1 时，表示第一个完整窗口已经形成（索引从 0 到 k-1）, 之后每移动一步，都会形成一个新的完整窗口
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
