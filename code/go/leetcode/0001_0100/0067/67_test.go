/*
67.简 二进制求和

给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。

示例 1：

输入:a = "11", b = "1"
输出："100"
示例 2：

输入：a = "1010", b = "1011"
输出："10101"

提示：

1 <= a.length, b.length <= 104
a 和 b 仅由字符 '0' 或 '1' 组成
字符串如果不是 "0" ，就不含前导零
*/
package demo

import (
	"fmt"
	"slices"
	"strconv"
	"testing"
)

func TestAddBinary(t *testing.T) {
	fmt.Println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	var result []byte

	as, bs := []byte(a), []byte(b)
	i, j := len(as)-1, len(bs)-1

	var x, y, sum, carry byte

	n := 0
	for i >= 0 || j >= 0 {
		n++
		x, y, sum = 0, 0, 0

		if i >= 0 {
			x = as[i] - '0'
			i--
		}
		if j >= 0 {
			y = bs[j] - '0'
			j--
		}

		//fmt.Println("1---", x, y, carry)

		sum = x + y + carry
		carry = sum / 2
		sum %= 2

		//fmt.Println("2---", carry, sum)

		result = append(result, sum+'0')
	}

	if carry == 1 {
		result = append(result, '1')
	}

	slices.Reverse(result)
	return string(result)
}

func addBinary2(a string, b string) string {
	var res string
	alen, blen, carry := len(a), len(b), 0
	n := max(alen, blen)
	for i := 0; i < n; i++ {
		if i < alen {
			carry += int(a[alen-i-1] - '0')
		}
		if i < blen {
			carry += int(b[blen-i-1] - '0')
		}
		res = strconv.Itoa(carry%2) + res
		carry /= 2
	}
	if carry > 0 {
		res = "1" + res
	}
	return res
}
