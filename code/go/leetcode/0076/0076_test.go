/*
76.困 最小覆盖子串 Minimum Window

给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
示例 2：

输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。
示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。
*/

package demo

import (
	"fmt"
	"math"
	"testing"
)

// 思路: 滑动窗口
func TestMinWindow(t *testing.T) {
	cs := []struct {
		s, t, want string
	}{{"t", "tt", ""}, {"abc", "def", ""}, {"abc", "ac", "abc"}}

	for _, v := range cs {
		if g := minWindow(v.s, v.t); g != v.want {
			fmt.Printf("%s %s ->%s != %s\n", v.s, v.t, g, v.want)
		}
		fmt.Println("---")
	}
}

func minWindow(s string, t string) string {
	if len(s) < len(t) || len(t) == 0 {
		return ""
	}

	at := [128]int{}
	mt := make(map[byte]struct{})
	for _, v := range t {
		at[v]++
		mt[byte(v)] = struct{}{}
	}

	st := [128]int{}

	var c byte
	max, window_begin := "", 0
	for i := range s {
		st[s[i]]++

		for window_begin < i {
			c = s[window_begin]

			if at[c] == 0 { // c不在t中
				window_begin++
			} else if st[c] > at[c] { // 窗口中c的个数大于t中指定的个数
				window_begin++
				st[c]--
			} else {
				break
			}
		}

		if isWindow(st, at, mt) {
			if max == "" || len(max) > i-window_begin+1 {
				max = string(s[window_begin : i+1])
			}
		}
	}

	return max
}

// 判断是否符合条件
func isWindow(st, at [128]int, m map[byte]struct{}) bool {
	for k := range m {
		if st[k] < at[k] {
			return false
		}
	}

	return true
}

// 官方解法
func minWindow2(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{} // ori统计t中字符出现的次数, cnt统计当前窗口中字符出现的次数
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32 // 最小覆盖子串的长度
	ansL, ansR := -1, -1 // 最小覆盖子串的start, end, 即目标窗口

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 { // s[r]出现在t中, 因此cnt[s[r]]++, 窗口右扩
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok { //  s[l]出现在t中, 因此cnt[s[l]] -= 1
				cnt[s[l]] -= 1
			}
			l++ // 窗口左缩
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
