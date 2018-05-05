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

// SortAges
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

	tmp := array[left]
	i, j := left, right

	for i != j {
		for array[j] >= tmp && i < j {
			j--
		}

		for array[i] <= tmp && i < j {
			i++
		}

		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}

	fmt.Println(array)

	array[left], array[i] = array[i], array[left]

	quickSort(array, left, i-1)
	quickSort(array, i+1, right)
}

// 快速排序方式
func QuickSort(array []int) {
	quickSort(array, 0, len(array)-1)
}

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
