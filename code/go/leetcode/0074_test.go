package leetcode

import (
	"fmt"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	fmt.Println(searchMatrix(nil, 10))
}

// 矩阵是有序的，从左下角来看，向上数字递减，向右数字递增，
// 因此从左下角开始查找，当要查找数字比左下角数字大时。右移
// 要查找数字比左下角数字小时，上移。这样找的速度最快
func searchMatrix(matrix [][]int, target int) bool {
	c, cn := 0, len(matrix[0]) // 横
	r := len(matrix) - 1       // 竖

	for r >= 0 && c < cn {
		if matrix[r][c] > target {
			r--
		} else if matrix[r][c] < target {
			c++
		} else {
			return true
		}
	}

	return false
}
