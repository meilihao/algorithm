/*
54.中 螺旋矩阵

给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/

package demo

import (
	"testing"
)

func TestSpiralOrder(t *testing.T) {

}

/*
			 t
	   ---------

l |       | r

	   ---------
			   b
*/
// 将矩阵看成若干层，首先输出最外层的元素，其次输出次外层的元素，直到输出最内层的元素
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, columns            = len(matrix), len(matrix[0])
		order                    = make([]int, rows*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)

	for left <= right && top <= bottom {
		for column := left; column <= right; column++ { // [左->右]
			order[index] = matrix[top][column]
			index++
		}
		for row := top + 1; row <= bottom; row++ { // (上->下]
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- { // (右->左)
				order[index] = matrix[bottom][column]
				index++
			}
			for row := bottom; row > top; row-- { // [左->上)
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}
