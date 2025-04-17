package leetcode

import (
	"fmt"
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	// s := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	s := []int{1, 100}
	fmt.Println(minCostClimbingStairs(s))
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1) // 创建长度为 n+1 的数组 dp，其中 dp[i] 表示达到下标 i 的最小花费

	for i := 2; i <= n; i++ {
		// 当 2≤i≤n 时，可以从下标 i−1 使用 cost[i−1] 的花费达到下标 i，或者从下标 i−2 使用 cost[i−2] 的花费达到下标 i
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[n]
}
