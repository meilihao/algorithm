package main

import (
	"fmt"
)

func main() {
	// a := []int{1, 5, 2, 6, 9, 0, 3, 5, 7, 8}

	// fmt.Println(BubbleSort(a))
	// fmt.Println(SortAges(a))
	sa := []int{5, 4, 3, 6, 1, 3, 2, 6, 5}
	QuickSort(sa)
	fmt.Println(sa)
	fmt.Println(BinarySearch(sa, 1))
}

// 冒泡排序
// 界定边界: 看最后一次比较a[length-2],a[length-1]
func BubbleSort(a []int) []int {
	length := len(a)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if a[i] > a[j] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}

	return a
}

// 我们将数组中的数据分为两个区间， 已排序区间和未排序区间. 初始已排序区间只有一个元素， 就是数组的第一个元素.
// 插入算法的核心思想是取未排序区间中的元素， 在已排序区间中找到合适的插入位置将其插入， 并保证已排序区间数据一直有序.
// 重复这个过程， 直到未排序区间中元素为空， 算法结束.
func InsertSort(a []int) {
	if len(a) <= 1 {
		return
	}

	n := len(a)
	var tmp, j int // 保存选中要进行插入的数据
	for i := 1; i < n; i++ {
		tmp = a[i]
		j = i - 1 //有序序列:0~i-1

		// 查找要插入的位置
		for ; j >= 0; j-- {
			if a[j] > tmp { // 要后挪
				a[j+1] = a[j]
			} else {
				break // 因为i前面的元素已有序,且i前面的元素已是有序区域的最大值
			}
		}

		a[j+1] = tmp
	}
}

// 选择排序算法的实现思路有点类似插入排序， 也分已排序区间和未排序区间.
// 但是选择排序每次会从未排序区间中找到最小的元素， 将其放到已排序区间的末尾
// 选择排序每次都要找剩余未排序元素中的最小值， 并和前面的元素交换位置， 这样破坏了稳定性, 比如`5， 8， 5， 2， 9`
func SelectSort(a []int) {
	if len(a) <= 1 {
		return
	}

	n := len(a)

	for i := 0; i < n; i++ {
		// 查找最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		// 交换
		if i != minIndex {
			a[i], a[minIndex] = a[minIndex], a[i]
		}
	}
}

func MergeSort(a []int) {
	if len(a) <= 1 {
		return
	}

	mergeSort(a, 0, len(a)-1)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ { //将前后部分数据依次按大小放入tmpArr
		if arr[i] < arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}

	// 其中一个子数组中的所有数据都放入临时数组中， 再把另一个数组中的数据依次加入到临时数组的末尾
	for ; i <= mid; i++ {
		tmpArr[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = arr[j]
		k++
	}

	copy(arr[start:end+1], tmpArr)
}

// SortAges, 计数排序
// age [0,149]
func SortAges(a []int) []int {
	maxAge := 150
	tmp := make([]int, maxAge)

	isInRange := true
	for _, v := range a {
		if v < 0 || v > maxAge-1 {
			isInRange = false

			break
		}

		tmp[v]++
	}

	if !isInRange {
		return nil
	}

	var index int
	for i, v := range tmp {
		if v > 0 {
			for j := 0; j < v; j++ {
				a[index] = i

				index++
			}
		}
	}

	return a
}

// best
func quickSort(array []int, left, right int) {
	if left >= right {
		return
	}

	tmp := array[left]  //基点,通常是前一个/最后一个元素
	i, j := left, right //i为什么不能是left+1(即基点也要参与比较), test case(5,6,5):不参与时不能保证array[left+1]<=tmp,即互换后不能保证左边不大于基点

	for i != j { // 两头向中间靠拢
		for array[j] >= tmp && i < j { // 从基点对应的另一端开始
			j--
		}
		// 碰到array[j]<tmp

		for array[i] <= tmp && i < j {
			i++
		}
		// 碰到array[i]>tmp

		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}

	fmt.Println(array)

	array[left], array[i] = array[i], array[left] // i为中间位置, 与基点交换, 交换后原基点数据在中间

	quickSort(array, left, i-1)
	quickSort(array, i+1, right)
}

// 推荐
func quickSort2(arr []int, start, end int) {
	if start >= end {
		return
	}

	// 选取第一位当对比数字
	pivot := arr[end]

	// 有点类似选择排序. 我们通过游标 i 把 A[p…r-1] 分成两部分. A[p…i-1] 的元素都是小于pivot 的，
	// 我们暂且叫它“已处理区间”，A[i…r-1] 是“未处理区间”.
	// 我们每次都从未处理的区间 A[i…r-1] 中取一个元素 A[j]， 与 pivot 对比，
	// 如果小于 pivot， 则将其加入到已处理区间的尾部， 也就是 A[i]的位置
	var i = start                  //查找a[i]>=pivot
	for j := start; j < end; j++ { //j 查找arr[j] < pivot的数
		if arr[j] < pivot {
			if i != j { //说明此时arr[i]>=pivot
				// 交换位置, 变成 arr[i]<pivot<= arr[j] 之后i进1:保证i的左边都小于pivot
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]

	quickSort2(arr, start, i-1)
	quickSort2(arr, i+1, end)
}

// 快速排序方式
func QuickSort(array []int) {
	quickSort(array, 0, len(array)-1)
}

// 二分查找(BinarySearch) : 针对**有序数据集合**的查找算法
func BinarySearch(array []int, target int) int {
	left, right := 0, len(array)-1
	var mid int

	for left <= right {
		mid = (left + right) / 2

		if array[mid] == target {
			return mid
		} else if array[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
