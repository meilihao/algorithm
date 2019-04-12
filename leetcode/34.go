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

// 类似问题:
// 如何快速定位出一个 IP地址的归属地 -> 在有序数组中， 查找最后一个小于等于某个给定值的元素
func searchIP(int []a, int n, int value) int {
	low := 0
	high := n - 1
	var mid int

	for low <= high {
		mid = (low + high) >> 1

		if a[mid] > value {
			high = mid - 1
		} else {
			if mid == n-1 || a[mid+1] > value {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}
