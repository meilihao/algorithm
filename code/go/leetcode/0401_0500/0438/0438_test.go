/*
438.中 找到字符串中所有字母异位词
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

示例 1:

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

示例 2:

输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
*/
package demo

import (
	"reflect"
	"testing"
)

// 字母异位词（Anagram，即字母相同但顺序不同的子串）
func TestFindAnagrams(t *testing.T) {
	cs := []struct {
		s, p string
		want []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
	}

	for i, v := range cs {
		if tmp := findAnagrams2(v.s, v.p); !reflect.DeepEqual(tmp, v.want) {
			t.Errorf("i: %d, %v, not match", i, tmp)
		}
	}
}

func findAnagrams(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	var sCount, pCount [26]int // pCount: 目标窗口
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] { // 仅表示移动次数
		// 右移窗口
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++

		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}

func findAnagrams2(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	var sCount, pCount [26]int // pCount: 目标窗口
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i := pLen; i < sLen; i++ { // 仅表示移动次数
		// 右移窗口
		sCount[s[i]-'a']++
		sCount[s[i-pLen]-'a']--

		if sCount == pCount {
			ans = append(ans, i-pLen+1)
		}
	}
	return
}
