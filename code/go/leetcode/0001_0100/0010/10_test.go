/*
10.困 正则表达式匹配

给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s 的，而不是部分字符串。

示例 1：

输入：s = "aa", p = "a"
输出：false
解释："a" 无法匹配 "aa" 整个字符串。
示例 2:

输入：s = "aa", p = "a*"
输出：true
解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3：

输入：s = "ab", p = ".*"
输出：true
解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。

提示：

1 <= s.length <= 20
1 <= p.length <= 20
s 只包含从 a-z 的小写字母。
p 只包含从 a-z 的小写字母，以及字符 . 和 *。
保证每次出现字符 * 时，前面都匹配到有效的字符
*/
package demo

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	cases := []struct {
		s    string
		p    string
		want bool
	}{
		{
			s:    "aa",
			p:    "a",
			want: false,
		},
		{
			s:    "aa",
			p:    "a*",
			want: true,
		},
		{
			s:    "ab",
			p:    ".*",
			want: true,
		},
		{
			s:    "aaa",
			p:    "a.a",
			want: true,
		},
		{
			s:    "aaa",
			p:    "ab*ac*a",
			want: true,
		},
		{
			s:    "aaa",
			p:    "aa.a",
			want: false,
		},
		{
			s:    "aaa",
			p:    "ab*a",
			want: false,
		},
		{
			s:    "a",
			p:    "ab*",
			want: true,
		},
		{
			s:    "aaaaaaaaaaaaaaaaaaab",
			p:    "a*a*a*a*a*a*a*a*a*a*",
			want: false,
		},
	}

	for i, v := range cases {
		if tmp := isMatch3(v.s, v.p); tmp != v.want {
			t.Errorf("%d not match(%v!=%v)", i, v.want, tmp)
		}
	}
}

// 时间复杂度：O(mn)，其中 m 和 n 分别是字符串 s 和 p 的长度
func isMatch2(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool { //用于判断 s 的第 i 个字符 (s[i-1]) 是否能与 p 的第 j 个字符 (p[j-1]) 匹配
		if i == 0 { // 如果 i 为 0，表示 s 字符串为空，此时无法匹配任何具体字符
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	// f[i][j] 表示字符串 s 的前 i 个字符是否能被模式串 p 的前 j 个字符匹配即s 的前 i 个字符（s[0:i]）是否能与 p 的前 j 个字符（p[0:j]）匹配
	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true            // f[0][0] 表示空字符串 s ("") 是否能匹配空模式串 p (""); 其余初始值为 false
	for i := 0; i <= m; i++ { // 遍历 s 的长度（从 0 到 m）
		// `f[i][j] || ...`, 因为循环会多次更新f[i][j]
		for j := 1; j <= n; j++ { // 遍历 p 的长度（从 1 到 n）, j 从 1 开始是因为 f[i][0]（即模式串为空）除了 f[0][0] 之外都应为 false（即空模式串无法匹配非空字符串）
			//old := f[i][j]

			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2] // `*`匹配零次 // "a"/"ab*"
				if matches(i, j-1) {           // 判断 s 的当前字符 (s[i-1]) 是否能与 * 前面的字符 (p[j-2]) 匹配      // `*`匹配1~N次 // "ab"/"ab*"(a->ab*), "abb"/"ab*"(ab->ab*, a -> ab*), "abbb"/"ab*"(abb->ab*, ab->ab*, a -> ab*)
					// 如果 s 的当前字符能匹配 * 前面的字符，那么 f[i][j] 的真假取决于 f[i-1][j]. f[i-1][j] 意味着 s 的前 i-1 个字符已经能匹配 p 的前 j 个字符（即 s[i-1] 被 * 匹配了，并且 * 继续发挥作用匹配 s 剩余的部分）
					// 简单理解就是: 匹配 s 末尾的一个字符，将该字符扔掉，而该组合还可以继续进行匹配
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1] // f[i][j] 的真假取决于 s 的前 i-1 个字符是否能匹配 p 的前 j-1 个字符
			}

			//fmt.Println(i, j, s[:i], p[:j], old, old != f[i][j])
		}
	}
	return f[m][n]
}

// 同isMatch2, 递推看起来更正向
func isMatch3(s string, p string) bool {
	m, n := len(s), len(p)

	dp := make([][]bool, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	// 处理模式可以匹配空字符串的情况
	for i := 1; i < n; i++ {
		if p[i] == '*' { // ""/"a*"
			// 即rule(p[:i+1])=rule(p[:i-1])=> rule(`a*`)=rule("")
			// 例如p = "a*b*c*"，可以让空串匹配它，逐步填充 dp[0][2], dp[0][4] 等
			dp[0][i+1] = dp[0][i-1]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 处理s[:i+1] + p[:j+1]

			// 处理'.'通配符
			if p[j] == '.' {
				dp[i+1][j+1] = dp[i][j]
			}

			// 处理普通字符匹配
			if p[j] == s[i] {
				dp[i+1][j+1] = dp[i][j]
			}

			if p[j] == '*' {
				if p[j-1] != s[i] && p[j-1] != '.' { // 如果前一个字符不匹配 "ab"/"ac*"->"ab"/"a"
					dp[i+1][j+1] = dp[i+1][j-1] // 当前字符不能匹配 * 前的字符 → 只能跳过整个 x*
				} else { // * 前面的字符 p[j-1] 匹配 s 的当前字符 s[i]（或者 p[j-1] 是 .）
					// dp[i+1][j]: 匹配一次 "ab"/"ab*"->"ab"/"ab"
					// dp[i][j+1]: 匹配多次 "ab"/"abb*"->"a"/"abb*"
					// dp[i+1][j-1]: 匹配零次 "ab"/"abb*"->"ab"/"ab"
					dp[i+1][j+1] = (dp[i+1][j] || dp[i][j+1] || dp[i+1][j-1])
				}
			}
		}
	}
	return dp[m][n]
}

// 用递归容易出现时间超时, 废弃
func isMatch(s string, p string) bool {
	if s == "" || p == "" {
		return false
	}

	return matchCore(s, p)
}

func matchCore(s, p string) bool {
	// 都到达字符串末尾，匹配成功
	if len(s) == 0 && len(p) == 0 {
		return true
	}

	// pattern 已结束，但 str 还有剩余，匹配失败
	if len(s) != 0 && len(p) == 0 {
		return false
	}

	//fmt.Println(s, p)
	// pattern 第二个字符是 '*'
	if len(p) > 1 && p[1] == '*' {
		if len(s) > 0 && (p[0] == s[0] || p[0] == '.') {
			// 进入下一个状态、留在当前状态、忽略 '*'
			return matchCore(s[1:], p[2:]) || matchCore(s[1:], p) || matchCore(s, p[2:])
		} else {
			// 略过一个'*'
			return matchCore(s, p[2:])
		}
	}

	// 当前字符匹配 或 pattern 是 '.'
	if len(s) > 0 && (p[0] == s[0] || p[0] == '.') {
		return matchCore(s[1:], p[1:])
	}

	return false
}
