/*
264.中 丑数 II

给你一个整数 n ，请你找出并返回第 n 个 丑数 。

丑数 就是质因子只包含 2、3 和 5 的正整数。

示例 1：

输入：n = 10
输出：12
解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
示例 2：

输入：n = 1
输出：1
解释：1 通常被视为丑数。

提示：

1 <= n <= 1690
*/
package demo

import (
	"fmt"
	"testing"
)

func TestNthUglyNumber(t *testing.T) {
	res := nthUglyNumber(10)
	fmt.Println(res)
}

// 较大丑数 = min(last丑数*[2,3,5])
func nthUglyNumber(n int) int {
	dp := make([]int, n+1) //  dp[i] 表示第 i 个丑数，第 n 个丑数即为 dp[n]
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1 // 维护三个指针 (p2, p3, p5)，分别指向 dp 数组中某个丑数，该丑数乘以 2、3 或 5 后，将是下一个待生成的丑数
	for i := 2; i <= n; i++ {
		// 当 2≤i≤n 时，令 dp[i]=min(dp[p2]*2,dp[p3]*3,dp[p5*5]),比较 dp[i]和dp[p2]*2,dp[p3]*3,dp[p5*5]是否相等，如果相等则将对应的指针加 1
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		// 如果某个丑数可以由多种方式生成（例如 6 既是 2*3 也是 3*2），那么所有参与生成它的指针都需要向前移动，以避免重复生成该丑数并保持序列的升序性
		if dp[i] == x2 {
			p2++ // 向前移动一位，指向下一个可能乘以 2 来生成丑数的 dp 数组中的元素
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}
