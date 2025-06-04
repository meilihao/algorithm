/*
322.中 零钱兑换

给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。

示例 1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
示例 2：

输入：coins = [2], amount = 3
输出：-1
示例 3：

输入：coins = [1], amount = 0
输出：0

提示：

1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104
*/
package main

import (
	"fmt"
	"testing"
)

/*
Coin Change 零钱兑换

思路(动态规划的递推):

	for _,v:=range coins {
		dp[i]= min(1, dp[amount-v]+1) // 1: 刚好有该面值的硬币, dp[amount-v]+1: 一枚面值v的硬币 + dp[amount-v]的最优解
	}
*/
func TestCoinChange(t *testing.T) {
	cases := []struct {
		coins  []int
		amount int
		want   int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1, 2, 5, 7, 10}, 14, 2},
	}

	for _, v := range cases {
		if coinChange(v.coins, v.amount) != v.want {
			fmt.Println("---", v)
		} else {
			fmt.Println("--- ok")
		}
	}
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) // dp[i] 表示凑齐钱数 i 需要的最少硬币数
	// dp[0] = 0 // 凑齐金额 0 需要 0 个硬币（空集）

	for i := 1; i <= amount; i++ {
		dp[i] = -1 // 将 dp[i] 设置为 -1，表示金额 i 暂时不可达

		for j := 0; j < len(coins); j++ {
			// 判断是否可以使用当前硬币 coin 来凑齐金额 i：
			// 1. `i >= coin`：当前金额 `i` 必须大于或等于当前硬币面额 `coin`，才能使用这枚硬币。
			// 2. `dp[i-coin] != -1`：表示金额 `i-coin` 是可以被凑齐的。如果 `i-coin` 都不可以凑齐，
			//    那么加上当前硬币 `coin` 也无法凑齐 `i`。
			if i >= coins[j] && dp[i-coins[j]] != -1 {
				if dp[i] == -1 || dp[i] > dp[i-coins[j]]+1 {
					//fmt.Println("-b:", dp[i] > dp[i-coins[j]]+1, i, coins[j], dp[i-coins[j]])
					dp[i] = dp[i-coins[j]] + 1
					//fmt.Println("-a:", dp[i])
				}
			}
		}
	}

	return dp[amount]
}

// dp[i] 表示交换到i的最小次数 不一定有1 但是可以用amount+1初始化
// 遍历coins数组, dp[i] = min(dp[i], dp[i-coin]+1)
// 官方
// best: 理解
func coinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1 //  amount + 1 : 不可能达到的方案
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				//fmt.Println("--:", i, coin, dp[i], dp[i-coin]+1)
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1) // dp[i] 表示凑齐钱数 i 需要的最少硬币数
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}

	for _, coin := range coins {
		dp[coin] = 1

		for i := coin; i <= amount; i++ {
			//fmt.Println("*", dp[i], dp[i-coin])
			if dp[i-coin] != -1 && dp[i] > dp[i-coin]+1 {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	fmt.Println(dp)
	if dp[amount] == 0x7ffffffe {
		return -1
	}

	return dp[amount]
}
