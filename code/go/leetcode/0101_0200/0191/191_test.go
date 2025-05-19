/*
191.简 位1的个数

给定一个正整数 n，编写一个函数，获取一个正整数的二进制形式并返回其二进制表达式中 设置位 的个数（也被称为汉明重量）。

示例 1：

输入：n = 11
输出：3
解释：输入的二进制串 1011 中，共有 3 个设置位。
示例 2：

输入：n = 128
输出：1
解释：输入的二进制串 10000000 中，共有 1 个设置位。
示例 3：

输入：n = 2147483645
输出：30
解释：输入的二进制串 1111111111111111111111111111101 中，共有 30 个设置位。

提示：

1 <= n <= 231 - 1
*/
package main

import (
	"fmt"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	var n int32 = 0b00000000000000000000000000001011
	fmt.Println(hammingWeight(n))
	//fmt.Println(hammingWeight2(n))
	fmt.Println(hammingWeight3(n))
}

func hammingWeight(num int32) int {
	res := 0
	for num > 0 {
		res++

		// x = x & (x-1) // 清除序列最后面(即最低位)的1
		// 6 & (6−1)=4, 110 & 101 = 100 运算结果 4 即为把 6 的二进制位中的最低位的 1 变为 0 之后的结果
		num &= (num - 1)
	}

	return res
}

// hammingWeight2有缺陷, 当num为负数, 移位时补符号1, 导致死循环
func hammingWeight2(num int32) int {
	res := 0
	for num > 0 {
		if num&1 > 0 {
			res++
		}
		num >>= 1
	}

	return res
}

func hammingWeight3(num int32) int {
	res := 0
	var flag int32 = 1
	for flag > 0 {
		if num&flag > 0 {
			res++
		}
		flag <<= 1
	}

	return res
}
