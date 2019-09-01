// 走M*N的格子, 左上角是入口, 右下角是出口, 只能Right/Down, 求走到出口的走法(不是格子数目)
// 思路: 动态规划
package main

import (
	"fmt"
)

func main() {
	fmt.Println(N3(5, 3))
}

func N(row, col int) int {
	if row == 0 && col == 0 {
		return 0
	}

	if row*col == 0 {
		return 1
	}

	dp := make([][]int, row+1)
	for i := range dp {
		dp[i] = make([]int, col+1)
	}

	for i := 1; i <= col; i++ {
		dp[0][i] = 1
	}

	for i := 1; i <= row; i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[row][col]
}

/*
 1, 1, 1,  1,  1,  1
 1, 2, 3,  4,  5,  6
 1, 3, 6,  10, 15, 21
 1, 4, 10, 20, 35, 56
 ......
 dp[i-1], dp[i]
 dp[...i-2]
*/
// cur空间 : O(col) => 最优O(min(col,row))
func N2(row, col int) int {
	if row == 0 && col == 0 {
		return 0
	}

	if row*col == 0 {
		return 1
	}

	dp := make([]int, col+1)
	for i := 0; i <= col; i++ {
		dp[i] = 1 // 此时dp[0]应该是1, 因为纵向开始迭代即row=1时, 起点是1
	}

	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			dp[j] = dp[j] + dp[j-1] // dp[0]始终是1, 不参与运算
		}
	}

	return dp[col]
}

// best
// 最优O(min(col,row))
// 二维dp -> 一维dp 的操作叫状态压缩
func N3(row, col int) int {
	if row == 0 && col == 0 {
		return 0
	}

	if row*col == 0 {
		return 1
	}

	m := min(row, col)

	dp := make([]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = 1
	}

	// 假设m是col
	if m != col { // 翻转格子
		row, col = col, row
	}

	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}

	return dp[m]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
