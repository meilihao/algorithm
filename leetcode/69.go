package main

import "fmt"

func main() {
	n := 5 // 1,4,8,9
	fmt.Println(mySqrt(n))
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
