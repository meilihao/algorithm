package leetcode

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSortArray(t *testing.T) {
	nums := []int{5, 2, 3, 1}

	fmt.Println(sortArray(nums))
	fmt.Println(sortArray2(nums))
	fmt.Println(sortArray3(nums))
}

func sortArray(nums []int) []int {
	quickSortAscBase(nums, 0, len(nums)-1)
	return nums
}

// 避免退化成n^2
// 三数取中: 有序/部分有序
// 随机: 随机数据/重复元素多
func quickSortMidOfThree(arr []int, start, end int) {
	mid := start + rand.Intn(end-start+1) // [0,end - start+1), max(mid)=start + end - start = end

	//fmt.Println("b", start, mid, end, arr[start], arr[mid], arr[end])

	/*
		if (arr[mid] >= arr[start] && arr[mid] <= arr[end]) || (arr[mid] >= arr[end] && arr[mid] <= arr[start]) { // mid在start, end中间

		} else if (arr[start] >= arr[mid] && arr[start] <= arr[end]) || (arr[start] >= arr[end] && arr[start] <= arr[mid]) { // start在mid,end中间
			mid = start
		} else {
			mid = end
		}
	*/

	// 作用同上, 可减少比较次数
	if (arr[mid]-arr[start])*(arr[end]-arr[mid]) >= 0 { // mid在start, end中间

	} else if (arr[start]-arr[mid])*(arr[end]-arr[start]) >= 0 { // start在mid,end中间
		mid = start
	} else {
		mid = end
	}

	//fmt.Println("a", start, mid, end, arr[start], arr[mid], arr[end])

	arr[end], arr[mid] = arr[mid], arr[end]
}

// O(nlogn)
// 快排 pivot 不随机选竟然超时: 在处理n个相同元素时，时间复杂度会退化为O（n^2）
func quickSortAscBase(arr []int, start, end int) {
	if start >= end {
		return
	}

	quickSortMidOfThree(arr, start, end)

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

// 堆排序
func sortArray3(nums []int) []int {
	var heapilfy func(root int, end int)

	heapilfy = func(root int, end int) {
		for {
			child := root*2 + 1 // 左子节点
			if child > end {    // 如果没有子节点，终止
				return
			}
			if child < end && nums[child+1] >= nums[child] { // 如果右子节点存在且比左子节点大，选择右子节点
				child++
			}
			if nums[child] < nums[root] { // 如果父节点已经 ≥ 子节点，堆结构已满足，终止
				return
			}
			nums[child], nums[root] = nums[root], nums[child]
			root = child // 继续向下调整
		}
	}
	length := len(nums) - 1
	for i := length / 2; i >= 0; i-- {
		heapilfy(i, length)
	}
	for i := length; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		length--
		heapilfy(0, length)
	}
	return nums
}
