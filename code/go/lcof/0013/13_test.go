/*
013 机器人的运动范围

地上有一个m行n列的方格。一个机器人从坐标(0, 0)的格子开始移动，它
每一次可以向左、右、上、下移动一格，但不能进入行坐标和列坐标的数位之和
大于k的格子。例如，当k为18时，机器人能够进入方格(35, 37)，因为3+5+3+7=18。
但它不能进入方格(35, 38)，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
*/
package demo

import (
	"fmt"
	"testing"
)

func TestMovingCount(t *testing.T) {
	fmt.Println(movingCount(2, 3, 1) == 3)
	fmt.Println(movingCount(3, 1, 0) == 1)
	fmt.Println(movingCount(36, 11, 15) == 362)
}

func TestMovingCount2(t *testing.T) {
	fmt.Println(movingCount2(2, 3, 1) == 3)
	fmt.Println(movingCount2(3, 1, 0) == 1)
	fmt.Println(movingCount2(36, 11, 15) == 362)
}

// 地图初始化
func InitBoard(m, n int) [][]int {
	var board [][]int
	for i := 0; i < m; i++ {
		var nums []int
		for j := 0; j < n; j++ {
			nums = append(nums, 0)
		}
		board = append(board, nums)
	}

	return board
}

// 时间复杂度：O(m*n)O(m*n)
func movingCount(m int, n int, k int) int {
	board := InitBoard(m, n)
	// 走
	find(m, n, 0, 0, k, board)
	// 统计足迹
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != 0 { // 被走到过的格子
				res++
			}
		}
		fmt.Println(board[i])
	}
	return res
}

// 深度遍历
// 优化原因: 发现随着限制条件 k 的增大，(0, 0) 所在的蓝色方格区域(可达)内新加入的非障碍方格都可以由上方或左方的格子移动一步得到。
// 而其他不连通的蓝色方格区域会随着 cnt 的增大而连通，且连通的时候也是由上方或左方的格子移动一步得到，因此可以将搜索方向缩减为向右或向下
func find(m, n, i, j, k int, board [][]int) {
	if i >= m || i < 0 || j >= n || j < 0 {
		return
	}
	if isright(i, j, k) {
		board[i][j]++
		if board[i][j] == 1 { // `board[i][j] == 1`第一次到这个格子. 加了这个判断极大缩短了时间
			find(m, n, i, j+1, k, board)
			find(m, n, i+1, j, k, board)
		}
	}
}

// 位数之和是否满足条件
func isright(num1, num2, k int) bool {
	sum := 0
	for num1 != 0 {
		sum += num1 % 10
		num1 /= 10
	}
	for num2 != 0 {
		sum += num2 % 10
		num2 /= 10
	}
	return sum <= k
}

// 递推
func movingCount2(m int, n int, k int) int {
	set := make(map[[2]int]int)
	set[[2]int{0, 0}] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (set[[2]int{i - 1, j}] == 1 || set[[2]int{i, j - 1}] == 1) && isright(i, j, k) {
				set[[2]int{i, j}] = 1
			}
		}
	}
	return len(set)
}
