// 121. 买卖股票的最佳时机
// 思路:
// 1. 求min,max. `prices[i]−minPrice` 的最大值，就是答案
package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(nums))
}

// best
// maxprofit=最大利润
func maxProfit(prices []int) (maxprofit int) {
	if len(prices) < 2 {
		return
	}

	minPrice := prices[0] // 最低买入价格
	for _, p := range prices {
		minPrice = min(minPrice, p)
		maxprofit = max(maxprofit, p-minPrice) // max(之前利润, 本次交易利润)
	}
	return
}
