/*
494.中 目标和

给你一个非负整数数组 nums 和一个整数 target 。

向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：

例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

示例 1：

输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
示例 2：

输入：nums = [1], target = 1
输出：1

提示：

1 <= nums.length <= 20
0 <= nums[i] <= 1000
0 <= sum(nums[i]) <= 1000
-1000 <= target <= 1000
*/
package leetcode

import (
	"testing"
)

func TestFindTargetSumWays(t *testing.T) {

}

// 添加"+"的数字之和为p, 添加"-"的数字之和为q, 要求p-q=target
/*
方法:
1. p+q=sum=> 2p=sum+target => p = (sum+target)/2 = 从数组中选出和为 (sum+target)/2的数字
2. p=sum-q => (sum-q) -q = target => q = (sum-target)/2 = 从数组中选出和为 (sum-target)/2的数字

f(i,j):
- i < nums[i] : f(i-1, j)
- i >= nums[i] : f(i-1, j) + f(i-1, j-nums[i])
*/
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	// diff < 0：意味着即使所有元素都取 +，和仍然小于 target，无法达到目标
	// diff % 2 == 1：意味着 target + sum 是奇数，那么 (target + sum) / 2 无法整除，q 就不是整数，这对于整数数组的子集求和来说是不可能的, 即无法将其分为两部分
	if diff < 0 || diff%2 == 1 {
		return 0
	}

	n, neg := len(nums), diff/2
	dp := make([][]int, n+1) // dp[i][j] 表示：从 nums 的前 i 个数字 (nums[0...i-1]) 中，有多少种方式可以选择一些数字，使得它们的和恰好等于 j
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1 // 当没有任何元素可以选取时，元素和只能是 0，对应的方案数是 1即选择空集
	for i, num := range nums {
		for j := 0; j <= neg; j++ { // j 代表目标和（从 0 到 neg）
			// 情况 1：不选择当前的数字 num
			// 那么从 nums 的前 i+1 个数字中凑出和 j 的方案数，
			// 就等于从 nums 的前 i 个数字中凑出和 j 的方案数。
			dp[i+1][j] = dp[i][j]

			// 情况 2：选择当前的数字 num, 剩于j-num
			// 这只有当目标和 j 至少大于等于当前数字 num 时才可能
			if j >= num {
				dp[i+1][j] += dp[i][j-num]
			}
		}
	}
	return dp[n][neg]
}
