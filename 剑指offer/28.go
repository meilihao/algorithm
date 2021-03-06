// 输入一个字符串，打印出该字符串中字符的所有排列
// 例：输入字符串abc，则输出由字符a、b、c 所能排列出来的所有字符串abc、acb、bac、bca、cab 和cba
// [《剑指offer》题解(二十)](https://zhuanlan.zhihu.com/p/34114529)
package main

import (
	"fmt"
	//"sort"
)

func main() {
	s := "abca"

	ss := []byte(s)

	if n := len(ss); n <= 1 {
		fmt.Println(s)
	} else {
		all := []string{}

		printAll(ss, 0, len(ss)-1, &all)

		fmt.Println(all)
		// sort.Strings(all)
		// fmt.Println(all)

		m := make(map[string]bool)

		for _, v := range all {
			if !m[v] { // 去重
				m[v] = true

				fmt.Println(v)
			}
		}
	}
}

func printAll(ss []byte, start, end int, p *[]string) {
	if end <= 0 {
		return
	}

	if start == end {
		*p = append(*p, string(ss))
	} else {
		for i := start; i <= end; i++ {
			ss[i], ss[start] = ss[start], ss[i]

			printAll(ss, start+1, end, p)

			ss[start], ss[i] = ss[i], ss[start]
		}
	}
}
