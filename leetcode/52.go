// 51. N皇后 : 摆皇后的位置，每行每列以及对角线只能出现 1 个皇后, 输出所有的情况
// 思路: 剪枝
// 题目要求，在以下地方都没有其他的Queen
// 1. Queen所在的行
// 2. Queen所在的列
// 3. Queen两条45度对角线上
package main

import "fmt"

func main() {
	n := 4

	res := totalNQueens2(n)
	fmt.Println(res)
}

func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}

	cols := make([]bool, n) // 记录Queen所在的列

	d1 := make([]bool, 2*n) // 记录 '\' 方向的对角线的占用情况
	d2 := make([]bool, 2*n) // 记录 '/' 方向的对角线的占用情况

	res := 0

	dfs(0, n, cols, d1, d2, &res)

	return res
}

func dfs(r, n int, cols, d1, d2 []bool, res *int) {
	if r >= n { // r 最多 n-1
		*res++
		return
	}

	var id1, id2 int
	for c := 0; c < n; c++ { // 尝试放queen
		id1 = r - c + n
		id2 = 2*n - r - c - 1

		if cols[c] || d1[id1] || d2[id2] { // 排除列和对角线
			continue
		}

		// 标记占用
		cols[c], d1[id1], d2[id2] = true, true, true

		dfs(r+1, n, cols, d1, d2, res)

		// 解除标记
		cols[c], d1[id1], d2[id2] = false, false, false
	}
}

// best
func totalNQueens2(n int) int {
	if n < 1 {
		return 0
	}

	// 可参考![](misc/img/n-queen.png) from [位运算解决八皇后问题](https://blog.csdn.net/kai_wei_zhang/article/details/8033194)中的两张图片理解pie,na
	var cols int // 这个好理解, 哪些列被皇后占用
	// 把 pie 和 na 理解为: 当前queen导致其所在的撇或捺(即相应行的哪些列)已被占用, 即在两个对角线方向的限制条件下这一行的哪些地方不能放
	var pie int // 记录 '/' 方向的对角线的占用情况, [1 = 被占据，0 = 未被占据]
	var na int  // 记录 '\' 方向的对角线的占用情况

	count := 0

	dfs2(0, n, cols, pie, na, &count)

	return count
}

func dfs2(row, n, cols, pie, na int, count *int) {
	if row >= n { // r 最多 n-1, row =n时表示换后已放完
		*count++
		return
	}

	// 一般语言中，按位取反是：~, 而go中是^
	// `^(cols | pie | na)` : 找到所有已占位的1, 再取反获得空位
	// `((1 << n) - 1)` : 取反时原有没用的最高位的0也变成了1, 因此要消掉. 它也可声明为常量`const upperlimit =  (1 << n)-1`便于操作
	bits := (^(cols | pie | na)) & ((1 << n) - 1) // 得到当前行所有空位

	var p int
	for bits > 0 { //遍历所有可以放置queen的位置
		p = bits & -bits // 取到序列最后一个1, 表示该位置可以放置queen. -bits = ^bit+1

		// cols|p, 表示占用对应的列
		// (pie|p)<<1, 表示下一行时, 原有p的左下角被占用, 通过递归并移位, 逐渐表示当前行相应的queen所在对角线上的格子已被占用
		// (na|p)>>1, 同理, 原有p的右下角被占用
		// pie,na移位操作后溢出或舍掉的位即表示不用管了
		dfs2(row+1, n, cols|p, (pie|p)<<1, (na|p)>>1, count) // 占用空位

		bits = bits & (bits - 1) // 去掉序列的最后一个1, 尝试下一个位置
	}
}
