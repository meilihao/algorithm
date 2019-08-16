package main

import "fmt"

func main() {
	// a := []int{1, 5, 2, 6, 9, 0, 3, 5, 7, 8}

	// fmt.Println(BubbleSort(a))
	// fmt.Println(SortAges(a))
	sa := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(sa)
	fmt.Println("result:", findKthLargest(sa, 2))
	fmt.Println(sa)
}

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	if !(n > 0 && k >= 1 && k <= n) {
		panic("invalid input")
	}

	return quickSort(nums, 0, n-1, n-k)
	//return quickSort2(nums, 0, n-1, k)
}

// 从小到大排
// 即求len(n)-k上的元素
func quickSort(arr []int, start, end, k int) int {
	if start >= end {
		return arr[end]
	}

	// 选取最后一位当对比数字
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

	fmt.Println(i, k)
	if i == k {
		return arr[i]
	} else if i < k { // k选中的值应该在i的右侧
		return quickSort(arr, i+1, end, k)
	} else {
		return quickSort(arr, start, i-1, k)
	}
}

// best
// 根据题意: 选择从大到小排序
func quickSort2(arr []int, start, end, k int) int {
	if start >= end {
		return arr[end]
	}

	// 选取最后一位当对比数字, 即遍历方向的最后一位
	pivot := arr[end]

	var i = start // 查找a[i]<=pivot
	for j := start; j < end; j++ {
		if arr[j] > pivot { // arr[j] > pivot => 合适时交换i, j
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++ // 不符合a[i]<=pivot的条件, 因此尝试下一位
		}
	}

	arr[i], arr[end] = arr[end], arr[i]

	//fmt.Println(i, k-1)
	if i == k-1 {
		return arr[i]
	} else if i > k-1 { // k选中的值应该在i的左侧
		return quickSort2(arr, start, i-1, k)
	} else {
		return quickSort2(arr, i+1, end, k)
	}
}

// todo : 利用最小堆
