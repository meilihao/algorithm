/*
69.简 x 的平方根

给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。

示例 1：

输入：x = 4
输出：2
示例 2：

输入：x = 8
输出：2
解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。

提示：

0 <= x <= 231 - 1
*/
package main

import (
	"fmt"
	"testing"
)

func TestMySqrtN(t *testing.T) {
	n := 5 // 1,4,8,9
	fmt.Println(mySqrtN(n))
	fmt.Println(mySqrt2(n))
}

// 使用浮点
func mySqrt(x int) int {
	if x <= 0 {
		return x
	}

	var low, up, target float64 = 0, float64(x), float64(x)
	var mid, tmp float64

	for (up - low) > 1e-6 {
		mid = (low + up) / 2
		tmp = mid * mid

		//fmt.Println(tmp, mid)

		if tmp > target {
			up = mid
		} else if tmp < target {
			low = mid
		} else {
			break
		}
	}

	targetInt := int(mid + 0.5)
	//fmt.Println(mid, targetInt)
	if targetInt*targetInt > x {
		targetInt--
	}

	return targetInt
}

// 使用整数
func mySqrt2(x int) int {
	if x <= 0 {
		return x
	}

	var low, up = 0, x
	var mid, tmp int

	for low <= up {
		mid = (low + up) >> 1 // 移位更快
		tmp = mid * mid

		fmt.Println(low, up, mid, tmp)

		if tmp == x {
			return mid
		} else if tmp > x { // if low==up && mid^2 > x , target = up = mid -1
			up = mid - 1
		} else {
			low = mid + 1
		}
	}

	return up
}

// 牛顿迭代法
func mySqrtN(x int) int {
	if x <= 0 {
		return x
	}

	r := x

	for r*r > x {
		r = (r + x/r) / 2
	}

	return r
}

// best
func mySqrt3(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
