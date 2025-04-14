/*
31. 栈的压入、弹出序列
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。
*/
package lcof

import (
	"fmt"
	"testing"
)

func TestValidateStackSequences(t *testing.T) {
	fmt.Println(ValidateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}

/*
栈模拟
*/
func ValidateStackSequences(pushed []int, popped []int) bool {
	stk := []int{}
	j := 0
	for _, v := range pushed {
		stk = append(stk, v)
		for len(stk) > 0 && stk[len(stk)-1] == popped[j] { // 模拟弹出
			stk = stk[:len(stk)-1]
			j++
		}
	}
	return j == len(pushed) // 检查压入次数是否等于预期的len(pushed)次
}
