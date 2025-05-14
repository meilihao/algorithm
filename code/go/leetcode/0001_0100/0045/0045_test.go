/*
45.中 跳跃游戏 II

给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

0 <= j <= nums[i]
i + j < n
返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

示例 1:

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。

	从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

示例 2:

输入: nums = [2,3,0,1,4]
输出: 2
*/
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
	last := 0        // 上次跳跃可达范围右边界（下次的最右起跳点）
	steps := 0       // 跳跃次数
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == last { // 到达上次跳跃能到达的右边界了, 需要重新跳了
			last = maxPosition // 目前能跳到的最远位置变成了下次起跳位置的有边界
			steps++
		}
	}
	return steps
}
