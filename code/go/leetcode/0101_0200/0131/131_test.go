/*
131.中 分割回文串

给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

示例 1：

输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
示例 2：

输入：s = "a"
输出：[["a"]]

提示：

1 <= s.length <= 16
s 仅由小写英文字母组成
*/
package leetcode

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	s := "aab" //"aaba"
	fmt.Println(partition(s))

	s2 := "google"
	fmt.Println(partition(s2))
}

func partition(s string) [][]string {
	list := [][]string{}

	var subs []string
	var helper func(string, int, *[][]string) // start int: 当前子串开始的索引。表示我们正在尝试从 s[start] 开始寻找回文子串
	helper = func(s string, start int, list *[][]string) {
		if start == len(s) { // start 索引已经到达了字符串 s 的末尾（即所有字符都已被分割）
			target := make([]string, len(subs))
			copy(target, subs)
			*list = append(*list, target)
			return
		}

		for i := start; i < len(s); i++ {
			if isPalindrome(s, start, i) { // 如果 s[start:i+1] 是回文串. 如果是回文串，将其加入 subs，并递归处理剩余字符串 s[i+1:]
				subs = append(subs, s[start:i+1])
				helper(s, i+1, list)      // 递归调用：从 i+1 处继续分割剩余字符串
				subs = subs[:len(subs)-1] // 回溯：撤销选择，从 subs 中移除当前回文子串
			}
		}
	}
	helper(s, 0, &list)

	return list
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}
