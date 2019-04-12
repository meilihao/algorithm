package main

import "fmt"

func main() {
	//nums := []int{5, 7, 7, 8, 8, 10}
	nums := []int{1}
	fmt.Println(searchRange(nums, 8))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	return []int{
		firstTarget(nums, 0, len(nums)-1, target),
		lastTarget(nums, 0, len(nums)-1, target),
	}
}

// 查找第一个值等于给定值的元素
func firstTarget(nums []int, low, up, target int) int {
	var mid int

	for low <= up {
		mid = (low + up) >> 1
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				up = mid - 1
			}
		} else if nums[mid] > target {
			up = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

// 查找最后一个值等于给定值的元素
func lastTarget(nums []int, low, up, target int) int {
	var mid int

	for low <= up {
		mid = (low + up) >> 1
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			} else {
				low = mid + 1
			}
		} else if nums[mid] > target {
			up = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
