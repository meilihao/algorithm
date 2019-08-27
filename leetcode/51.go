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

	res0 := solveNQueens(n)
	fmt.Println(res0)

	// test generate_result
	res := [][]int{
		[]int{1, 3, 0, 2},
		[]int{2, 0, 3, 1},
	}
	fmt.Println(generate_result(res, n))
}

func solveNQueens(n int) [][]string {
	if n == 0 {
		return nil
	}

	cols := make([]bool, n) // 记录Queen所在的列

	// 斜线数量是2n-1, 因为斜线穿过横向n个格子,纵向n个格子, 但其中一个格子重合,因此是2n-1
	d1 := make([]bool, 2*n-1) // 记录 '\' 方向的对角线的占用情况
	d2 := make([]bool, 2*n-1) // 记录 '/' 方向的对角线的占用情况

	board := make([]int, n) // 保存结果

	res := [][]int{}

	dfs(0, n, cols, d1, d2, board, &res)

	return generate_result(res, n)
}

// r 表示行数的索引, 从0开始
// 对于所有的主对角线有 行号 + 列号 = 常数，对于所有的次对角线有 行号 - 列号 = 常数
// d1, d2 也可用map代替
func dfs(r, n int, cols, d1, d2 []bool, board []int, res *[][]int) {
	if r >= n { // r 最多 n-1
		// 使用copy, 因为其他的解会共用该board
		tmp := make([]int, len(board))
		copy(tmp, board)
		*res = append(*res, tmp)
		return
	}

	var id1, id2 int
	for c := 0; c < n; c++ { // 尝试放queen
		id1 = r - c + n       //  '\' 的方程式
		id2 = 2*n - r - c - 1 // '/' 的方程式

		fmt.Println(r, c, n, id1, id2)
		if cols[c] || d1[id1] || d2[id2] { // 排除列和对角线
			continue
		}

		// 放皇后
		board[r] = c

		// 标记占用
		cols[c], d1[id1], d2[id2] = true, true, true

		dfs(r+1, n, cols, d1, d2, board, res)

		// 解除标记
		cols[c], d1[id1], d2[id2] = false, false, false
	}
}

// 根据queen的位置索引生成字符串
func generate_result(input [][]int, n int) [][]string {
	result := make([][]string, len(input))
	buf := make([]byte, n)

	for i, v := range input {
		ss := make([]string, n)

		for j, vv := range v {
			for j := range buf { // 填充"."
				buf[j] = '.'
			}

			buf[vv] = 'Q' // 设置Queen

			ss[j] = string(buf)
		}

		result[i] = ss
	}

	return result
}
