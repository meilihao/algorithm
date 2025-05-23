/*
70.简 爬楼梯

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

示例 1：

输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
示例 2：

输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶
*/
package leetcode

import (
	"fmt"
	"testing"
)

// 思路:
// 1. 递归求解
// 2. 动态规划
func TestClimbStairs(t *testing.T) {
	n := 5

	fmt.Println(climbStairs(n))
	//fmt.Println(climbStairs3(n))
}

// 动态规划

// 算法: f(n) = f(n-1) + f(n-2)
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
