// 231. 2的幂
package main

import "fmt"

func main() {
	n := 1

	fmt.Println(isPowerOfTwo2(n))
}

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	for n > 1 {
		if n%2 == 1 {
			return false
		}
		n /= 2
	}

	return true
}

// best
func isPowerOfTwo2(n int) bool {
	// x = x & (x-1) // 清除序列最后面(即最低位)的1
	return n > 0 && n&(n-1) == 0
}
