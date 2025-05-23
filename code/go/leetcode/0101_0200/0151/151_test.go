/*
151.中 反转字符串中的单词

给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

示例 1：

输入：s = "the sky is blue"
输出："blue is sky the"
示例 2：

输入：s = "  hello world  "
输出："world hello"
解释：反转后的字符串中不能存在前导空格和尾随空格。
示例 3：

输入：s = "a good   example"
输出："example good a"
解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。

提示：

1 <= s.length <= 104
s 包含英文大小写字母、数字和空格 ' '
s 中 至少存在一个 单词
*/
package main

import (
	"fmt"
	"testing"
)

// 非原地: 栈
func TestReverseWords(t *testing.T) {
	s := "  hello world  "
	fmt.Println(reverseWords(s) == "world hello")

	s2 := "a good   example"
	fmt.Println(reverseWords(s2) == "example good a")
}

// 原地

func reverseWords(s string) string {
	slow, fast := 0, 0 //处理冗余空格的快慢指针, slow=放置非空格字符的索引
	b := []byte(s)     //将string类型转换为字节数组方便处理
	//处理开头的冗余空格
	for fast < len(s) && b[fast] == ' ' {
		fast++
	}
	// 挪动words时去掉多余空格
	//处理单词间的冗余空格，可能会造成最后一个单词末尾有一个多余空格
	for ; fast < len(s); fast++ {
		if fast > 0 && b[fast-1] == ' ' && b[fast] == ' ' {
			continue
		}
		b[slow] = b[fast]
		slow++
	}
	//处理最后一个单词后面的冗余空格
	if b[slow-1] == ' ' {
		b = b[:slow-1]
	} else {
		b = b[:slow]
	}

	//翻转字节数组形式字符串的函数
	reverse := func(start, end int) {
		for ; start < end; start, end = start+1, end-1 {
			b[start], b[end] = b[end], b[start]
		}
	}
	//先翻转全部单词
	reverse(0, len(b)-1)

	//再逐个翻转每个单词
	for start := 0; start < len(b); {
		j := start
		//判断是否为单词末尾
		for j < len(b) && b[j] != ' ' {
			j++
		}
		//翻转单词
		reverse(start, j-1)
		start = j + 1
	}
	//将字节数组类型的字符串转换成string返回
	return string(b)
}
