/*
377.中 组合总和 Ⅳ

给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。

题目数据保证答案符合 32 位整数范围。

示例 1：

输入：nums = [1,2,3], target = 4
输出：7
解释：
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。
示例 2：

输入：nums = [9], target = 3
输出：0

提示：

1 <= nums.length <= 200
1 <= nums[i] <= 1000
nums 中的所有元素 互不相同
1 <= target <= 1000
*/
package leetcode

import (
	"testing"
)

func TestCombinationSum4(t *testing.T) {

}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1) // dp[x] 表示选取的元素之和等于 x 的方案数
	dp[0] = 1                   // 只有当不选取任何元素时，元素之和才为 0，因此只有 1 种方案
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				// dp[i] += dp[i-num] 的逻辑：
				// 如果想凑出和 i，并且选择 num 作为“最后一个”数字，
				// 那么剩下的部分 (i - num) 必须由之前的数字凑出。
				// 凑出 (i - num) 的方案数就是 dp[i-num]。
				// 将这些方案数累加到 dp[i] 中。
				dp[i] += dp[i-num] // 使用`dp[i] +=`: dp[i]由多种方案组成
			}
		}
	}
	return dp[target]
}
