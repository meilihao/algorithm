/*
74.中 搜索二维矩阵

给你一个满足下述两条属性的 m x n 整数矩阵：

每行中的整数从左到右按非严格递增顺序排列。
每行的第一个整数大于前一行的最后一个整数。
给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。

示例 1：

输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出：true
示例 2：

输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
输出：false
*/
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
	c, cn := 0, len(matrix[0]) // 列
	r := len(matrix) - 1       // 行

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
