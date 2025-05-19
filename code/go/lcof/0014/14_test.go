/*
14. 剪绳子

给你一根长度为n绳子，请把绳子剪成m段（m、n都是整数，n>1并且m≥1）。
每段的绳子的长度记为k[0]、k[1]、……、k[m]。k[0]*k[1]*…*k[m]可能的最大乘
积是多少？例如当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此
时得到最大的乘积18。
*/
package demo

import (
	"fmt"
	"testing"
)

func TestCuttingRope(t *testing.T) {
	fmt.Println(cuttingRope(50) == 86093442)
	fmt.Println(cuttingRope2(50) == 86093442)
}

func cuttingRope(n int) int {
	var r = make(map[int]int, n+1)
	return dp(n, &r)
}

// f(n)=max(i * max(f(n-1), n-i))
func dp(n int, r *map[int]int) int {
	//定义0，1，2，3对应的结果值
	/*
		长度为 0：无法剪，定义为 1（不影响乘积）
		长度为 1：无法剪成两段，定义为 1
		长度为 2：只能剪成 1 和 1，乘积为 1
		长度为 3：可以剪成 1 和 2，乘积为 2
	*/
	(*r)[0], (*r)[1], (*r)[2], (*r)[3] = 1, 1, 1, 2
	var tmp int
	//从2开始，循环
	for i := 2; i <= n-2; i++ {
		//如果r中有结果，说明该值已经计算过，不再重复计算
		if _, ok := (*r)[n-i]; ok { // 检查长度为 n-i 的绳子的最大乘积是否已经计算过并存储在 map r 中
			tmp = (*r)[n-i]
		} else {
			tmp = dp(n-i, r) // 递归计算
		}

		// 比较两种情况：继续剪(i*f(n-i)) vs 不继续剪(i*(n-i))
		if i*tmp > i*(n-i) {
			tmp = i * tmp
		} else {
			tmp = i * (n - i)
		}
		// 更新最大值
		if tmp > (*r)[n] {
			(*r)[n] = tmp
		}
	}
	return (*r)[n]
}

func cuttingRope2(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1 // min(i-i/2) = 2

	// i=绳子长度, j=被减长度
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ { // `j <= i/2`=> a*b = b*a
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j])) // max(dp[i],...) = dp[i]在遍历j的过程中会被多次更新, 即多种切法中取所有j中的最大值
		}
	}

	//fmt.Println(dp)
	return dp[n]
}

// [0 0 1 2 4 6 9 12 18 27 36 54
// [0 0 1 2 4 6 9 12 16 24 30 45
