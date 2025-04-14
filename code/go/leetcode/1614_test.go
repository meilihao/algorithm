package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
}

func maxDepth(seq string) int {
	cur := 0 // 当前深度
	max := 0 // 最大深度

	for i := 0; i < len(seq); i++ {
		if seq[i] == '(' {
			cur++

			if cur > max {
				max = cur
			}
		} else if seq[i] == ')' {
			cur--
		}
	}

	return max
}
