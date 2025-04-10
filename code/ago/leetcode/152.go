// 152. 乘积最大子序列
package main

import "fmt"

func main() {
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

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// todo
// 根据符号的个数:
// 1. 当负数个数为偶数时候, 全部相乘一定最大
// 1. 当负数个数为奇数时候, 它的左右两边的负数个数一定为偶数, 只需求两边最大值
// 1. 当有0情况,重置就可以了
