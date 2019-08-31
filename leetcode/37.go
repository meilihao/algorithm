// 37.  解数独
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
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
