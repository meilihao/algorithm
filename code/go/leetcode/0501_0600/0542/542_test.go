/*
542.中 01 矩阵

给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。

两个相邻元素间的距离为 1 。

示例 1：

输入：mat = [[0,0,0],[0,1,0],[0,0,0]]
输出：[[0,0,0],[0,1,0],[0,0,0]]
示例 2：

输入：mat = [[0,0,0],[0,1,0],[1,1,1]]
输出：[[0,0,0],[0,1,0],[1,2,1]]

提示：

m == mat.length
n == mat[i].length
1 <= m, n <= 104
1 <= m * n <= 104
mat[i][j] is either 0 or 1.
mat 中至少有一个 0
*/
package leetcode

import (
	"testing"
)

func TestUpdateMatrix(t *testing.T) {

}

// 修改了原图
func updateMatrix(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	queue := make([][]int, 0)
	for i := 0; i < n; i++ { // 把0全部存进队列，后面从队列中取出来，判断每个访问过的节点的上下左右，直到所有的节点都被访问过为止。
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				point := []int{i, j}

				// 将所有 0 作为起点，并同时向外扩展，确保了当任何一个 1 的单元格第一次被访问时，所计算的距离就是其到最近的 0 的最短距离
				queue = append(queue, point) // 所有值为 0 的单元格加入队列, 这些是 BFS 的起始点
			} else {
				matrix[i][j] = -1 // 表示待访问即距离是未知的
			}
		}
	}
	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // 定义四个方向：右、左、下、上

	for len(queue) > 0 { // 这里就是 BFS 模板操作了
		point := queue[0]
		queue = queue[1:]

		// // 遍历当前单元格 (point) 的所有四个邻居
		for _, v := range direction {
			x := point[0] + v[0] // 邻居的行坐标
			y := point[1] + v[1] // 邻居的列坐标
			if x >= 0 && x < n && y >= 0 && y < m && matrix[x][y] == -1 {
				matrix[x][y] = matrix[point[0]][point[1]] + 1
				queue = append(queue, []int{x, y}) //  将邻居加入队列，以便继续扩展
			}
		}
	}

	return matrix
}
