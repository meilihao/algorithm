/*
509.简 斐波那契数

斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

F(0) = 0，F(1) = 1
F(n) = F(n - 1) + F(n - 2)，其中 n > 1
给定 n ，请计算 F(n)

示例 1：

输入：n = 2
输出：1
解释：F(2) = F(1) + F(0) = 1 + 0 = 1
示例 2：

输入：n = 3
输出：2
解释：F(3) = F(2) + F(1) = 1 + 1 = 2
示例 3：

输入：n = 4
输出：3
解释：F(4) = F(3) + F(2) = 2 + 1 = 3
*/

package leetcode

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	fmt.Println(fib(5))
}

/*
F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
*/
func fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	a, b := 1, 1
	result := 0
	for i := 3; i <= n; i++ {
		result = a + b
		a, b = b, result
	}
	return result
}

// best
func fib2(n int) int {
	p, q, r := 0, 1, 1
	for i := 0; i < n; i++ {
		p, q, r = q, r, r+q

	}
	return p
}
