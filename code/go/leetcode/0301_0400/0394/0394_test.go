/*
394.中 字符串解码

给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例 1：

输入：s = "3[a]2[bc]"
输出："aaabcbc"
示例 2：

输入：s = "3[a2[c]]"
输出："accaccacc"
示例 3：

输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"
示例 4：

输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"

提示：

1 <= s.length <= 30
s 由小写英文字母、数字和方括号 '[]' 组成
s 保证是一个 有效 的输入。
s 中所有整数的取值范围为 [1, 300]
*/
package leetcode

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestDecodeString(t *testing.T) {
	//s := "3[a]2[bc]"
	s := "2[a2[bc]]"

	fmt.Println(decodeString(s))
}

/*
时间复杂度：记解码后得出的字符串长度为 S，除了遍历一次原字符串 s，还需要将解码后的字符串中的每个字符都入栈，并最终拼接进答案中，故渐进时间复杂度为 O(S+∣s∣)，即 O(S)。
空间复杂度：记解码后得出的字符串长度为 S，这里用栈维护 TOKEN，栈的总大小最终与 S 相同，故渐进空间复杂度为 O(S)
*/
func decodeString(s string) string {
	stk := []string{}
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stk = append(stk, string(cur))
			ptr++
		} else { // 遇到"]"
			ptr++

			sub := []string{} // 获取需要处理的字符, sub中字符当前顺序是反的
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}

			// 反转
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}

			stk = stk[:len(stk)-1]                      // 舍弃"["
			repTime, _ := strconv.Atoi(stk[len(stk)-1]) // 获取"重复次数"
			stk = stk[:len(stk)-1]                      // 舍弃"重复次数"
			t := strings.Repeat(getString(sub), repTime)
			stk = append(stk, t)
		}
	}
	return getString(stk)
}

func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}
