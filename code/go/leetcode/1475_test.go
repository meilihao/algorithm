package leetcode

import (
	"fmt"
	"testing"
)

// 期望min(j)>i & prices[j] <= prices[i]
func TestFinalPrices(t *testing.T) {
	s := []int{8, 4, 6, 2, 3} // [4 2 4 2 3]

	fmt.Println(finalPrices1475_2(s))
}

/*
两层循环:
1. 第一层遍历数组
2. 第二层找到符合条件的j

时间复杂度：O(n^2)，其中 n 为数组的长度。对于每个商品，我们需要遍历一遍数组查找符合题目要求的折扣。

空间复杂度：O(1)。返回值不计入空间复杂度
*/
func finalPrices(prices []int) []int {
	ans := make([]int, len(prices))
	for i, p := range prices {
		discount := 0
		for _, q := range prices[i+1:] {
			if q <= p {
				discount = q
				break
			}
		}
		ans[i] = p - discount
	}
	return ans
}

/*
解法的重点在于考虑如何更高效地计算 prices 中每个元素右边第一个更小的值即每个元素的下一个小于等于它的元素

单调栈中维护当前位置右边的更小的元素列表，从栈底到栈顶的元素是单调递增的

时间复杂度：O(n)，其中 n 为数组的长度。只需遍历一遍数组即可。

空间复杂度：O(n) 需要栈空间存储中间变量，需要的空间为 O(n)
*/
func finalPrices1475(prices []int) []int {
	n := len(prices)
	ans := make([]int, n)

	st := []int{0} // 最后一个元素没有折扣, 处理i=n-1时的情况
	for i := n - 1; i >= 0; i-- {
		p := prices[i]
		//	fmt.Println(p)
		for len(st) > 1 && st[len(st)-1] > p { // 将单调栈中所有大于 prices[i] 的元素弹出单调栈，当前位置右边的第一个小于等于 prices[i] 的元素即为栈顶元素
			st = st[:len(st)-1]
		}
		ans[i] = p - st[len(st)-1] // 原价-折扣
		st = append(st, p)

		//fmt.Println(st)
	}
	return ans
}

// best
func finalPrices1475_2(prices []int) []int {
	st := []int{}
	for i, v := range prices {
		//fmt.Println("p:", i, st)

		for len(st) > 0 && prices[st[len(st)-1]] >= v { // 处理找到的j
			//fmt.Println("s:", i, st[len(st)-1])

			prices[st[len(st)-1]] -= v
			st = st[:len(st)-1]
		}

		st = append(st, i)
	}

	return prices
}
