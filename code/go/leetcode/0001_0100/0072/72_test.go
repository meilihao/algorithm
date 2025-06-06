/*
72.中 编辑距离

给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符


示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')


提示：

0 <= word1.length, word2.length <= 500
word1 和 word2 由小写英文字母组成
*/

package main

import (
	"fmt"
	"testing"
)

// 思路: 最小编辑距离
func TestMinDistance(t *testing.T) {
	a := "horse"
	b := "ros"

	fmt.Println(minDistance(a, b))
}

// 动态规划
// dp[i][j] 表示 a 的前 i 个字母和 b 的前 j 个字母之间的最小编辑距离
// 当 a[i] == b[j], 此时两个子串的最后一个字母相同, 因此比较它们前面的子串即可 => dp[i][j] = dp[i-1][j-1]
// 当 a[i] != b[j], 此时有三种情况:
// 1. dp[i-1][j]+1 : 删除 a[i]， 比如 fxy -> fab 的编辑距离 = fx -> fab 的编辑距离 + 1
// 2. dp[i][j-1]+1 : 插入 b[j]， 比如 fxy -> fab 的编辑距离 = fxy -> fa 的编辑距离 + 1
// 3. dp[i-1][j-1]+1 : 将 a[i] 替换为 b[j]， 比如 fxy -> fab 的编辑距离 = fxb -> fab 的编辑距离 + 1 = fx -> fa 的编辑距离 + 1
func minDistance(A string, B string) int {
	// if one of the strings is empty
	if len(A)*len(B) == 0 {
		return len(A) + len(B)
	}

	if A == B {
		return 0
	}

	as := []rune(A)
	bs := []rune(B)
	al, bl := len(as), len(bs)

	dp := make([][]rune, bl+1) // 纵向
	for i := range dp {
		dp[i] = make([]rune, al+1) // 横向
	}

	// 初始化原始字符串与空串之间的编辑距离
	// dp[0][0]表示空串
	for i := 0; i <= bl; i++ { // 空串与B的编辑距离
		dp[i][0] = rune(i)
	}
	for i := 0; i <= al; i++ { // // 空串与A的编辑距离
		dp[0][i] = rune(i)
	}

	for i := 1; i <= bl; i++ {
		for j := 1; j <= al; j++ {
			if bs[i-1] == as[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1 // min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}

	return int(dp[bl][al])
}

func min(a, b, c rune) rune {
	// 假设min是a
	if b < a {
		a = b
	}

	if c < a {
		a = c
	}

	return a
}
