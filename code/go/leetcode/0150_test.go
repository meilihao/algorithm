package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

// 逆波兰表达式是一种后缀表达式
func TestEvalRPN(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens))
}

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			default:
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}
