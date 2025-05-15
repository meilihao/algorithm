/*
04 二维数组中的查找

在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
*/

package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	board  [][]int
	target int
}

type ans struct {
	find bool
}

type question struct {
	p para
	a ans
}

func TestFind(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				board: [][]int{
					[]int{1, 2, 8, 9},
					[]int{2, 4, 9, 12},
					[]int{4, 7, 10, 13},
					[]int{6, 8, 11, 15},
				},
				target: 7,
			},
			a: ans{
				find: true,
			},
		},
		question{
			p: para{
				board: [][]int{
					[]int{1, 2, 8, 9},
					[]int{2, 4, 9, 12},
					[]int{4, 6, 10, 13},
					[]int{6, 8, 11, 15},
				},
				target: 7,
			},
			a: ans{
				find: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.find, Find(p.board, p.target), "输入:%v", p)
	}
}

// 不能从左上角或右下角开始, 因为每次移动方向有两个可选择, 无法抉择. 因此要从左下或右上
// 这里选择左下
func Find(board [][]int, target int) bool {
	r := len(board) - 1
	c, cm := 0, len(board[0])

	for r >= 0 && c < cm {
		if board[r][c] > target {
			r--
		} else if board[r][c] < target {
			c++
		} else {
			return true
		}
	}

	return false

}
