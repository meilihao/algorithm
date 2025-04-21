package leetcode

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	s := "(10+(4+5+2)-3)+(6+8)"

	fmt.Println(calculate(s))
}

func calculate(s string) (ans int) {
	ops := []int{1}
	sign := 1
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}

		fmt.Println(i, sign, ops)
	}
	return
}
