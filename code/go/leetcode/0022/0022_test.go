/*
22.中 括号生成

数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]
*/
package main

import (
	"fmt"
	"testing"
)

// 思路: 回溯法
func TestGenerateParenthesis(t *testing.T) {
	n := 3

	fmt.Println(generateParenthesis(n))
}

func generateParenthesis(n int) []string {
	list := []string{}

	gen(0, 0, n, "", &list)

	return list
}

// 只有在知道序列仍然保持有效时才添加 '(' or ')', 我们可以通过跟踪到目前为止放置的左括号和右括号的数目来做到这一点.
// left : 已用左括号个数
// right : 已用右括号个数
// result : 原始字符串
func gen(left, right, n int, result string, list *[]string) {
	if left == n && right == n {
		*list = append(*list, result)
	}

	if left < n { // 加左括号的条件
		gen(left+1, right, n, result+"(", list)
	}

	if left > right && right < n { // 加右括号的条件
		gen(left, right+1, n, result+")", list)
	}
}
