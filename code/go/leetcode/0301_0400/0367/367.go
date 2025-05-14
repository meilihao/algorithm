/*
367.简 有效的完全平方数

给你一个正整数 num 。如果 num 是一个完全平方数，则返回 true ，否则返回 false 。

完全平方数 是一个可以写成某个整数的平方的整数。换句话说，它可以写成某个整数和自身的乘积。

不能使用任何内置的库函数，如  sqrt 。

示例 1：

输入：num = 16
输出：true
解释：返回 true ，因为 4 * 4 = 16 且 4 是一个整数。
示例 2：

输入：num = 14
输出：false
解释：返回 false ，因为 3.742 * 3.742 = 14 但 3.742 不是一个整数。

提示：

1 <= num <= 231 - 1
*/
package main

import (
	"fmt"
	"testing"
)

func TestIsPerfectSquare(t *testing.T) {
	n := 14

	fmt.Println(isPerfectSquare(n))
}

func isPerfectSquare(num int) bool {
	r := mySqrt2(num)
	return r*r == num
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

		//	fmt.Println(low, up, mid, tmp)

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
