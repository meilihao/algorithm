/*
58（二）：左旋转字符串

字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
请定义一个函数实现字符串左旋转操作的功能。比如输入字符串"abcdefg"和数
字2，该函数将返回左旋转2位得到的结果"cdefgab"。
*/
package main

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	s := "abcdefg"
	//fmt.Println(leftRotateString(s, 2) == "cdefgab")
	fmt.Println(leftRotateString2(s, 2) == "cdefgab")
}

func leftRotateString(str string, k int) string {
	length := len(str)
	if str == "" || length == 0 {
		return str
	}

	k = k % length

	str = str[k:length] + str[0:k] //Go语言直接根据index截取
	return str
}

func leftRotateString2(str string, k int) string {
	n := len(str)

	ss := []byte(str)
	if n > 0 && k > 0 && k < n {
		p1Start := 0
		p1End := k
		p2Start := k
		p2End := n

		reverse(ss[p1Start:p1End]) // 翻转字符串的前面n个字符
		reverse(ss[p2Start:p2End]) // 翻转字符串的后面部分
		reverse(ss[:])             // 翻转整个字符串

		return string(ss)
	}

	return ""
}

func reverse(ss []byte) {
	i, j, n := 0, len(ss)/2, len(ss)-1
	for i < j {
		ss[i], ss[n] = ss[n], ss[i]

		i++
		n--
	}
}
