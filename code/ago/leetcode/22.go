// 22. 括号生成
// 思路: 回溯法
package main

import "fmt"

func main() {
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
