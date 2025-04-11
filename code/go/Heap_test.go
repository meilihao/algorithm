// https://en.wikipedia.org/wiki/Heap_%28data_structure%29
package ago

import (
	"fmt"
	"testing"
)

type Heap struct {
	a     []int
	max   int // 堆可以存储的最大个数
	count int // 堆中已存储的个数
}

func NewHeap(max int) *Heap {
	return &Heap{
		a:     make([]int, max+1), // a[0]不用
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
	h.a[h.count] = data // 将元素放至数组末尾

	// 自底向上堆化
	i := h.count
	for i/2 > 0 && h.a[i] > h.a[i/2] { // 子节点大于父节点
		h.a[i], h.a[i/2] = h.a[i/2], h.a[i]

		i /= 2
	}
}

func (h *Heap) RemoveMax() {
	if h.count == 0 {
		fmt.Println("heap is empty")
		return
	}

	h.a[1] = h.a[h.count] // 将最后一个元素移动到堆顶
	//h.a[h.count] = 0
	h.count--

	heapify(h.a, h.count, 1)
}

func (h *Heap) Sort() {
	k := h.count

	for k > 1 {
		h.a[1], h.a[k] = h.a[k], h.a[1] // 将堆顶和末尾元素调换位置，从而将取出堆顶元素和堆化的第一步(将末尾元素放至根结点位置)进行合并
		k--

		heapify(h.a, k, 1)
	}
}

func TestHeap(t *testing.T) {
	HeapOps()
}

func TestHeapSort(t *testing.T) {
	nums := []int{0, 7, 5, 19, 8, 4, 1, 20, 13, 16}
	h := NewHeap(10)
	for _, v := range nums[1:] {
		h.Insert(v)
	}

	fmt.Println(h.a[1 : h.count+1])
	//h.RemoveMax()
	h.Sort()
	fmt.Println(h.a[1 : h.count+1])
}

func HeapOps() {
	nums := []int{0, 7, 5, 19, 8, 4, 1, 20, 13, 16}
	//nums := []int{0, 7, 8}
	fmt.Println(nums)

	buildHead(nums, len(nums)-1)
	fmt.Println(nums)
}

// n为堆中元素的个数
func buildHead(a []int, n int) {
	if n <= 1 { // 因为a[0]是空位, 1个元素也无需堆化
		return
	}

	for i := n / 2; i >= 1; i-- { // 顺序是从后往前堆化
		heapify(a, n, i)
	}
}

// 从上往下堆化
// 不停与左右子节点的值进行比较，和较大的子节点交换位置，直到无法交换位置
func heapify(a []int, n, i int) {
	var maxPos int

	for {
		maxPos = i

		if i*2 <= n && a[i] < a[i*2] { // 父节点<左节点
			maxPos = i * 2 // 选出max(父, 左)
		}
		if i*2+1 <= n && a[maxPos] < a[i*2+1] { // 三个数比较大小, 这里是a[maxPos]而不是a[i]
			maxPos = i*2 + 1
		}
		if maxPos == i {
			break
		}

		a[i], a[maxPos] = a[maxPos], a[i]

		// fmt.Println("--", i, maxPos, a)

		i = maxPos
	}
}
