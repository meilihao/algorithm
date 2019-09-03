// 200. 岛屿数量
// 给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
// 思路:
// 1. 染色: bfs/dfs, dfs简单. 时间复杂度 : O(M*N)
// 1. 并查集 : 从并查集的功能入手， 显然题目要求有 “查询”， “合并” 操作时可以考虑使用
package main

import "fmt"

func main() {
	// as := [][]byte{
	// 	[]byte{'1', '1', '1', '1', '0'},
	// 	[]byte{'1', '1', '0', '1', '0'},
	// 	[]byte{'1', '1', '0', '0', '0'},
	// 	[]byte{'0', '0', '0', '0', '0'},
	// }
	// as := [][]byte{
	// 	[]byte{'1', '1', '0', '0', '0'},
	// 	[]byte{'1', '1', '0', '0', '0'},
	// 	[]byte{'0', '0', '1', '0', '0'},
	// 	[]byte{'0', '0', '0', '1', '1'},
	// }
	as := [][]byte{
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '0', '1', '0', '1'},
		[]byte{'1', '1', '1', '0', '1'},
	}
	// as3 := [][]byte{
	// 	[]byte{'1', '1', '1'},
	// 	[]byte{'0', '1', '0'},
	// 	[]byte{'1', '1', '1'}, // as3[2][0]必须合并后面, 此时发现UnionFind要先扫一遍才行
	// }

	fmt.Println(numIslands(as))
}

// dfs, 简洁
// 染色
// 遍历节点
// 	if node =='1'{
// 		count++

// 		递归染色相连的节点成'0' // 染色也可用map代替, 避免修改grid
// 	}
// best
func numIslands2(grid [][]byte) int {
	var count int
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				count++

				dfs(grid, i, j, row, col)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, i, j, row, col int) {
	if i < 0 || i == row || j < 0 || j == col || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'

	dfs(grid, i-1, j, row, col) // 上面
	dfs(grid, i, j-1, row, col) // 前面
	dfs(grid, i+1, j, row, col) // 下
	dfs(grid, i, j+1, row, col) // 后
}

// 时间/空间复杂度：O(M*N)
func numIslands(grid [][]byte) int {
	var count int
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])

	n := row * col
	UnionFind := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				UnionFind[i*col+j] = i*col + j
				count++
			} else {
				UnionFind[i*col+j] = -1 // 因为grid[0][0]可能岛屿,此时UnionFind的默认值也是0, 因此需要区分0
			}
		}
	}

	fmt.Println(UnionFind, count)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				// union 与 union2 雷同
				union(UnionFind, rank, grid, i, j, i-1, j, &count) // 上面
				union(UnionFind, rank, grid, i, j, i, j-1, &count) // 前面
				union(UnionFind, rank, grid, i, j, i+1, j, &count) // 下
				union(UnionFind, rank, grid, i, j, i, j+1, &count) // 后
			}
		}
	}

	//fmt.Println(UnionFind)
	//fmt.Println(rank)

	return count
}

func findRoot(unionFind []int, i int) int {
	// fmt.Println(unionFind, i)
	for unionFind[i] != i { // 不是根节点
		i = unionFind[i]
	}

	return i
}

func union2(unionFind, rank []int, grid [][]byte, i, j, a, b int, count *int) {
	row := len(grid)
	col := len(grid[0])

	var rootA, rootB int
	ia, ib := i*col+j, a*col+b

	if a >= 0 && a < row && b >= 0 && b < col && grid[a][b] == '1' {
		rootA = findRoot(unionFind, ia)
		rootB = findRoot(unionFind, ib)

		if rootA != rootB { // 合并根, 否则排除
			unionFind[rootB] = rootA
			*count--
		}

		fmt.Println(i, j, a, b, ia, ib, rootA, rootB, rootA != rootB, unionFind)
	}
}

func union(unionFind, rank []int, grid [][]byte, i, j, a, b int, count *int) {
	row := len(grid)
	col := len(grid[0])

	ia, ib := i*col+j, a*col+b

	if a >= 0 && a < row && b >= 0 && b < col && grid[a][b] == '1' {
		rootA := findRoot(unionFind, ia)
		rootB := findRoot(unionFind, ib)

		// 新节点对应的rank总是为0
		// 合并根
		if rootA != rootB { // 已合并的话, rootA==rootB
			if rank[rootA] > rank[rootB] {
				unionFind[rootB] = rootA
			} else if rank[rootA] < rank[rootB] {
				unionFind[rootA] = rootB
			} else {
				unionFind[rootB] = rootA
				rank[rootA] += 1 //标记是否为根
			}

			*count--
		}

		fmt.Println(i, j, a, b, ia, ib, rootA, rootB, rootA != rootB, rank[rootA], rank[rootB], unionFind)
	}
}
