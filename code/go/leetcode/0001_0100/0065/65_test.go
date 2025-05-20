/*
65.困 有效数字

给定一个字符串 s ，返回 s 是否是一个 有效数字。

例如，下面的都是有效数字："2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"，而接下来的不是："abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"。

一般的，一个 有效数字 可以用以下的规则之一定义：

一个 整数 后面跟着一个 可选指数。
一个 十进制数 后面跟着一个 可选指数。
一个 整数 定义为一个 可选符号 '-' 或 '+' 后面跟着 数字。

一个 十进制数 定义为一个 可选符号 '-' 或 '+' 后面跟着下述规则：

数字 后跟着一个 小数点 .。
数字 后跟着一个 小数点 . 再跟着 数位。
一个 小数点 . 后跟着 数位。
指数 定义为指数符号 'e' 或 'E'，后面跟着一个 整数。

数字 定义为一个或多个数位。

示例 1：

输入：s = "0"

输出：true

示例 2：

输入：s = "e"

输出：false

示例 3：

输入：s = "."

输出：false

提示：

1 <= s.length <= 20
s 仅含英文字母（大写和小写），数字（0-9），加号 '+' ，减号 '-' ，或者点 '.' 。
*/
package demo

import (
	"fmt"
	"testing"
)

func TestIsNum(t *testing.T) {
	Case("Test1", "100", true)
	Case("Test2", "123.45e+6", true)
	Case("Test3", "+500", true)
	Case("Test4", "5e2", true)
	Case("Test5", "3.1416", true)
	Case("Test6", "600.", true)
	Case("Test7", "-.123", true)
	Case("Test8", "-1E-16", true)
	Case("Test9", "1.79769313486232E+308", true)
	Case("Test10", "0.8", true)

	fmt.Println("---")

	Case("Test10", "12e", false)
	Case("Test11", "1a3.14", false)
	Case("Test12", "1+23", false)
	Case("Test13", "1.2.3", false)
	Case("Test14", "+-5", false)
	Case("Test15", "12e+5.4", false)
	Case("Test16", ".", false)
	Case("Test17", ".e1", false)
	Case("Test18", "e1", false)
	Case("Test19", "+.", false)
	Case("Test20", "", false)
}

func Case(name, in string, want bool) {
	if tmp := IsNum(in); tmp != want {
		panic(fmt.Sprintf("%s in: %s , get %v!=%v", name, in, tmp, want))
	}
}

// 数字格式可用A[.[B]][e|EC]或者.B[e|EC]表示, 其中A, C都是整数, B是无符号整数
func IsNum(s string) bool {
	if s == "" {
		return false
	}

	var idx int
	numeric := scanIntger(s, &idx)

	if idx < len(s) && s[idx] == '.' {
		idx += 1

		// 下面一行代码用||的原因：
		// 1. 小数可以没有整数部分，例如.123等于0.123；
		// 2. 小数点后面可以没有数字，例如233.等于233.0；
		// 3. 当然小数点前面和后面可以有数字，例如233.666
		numeric = scanUnsignedIntger(s, &idx) || numeric // 这里顺序不能错, 要先检查小数部分, 否则case "0.8"会返回false
	}

	if idx < len(s) && (s[idx] == 'e' || s[idx] == 'E') {
		idx += 1

		// 下面一行代码用&&的原因：
		// 1. 当e或E前面没有数字时，整个字符串不能表示数字，例如.e1、e1；
		// 2. 当e或E后面没有整数时，整个字符串不能表示数字，例如12e、12e+5.4
		numeric = numeric && scanIntger(s, &idx)
	}

	return numeric && idx == len(s)
}

func scanUnsignedIntger(s string, idx *int) bool {
	start := *idx

	for *idx != len(s) && s[*idx] >= '0' && s[*idx] <= '9' {
		*idx += 1
	}

	return *idx > start // 存在数字
}

// 整数的格式可以用[+|-]B表示, 其中B为无符号整数
func scanIntger(s string, idx *int) bool {
	if *idx == len(s) { // 想找整数却没有
		return false
	}

	if s[*idx] == '+' || s[*idx] == '-' {
		*idx += 1
	}
	return scanUnsignedIntger(s, idx)
}
