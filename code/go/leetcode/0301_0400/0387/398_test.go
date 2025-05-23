/*
387.简 字符串中的第一个唯一字符

给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。

示例 1：

输入: s = "leetcode"
输出: 0
示例 2:

输入: s = "loveleetcode"
输出: 2
示例 3:

输入: s = "aabb"
输出: -1

提示:

1 <= s.length <= 105
s 只包含小写字母
*/
package demo

import (
	"fmt"
	"testing"
)

func TestNthUglyNumber(t *testing.T) {
	res := firstUniqChar("leetcode")
	fmt.Println(res == 0)

	res = firstUniqChar("aabb")
	fmt.Println(res == -1)
}

func firstUniqChar(s string) int {
	n := len(s)
	pos := [26]int{}
	for i := range pos[:] {
		pos[i] = n // n表示不可能出现的索引
	}
	for i, ch := range s {
		ch -= 'a'
		if pos[ch] == n { // 第一次遇到这个字符 ch
			pos[ch] = i
		} else {
			pos[ch] = n + 1
		}
	}
	ans := n
	// 获取最小的索引
	for _, p := range pos[:] {
		if p < ans {
			ans = p
		}
	}
	if ans < n {
		return ans
	}
	return -1
}
