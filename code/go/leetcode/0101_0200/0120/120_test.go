/*
120.中 三角形最小路径和

给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。

示例 1：

输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：

	  2
	 3 4
	6 5 7

4 1 8 3
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
示例 2：

输入：triangle = [[-10]]
输出：-10

提示：

1 <= triangle.length <= 200
triangle[0].length == 1
triangle[i].length == triangle[i - 1].length + 1
-104 <= triangle[i][j] <= 104
*/
package main

import (
	"fmt"
	"strconv"
	"testing"

	. "al/leetcode"
)

// 思路:
// 1. 递归, 0(2^n)
// 2. dp,0(m*n) : dp[i,j]=min(dp[i+1,j],dp[i+1,j+1]) + triangle[i,j]
func TestMinimumTotal(t *testing.T) {
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

	triangleTmp := String2TwoDimensionalByteArray(triangleStr, ",", f, 0, 0)
	PrintTwoDimensionalByteArray(triangleTmp)
	triangle := TwoDimensionalByteArray2Int(triangleTmp)
	//fmt.Println(triangle)

	//fmt.Println(minimumTotal(triangle))
	fmt.Println(minimumTotalR(triangle))
}

// **对其子数组再看比较直观**
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

	// dp[j] 表示：从三角形最底层 triangle[n-1] 的 j 列开始，到达 triangle[n-1] 的 j 列的最小路径和
	dp := triangle[n-1] // 取最下面一行作为起始值
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			/*
							dp[j] 的更新逻辑：
							当前考虑的节点是 triangle[i][j]。
							从这个节点到三角形底部，有两条可能的路径：
							1. 移动到下一行的 dp[j] (对应 triangle[i+1][j])
							2. 移动到下一行的 dp[j+1] (对应 triangle[i+1][j+1])

				            dp[j] 和 dp[j+1] 在当前循环迭代之前，已经存储了从下一行对应位置到最底部的最小路径和。
				            例如，当 i 是 n-2 时：
				            dp[j] 存储的是 triangle[n-1][j] 的值。
				            dp[j+1] 存储的是 triangle[n-1][j+1] 的值。

				            所以，min(dp[j], dp[j+1]) 找到了从下一行两个相邻点中选择一个，并到最底部的最小和。
				            再加上当前节点 triangle[i][j] 的值，就是从 triangle[i][j] 到最底部的最小路径和。
				            dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
			*/
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j] // triangle[i][j] + min(triangle[i+1][j], triangle[i+1][j+1])
		}

		fmt.Println(dp)
	}

	return dp[0]
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
