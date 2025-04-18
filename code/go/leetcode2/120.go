// 120. 三角形最小路径和
// 思路:
// 1. 递归, 0(2^n)
// 2. dp,0(m*n) : dp[i,j]=min(dp[i+1,j],dp[i+1,j+1]) + triangle[i,j]
package main

import (
	"fmt"
	"strconv"

	"ago/helper"
)

func main() {
	triangleStr := `
	[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]`

	f := func(s string) string {
		return s
	}

	triangleTmp := helper.String2TwoDimensionalByteArray(triangleStr, ",", f, 0, 0)
	helper.PrintTwoDimensionalByteArray(triangleTmp)
	triangle := helper.TwoDimensionalByteArray2Int(triangleTmp)
	//fmt.Println(triangle)

	//fmt.Println(minimumTotal(triangle))
	fmt.Println(minimumTotalR(triangle))
}

// 原始递推方程(从下往上): dp[i][j] = min(dp[i+1,j],dp[i+1,j+1]) + triangle[i][j]
// 选择从下往上的原因: 从下往上是分散的; 而从下往上是合并的过程, 符合递推模式
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return triangle[0][0]
	}

	dp := triangle[n-1] // 取最下面一行作为起始值
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}

		fmt.Println(dp)
	}

	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// // 递归
//
//	r(i,j int){
//		r(i+1,j)
//		r(i+1,j+1)
//	}
func minimumTotalR(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}

	min := 1 << 20
	_dfs(triangle, 0, 0, 0, &min, "")

	return min
}

func _dfs(triangle [][]int, i, j, sum int, min *int, path string) {
	if i == len(triangle)-1 { // 到倒数第二层完成即结束了
		sum += triangle[i][j]
		path += strconv.Itoa(triangle[i][j])

		fmt.Println(path, "#", sum)
		if sum < *min {
			*min = sum
		}

		return
	}

	sum += triangle[i][j]
	path += strconv.Itoa(triangle[i][j]) + "->"

	_dfs(triangle, i+1, j, sum, min, path)   //向下结合
	_dfs(triangle, i+1, j+1, sum, min, path) // 向右下结合
}
