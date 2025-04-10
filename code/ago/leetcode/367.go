// 367. 有效的完全平方数
package main

import "fmt"

func main() {
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
