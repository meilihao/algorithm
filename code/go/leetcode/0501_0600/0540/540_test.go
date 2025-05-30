/*
540.中 有序数组中的单一元素

给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。

请你找出并返回只出现一次的那个数。

你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。

示例 1:

输入: nums = [1,1,2,3,3,4,4,8,8]
输出: 2
示例 2:

输入: nums =  [3,3,7,7,10,11,11]
输出: 10

提示:

1 <= nums.length <= 105
0 <= nums[i] <= 105
*/
package leetcode

import (
	"fmt"
	"testing"
)

func TestSingleNonDuplicate(t *testing.T) {
	nums := []int{3, 3, 7, 7, 10, 11, 11}
	fmt.Println(singleNonDuplicate(nums))

	nums1 := []int{3, 4, 4, 7, 7, 11, 11}
	fmt.Println(singleNonDuplicate(nums1))

	nums2 := []int{3, 3, 7, 7, 10, 10, 11}
	fmt.Println(singleNonDuplicate(nums2))
}

/*
如果规律存在, 那么根据 mid 的奇偶性决定和左边或右边的相邻元素比较：
如果 mid 是偶数，则比较 nums[mid] 和 nums[mid+1] 是否相等；
如果 mid 是奇数，则比较 nums[mid−1] 和 nums[mid] 是否相等
*/
func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1

	var mid int
	for l < r { // 用`l <= r`会变成死循环, 条件见二分查找模板
		mid = l + (r-l)>>1

		mid -= mid & 1 // mid变成偶数
		//fmt.Println(mid, l, r)
		// 画图笔记直观
		if nums[mid] == nums[mid+1] { // num[:mid+2]前都是规律的
			l = mid + 2
		} else {
			r = mid
		}
	}

	return nums[l]
}
