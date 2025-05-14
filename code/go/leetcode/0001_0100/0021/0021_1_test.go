// 归并排序. 给定两个有序数组（整数、递增），请将它们归并成一个有序数组。
// 例：
// 输入：[1,3,5,7], [2,4,6,8]
// 输出：[1,2,3,5,6,7,8]
package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeArray(t *testing.T) {
	a := []int{1, 3, 5, 7}
	b := []int{2, 4, 6, 8}

	fmt.Println(mergeArray(a, b))
}

func mergeArray(a, b []int) []int {
	maxA := len(a)
	maxB := len(b)

	if maxA == 0 {
		return b
	}
	if maxB == 0 {
		return a
	}

	newArray := make([]int, 0, maxA+maxB)
	i, j, max := 0, 0, min(maxA, maxB)

	for i < max && j < max {
		if a[i] > b[j] {
			newArray = append(newArray, b[j])
			j++
		} else {
			newArray = append(newArray, a[i])
			i++
		}
	}

	if i < maxA {
		newArray = append(newArray, a[i:]...)
	}
	if j < maxB {
		newArray = append(newArray, b[j:]...)
	}

	return newArray
}
