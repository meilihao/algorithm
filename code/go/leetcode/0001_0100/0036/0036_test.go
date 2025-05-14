/*
36.中 有效的数独

请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）

注意：

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
空白格用 '.' 表示。

示例 1：

输入：board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
输出：true
示例 2：

输入：board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
输出：false
解释：除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。 但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。

提示：

board.length == 9
board[i].length == 9
board[i][j] 是一位数字（1-9）或者 '.'
*/
package demo

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestIsValidSudoku(t *testing.T) {
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
		`[
		["8","3",".",".","7",".",".",".","."],
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
		a := parse(v)

		//fmt.Println(a)

		fmt.Println(isValidSudoku(a))
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

func isValidSudoku(board [][]byte) bool {
	if len(board) != 9 {
		return false
	}

	rows := make([][9]bool, 9) // 横向检查 也可使用map代替
	cols := make([][9]bool, 9) // 纵向检查
	boxs := make([][9]bool, 9) // 3 x 3 子数独检查

	for i := 0; i < 9; i++ {
		rows[i] = [9]bool{}
		cols[i] = [9]bool{}
		boxs[i] = [9]bool{}
	}

	var num, boxIndex int

	for i := 0; i < 9; i++ {
		if len(board[i]) != 9 {
			return false
		}

		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			num = int(board[i][j] - '1')
			boxIndex = (i/3)*3 + j/3 // boxIndex = (row / 3) * 3 + columns / 3,  将`3 x 3`子数独看成一个元素,整个数独就是一个3x3的新数独 => index = newIndex * 3 + newColIndex

			//fmt.Println(i, j, num, rows[i][num], cols[j][num], boxs[boxIndex][num])
			if rows[i][num] || cols[j][num] || boxs[boxIndex][num] { // 判断数字是否重复出现
				return false
			}

			rows[i][num] = true
			cols[j][num] = true
			boxs[boxIndex][num] = true
		}
	}

	return true
}
