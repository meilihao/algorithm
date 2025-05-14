/*
50.中 Pow(x, n)

实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。

示例 1：

输入：x = 2.00000, n = 10
输出：1024.00000
示例 2：

输入：x = 2.10000, n = 3
输出：9.26100
示例 3：

输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25
*/
package leetcode

import (
	"fmt"
	"testing"
)

// 思路分治
// 快速幂算法的核心思想是将幂指数 n 拆分为若干个二进制位上的 1 的和，然后将x 的n 次幂转化为 x 的若干个幂的乘积
// 时间复杂度 O(logn)，空间复杂度O(1), n 为幂指数
func TestMyPow(t *testing.T) {
	x := 2.0
	n := 15

	//fmt.Println(myPow(x, n))
	fmt.Println(myPow2(x, n))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	if n < 0 {
		return 1.0 / pow(x, -n)
	}

	return pow(x, n)
}

// 递归
func pow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	if n == 1 {
		return x
	}

	res := pow(x, n>>1)
	if n&1 == 0 { // n&1 判断奇偶: 0,偶数; 1, 奇数
		return res * res
	}

	return res * res * x
}

// 非递归
// 由低次幂逐渐升级即：x, x^2, x^4, x^8, ...
// x = 2, n =15, 推算轮次:
// 1. 2 * 4^7
// 2. 2*4*16^3
// 3. 2*4*16*32^1
func myPow2(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	final := 1.0

	for n > 0 {
		if n&1 != 0 { // n是奇数
			final *= x // 取出一个x先处理
		}

		x *= x
		n >>= 1

		fmt.Println(x, n)
	}

	return final
}
