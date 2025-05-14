/*
204.中 计数质数

给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。

示例 1：

输入：n = 10
输出：4
解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
示例 2：

输入：n = 0
输出：0
示例 3：

输入：n = 1
输出：0

提示：

0 <= n <= 5 * 106
*/
package leetcode

import (
	"fmt"
	"testing"
)

func TestCountPrimes(t *testing.T) {
	n := 10 //4
	fmt.Println(countPrimes(n))
	fmt.Println(countPrimes2(n))
}

// 埃氏筛
// 从2开始，2的倍数全部去掉；接下来3是质数，3的倍数全部去掉，如此下去，留下的都是质数
func countPrimes(n int) (cnt int) {
	isPrime := make([]bool, n)
	for i := range isPrime { // isPrime[i] 表示数 i 是不是质数
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
			for j := 2 * i; j < n; j += i { // 2x,3x...不是整数
				isPrime[j] = false
			}
		}
	}
	return
}

// 如果n是合数, 那么这对约数必须一个在根号n之前，一个在根号n之后. 因为都在根号n之前的话，乘积一定小于n, 同样都在根号n之后的话，乘积一定大于n.
// 容易超时
func countPrimes2(n int) (cnt int) {
	isPrime := func(x int) bool {
		for i := 2; i*i <= x; i++ { // i in [2, n^0.5]
			if x%i == 0 {
				return false
			}
		}
		return true
	}

	for i := 2; i < n; i++ {
		if isPrime(i) {
			cnt++
		}
	}
	return
}
