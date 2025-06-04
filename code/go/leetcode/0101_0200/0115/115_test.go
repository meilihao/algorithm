/*
115.困 不同的子序列

给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数。

测试用例保证结果在 32 位有符号整数范围内。

示例 1：

输入：s = "rabbbit", t = "rabbit"
输出：3
解释：
如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
rabbbit
rabbbit
rabbbit
示例 2：

输入：s = "babgbag", t = "bag"
输出：5
解释：
如下所示, 有 5 种可以从 s 中得到 "bag" 的方案。
babgbag
babgbag
babgbag
babgbag
babgbag

提示：

1 <= s.length, t.length <= 1000
s 和 t 由英文字母组成
*/
package leetcode

import (
	"fmt"
	"testing"
)

func TestNumDistinct(t *testing.T) {
	s := "rabbbit"
	s0 := "rabbit"

	fmt.Println(numDistinct(s, s0) == 3)
	fmt.Println(numDistinct2(s, s0) == 3)
}

func numDistinct(s, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	dp := make([][]int, m+1) //  dp[i][j] 表示在 s[i:] 的子序列中 t[j:] 出现的个数, s[i:] 表示 s 从下标 i 到末尾的子字符串，t[j:] 表示 t 从下标 j 到末尾的子字符串
	for i := range dp {
		dp[i] = make([]int, n+1)
		// dp[i][n] = 1 的含义： 当 j 等于 n 时，t[n...] 表示一个空字符串. 任何字符串 s[i...] 都可以通过删除所有字符来形成一个空字符串。
		// 所以，从 s 的任何位置 i 开始的子串，都可以用 1 种方式形成一个空字符串 t[n...]（即，不选择任何字符）
		dp[i][n] = 1
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				// 两种情况加起来：
				// 1. 选中 s[i] 来匹配 t[j]：
				//    那么 s[i+1...] 必须能够形成 t[j+1...]。这由 dp[i+1][j+1] 表示。
				// 2. 不选中 s[i]（跳过 s[i]）：
				//    那么 s[i+1...] 必须能够形成 t[j...]。这由 dp[i+1][j] 表示。
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				// 如果 s 的当前字符 s[i] 和 t 的当前字符 t[j] 不匹配
				// 那么 s[i] 肯定不能用来匹配 t[j]。
				// 我们只能选择跳过 s[i]：
				// s[i+1...] 必须能够形成 t[j...]。这由 dp[i+1][j] 表示。
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}

/*
正向dp:

f(i,j):
- s[i]!=t[j] : f(i-1,j) # s[i]没有贡献

	s="acab", t="ac", i=3, j=1

- s[i]=t[j] : f(i-1,j) + f(i-1,j-1)

	s="acac", t="ac", i=3, j=1
	- 使用s[3]=c, 此时消耗t[1]=c: f(i-1,j-1)
	- 不使用s[3]=c, 看"aca"是否有t子序列: f(i-1,j)
*/
func numDistinct2(s, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}

	// dp[i][j]:以i-1为结尾的s子序列(s[:i])中出现以j-1为结尾的t(t[:j])的个数
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1 // 空字符串 t 是任何字符串 s 的一个子序列
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j] // 不使用 s[i-1] 时的方案数

			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}
