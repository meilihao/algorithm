package leetcode

import (
	"fmt"
	"testing"
)

/*
时间复杂度：O(mn)。

空间复杂度：O(mn)，即为存储所有状态需要的空间. 注意到 f(i,j) 仅与`第 i 行`和`第 i−1 行`的状态有关，因此可以使用滚动数组代替代码中的二维数组，使空间复杂度降低为 O(n)。此外，由于我们交换行列的值并不会对答案产生影响，因此我们总可以通过交换 m 和 n 使得 m≤n，这样空间复杂度降低至 O(min(m,n))
*/
func TestUniquePaths(t *testing.T) {
	m := 3
	n := 7

	//fmt.Println(uniquePaths(m, n))
	fmt.Println(uniquePaths2(m, n))
}

func uniquePaths(m, n int) int {
	dp := make([][]int, m) // 行
	for i := range dp {
		dp[i] = make([]int, n) // 列
		dp[i][0] = 1           // 到第一列的每格都只有1种走法
	}
	for j := 0; j < n; j++ { // // 到第一行的每格都只有1种走法
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] // f(i,j)=f(i−1,j)+f(i,j−1)
		}
	}
	return dp[m-1][n-1]
}

func uniquePaths2(m, n int) int {
	dp := make([]int, n)
	for i := range dp { // 到第一行的每格都只有1种走法
		dp[i] = 1
	}

	// 遍历列，dp[j]表示：到每行j列的路径数: 扫描行：本列dp[j] = 左列dp[j]→ + 上行dp[j - 1]
	for i := 1; i < m; i++ {
		fmt.Println("b", i, dp)

		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}

		fmt.Println("r", i, dp)
	}
	return dp[n-1]
}
