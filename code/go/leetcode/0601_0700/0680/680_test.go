/*
680.简 验证回文串 II

给你一个字符串 s，最多 可以从中删除一个字符。

请你判断 s 是否能成为回文字符串：如果能，返回 true ；否则，返回 false 。

示例 1：

输入：s = "aba"
输出：true
示例 2：

输入：s = "abca"
输出：true
解释：你可以删除字符 'c' 。
示例 3：

输入：s = "abc"
输出：false


提示：

1 <= s.length <= 105
s 由小写英文字母组成
*/

package demo

import (
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	cs := []struct {
		s1   string
		want bool
	}{
		{"aba", true},
		{"abca", true},
		{"abc", true},
	}

	for i, v := range cs {
		if tmp := validPalindrome(v.s1); tmp != v.want {
			t.Errorf("i: %d, not match", i)
		}
	}
}

func validPalindrome(s string) bool {
	low, high := 0, len(s)-1
	for low < high {
		if s[low] == s[high] { // 无需处理
			low++
			high--
		} else {
			flag1, flag2 := true, true
			for i, j := low, high-1; i < j; i, j = i+1, j-1 { // 尝试删除右侧
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}
			for i, j := low+1, high; i < j; i, j = i+1, j-1 { // 尝试删除左侧
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}
