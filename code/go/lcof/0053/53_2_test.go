/*
53（二）：0到n-1中缺失的数字

一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字
都在范围0到n-1之内。在范围0到n-1的n个数字中有且只有一个数字不在该数组
中，请找出这个数字。
*/
package demo

import (
	"fmt"
	"testing"
)

func TestGetMissingNumber(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(GetMissingNumber(nums) == 0)

	nums2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println(GetMissingNumber(nums2) == 8)
}

func GetMissingNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	return getMissingNumber(nums, 0, len(nums)-1)
}

/*
二分法
初始化： 左边界 left = 0，右边界 right = len(nums) - 1；代表闭区间 [i, j] 。
循环二分： 当 i ≤ j 时循环 （即当闭区间 [i, j] 为空时跳出）；计算中点 mid = (i + j) / 2
若 nums[mid] = mid ，则 “右子数组的首位元素” 一定在闭区间 [mid + 1, j] 中，因此执行 left = mid + 1；
若 nums[mid] != mid ，则 “左子数组的末位元素” 一定在闭区间 [i, mid - 1] 中，因此执行 right = mid - 1；
返回值： 跳出时，变量 i 和 j 分别指向 “右子数组的首位元素” 和 “左子数组的末位元素” 。因此返回 i 即可。
*/
func getMissingNumber(nums []int, start int, end int) int {
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] != mid {
			//nums是有序数组，如果mid和数字不相同就在左边查找
			end = mid - 1
		} else {
			//如果mid和数字相同，说明左边是连续的有序数组
			//缺失的数字就在右边查找，start向上取整+1
			start = mid + 1
		}
	}

	return start
}
