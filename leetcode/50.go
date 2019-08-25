// 50. Pow(x, n)
// 思路分治
package main

import "fmt"

func main() {
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
