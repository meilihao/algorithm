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
