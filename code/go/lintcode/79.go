// 79. 最长公共子串 : 给出两个字符串，找到最长公共子串，并返回其长度
// 思路: 最小编辑距离
package main

import "fmt"

func main() {
	a := "aaaaaaaaaaaabbbbbcdd" //"acd" // "abcd"
	b := "abcdd"                // "ac"

	fmt.Println(longestCommonSubstring2(b, a))
}

// best, 做性能测试发现, 虽然比longestCommonSubstring2浪费了空间, 但却更快 => 空间优化和实际效率要具体情况具体分析
// dp[i][j] 表示 a 的前 i 个字母和 b 的前 j 个字母之间的最长公共子串长度
// 当 a[i] == b[j], a[i]和b[j]可以作为公共子串最后一个字符 => dp[i][j] = dp[i-1][j-1]+1
// 当 a[i] != b[j], dp[i][j]什么也不做 => dp[i][j] = 0, 其实该语句也可直接省略
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

// 空间复杂度 : O(min(len(A),len(B))
func longestCommonSubstring2(A string, B string) int {
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

	m := getMin(al, bl)
	dp := make([]rune, m+1) // dp 与 min(al,bl) 对应

	if m != al { // 假设al是min
		as, bs = bs, as
		al, bl = bl, al
	}

	fmt.Println(al, bl, string(as), string(bs))

	// 用了对角线上值, 导致无法直接计算
	var pre, cur rune // 之前的dp[j-1](即对角线的值), 之前的dp[j]

	// 虽然使用了对角线的值, 但因为匹配位置不定(前面,中部, 后部等等), 还是需要遍历i
	for i := 1; i <= bl; i++ {
		pre = 0 // 因为每行的dp[0]=0

		for j := 1; j <= al; j++ { // j按最小值列遍历
			cur = dp[j]

			if bs[i-1] == as[j-1] {
				dp[j] = pre + 1
			} else {
				dp[j] = 0 // 二维数组时不用这一步是因为它不会重复利用空间, 但使用一维数组时会重复使用
			}

			pre = cur

			max = getMax(max, dp[j])
		}
	}

	return int(max)
}

func getMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}
