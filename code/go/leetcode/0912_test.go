package leetcode

import (
	"fmt"
	"testing"
)

func TestSortArray(t *testing.T) {
	nums := []int{5, 2, 3, 1}

	fmt.Println(sortArray(nums))
	fmt.Println(sortArray2(nums))
}

func sortArray(nums []int) []int {
	quickSortAscBase(nums, 0, len(nums)-1)
	return nums
}

// O(nlogn)
// 快排 pivot 不随机选竟然超时: 在处理n个相同元素时，时间复杂度会退化为O（n^2）
func quickSortAscBase(arr []int, start, end int) {
	if start >= end {
		return
	}

	// 选取最后一位当对比数字, 即遍历方向的最后一位
	pivot := arr[end]

	var i = start
	for j := start; j < end; j++ {
		if arr[j] < pivot { // 合适时交换i, j
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++ // 不符合条件, 因此尝试下一位
		}
	}

	arr[i], arr[end] = arr[end], arr[i]

	quickSortAscBase(arr, start, i-1)
	quickSortAscBase(arr, i+1, end)
}

func sortArray2(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

// 归并排序
func mergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	mid := (l + r) >> 1
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)

	inPlaceMerge(arr, l, mid, r)
}

// 原地归并（关键函数）
func inPlaceMerge(arr []int, left, mid, right int) {
	start1 := left
	start2 := mid + 1

	// mid是arr左半部最后的一个数
	// 如果两个子数组已经有序，无需合并
	if arr[mid] <= arr[start2] {
		return
	}

	for start1 <= mid && start2 <= right {
		if arr[start1] <= arr[start2] {
			start1++
		} else {
			// arr[start2] 比 arr[start1] 小，需要插入到前面
			// 步骤:
			// 1. value暂时arr[start2]
			// 2. 将arr[start1:start2]后移一位
			// 3. 将value放入空出的位置
			// 4. 更新索引
			value := arr[start2]
			index := start2

			// 所有元素向后移动一位
			for index != start1 {
				arr[index] = arr[index-1]
				index--
			}
			arr[start1] = value

			// 更新所有指针位置
			start1++
			mid++ // 因为后移了一位
			start2++
		}
	}
}
