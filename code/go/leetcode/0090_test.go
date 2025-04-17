package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestXxx(t *testing.T) {
	s := []int{1, 2, 2, 2}

	fmt.Println(subsetsWithDup(s))
}

func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)

outer:
	for mask := 0; mask < 1<<n; mask++ {
		t := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				if i > 0 && mask>>(i-1)&1 == 0 && v == nums[i-1] { // 若发现没有选择上一个数，且当前数字与上一个数相同，则可以跳过当前生成的子集
					continue outer
				}
				t = append(t, v)
			}
		}
		ans = append(ans, append([]int(nil), t...))
	}
	return
}
