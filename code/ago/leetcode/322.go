/*
	Coin Change 零钱兑换

	思路(动态规划的递推):
	for _,v:=range coins {
		dp[i]= min(1, dp[amount-v]+1) // 1: 刚好有该面值的硬币, dp[amount-v]+1: 一枚面值v的硬币 + dp[amount-v]的最优解
	}
*/
package main

import "fmt"

func main() {
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
		if coinChange2(v.coins, v.amount) != v.want {
			fmt.Println("---", v)
		}
	}
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) // dp[i] 表示凑齐钱数 i 需要的最少硬币数
	// dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = -1

		for j := 0; j < len(coins); j++ {
			// i>=coins[j]: i应该不小于最小硬币
			// dp[i-coins[j]] != -1:
			if i >= coins[j] && dp[i-coins[j]] != -1 {
				if dp[i] == -1 || dp[i] > dp[i-coins[j]]+1 {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
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
