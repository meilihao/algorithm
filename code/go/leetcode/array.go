package leetcode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// 解析二维数组
// arraySplit是数组分隔符,f是数组元素的处理函数
// rows,cols为0时不检查
// https://regex101.com/ 通过正则提取
func String2TwoDimensionalByteArray(input, arraySplit string, f func(string) string, rows, cols int) [][]string {
	a := make([][]string, 0)

	re := regexp.MustCompile(`(?m)\[(.*)\]`)

	for i, match := range re.FindAllStringSubmatch(input, -1) {
		//fmt.Println(match)
		if len(match) != 2 {
			panic(fmt.Sprintf("invalid input(%d) %v", i, match))
		}

		a = append(a, parseArray(match[1], arraySplit, f, cols))
	}

	if rows > 0 && len(a) != rows {
		panic(fmt.Sprintf("invalid sudo: %v", a))
	}

	return a
}

func parseArray(s, arraySplit string, f func(string) string, cols int) []string {
	ss := strings.Split(s, arraySplit)
	if cols > 0 && len(ss) != cols {
		panic(fmt.Sprintf("invalid row: %s", s))
	}

	a := make([]string, len(ss))
	for i, v := range ss {
		a[i] = f(v)
	}

	return a
}

func String2Byte(s string) byte {
	ss := []byte(s)

	if len(ss) != 3 {
		panic(fmt.Sprintf("invalid raw number: %s", s))
	}

	return ss[1]
}

func PrintTwoDimensionalByteArray(board [][]string) {
	var l int
	for _, v := range board {
		l = len(v)

		for i, vv := range v {
			if i == l-1 {
				fmt.Printf("%s\n", vv)
			} else {
				fmt.Printf("%s ", vv)
			}
		}
	}
}

func TwoDimensionalByteArray2Int(board [][]string) [][]int {
	a := make([][]int, 0, len(board))

	for _, v := range board {
		tmp := make([]int, 0, len(v))

		for _, vv := range v {
			n, err := strconv.Atoi(vv)
			if err != nil {
				panic(err)
			}

			tmp = append(tmp, n)
		}

		a = append(a, tmp)
	}

	return a
}
