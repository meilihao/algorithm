/*
400.中 第 N 位数字

给你一个整数 n ，请你在无限的整数序列 [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...] 中找出并返回第 n 位上的数字。

示例 1：

输入：n = 3
输出：3
示例 2：

输入：n = 11
输出：0
解释：第 11 位数字在序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... 里是 0 ，它是 10 的一部分。

提示：

1 <= n <= 231 - 1
*/
package demo

import (
	"math"
	"testing"
)

func TestFindNthDigit(t *testing.T) {

}

/*
规律:
1. 1~9 -> 9 * 1 -> 9 *1 *1
2. 10~99 -> 90 * 2 -> 9 *10 *2
3. 100~999 -> 900 * 3 -> 9 *100 *3
...
*/
func findNthDigit(n int) int {
	d := 1
	for count := 9; n > d*count; count *= 10 { // n=减去指定个数的剩余个数
		n -= d * count
		d++
	}
	index := n - 1                                    // 为了方便计算目标数字，使用目标数字在所有 d 位数中的下标进行计算，下标从 0 开始计数。令 index=n−1，则 index 即为目标数字在所有 d 位数中的下标，index 的最小可能取值是 0
	start := int(math.Pow10(d - 1))                   // 当前 d 位数的起始数字
	num := start + index/d                            // 第 n 位数字所在的那个具体数字 = start + 所有 d 位数字中的总索引
	digitIndex := index % d                           // 第 n 位数字在 num 这个数中是从左往右的第几位（从 0 开始计数）
	return num / int(math.Pow10(d-digitIndex-1)) % 10 // d - digitIndex - 1 = 计算的是要将 num 除以 10 的多少次幂
}
