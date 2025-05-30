/*
35.简 搜索插入位置

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。

示例 1:

输入: nums = [1,3,5,6], target = 5
输出: 2
示例 2:

输入: nums = [1,3,5,6], target = 2
输出: 1
示例 3:

输入: nums = [1,3,5,6], target = 7
输出: 4

提示:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 为 无重复元素 的 升序 排列数组
-104 <= target <= 104
*/
package leetcode

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 5) == 2)
	fmt.Println(searchInsert(nums, 2) == 1)
	fmt.Println(searchInsert(nums, 7) == 4)
}

// 用二分法逼近查找第一个大于等于 target 的下标
//
// 考虑这个插入的位置 pos，它成立的条件为：nums[pos−1]<target≤nums[pos]
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	ans := len(nums)

	var mid int
	for l <= r {
		mid = l + (r-l)>>1

		//fmt.Println(mid, nums[mid], target)
		if nums[mid] >= target {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ans
}
