/*
37.困 解数独

编写一个程序，通过填充空格来解决数独问题。

数独的解法需 遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
数独部分空格内已填入了数字，空白格用 '.' 表示。

示例 1：

输入：board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
输出：[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
解释：输入的数独如上图所示，唯一有效的解决方案如下所示：

提示：

board.length == 9
board[i].length == 9
board[i][j] 是一位数字或者 '.'
题目数据 保证 输入数独仅有一个解
*/
package main

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	s := []string{
		`[
		["5","3",".",".","7",".",".",".","."],
		["6",".",".","1","9","5",".",".","."],
		[".","9","8",".",".",".",".","6","."],
		["8",".",".",".","6",".",".",".","3"],
		["4",".",".","8",".","3",".",".","1"],
		["7",".",".",".","2",".",".",".","6"],
		[".","6",".",".",".",".","2","8","."],
		[".",".",".","4","1","9",".",".","5"],
		[".",".",".",".","8",".",".","7","9"]
	  ]`,
	}

	for _, v := range s {
		board := parse(v)

		//fmt.Println(a)

		solveSudoku(board)

		printBoad(board)
	}
}

func printBoad(board [][]byte) {
	for _, v := range board {
		for i, vv := range v {
			fmt.Printf("%s ", string(vv))
			if i == 8 {
				fmt.Println()
			}
		}
	}
}

// https://regex101.com/ 通过正则提取
func parse(input string) [][]byte {
	a := make([][]byte, 0)

	re := regexp.MustCompile(`(?m)\[(.*)\]`)

	for i, match := range re.FindAllStringSubmatch(input, -1) {
		if len(match) != 2 {
			panic(fmt.Sprintf("invalid input(%d) %v", i, match))
		}

		a = append(a, parseArray(match[1]))
	}

	if len(a) != 9 {
		panic(fmt.Sprintf("invalid sudo: %v", a))
	}

	return a
}

func parseArray(s string) []byte {
	ss := strings.Split(s, ",")
	if len(ss) != 9 {
		panic(fmt.Sprintf("invalid row: %s", s))
	}

	a := make([]byte, 9)
	for i, v := range ss {
		if v == `"."` {
			a[i] = '.'
		} else {
			a[i] = getNumber(v)
		}
	}

	return a
}

func getNumber(s string) byte {
	ss := []byte(s)

	if len(ss) != 3 {
		panic(fmt.Sprintf("invalid raw number: %s", s))
	}

	return ss[1]
}

func solveSudoku(board [][]byte) {
	// 已使用的数字们
	rows := make([][9]bool, 9) // 也可使用map代替
	cols := make([][9]bool, 9)
	boxs := make([][9]bool, 9) // 3 x 3 子数独

	for i := 0; i < 9; i++ {
		rows[i] = [9]bool{}
		cols[i] = [9]bool{}
		boxs[i] = [9]bool{}
	}

	var num, boxIndex int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num = int(board[i][j] - '1')
				boxIndex = (i/3)*3 + j/3

				rows[i][num] = true
				cols[j][num] = true
				boxs[boxIndex][num] = true
			}
		}
	}

	// 递归尝试填充数组
	recusiveSolveSudoku(board, rows, cols, boxs, 0, 0)
}

func recusiveSolveSudoku(board [][]byte, rows, cols, boxs [][9]bool, row, col int) bool {
	if col == 9 {
		col = 0
		row++
	}

	// 边界校验, 如果已经填充完成, 返回true, 表示一切结束
	if row == 9 {
		return true
	}

	// 是空则尝试填充, 否则跳过继续尝试填充下一个位置
	if board[row][col] == '.' {
		boxIndex := (row/3)*3 + col/3
		var num byte

		for num = 0; num <= 8; num++ { // todo 仅尝试可以遍历的空间
			if rows[row][num] || cols[col][num] || boxs[boxIndex][num] {
				continue
			}

			board[row][col] = '1' + num
			rows[row][num] = true
			cols[col][num] = true
			boxs[boxIndex][num] = true

			// board[row][col]填好后, 填充下一个
			if recusiveSolveSudoku(board, rows, cols, boxs, row, col+1) {
				return true
			}

			// 恢复现场
			board[row][col] = '.'
			rows[row][num] = false
			cols[col][num] = false
			boxs[boxIndex][num] = false
		}

	} else {
		return recusiveSolveSudoku(board, rows, cols, boxs, row, col+1)
	}

	return false
}
