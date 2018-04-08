/*
	最小窗口子字符串 Minimum Window Substring
*/

package main

import "fmt"

func main() {
	cs := []struct {
		s, t, want string
	}{{"t", "tt", ""}, {"abc", "def", ""}, {"abc", "ac", "abc"}}

	for _, v := range cs {
		if g := minWindow2(v.s, v.t); g != v.want {
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
