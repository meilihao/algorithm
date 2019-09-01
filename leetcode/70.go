// 70.  爬楼梯
// 思路:
// 1. 递归求解
// 2. 动态规划
package main

import "fmt"

func main() {
	n := 5

	fmt.Println(climbStairs3(n))
}

// 动态规划

// 算法:
// 第 i 阶可以由以下两种方法得到：
// 1. 在第 (i−1) 阶后向上爬一阶
// 2. 在第 (i−2) 阶后向上爬 2 阶
// 即 dp[i]=dp[i−1]+dp[i−2]
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 斐波那契数 : dp[i]=dp[i−1]+dp[i−2]
func climbStairs11(n int) int {
	if n == 1 {
		return 1
	}

	first, second := 1, 2
	for i := 3; i <= n; i++ {
		first, second = second, first+second
	}

	return second
}

// 容易超时: o(n^2)
// bad
func climbStairs2(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

// 记忆化递归 : o(n)
// 容易超时
func climbStairs3(n int) int {
	if n < 1 {
		return -1
	}

	a := make([]int, n+1)

	return climbStairsWithMemory(n, a)
}

func climbStairsWithMemory(n int, a []int) int {
	if n <= 1 { // 即a[0]=a[1]=1
		return 1
	} else if n == 2 {
		return 2
	} else if a[n] == 0 {
		a[n] = climbStairsWithMemory(n-1, a) + climbStairsWithMemory(n-2, a)
	}

	return a[n]
}
