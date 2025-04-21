// 20. 有效的括号
package leetcode

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	s := "(){}[]"

	fmt.Println(isValid(s))
}

// 使用stack : 时间O(n) + 空间O(n)
func isValid(s string) bool {
	//空串认为是合法的
	if s == "" {
		return true
	}

	//如果s长度为奇数是不可能合法的
	if len(s)%2 == 1 {
		return false
	}

	brackets := map[byte]byte{')': '(', ']': '[', '}': '{'}
	var stack []byte

	for _, c := range []byte(s) {
		if c == '(' || c == '{' || c == '[' {
			// 入栈
			stack = append(stack, c)
		} else if len(stack) > 0 && brackets[c] == stack[len(stack)-1] { // 栈中有数据，且此元素与栈顶元素相同
			stack = stack[:len(stack)-1] // pop
		} else {
			return false
		}
	}

	// 循环结束，栈中还有数据则 false
	return len(stack) == 0
}
