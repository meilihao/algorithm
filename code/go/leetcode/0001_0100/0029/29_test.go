/*
29.中 两数相除

给你两个整数，被除数 dividend 和除数 divisor。将两数相除，要求 不使用 乘法、除法和取余运算。

整数除法应该向零截断，也就是截去（truncate）其小数部分。例如，8.345 将被截断为 8 ，-2.7335 将被截断至 -2 。

返回被除数 dividend 除以除数 divisor 得到的 商 。

注意：假设我们的环境只能存储 32 位 有符号整数，其数值范围是 [−231,  231 − 1] 。本题中，如果商 严格大于 231 − 1 ，则返回 231 − 1 ；如果商 严格小于 -231 ，则返回 -231 。

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = 3.33333.. ，向零截断后得到 3 。
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = -2.33333.. ，向零截断后得到 -2 。

提示：

-231 <= dividend, divisor <= 231 - 1
divisor != 0
*/
package demo

import (
	"fmt"
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	fmt.Println(divide(10, 3) == 3)
	fmt.Println(divide(7, -3) == -2)
	fmt.Println(divide(0, -1) == 0)
	fmt.Println(divide(0, 1) == 0)
	fmt.Println(divide(-2147483648, -1))
	fmt.Println(divide(2147483647, 3))
	fmt.Println(0xc0000000)
	fmt.Println(math.MinInt32)
	fmt.Println(math.MinInt32 / 2)
}

// int32: -2^31 ~ 2^31-1
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == -1 {
			return math.MaxInt32
		}
		if divisor == 1 {
			return math.MinInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	negative := 2
	if dividend > 0 { // 因-2^31转正式时会溢出, 因此转成负数再计算
		negative--
		dividend = -dividend
	}

	if divisor > 0 {
		negative--
		divisor = -divisor
	}

	result := divideCore(dividend, divisor)
	if negative == 1 {
		return -result
	} else {
		return result
	}
}

func divideCore(dividend int, divisor int) int {
	limit := math.MinInt32 / 2 // limit = -2^30, 为了防止 value 在执行 value += value（即 value *= 2）时溢出变成正数. 不能用limit=0xc0000000, 因为其类型是int, 此时limit是正整数

	var result, value, quotient int // result, 累积的商; value, 临时存储 divisor * 2^k; quotient, 存储当前减法步骤的 2^k
	for dividend <= divisor {       // 因为dividend, divisor都是负数, 因此 dividend <= divisor 实际上意味着 |dividend| >= |divisor|（被除数的绝对值大于等于除数的绝对值）即还能整除
		value = divisor
		quotient = 1

		for value >= limit && dividend <= value+value { // 检查 2 * value（即 divisor * 2^(k+1)）是否仍然可以从 dividend 中减去. |dividend| >= |2*value|
			quotient += quotient // 将 quotient 加倍（例如，从 1 到 2，2 到 4，以此类推）
			value += value       // 将 value 加倍（例如，从 divisor 到 2*divisor，4*divisor，以此类推）
		}

		result += quotient
		dividend -= value
	}

	return result
}
