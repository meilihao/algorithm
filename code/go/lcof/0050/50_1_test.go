/*
50（二）：字符流中第一个只出现一次的字符

请实现一个函数用来找出字符流中第一个只出现一次的字符。例如，当从
字符流中只读出前两个字符"go"时，第一个只出现一次的字符是'g'。当从该字
符流中读出前六个字符"google"时，第一个只出现一次的字符是'l'。
*/
package demo

import (
	"fmt"
	"testing"
)

func TestFirstAppearingOnce(t *testing.T) {
	s := "google"
	for _, v := range []byte(s) {
		fmt.Printf("%c\n", firstAppearingOnce(v))
	}
}

var (
	q     = []byte{}       // 到目前为止只出现过一次的字符，并且这些字符按照它们在流中出现的顺序排列
	count = map[byte]int{} // 统计频率
)

func insert(c byte) {
	q = append(q, c)
	count[c] += 1

	if count[c] > 1 {
		for len(q) > 0 && count[q[0]] > 1 { // 清除队列头部所有已经重复的字符, 因此len(q)>0, q[0]即为符合条件的char
			q = q[1:]
		}
	}
}

func firstAppearingOnce(s byte) byte {
	insert(s)

	if len(q) == 0 {
		return '#'
	}

	return q[0]
}
