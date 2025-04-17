package leetcode

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {
	//s := []int{2, 3, 1, 1, 4}
	s := []int{2, 3, 1, 1, 0, 4, 1, 2, 3, 4} // 按每次找到可到达的最远位置跳法, 会卡在s[4]

	fmt.Println(jump(s))
}

func jump(nums []int) int {
	maxPosition := 0 // 目前能跳到的最远位置
	end := 0         // 上次跳跃可达范围右边界（下次的最右起跳点）
	steps := 0       // 跳跃次数
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end { // 到达上次跳跃能到达的右边界了
			end = maxPosition // 目前能跳到的最远位置变成了下次起跳位置的有边界
			steps++
		}
	}
	return steps
}
