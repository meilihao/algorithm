/*
85.困 最大矩形

给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

示例 1：

输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：6
解释：最大矩形如上图所示。
示例 2：

输入：matrix = [["0"]]
输出：0
示例 3：

输入：matrix = [["1"]]
输出：1

提示：

rows == matrix.length
cols == matrix[0].length
1 <= row, cols <= 200
matrix[i][j] 为 '0' 或 '1'
*/
package leetcode

import (
	"testing"
)

func TestMovingAverage(t *testing.T) {

}

type MovingAverage struct {
	size, sum int
	q         []int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{size: size}
}

func (m *MovingAverage) Next(val int) float64 {
	if len(m.q) == m.size {
		m.sum -= m.q[0]
		m.q = m.q[1:]
	}
	m.sum += val
	m.q = append(m.q, val)
	return float64(m.sum) / float64(len(m.q))
}
