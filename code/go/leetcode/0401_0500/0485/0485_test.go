/*
485.简 最大连续 1 的个数

给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。

示例 1：

输入：nums = [1,1,0,1,1,1]
输出：3
解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.
示例 2:

输入：nums = [1,0,1,1,0,1]
输出：2

提示：

1 <= nums.length <= 105
nums[i] 不是 0 就是 1.
*/
package leetcode

import (
	"fmt"
	"testing"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	nums := []int{1, 1, 0, 1, 1, 1}

	fmt.Println(findMaxConsecutiveOnes(nums))
	fmt.Println(findMaxConsecutiveOnes2(nums))
}

func findMaxConsecutiveOnes(nums []int) int {
	m := 0   // 最大连续 1 的个数
	cur := 0 // 当前遇到连续 1 的个数
	lastIdx := len(nums) - 1

	for i, v := range nums {
		if v == 0 {
			m = max(m, cur)

			cur = 0
		} else {
			cur++

			if i == lastIdx { // 最后一个元素是1的情况
				m = max(m, cur)
			}
		}
	}

	return m
}

func findMaxConsecutiveOnes2(nums []int) int {
	ans := 0
	cnt := 0

	for _, v := range nums {
		if v == 1 {
			cnt++
			ans = max(ans, cnt)
		} else {
			cnt = 0
		}
	}

	return ans
}
