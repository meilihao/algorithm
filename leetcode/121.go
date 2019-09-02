// 121. 买卖股票的最佳时机
// 思路:
// 1. 求min,max
package main

import "fmt"

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(nums))
}

func maxProfit(prices []int) int {
	res := 0

	for i := 1; i < len(prices); i++ { // 已能排除len(prices) < 2的情况
		if prices[i] > prices[i-1] { // 只要当天的股价比前一天高就买卖一次
			res += prices[i] - prices[i-1]
		}
	}

	return res
}
