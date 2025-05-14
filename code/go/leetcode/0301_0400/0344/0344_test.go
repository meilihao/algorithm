/*
344.简 反转字符串

编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题

示例 1：

输入：s = ["h","e","l","l","o"]
输出：["o","l","l","e","h"]
示例 2：

输入：s = ["H","a","n","n","a","h"]
输出：["h","a","n","n","a","H"]
*/
package demo

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	cases := []struct {
		in   []byte
		want []byte
	}{
		{
			in:   []byte{},
			want: []byte{},
		},
		{
			in:   []byte{'a'},
			want: []byte{'a'},
		},
		{
			in:   []byte{'h', 'e', 'l', 'l', 'o'},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			in:   []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}

	for i := range cases {
		reverseString(cases[i].in)

		if !reflect.DeepEqual(cases[i].in, cases[i].want) {
			t.Errorf("i: %d, not match", i)
		}
	}
}

func reverseString(s []byte) {
	if len(s) <= 1 {
		return
	}

	l, r := 0, len(s)-1

	for l < r {
		s[l], s[r] = s[r], s[l]

		l++
		r--
	}
}
