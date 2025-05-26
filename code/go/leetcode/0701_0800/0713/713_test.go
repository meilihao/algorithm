/*
713.中 乘积小于 K 的子数组

给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。

示例 1：

输入：nums = [10,5,2,6], k = 100
输出：8
解释：8 个乘积小于 100 的子数组分别为：[10]、[5]、[2]、[6]、[10,5]、[5,2]、[2,6]、[5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于 100 的子数组。
示例 2：

输入：nums = [1,2,3], k = 0
输出：0

提示:

1 <= nums.length <= 3 * 104
1 <= nums[i] <= 1000
0 <= k <= 106
*/
package leetcode

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{10, 5, 2, 6}
	fmt.Println(numSubarrayProductLessThanK(nums, 100) == 8)
	fmt.Println("---")
	nums2 := []int{1, 2, 3}
	fmt.Println(numSubarrayProductLessThanK(nums2, 0) == 0)
}

func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	prod, i := 1, 0 // prod=子数组 [i,j] 的元素乘积
	for j, num := range nums {
		prod *= num
		for i <= j && prod >= k {
			prod /= nums[i]
			i++
		}
		ans += j - i + 1 // 对于每一个有效的窗口 [i, j]，所有以 j 结尾(可以理解为固定j)且起始于 i 到 j 之间的子数组都是有效的, 这些新有效的子数组的数量是 j - i + 1
		// [10] 1
		// [10, 5] 2
		// [5,2] 2
		// [5, 2, 6] 3
		// fmt.Println(i, j, ans)
	}
	return
}
