/*
242.简 有效的字母异位词

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的 字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false

提示:

1 <= s.length, t.length <= 5 * 104
s 和 t 仅包含小写字母
*/
package main

import (
	"fmt"
	"testing"
)

// 思路:
// 1. 排序字符串后再比较
// 2. map计数(best)
func TestIsAnagram(t *testing.T) {
	s := "中国"  //"anagram"
	s2 := "国中" //"nagaram"

	fmt.Println(isAnagram(s, s2))
}

func isAnagram(s string, t string) bool {
	ms := make(map[rune]int32)
	mt := make(map[rune]int32)

	for _, v := range s {
		ms[v] += 1
	}

	for _, v := range t {
		mt[v] += 1
	}

	if len(ms) != len(mt) {
		return false
	}

	for k, v := range ms {
		if v != mt[k] {
			return false
		}
	}

	return true
}

// s 仅包含小写字母时可用make([]byte,26)代替map
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var (
		freq [26]int
		i    int
	)

	for i = 0; i < len(s); i++ { // 设置出现次数
		freq[s[i]-'a']++
	}

	for i = 0; i < len(t); i++ {
		freq[t[i]-'a']--
		if freq[t[i]-'a'] < 0 { // 非字母异位词的情况下, t中部分字母比s少, 同时t中部分字母比s多, 此时判断少的情况比较方便
			return false
		}
	}

	return true
}
