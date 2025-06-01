/*
213.中 打家劫舍 II

你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

示例 1：

输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
示例 2：

输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。

	偷窃到的最高金额 = 1 + 3 = 4 。

示例 3：

输入：nums = [1,2,3]
输出：3

提示：

1 <= nums.length <= 100
0 <= nums[i] <= 1000
*/
package leetcode

import (
	"testing"
)

func TestRob(t *testing.T) {

}

/*
_rob 函数解决的是 线性抢劫问题，通过维护两个状态变量来动态计算最大金额:
- first：表示考虑到第 i 个房屋时，不抢当前房屋的最大金额
- second：表示考虑到第 i 个房屋时，抢当前房屋的最大金额
*/
func _rob(nums []int) int {
	// dp[i] = max(dp[i-2] + nums[i], dp[i-1]) = max(偷当前房屋 + dp[i-2], 不偷当前房屋 + dp[i-1])
	// 由于 dp[i-2] 就是 first，dp[i-1] 就是 second
	// 所以新的 dp[i] 就是 max(first + v, second)
	// 然后更新 first 和 second，为下一次循环做准备
	// first 变为之前的 second (即 dp[i-1])
	// second 变为当前的 dp[i]
	first, second := nums[0], max(nums[0], nums[1])
	for _, v := range nums[2:] {
		first, second = second, max(first+v, second)
	}
	return second
}

// rob 函数解决的是 环形抢劫问题，通过分成两个子问题来避免考虑首尾相邻的房屋，同时调用 _rob 来处理这两个子问题
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	// 将环形问题分解为两个线性问题，并取最大值
	// _rob(nums[:n-1]): 解决不偷窃最后一间房屋的问题，即偷窃 nums[0] 到 nums[n-2]
	// _rob(nums[1:]):  解决不偷窃第一间房屋的问题，即偷窃 nums[1] 到 nums[n-1]
	return max(_rob(nums[:n-1]), _rob(nums[1:]))
}
