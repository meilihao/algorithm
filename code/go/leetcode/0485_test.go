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
