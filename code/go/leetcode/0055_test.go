package leetcode

import (
	"fmt"
	"testing"
)

func TestCanJump(t *testing.T) {
	s := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(s))
}

func canJump(nums []int) bool {
	mx := 0 // 当前能跳到的最远索引
	for i, jump := range nums {
		if i > mx { // 无法到达 i. i<=mx说明可跳到该位置
			return false
		}
		mx = max(mx, i+jump) // 从 i 最右可以跳到 i + jump
	}
	return true
}
