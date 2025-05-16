/*
79.中 单词搜索

给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
输出：true
示例 3：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
输出：false

提示：

m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board 和 word 仅由大小写英文字母组成
*/
package demo

// 时间复杂度: O(MN⋅3^L), L=len(word).
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	var dfs func(i, j, k int) bool // i,j 网格的开始位置, k是word的索引
	dfs = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[k] { // i, j越界+字符不相等
			return false
		}
		// 发现board[i][j]满足条件
		board[i][j] = ' ' // 修改内容即不允许走重复位置
		dirs := []int{-1, 0, 1, 0, -1}
		ans := false
		for l := 0; l < 4; l++ { // 走四个方向
			ans = ans || dfs(i+dirs[l], j+dirs[l+1], k+1) // ans = dfs(i-1, j, k+1) ||  dfs(i, j+1, k+1) || dfs(i+1, j, k+1) || dfs(i, j-1, k+1) // 左, 下, 右,  上
		}
		board[i][j] = word[k] // 回溯位置, 确保下一次走还是有数据的
		return ans
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
