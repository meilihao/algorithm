/*
154.困 寻找旋转排序数组中的最小值 II

已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,4,4,5,6,7] 在变化后可能得到：
若旋转 4 次，则可以得到 [4,5,6,7,0,1,4]
若旋转 7 次，则可以得到 [0,1,4,4,5,6,7]
注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。

给你一个可能存在 重复 元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。

你必须尽可能减少整个过程的操作步骤。



示例 1：

输入：nums = [1,3,5]
输出：1
示例 2：

输入：nums = [2,2,2,0,1]
输出：0


提示：

n == nums.length
1 <= n <= 5000
-5000 <= nums[i] <= 5000
nums 原来是一个升序排序的数组，并进行了 1 至 n 次旋转
*/

package demo

import (
	"fmt"
	"testing"
)

/*
假设最小值的索引是p, p将数组分成了两部分, 每部分都是非严格递增的数组, 且前半部的每个值都>=后半部的每个值
*/
func TestMinArray(t *testing.T) {
	nums1 := []int{0, 1, 2, 3, 4, 5} // 旋转n次变回原数组
	nums2 := []int{3, 4, 5, 0, 1, 2}
	nums3 := []int{1, 1, 1, 0, 1}
	nums4 := []int{1, 1, 1, 1, 1}

	fmt.Println(minArray(nums1) == 0)
	fmt.Println(minArray(nums2) == 0)
	fmt.Println(minArray(nums3) == 0)
	fmt.Println(minArray(nums4) == 1)
}

func minArray(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l < r { // 当 l == r 时，表示已找到, 循环结束，nums[l] 就是最小值
		mid := l + (r-l)/2
		fmt.Println(l, r, mid)
		// nums[mid]与后半部的最大值比较
		if nums[mid] < nums[r] { // 因为every(前半)>=every(后半) , 因此mid在后半部分, 又因为后半也有序, 因此nums[mid]是最小值右侧的元素或本身
			r = mid
		} else if nums[mid] > nums[r] { // mid在前半部分, p肯定在mid后面
			l = mid + 1
		} else {
			// 由于重复元素的存在，并不能确定 mid 究竟在最小值的左侧还是右侧.
			// 由于它们的值相同，所以无论 nums[r] 是不是最小值，都有一个它的「替代品」nums[mid]，因此可以忽略二分查找区间的右端点
			r--
		}
	}
	return nums[l]
}
