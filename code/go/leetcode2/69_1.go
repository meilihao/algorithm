// 已知sqrt(2)约等于1.414，要求不用数学库，求sqrt(2)精确到小数点后10位
// 在(1.4,1.5)区间做二分查找
package main

import "fmt"

func main() {
	fmt.Println(sqrt2())
}

func sqrt2() float64 {
	var low, up, epsinon float64 = float64(1.4), float64(1.5), float64(0.0000000001)
	var mid float64

	for up-low > epsinon {
		mid = (low + up) / 2

		if mid*mid > 2 {
			up = mid
		} else {
			low = mid
		}
	}

	return mid
}
