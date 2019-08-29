// 79. 最长公共子串 : 给出两个字符串，找到最长公共子串，并返回其长度
// 思路: 最小编辑距离
package main

import "fmt"

func main() {
	a := "ABCD"
	b := "EACD"

	fmt.Println(longestCommonSubstring(a, b))
}

// dp[i][j] 表示 a 的前 i 个字母和 b 的前 j 个字母之间的最长公共子串长度
// 当 a[i] == b[j], a[i]和b[j]可以作为公共子串最后一个字符 => dp[i][j] = dp[i-1][j-1]+1
// 当 a[i] != b[j], dp[i][j]什么也不做 => dp[i][j] = 0, 其实该语句也可直接省略
// todo : 可再优化
func longestCommonSubstring(A string, B string) int {
	if len(A)*len(B) == 0 {
		return 0
	}
	if A == B {
		return len(A)
	}

	as := []rune(A)
	bs := []rune(B)
	al, bl := len(as), len(bs)

	var max rune

	dp := make([][]rune, bl+1)
	for i := range dp {
		dp[i] = make([]rune, al+1)
	}

	// 因为原始字符串与空串之间的最长公共子串就是0, 因此不用初始dp[i][0]和dp[0][j]

	for i := 1; i <= bl; i++ {
		for j := 1; j <= al; j++ {
			if bs[i-1] == as[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			// } else {
			// 	dp[i][j] = 0
			// }

			max = getMax(max, dp[i][j])
		}
	}

	fmt.Println(dp)

	return int(max)
}

func getMax(a, b rune) rune {
	if a > b {
		return a
	}

	return b
}
