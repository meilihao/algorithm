/*
11. 旋转数组的最小数字
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个非递减序列的一个旋转，输出旋转数组的最小元素。

例如

数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，

该数组的最小值为1。


*/

package demo

import (
	"fmt"
	"testing"
)

func TestMinArray(t *testing.T) {
	nums1 := []int{0, 1, 2, 3, 4, 5}
	nums2 := []int{3, 4, 5, 0, 1, 2}

	fmt.Println(minArray(nums1) == 0)
	fmt.Println(minArray(nums2) == 0)
}

func minArray(numbers []int) int {
	l := 0
	r := len(numbers) - 1
	for l < r {
		mid := l + (r-l)/2
		fmt.Println(l, r, mid)
		if numbers[mid] < numbers[r] {
			r = mid
		} else if numbers[mid] > numbers[r] {
			l = mid + 1
		} else {
			r--
		}
	}
	return numbers[l]
}
