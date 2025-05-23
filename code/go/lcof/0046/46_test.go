/*
46：把数字翻译成字符串

给定一个数字，我们按照如下规则把它翻译为字符串：0翻译成"a"，1翻
译成"b"，……，11翻译成"l"，……，25翻译成"z"。一个数字可能有多个翻译。例
如12258有5种不同的翻译，它们分别是"bccfi"、"bwfi"、"bczi"、"mcfi"和
"mzi"。请编程实现一个函数用来计算一个数字有多少种不同的翻译方法。
*/
package demo

import (
	"fmt"
	"strconv"
	"testing"
)

func TestTranslateNum(t *testing.T) {
	res := translateNum(12258)
	fmt.Println(res)
	res2 := translateNum2(12258)
	fmt.Println(res2)
	res3 := translateNum2(06)
	fmt.Println(res3)
}

func translateNum(num int) int {
	str := strconv.Itoa(num)
	l := len(str)
	var dp = make([]int, l)
	dp[0] = 1
	for i := 1; i < l; i++ {
		tmp, _ := strconv.Atoi(string(str[i-1]) + string(str[i]))
		if tmp <= 25 && tmp >= 10 {
			if i-2 < 0 {
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = dp[i-1] + dp[i-2]
			}
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[l-1]
}

/*
设 dp[i] 表示字符串 s 前 i 个数字 s[0: i] 的翻译方案数。dp[i] 的来源有两种情况：
- 第 i - 1、i - 2 构成的数字在 [10, 25]之间，则 dp[i] 来源于： s[i - 1] 单独翻译的方案数（即 dp[i - 1]） + s[i - 2] 和 s[i - 1] 连起来进行翻译的方案数（即 dp[i - 2]）
- 第 i - 1、i - 2 构成的数字在 [10, 25]之外，则 dp[i] 来源于：s[i] 单独翻译的方案数
*/
func translateNum2(num int) int {
	str := strconv.Itoa(num)
	l := len(str)
	var dp = make([]int, l+1)
	dp[0] = 1 // 空字符串只有一种解码方式
	dp[1] = 1 // 第一个数字只能被解码成一种方式

	for i := 2; i <= l; i++ {
		tmp, _ := strconv.Atoi(string(str[i-2 : i]))
		if tmp >= 10 && tmp <= 25 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[l]
}
