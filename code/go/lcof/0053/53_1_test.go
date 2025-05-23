/*
53（一）：数字在排序数组中出现的次数

统计一个数字在排序数组中出现的次数。例如输入排序数组{1, 2, 3, 3, 3, 3, 4, 5}和数字3，由于3在这个数组中出现了4次，因此输出4
*/
package demo

import (
	"fmt"
	"testing"
)

func TestGetNumberOfK(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 3, 4, 5}
	fmt.Println(getNumberOfK(nums, 3) == 4)
}

func getNumberOfK(num []int, k int) int {
	length := len(num)
	firstK := getFirstK(num, k, 0, length-1)
	lastK := getLastK(num, k, 0, length-1)
	if firstK != -1 && lastK != -1 {
		return lastK - firstK + 1
	}
	return 0
}

// 递归
func getFirstK(num []int, k int, start int, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if num[mid] > k {
		return getFirstK(num, k, start, mid-1)
	} else if num[mid] < k {
		return getFirstK(num, k, mid+1, end)
	} else if mid-1 >= 0 && num[mid-1] == k { //mid为k，mid-1也为k即前一个也是k
		return getFirstK(num, k, start, mid-1) //找到最前面的K
	} else {
		return mid
	}
}

// 循环
func getLastK(num []int, k int, start int, end int) int {
	length := len(num)
	mid := (start + end) / 2
	for start <= end {
		if num[mid] > k {
			end = mid - 1
		} else if num[mid] < k {
			start = mid + 1
		} else if mid+1 <= length-1 && num[mid+1] == k { //mid为k，mid+1也为k
			start = mid + 1 //找到最后面的K
		} else {
			return mid
		}
		mid = (start + end) / 2
	}
	return -1
}
