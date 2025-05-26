/*
304.中 二维区域和检索 - 矩阵不可变

给定一个二维矩阵 matrix，以下类型的多个请求：

计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
实现 NumMatrix 类：

NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。

示例 1：

输入:
["NumMatrix","sumRegion","sumRegion","sumRegion"]
[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]
输出:
[null, 8, 11, 12]

解释:
NumMatrix numMatrix = new NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]);
numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素总和)
numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素总和)
numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素总和)

提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 200
-105 <= matrix[i][j] <= 105
0 <= row1 <= row2 < m
0 <= col1 <= col2 < n
最多调用 104 次 sumRegion 方法
*/
package leetcode

import (
	"testing"
)

func TestNumMatrix(t *testing.T) {

}

type NumMatrix struct {
	sums [][]int
}

// 将前缀和数组 sums 的长度设为 n+1 的目的是为了方便计算 不需要对 i=0 的情况特殊处理
// f(i,j)=原始矩阵中从 (0, 0) 到 (i, j) 形成的矩形区域的所有元素之和
// f(i−1,j) = matrix[0...i-1][0...j] 的和（即当前元素正上方矩形的和）
// f(i,j−1) = matrix[0...i][0...j-1] 的和（即当前元素正左方矩形的和）
// f(i−1,j−1) = matrix[0...i-1][0...j-1] 的和（即左上角重叠部分的和）
// f(i,j)=f(i−1,j)+f(i,j−1)−f(i−1,j−1)+matrix[i][j]
// NumMatrix.sums 矩阵的大小会比原矩阵大一圈，因为用 sums[i+1][j+1] 来对应 matrix[i][j]，这样 sums[0][...] 和 sums[...][0] 就可以表示边界情况（即和为 0）. 否则f(0,0)=f(−1,0)+f(0,−1)−f(−1,−1)+matrix[0][0], -1不是有效坐标
func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix) // 行
	if m == 0 {
		return NumMatrix{}
	}

	n := len(matrix[0]) // 列
	sums := make([][]int, m+1)
	sums[0] = make([]int, n+1)
	for i, row := range matrix {
		sums[i+1] = make([]int, n+1)
		for j, v := range row {
			sums[i+1][j+1] = sums[i+1][j] + sums[i][j+1] - sums[i][j] + v
		}
	}
	return NumMatrix{sums}
}

func (nm *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	// 画图比较清楚
	return nm.sums[row2+1][col2+1] - nm.sums[row1][col2+1] - nm.sums[row2+1][col1] + nm.sums[row1][col1]
}
