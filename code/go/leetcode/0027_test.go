package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3

	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {
	left, right, n := 0, 0, len(nums)

	for right < n {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}

		right++
	}
	return left
}
