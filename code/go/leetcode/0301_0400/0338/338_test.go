/*
338.简 比特位计数

给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。

示例 1：

输入：n = 2
输出：[0,1,1]
解释：
0 --> 0
1 --> 1
2 --> 10
示例 2：

输入：n = 5
输出：[0,1,1,2,1,2]
解释：
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101

提示：

0 <= n <= 105
*/
package main

import (
	"fmt"
	"testing"
)

func TestCountBits(t *testing.T) {
	num := 5

	fmt.Println(countBits(num))
}

// best
func countBits(num int) []int {
	res := make([]int, num+1)

	for i := 1; i <= num; i++ {
		//fmt.Println(i & (i - 1))
		// i 比 i&(i-1) 多一个1, 同时 i > i&(i-1), 因此递推公式成立
		res[i] = res[i&(i-1)] + 1

		// 或者: i >> 1去掉i的最低位；因(i >> 1) < i，故res[i >> 1]已计算，因此i中1的个数为i >> 1中1的个数加最后一位1的个数，即为res[i >> 1] + (i & 1)
		// res[i] = res[i>>1]+(i&1)
	}

	return res
}

func countBits3(num int) []int {
	res := make([]int, num+1)

	for i := 1; i <= num; i++ {
		if i&1 > 0 { // 奇数
			// 二进制表示中，奇数一定比前面那个偶数多一个 1，因为多的就是最低位的 1
			// 比如: 2 = 10      3 = 11
			res[i] = res[i-1] + 1
		} else {
			// 二进制表示中，偶数中 1 的个数一定和除以 2 之后的那个数一样多. 因为最低位是 0，除以 2 就是右移一位，也就是把那个 0 抹掉而已，所以 1 的个数是不变的
			// 比如: 2 = 10       4 = 100       8 = 1000
			//       3 = 11       6 = 110       12 = 1100
			res[i] = res[i/2]
		}
	}

	return res
}

func countBits2(num int) []int {
	res := make([]int, num+1)

	for i := 0; i <= num; i++ {
		res[i] = hammingWeight(i)
	}

	return res
}

func hammingWeight(num int) int {
	if num == 0 {
		return 0
	}

	res := 0
	for num > 0 {
		res++

		// x = x & (x-1) // 清除序列最后面(即最低位)的1
		num &= (num - 1)
	}

	return res
}
