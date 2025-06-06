/*
97.中 交错字符串

给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。

两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：

s = s1 + s2 + ... + sn
t = t1 + t2 + ... + tm
|n - m| <= 1
交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
注意：a + b 意味着字符串 a 和 b 连接。

示例 1：

输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出：true
示例 2：

输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出：false
示例 3：

输入：s1 = "", s2 = "", s3 = ""
输出：true

提示：

0 <= s1.length, s2.length <= 100
0 <= s3.length <= 200
s1、s2、和 s3 都由小写英文字母组成
*/
package leetcode

import (
	"testing"
)

func TestIsInterleave(t *testing.T) {

}

// todo: 空间效率
func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if (n + m) != t {
		return false
	}

	// f[i][j] 的含义是：s1 的前 i 个字符 (s1[0...i-1]) 和 s2 的前 j 个字符 (s2[0...j-1]) 是否能交错形成 s3 的前 i + j 个字符 (s3[0...i+j-1])
	f := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]bool, m+1)
	}

	f[0][0] = true // 表示 s1 的前 0 个字符（空字符串）和 s2 的前 0 个字符（空字符串）可以交错形成 s3 的前 0 个字符（空字符串）
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j - 1 // s3 中对应的字符索引, 从 0 开始. p=len(s1[:i-1])+len(s1[:j-1])-1

			// f[i][j] = f[i][j] || (...)：使用 || 意味着 f[i][j] 为 true，只要“情况 1”或“情况 2”中有一个成立即可

			// 情况 1：s3[p] 是否来自 s1 的最后一个字符 s1[i-1]？
			// 只有满足以下所有条件才可能：
			// 1. 至少从 s1 中取出了一个字符 (i > 0)。
			// 2. `s1` 的前 `i-1` 个字符和 `s2` 的前 `j` 个字符已经能够交错形成 `s3` 的前 `p-1` 个字符 (即 `f[i-1][j]` 为 true)。
			// 3. `s1` 的当前字符 `s1[i-1]` 确实等于 `s3` 的当前字符 `s3[p]`。
			if i > 0 {
				f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[p])
			}

			// 情况 2：s3[p] 是否来自 s2 的最后一个字符 s2[j-1]？
			// 只有满足以下所有条件才可能：
			// 1. 至少从 s2 中取出了一个字符 (j > 0)。
			// 2. `s1` 的前 `i` 个字符和 `s2` 的前 `j-1` 个字符已经能够交错形成 `s3` 的前 `p-1` 个字符 (即 `f[i][j-1]` 为 true)。
			// 3. `s2` 的当前字符 `s2[j-1]` 确实等于 `s3` 的当前字符 `s3[p]`。
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}
	return f[n][m]
}
