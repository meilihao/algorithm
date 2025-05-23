/*
LCR 166.中 珠宝的最高价值

现有一个记作二维矩阵 frame 的珠宝架，其中 frame[i][j] 为该位置珠宝的价值。拿取珠宝的规则为：

只能从架子的左上角开始拿珠宝
每次可以移动到右侧或下侧的相邻位置
到达珠宝架子的右下角时，停止拿取
注意：珠宝的价值都是大于 0 的。除非这个架子上没有任何珠宝，比如 frame = [[0]]。

示例 1：

输入：frame = [[1,3,1],[1,5,1],[4,2,1]]
输出：12
解释：路径 1→3→5→2→1 可以拿到最高价值的珠宝

提示：

0 < frame.length <= 200
0 < frame[0].length <= 200
*/
package demo

import (
	"fmt"
	"testing"
)

func TestJewelleryValue(t *testing.T) {
	frame := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}

	fmt.Println(jewelleryValue(frame) == 12)
}

func jewelleryValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i, g := range grid {
		for j, x := range g {
			if i > 0 {
				f[i][j] = max(f[i][j], f[i-1][j]) // 从上到cur
			}
			if j > 0 {
				f[i][j] = max(f[i][j], f[i][j-1]) // 从左到cur
			}
			f[i][j] += x
		}
	}
	return f[m-1][n-1]
}

// f(i,j) 只会从 f(i−1,j) 和 f(i,j−1) 转移而来，而与 f(i−2,⋯) 以及更早的状态无关，因此
// 同一时刻只需要存储最后两行的状态，即使用两个长度为 n 的一位数组代替 m×n 的二维数组 f，交替地进行状态转移，减少空间复杂度
func jewelleryValue2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, 2)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i, g := range grid {
		pos := i % 2
		for j, x := range g {
			f[pos][j] = 0
			if i > 0 {
				f[pos][j] = max(f[pos][j], f[1-pos][j])
			}
			if j > 0 {
				f[pos][j] = max(f[pos][j], f[pos][j-1])
			}
			f[pos][j] += x
		}
	}
	return f[(m-1)%2][n-1]
}
