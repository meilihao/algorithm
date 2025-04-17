// 122. 买卖股票的最佳时机 II
// 思路:
// 1. 贪心
// 2. dp(动态规划)
package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxProfit122(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit122(nums))
}

// 贪心
func maxProfit122(prices []int) int {
	res := 0

	for i := 1; i < len(prices); i++ { // 已能排除len(prices) < 2的情况
		if prices[i] > prices[i-1] { // 只要当天的股价比前一天高就买卖一次
			res += prices[i] - prices[i-1]
		}
	}

	return res
}
