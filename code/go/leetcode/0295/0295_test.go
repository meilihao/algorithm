package demo

import (
	"container/heap"
	"sort"
	"testing"
)

// todo: others
func TestMedianFinder(t *testing.T) {

}

// 当数据流元素数量为偶数：l 和 r 大小相同，此时动态中位数为两者堆顶元素的平均值；
// 当数据流元素数量为奇数：l 比 r 多一，此时动态中位数为 l
// Go 的 heap 包默认实现是小顶堆，因此 queMin 通过存储负数模拟大顶堆
type MedianFinder struct {
	queMin hp // l, 小于等于中位数的数, 维护数据流左半边数据的优先队列. 本来需要大顶堆, 通过取反也转成小顶堆, 即可复用小顶堆代码
	queMax hp // r, 大于中位数的数, 维护数据流右半边数据的优先队列, 小顶堆
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	minQ, maxQ := &mf.queMin, &mf.queMax

	// 如果 l 为空，说明当前插入的是首个元素，直接添加到 l
	// num 小于等于中位数, 也插入l. 新的中位数将小于等于原来的中位数，因此可能需要将 queMin 中最大的数移动到 queMax 中
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {
		heap.Push(minQ, -num)
		if maxQ.Len()+1 < minQ.Len() { // 确保 minQ 的大小不超过 maxQ + 1
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		heap.Push(maxQ, num)
		if maxQ.Len() > minQ.Len() { // 新的中位数将大于等于原来的中位数，因此可能需要将 queMax 中最小的数移动到 queMin 中. 确保 maxQ 的大小不超过 minQ
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	minQ, maxQ := mf.queMin, mf.queMax
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	}
	return float64(maxQ.IntSlice[0]-minQ.IntSlice[0]) / 2
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
