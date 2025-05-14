/*
152.中 乘积最大子数组

给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续 子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

测试用例的答案是一个 32-位 整数。

示例 1:

输入: nums = [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: nums = [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

提示:

1 <= nums.length <= 2 * 104
-10 <= nums[i] <= 10
nums 的任何子数组的乘积都 保证 是一个 32-位 整数
*/
package main

import (
	"fmt"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	c1 := []int{-2, 0, -1, -2, 4}

	fmt.Println(maxProduct(c1))
}

// dp, 因为存在负数, 因此dp[i]应该有两种状态
// 由于存在负数，那么会导致最大的变最小的，最小的变最大的. 因此还需要维护当前最小值_min
// dp[i] = max(nums[i] * pre_max, nums[i] * pre_min, nums[i]), 这里0 不需要单独考虑, 因为当相乘不管最大值和最小值,都会置0
func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	res, _min, _max := nums[0], nums[0], nums[0]

	for i := 1; i < n; i++ {
		if nums[i] < 0 {
			_min, _max = _max, _min
		}

		_min = min(nums[i], nums[i]*_min)
		_max = max(nums[i], nums[i]*_max)

		//fmt.Println(_min, _max)
		res = max(res, _max)
	}

	return res
}

// todo
// 根据符号的个数:
// 1. 当负数个数为偶数时候, 全部相乘一定最大
// 1. 当负数个数为奇数时候, 它的左右两边的负数个数一定为偶数, 只需求两边最大值
// 1. 当有0情况,重置就可以了
