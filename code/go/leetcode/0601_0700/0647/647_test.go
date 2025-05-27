/*
647.中 回文子串

给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。

回文字符串 是正着读和倒过来读一样的字符串。

子字符串 是字符串中的由连续字符组成的一个序列。

示例 1：

输入：s = "abc"
输出：3
解释：三个回文子串: "a", "b", "c"
示例 2：

输入：s = "aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"

提示：

1 <= s.length <= 1000
s 由小写英文字母组成
*/

package demo

import (
	"testing"
)

// other: Manacher 算法
func TestCountSubstrings(t *testing.T) {
	cs := []struct {
		s1   string
		want int
	}{
		{"abc", 3},
		{"aaa", 6},
	}

	for i, v := range cs {
		if tmp := countSubstrings(v.s1); tmp != v.want {
			t.Errorf("i: %d, not match", i)
		}
	}
}

// 中心拓展
func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := 0
	for i := 0; i < len(s); i++ {
		count += countPalindrome(s, i, i)   // 以 i 为中心（奇数长度回文）
		count += countPalindrome(s, i, i+1) // 以 i 和 i+1 为中心（偶数长度回文）
	}

	return count
}

func countPalindrome(s string, start, end int) int {
	count := 0

	for start >= 0 && end < len(s) && s[start] == s[end] { // 向两边扩展
		count++
		start--
		end++
	}

	return count
}
