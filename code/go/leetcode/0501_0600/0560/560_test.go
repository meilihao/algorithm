/*
560.中 和为 K 的子数组

给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

子数组是数组中元素的连续非空序列。

示例 1：

输入：nums = [1,1,1], k = 2
输出：2
示例 2：

输入：nums = [1,2,3], k = 3
输出：2

提示：

1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107
*/
package leetcode

import (
	"fmt"
	"testing"
)

/*
注意:
1.  r右移不一定让窗口和变大, 同理l左移不一定让窗口和变小, 因此不适用滑动窗口
*/
func TestThreeSum(t *testing.T) {
	var f func(nums []int, k int) int
	f = subarraySum2

	nums := []int{1, 1, 1}
	fmt.Println(f(nums, 2) == 2)

	fmt.Println("---")
	nums2 := []int{1, 2, 3}
	fmt.Println(f(nums2, 3) == 2)

	fmt.Println("---")
	nums3 := []int{1, 2, 1, 2, 1}
	fmt.Println(f(nums3, 3) == 4)

	fmt.Println("---")
	nums4 := []int{1}
	fmt.Println(f(nums4, 0) == 0)
}

// 时间复杂度：O(n^2)，其中 n 为数组的长度
func subarraySum(nums []int, k int) int {
	count := 0
	// nums[j...i], 固定i, 推算j
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

// 前缀和=从第0项到当前项的总和:= preSum[x] = nums[0]+...+num[x]=> num[x] = preSum[x]-preSum[x-1] & num[i]+...+num[j] = preSum[j]-preSum[i-1]
// 题目转为: i<j && preSum[j] - preSum[i-1] = k (preSum[j] - k = preSum[i-1])
func subarraySum2(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{} // 前缀和 -> 属性次数
	m[0] = 1           // 初始前缀和为 0 出现了 1 次, 例如`1, 1, 1`+i=1时m[pre-k]会用到
	for i := 0; i < len(nums); i++ {
		pre += nums[i]

		if _, ok := m[pre-k]; ok {
			//fmt.Println("---m:", i, pre-k == 0)
			count += m[pre-k]
		}

		m[pre] += 1 // 顺序在`if _, ok := m[pre-k]; ok {`后: 因为start需要小于i
	}
	return count
}
