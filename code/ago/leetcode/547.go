// 547. 朋友圈
// 思路: 并查集, 与200有区别, 不能直接用dfs求解
package main

import (
	"fmt"

	"ago/helper"
)

func main() {
	// as := [][]int{
	// 	[]int{1, 0, 0, 1},
	// 	[]int{0, 1, 1, 0},
	// 	[]int{0, 1, 1, 1},
	// 	[]int{1, 0, 1, 1},
	// }
	asStr := `
	[
		[1,1,0,0,0,0,0,1,0,0,0,0,0,0,0],
		[1,1,0,0,0,0,0,0,0,0,0,0,0,0,0],
		[0,0,1,0,0,0,0,0,0,0,0,0,0,0,0],
		[0,0,0,1,0,1,1,0,0,0,0,0,0,0,0],
		[0,0,0,0,1,0,0,0,0,1,1,0,0,0,0],
		[0,0,0,1,0,1,0,0,0,0,1,0,0,0,0],
		[0,0,0,1,0,0,1,0,1,0,0,0,0,1,0],
		[1,0,0,0,0,0,0,1,1,0,0,0,0,0,0],
		[0,0,0,0,0,0,1,1,1,0,0,0,0,1,0],
		[0,0,0,0,1,0,0,0,0,1,0,1,0,0,1],
		[0,0,0,0,1,1,0,0,0,0,1,1,0,0,0],
		[0,0,0,0,0,0,0,0,0,1,1,1,0,0,0],
		[0,0,0,0,0,0,0,0,0,0,0,0,1,0,0],
		[0,0,0,0,0,0,1,0,1,0,0,0,0,1,0],
		[0,0,0,0,0,0,0,0,0,1,0,0,0,0,1]
		]
`
	f := func(s string) string {
		return s
	}

	tmp := helper.String2TwoDimensionalByteArray(asStr, ",", f, 0, 0)
	//helper.PrintTwoDimensionalByteArray(tmp)
	as := helper.TwoDimensionalByteArray2Int(tmp)

	fmt.Println(findCircleNum2(as))
}

// 并查集
func findCircleNum(M [][]int) int {
	n := len(M)
	if n == 0 {
		return 0
	}

	res := n
	friend := make([]int, res) // n个人, 最多n个朋友圈
	for i := 0; i < res; i++ { // 每个人的root是自己
		friend[i] = i
	}

	fmt.Println(friend)

	// s 和 d 是朋友关系
	// 所以，s 的所有朋友都是 d 的朋友
	union := func(s, d int) {
		for i := range friend {
			if friend[i] == s { // s <- d, 通过覆盖, 是朋友关系的最终会变成一个值, 更直观
				friend[i] = d
			}
		}

		// fmt.Println(s, d, friend)
	}

	// 图形关于'\'对称, 仅遍历一半即可
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if M[i][j] == 1 {
				if friend[i] != friend[j] {
					res--
					union(friend[i], friend[j])
				}
			}
		}
	}

	fmt.Println(friend)

	return res
}

func findCircleNum2(M [][]int) int {
	n := len(M)
	if n == 0 {
		return 0
	}

	res := n
	friend := make([]int, res) // n个人, 最多n个朋友圈
	for i := 0; i < res; i++ { // 每个人的root是自己
		friend[i] = i
	}

	fmt.Println(friend)

	find := func(i int) int {
		for friend[i] != i {
			i = friend[i]
		}

		return i
	}

	// 图形关于'\'对称, 仅遍历一半即可
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if M[i][j] == 1 {
				xi, xj := find(i), find(j)
				if xi != xj { // 算的是合并次数 : group = n - 合并次数
					//如果不属于同个朋友圈的话就把i归为j的组
					friend[xj] = xi
					res--
				}
				fmt.Println(i, j, xi, xj, xi != xj, friend)
			}
		}
	}

	fmt.Println(friend)

	return res
}
