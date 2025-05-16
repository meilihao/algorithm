/*
153.中 搜索旋转排序数组

整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。


示例 1：

输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例 2：

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：

输入：nums = [1], target = 0
输出：-1


提示：

1 <= nums.length <= 5000
-104 <= nums[i] <= 104
nums 中的每个值都 独一无二
题目数据保证 nums 在预先未知的某个下标上进行了旋转
-104 <= target <= 104
*/

package demo

import (
	"fmt"
	"testing"
)

/*
假设最小值的索引是p, p将数组分成了两部分, 每部分都是严格递增的数组, 且前半部的每个值都>后半部的每个值
*/
func TestFindMin(t *testing.T) {
	nums1 := []int{0, 1, 2, 3, 4, 5} // 旋转n次变回原数组
	nums2 := []int{3, 4, 5, 0, 1, 2}
	nums3 := []int{4, 5, 6, 7, 0, 1, 2}
	nums4 := []int{3, 1}

	fmt.Println(search(nums1, 1) == 1)
	fmt.Println(search(nums2, 6) == -1)
	fmt.Println(search(nums3, 0) == 4)
	fmt.Println(search(nums4, 1) == 1)
}

func TestFindMin2(t *testing.T) {
	nums1 := []int{0, 1, 2, 3, 4, 5} // 旋转n次变回原数组
	nums2 := []int{3, 4, 5, 0, 1, 2}
	nums3 := []int{4, 5, 6, 7, 0, 1, 2}
	nums4 := []int{3, 1} // 验证需要`nums[l] <= nums[mid]`而不是`nums[l] < nums[mid]`

	fmt.Println(search2(nums1, 1) == 1)
	fmt.Println(search2(nums2, 6) == -1)
	fmt.Println(search2(nums3, 0) == 4)
	fmt.Println(search2(nums4, 1) == 1)
}

func TestFindMin3(t *testing.T) {
	nums1 := []int{0, 1, 2, 3, 4, 5} // 旋转n次变回原数组
	nums2 := []int{3, 4, 5, 0, 1, 2}
	nums3 := []int{4, 5, 6, 7, 0, 1, 2}
	nums4 := []int{3, 1}

	fmt.Println(search3(nums1, 1) == 1)
	fmt.Println(search3(nums2, 6) == -1)
	fmt.Println(search3(nums3, 0) == 4)
	fmt.Println(search3(nums4, 1) == 1)
}

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	end := len(nums) - 1

	check := func(mid int) bool {
		if nums[mid] > nums[end] { // mid左侧单调有序=旋转点在右侧
			return target < nums[mid] && target > nums[end] // 目标值应该在左侧的条件
		} else { // mid右侧单调有序=旋转点在左侧或未旋转
			return target < nums[mid] || target > nums[end] // 目标值应该在左侧的条件
		}
	}

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return mid
		} else if check(mid) { // check是判断目标值是否位于 mid 的左侧
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

// 假设使用`nums[l] < nums[mid]`, search2([3, 1],1)会报错
func search2(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return mid
		}

		fmt.Println(l, mid, r)
		if nums[l] <= nums[mid] { // [l, mid]是有序的
			if nums[l] <= target && target < nums[mid] { // target在[l,mid)即在升序部分
				r = mid - 1
			} else { // 不在[l,mid)即如果目标值不在这个升序子数组的范围内，那么它可能在右半部分
				l = mid + 1
			}
		} else { // [mid, r]是无序的即中间有旋转
			if target > nums[mid] && target <= nums[r] { // target在(mid,r]即在升序部分
				l = mid + 1
			} else { // 不在(mid,r]
				r = mid - 1
			}
		}
	}

	return -1
}

func search3(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return mid
		}

		fmt.Println(l, mid, r)
		if nums[mid] > nums[r] { // [l, mid]是有序的
			if nums[l] <= target && target < nums[mid] { // target在[l,mid)
				r = mid - 1
			} else { // 不在[l,mid)
				l = mid + 1
			}
		} else { // [mid, r]是无序的即中间有旋转
			if target > nums[mid] && target <= nums[r] { // target在(mid,r]
				l = mid + 1
			} else { // 不在(mid,r]
				r = mid - 1
			}
		}
	}

	return -1
}
