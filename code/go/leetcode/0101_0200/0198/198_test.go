/*
198.中 打家劫舍

你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

示例 1：

输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。

	偷窃到的最高金额 = 1 + 3 = 4 。

示例 2：

输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。

	偷窃到的最高金额 = 2 + 9 + 1 = 12 。

提示：

1 <= nums.length <= 100
0 <= nums[i] <= 400
*/
package demo

import (
	"fmt"
	"testing"
)

func TestTranslateNum(t *testing.T) {
	// res := rob([]int{1, 2, 3, 1})
	// fmt.Println(res == 4)
	// res2 := rob([]int{2, 7, 9, 3, 1})
	// fmt.Println(res2 == 12)
	res3 := rob([]int{2, 1, 1, 2})
	fmt.Println(res3 == 4)
}

/*
对于第 k (k>2) 间房屋，有两个选项：

偷窃第 k 间房屋，那么就不能偷窃第 k−1 间房屋，偷窃总金额为前 k−2 间房屋的最高总金额与第 k 间房屋的金额之和。

不偷窃第 k 间房屋，偷窃总金额为前 k−1 间房屋的最高总金额。

边界条件为：

dp[0]=nums[0] : 只有一间房屋，则偷窃该房屋
dp[1]=max(nums[0],nums[1]) :只有两间房屋，选择其中金额较高的房屋进行偷窃
*/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	n := len(nums)
	dp := make([]int, n+1) // 定义 dp[i] 表示偷了i+1家的最高金额
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
		//fmt.Println(i, dp, dp[i-2]+nums[i], dp[i-1], dp[i])
	}

	return dp[n-1]
}
