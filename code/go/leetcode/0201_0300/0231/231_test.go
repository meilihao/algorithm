// 231. 2的幂
/*
231.简 2 的幂

给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。

如果存在一个整数 x 使得 n == 2x ，则认为 n 是 2 的幂次方。

示例 1：

输入：n = 1
输出：true
解释：20 = 1
示例 2：

输入：n = 16
输出：true
解释：24 = 16
示例 3：

输入：n = 3
输出：false


提示：

-231 <= n <= 231 - 1
*/
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
