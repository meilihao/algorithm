/*
3.中 无重复字符的最长子串 : Longest Substring Without Repeating Characters

思路:
假设L[i] = s[start … i]，表示最长的子字符串，其中没有重复的元素，我们以map保存< 字符，索引>；然后访问s[i + 1]：

	1）如果s[i + 1]没有出现在map中，我们可以将s[i + 1]添加到map中，则有L[i + 1] = s[start … i + 1]；
	2）如果s[i + 1]存在于map中，并且map中对应的索引是k；则令start = max(start，k)，则L[i + 1] = s[start … i + 1]，记录此时的最大长度；

	start = max(start，k)是为了避免start出现回退,比如"abba"的最后的字符"a"
*/
package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	test()
	test2()
	test3()
	test4()
	test5()
	test6()
}

func test() {
	s := "abcabcbb"

	n := lengthOfLongestSubstring4(s)
	fmt.Println(s, n, n == 3)
}

func test2() {
	s := "bbbbb"

	n := lengthOfLongestSubstring4(s)
	fmt.Println(s, n, n == 1)
}

func test3() {
	s := "pwwkew"

	n := lengthOfLongestSubstring4(s)
	fmt.Println(s, n, n == 3)
}

func test4() {
	s := "a"

	n := lengthOfLongestSubstring4(s)
	fmt.Println(s, n, n == 1)
}

func test5() {
	s := "aabab"

	n := lengthOfLongestSubstring4(s)
	fmt.Println(s, n, n == 2)
}

func test6() {
	s := "abba"

	n := lengthOfLongestSubstring4(s)
	fmt.Println(s, n, n == 2)
}

func printBytes(cs []byte) {
	for _, v := range cs {
		fmt.Printf("%c", v)
	}

	fmt.Println()
}

// 最笨
// O（n ^ 2）
func lengthOfLongestSubstring(s string) int {
	cs := []byte(s)
	l := len(cs)
	if l == 0 {
		return 0
	}

	max := 1 // Test4

	for i := 0; i < l; i++ {
		m := make(map[byte]struct{})
		m[cs[i]] = struct{}{}

		for j := i + 1; j < l; j++ {
			if _, ok := m[cs[j]]; !ok {
				m[cs[j]] = struct{}{}
			} else {
				if len(m) > max {
					max = len(m)
				}

				break
			}

			if j == l-1 && len(m) > max { // for Test5
				max = len(m)
			}
		}
	}

	return max
}

// 用v数组记录字母出现的位置，start指上次子串查找中, 重复出现的那个字符的位置即(本次查找子串的第一个字符的前一个字符)???
// O（n）
func lengthOfLongestSubstring2(s string) int {
	var a [128]int
	for i := range a { // 排除a[i]为0的影响
		a[i] = -1
	}

	max, start := 0, -1

	for i, v := range s {
		//if a[v] > -1 { // for Test6
		if a[v] > start { // 本次出现的位置a[v]应在满足上一次出现重复字符的位置(start)的后面
			start = a[v] // 上一次v出现时的index 对应下面的`a[v] = i`
		}

		a[v] = i // 当前v出现的index

		if i-start > max {
			max = i - start
		}

		fmt.Println(v, a[v], start)
	}

	return max
}

// 最好理解
// O（n）
func lengthOfLongestSubstring3(s string) int {
	max, start := 0, 0      // max: 记录当前找到的最长子串长度; start: 当前无重复子串的起始位置, start为0是指从s[0]开始查找
	m := make(map[rune]int) // 记录每个字符最后一次出现的位置

	for i, v := range s {
		// start <= j -> start<j:本次出现的起点比上一次靠后, start==j处理连续且重复的字符     // example: "aab" 或 "aaaa" for start == j
		if j, ok := m[v]; ok && start <= j { // j 大于等于 start，说明这个字符在当前子串中重复了
			start = j + 1 // 从上一次出现重复字符的下一个位置开始重新计算
		} else {
			if i-start+1 > max { // for 求max
				max = i - start + 1
			}
		}

		m[v] = i
	}

	return max
}

// best
func lengthOfLongestSubstring4(s string) int {
	max, start := 0, 0 // start为0是指从s[0]开始查找
	m := [128]int{}
	for i := range m {
		m[i] = -1
	}

	for i, v := range s {
		// start <= j -> start<j:本次出现的起点比上一次靠后, start==j处理连续且重复的字符     // example: "aab" 或 "aaaa" for start == j
		if m[v] > -1 && start <= m[v] {
			start = m[v] + 1 // 从上一次出现重复字符的下一个位置开始重新计算
		} else {
			if i-start+1 > max { // for 求max
				max = i - start + 1
			}
		}

		m[v] = i
	}

	return max
}
