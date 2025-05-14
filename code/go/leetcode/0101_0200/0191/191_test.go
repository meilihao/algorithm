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
	var n uint32
	n = 0b00000000000000000000000000001011
	fmt.Println(hammingWeight(n))
}

func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		res++

		// x = x & (x-1) // 清除序列最后面(即最低位)的1
		num &= (num - 1)
	}

	return res
}
