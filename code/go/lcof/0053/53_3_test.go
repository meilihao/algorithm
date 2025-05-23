/*
53（三）：数组中数值和下标相等的元素

假设一个单调递增的数组里的每个元素都是整数并且是唯一的。请编程实
现一个函数找出数组中任意一个数值等于其下标的元素。例如，在数组{-3, -1,1, 3, 5}中，数字3和它的下标相等。
*/
package demo

import (
	"fmt"
	"testing"
)

func TestGetNumberSameAsIndex(t *testing.T) {
	nums := []int{-3, -1, 1, 3, 5}
	fmt.Println(GetNumberSameAsIndex(nums) == 3)
}

func GetNumberSameAsIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	return getNumberSameAsIndex(nums, 0, len(nums)-1)
}

// 函数 f(i) = nums[i] - i, 要找 f(i) = 0 的情况
// 当 nums[mid] > mid 时，意味着 f(mid) = nums[mid] - mid > 0. 因为nums单调递增, nums[mid]增大趋势(step>=1)>=mid增大的趋势(step=1)
func getNumberSameAsIndex(nums []int, start int, end int) int {
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == mid {
			return mid
		}

		if nums[mid] > mid { // nums[mid] 的值大于它的索引 mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}
