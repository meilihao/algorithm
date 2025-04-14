package leetcode

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	fmt.Println(search(nil, 1))
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left // 直接写 (left + right)/2 在 left 和 right 很大时可能溢出
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
