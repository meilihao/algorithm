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
