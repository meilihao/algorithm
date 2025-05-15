/*
03.01. 不修改数组找出重复的数字

在一个长度为 n+1 的数组 nums 里的所有数字都在 1～n 的范围内, 所以数组中至少有一个数字是重复的. 请找出数组中任意一个重复的数字, 但不能修改输入的数组.

示例 1：
输入：
[2, 3, 5, 4, 3, 2, 6, 7]
输出：2 或 3
*/

package lcof

import (
	"fmt"
	"testing"
)

func TestFindDuplicateNoChange(t *testing.T) {
	numLists := [][]int{
		// []int{1, 3, 4, 2, 2},
		// []int{0, 0, 1, 3, 4, 5, 2},
		// []int{0, 1, 2, 3, 3, 5, 2},
		[]int{2, 3, 5, 4, 3, 2, 6, 7},
	}
	cases := []struct {
		f    func([]int) int
		want int
	}{
		{
			f:    findDuplicateNoChange,
			want: -1,
		},
	}

	for x, nums := range numLists {
		for y, v := range cases {
			nums2 := make([]int, len(nums))
			copy(nums2, nums)

			if tmp := v.f(nums2); tmp == v.want {
				t.Errorf("idx: (%d, %d), err get: %d", x, y, tmp)
			} else {
				fmt.Printf("idx: (%d, %d), get: %d", x, y, tmp)
			}
		}
	}
}

// 它使用了二分查找的思想，结合计数统计的方法来定位重复数字
// 该算法只找到第一个重复的数
// 时间复杂度: O(logN, 调用countRange的次数) *  O(N, countRange中遍历nums)
func findDuplicateNoChange(nums []int) int {
	start, end, n := 1, len(nums)-1, len(nums)

	for start <= end {
		middle := (end-start)>>1 + start

		// 对于中间值 middle，统计数组中在 [start, middle] 范围内的数字个数
		count := countRange(nums, n, start, middle)
		if end == start {
			if count > 1 { // 因为重复了
				return start
			} else {
				break
			}
		}

		if count > (middle - start + 1) { // [start, middle]中出现重复, 因为如果没有重复那么count应该等于middle - start + 1
			end = middle
		} else {
			start = middle + 1
		}
	}

	return -1
}

func countRange(nums []int, length, start, end int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 0
	for i := 0; i < length; i++ {
		if nums[i] >= start && nums[i] <= end {
			count += 1
		}
	}
	return count
}
