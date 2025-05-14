/*
121.简 买卖股票的最佳时机

给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

示例 1：

输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。

	注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。

示例 2：

输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
*/
package leetcode

import (
	"fmt"
	"testing"
)

// 思路:
// 1. 求min,max. `prices[i]−minPrice` 的最大值，就是答案
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
