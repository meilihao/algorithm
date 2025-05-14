/*
34.中 在排序数组中查找元素的第一个和最后一个位置

给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]
*/
package demo

import (
	"fmt"
	"testing"
)

func TestSearchRange(t testing.T) {
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
func searchIP(a []int, n, value int) int {
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
