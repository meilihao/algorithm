// 51. N皇后 : 摆皇后的位置，每行每列以及对角线只能出现 1 个皇后, 输出所有的情况
// 思路: 剪枝
// 题目要求，在以下地方都没有其他的Queen
// 1. Queen所在的行
// 2. Queen所在的列
// 3. Queen两条45度对角线上
package main

import "fmt"

func main() {
	n := 4

	res := totalNQueens(n)
	fmt.Println(res)
}

func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}

	cols := make([]bool, n) // 记录Queen所在的列

	d1 := make([]bool, 2*n) // 记录 '\' 方向的对角线的占用情况
	d2 := make([]bool, 2*n) // 记录 '/' 方向的对角线的占用情况

	res := 0

	dfs(0, n, cols, d1, d2, &res)

	return res
}

func dfs(r, n int, cols, d1, d2 []bool, res *int) {
	if r >= n { // r 最多 n-1
		*res++
		return
	}

	var id1, id2 int
	for c := 0; c < n; c++ { // 尝试放queen
		id1 = r - c + n
		id2 = 2*n - r - c - 1

		if cols[c] || d1[id1] || d2[id2] { // 排除列和对角线
			continue
		}

		// 标记占用
		cols[c], d1[id1], d2[id2] = true, true, true

		dfs(r+1, n, cols, d1, d2, res)

		// 解除标记
		cols[c], d1[id1], d2[id2] = false, false, false
	}
}
