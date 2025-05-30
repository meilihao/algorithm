/*
852.中 山脉数组的峰顶索引

给定一个长度为 n 的整数 山脉 数组 arr ，其中的值递增到一个 峰值元素 然后递减。

返回峰值元素的下标。

你必须设计并实现时间复杂度为 O(log(n)) 的解决方案。

示例 1：

输入：arr = [0,1,0]
输出：1
示例 2：

输入：arr = [0,2,1,0]
输出：1
示例 3：

输入：arr = [0,10,5,2]
输出：1

提示：

3 <= arr.length <= 105
0 <= arr[i] <= 106
题目数据 保证 arr 是一个山脉数组
*/
package leetcode

import (
	"fmt"
	"testing"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	//nums := []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}
	nums := []int{3, 5, 3, 2, 0}
	fmt.Println(peakIndexInMountainArray(nums))
}

// 最小的满足 nums[i]>nums[i+1]的下标i
func peakIndexInMountainArray(nums []int) int {
	l, r := 0, len(nums)-1
	ans := 0

	var mid int
	for l <= r {
		mid = l + (r-l)>>1

		fmt.Println(mid, l, r)

		if nums[mid] > nums[mid+1] {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ans
}
