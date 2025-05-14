/*
703.简 数据流中的第 K 大元素

设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。

请实现 KthLargest 类：

KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。

示例 1：

输入：
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]

输出：[null, 4, 5, 5, 8, 8]

解释：

KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3); // 返回 4
kthLargest.add(5); // 返回 5
kthLargest.add(10); // 返回 5
kthLargest.add(9); // 返回 8
kthLargest.add(4); // 返回 8

示例 2：

输入：
["KthLargest", "add", "add", "add", "add"]
[[4, [7, 7, 7, 7, 8, 3]], [2], [10], [9], [9]]

输出：[null, 7, 7, 7, 8]

解释：

KthLargest kthLargest = new KthLargest(4, [7, 7, 7, 7, 8, 3]);
kthLargest.add(2); // 返回 7
kthLargest.add(10); // 返回 7
kthLargest.add(9); // 返回 7
kthLargest.add(9); // 返回 8

提示：
0 <= nums.length <= 104
1 <= k <= nums.length + 1
-104 <= nums[i] <= 104
-104 <= val <= 104
最多调用 add 方法 104 次
*/
package main

import "fmt"

// 思路: minHeap
// 复杂度: N*log2(k)
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
