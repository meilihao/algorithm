package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	testCases := []struct {
		in       string
		expected int
	}{
		{"))()())", 4},
		{"(()", 2},
		{"()((())", 4},
		{"(())", 4},
		{"()()", 4},
		{"()()))", 4},
	}

	for i, v := range testCases {
		tmp := longestValidParenthesesWithValue(v.in)
		fmt.Println("start:", i, v.in, tmp)
		if tmp != v.expected {
			t.Errorf("error: %s, %d != %d", v.in, tmp, v.expected)
		}
		fmt.Println("end:", i, v.in)
	}
}

// 栈顶是匹配序列前的")"
func longestValidParenthesesWithValue(s string) int {
	type item struct {
		Val byte
		Idx int
	}

	maxAns := 0

	stack := []*item{&item{Val: ')', Idx: -1}} // 避免"()()"
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, &item{Val: s[i], Idx: i})
		} else { // get ')'
			if stack[len(stack)-1].Val == ')' { // 空或没法匹配
				stack = append(stack, &item{Val: s[i], Idx: i})
				continue
			}

			// fmt.Println("stack:")
			// for i := len(stack) - 1; i >= 0; i-- {
			// 	fmt.Println(stack[i].Val)
			// }

			// match
			stack = stack[:len(stack)-1]
			maxAns = max(maxAns, i-stack[len(stack)-1].Idx)
		}
	}

	return maxAns
}

// longestValidParenthesesWithValue的简化版
func longestValidParentheses(s string) int {
	maxAns := 0

	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i) // 压入'('下标
		} else { // get ')'
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}
	return maxAns
}

// 每当 left 计数器与 right 计数器相等时，计算当前有效字符串的长度，并且记录目前为止找到的最长子字符串
// 反向计算的原因: 避免`(()`即遍历的时候左括号的数量始终大于右括号的数量, 此时最长有效括号是求不出来的
func longestValidParentheses2(s string) int {
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}
