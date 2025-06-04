/*
1143.中 最长公共子序列

给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

示例 1：

输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace" ，它的长度为 3 。
示例 2：

输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc" ，它的长度为 3 。
示例 3：

输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0 。

提示：

1 <= text1.length, text2.length <= 1000
text1 和 text2 仅由小写英文字符组成。
*/
package leetcode

import (
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {

}

// 最长公共子序列问题是典型的二维动态规划问题
/*
状态方程:
1. s1[i]==s2[j] : f(i,j) = f(i-1, j-1) + 1
1. s1[i]!=s2[j] : f(i,j) = max(f(i-1, j), f(i, j-1))
*/
// todo: 空间效率
func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1) // dp[i][j] 表示 text1[0:i]和text2[0:j]的最长公共子序列的长度
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 { // 如果两个字符串的最后一个字符相同，那么这个共同的字符就可以作为公共子序列的一部分。此时，LCS 的长度就等于前一个子问题（不包含这两个字符的子串的 LCS 长度）加 1
				dp[i+1][j+1] = dp[i][j] + 1
			} else { // 如果两个字符串的最后一个字符不相同，那么它们不可能同时作为 LCS 的一部分。这时，有两种选择，并且需要取这两种选择中能得到最大 LCS 长度的那一个
				// 1. 排除 text2 中的当前字符 c2：对应 dp[i][j+1]
				// 1. 排除 text1 中的当前字符 c1：对应 dp[i+1][j]
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}
