/*
215.中 数组中的第K个最大元素

给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1:

输入: [3,2,1,5,6,4], k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4

提示：

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
*/
package leetcode

import (
	"fmt"
	"math/rand"
	"testing"
)

// 其实就是top k的问题, 相对好的解题思路:
// 1. 最小堆
// 2. 快排思路(最优)
func TestFindKthLargest(t *testing.T) {
	// a := []int{1, 5, 2, 6, 9, 0, 3, 5, 7, 8}

	// fmt.Println(BubbleSort(a))
	// fmt.Println(SortAges(a))
	sa := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(sa)
	fmt.Println("result:", findKthLargest(sa, 1))
	fmt.Println(sa)
}

func TestFindKthLargest2(t *testing.T) {
	// a := []int{1, 5, 2, 6, 9, 0, 3, 5, 7, 8}

	// fmt.Println(BubbleSort(a))
	// fmt.Println(SortAges(a))
	sa := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(sa)
	fmt.Println("result:", findKthLargest2(sa, 1))
	fmt.Println(sa)
}

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	if !(n > 0 && k >= 1 && k <= n) {
		panic("invalid input")
	}

	//return quickSort(nums, 0, n-1, n-k)
	return quickSort2(nums, 0, n-1, k-1)
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
// k=目标位置的索引
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
	if i == k {
		return arr[i]
	} else if i > k { // k选中的值应该在i的左侧
		return quickSort2(arr, start, i-1, k)
	} else {
		return quickSort2(arr, i+1, end, k)
	}
}

// todo : 利用最小堆

func findKthLargest2(nums []int, k int) int {
	target := len(nums) - k
	start := 0
	end := len(nums) - 1

	index := partition(nums, start, end)
	for index != target {
		if index > target {
			end = index - 1
		} else {
			start = index + 1
		}

		index = partition(nums, start, end)
	}

	return nums[index]
}

// 小->大
func partition(nums []int, start, end int) int {
	mid := start + rand.Intn(end-start+1)
	nums[mid], nums[end] = nums[end], nums[mid]

	small := start - 1 // small = -1, 指向用来放小数的位置
	for i := start; i < end; i++ {
		if nums[i] < nums[end] {
			small++
			nums[i], nums[small] = nums[small], nums[i]
		}
	}

	small++
	nums[small], nums[end] = nums[end], nums[small]

	return small
}
