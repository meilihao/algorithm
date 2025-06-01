/*
93.中 复原 IP 地址

有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。

示例 1：

输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]
示例 2：

输入：s = "0000"
输出：["0.0.0.0"]
示例 3：

输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

提示：

1 <= s.length <= 20
s 仅由数字组成
*/
package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"25525511135", 2},
		// {"0000", 1},
		// {"101023", 5},
	}

	for i, v := range cases {
		if tmp := restoreIpAddresses(v.in); len(tmp) != v.want {
			t.Errorf("%d not match", i)
		} else {
			fmt.Println(tmp)
		}
	}
}

func restoreIpAddresses(s string) []string {
	list := []string{}

	var helper func(string, int, int, string, string, *[]string)
	/*
		- i：当前在字符串 s 中的索引，表示正在处理 s[i]
		- segI：当前处理的段数（segI 是 IP 地址的第几个段，最多有 4 段）
		- seg：当前正在构建的 IP 地址段
		- ip：到目前为止已经完成的 IP 地址部分（已经加上了前面的段）
	*/
	helper = func(s string, i, segI int, seg, ip string, list *[]string) {
		if i == len(s) && segI == 3 && isValidSeg(seg) {
			*list = append(*list, ip+seg)
		} else if i < len(s) && segI <= 3 {
			ch := s[i : i+1] // 获取当前字符

			// 决策 1: 将当前字符 ch 添加到当前正在构建的段 seg 中
			// 只有当 seg + ch 仍然是一个有效的 IP 段时，才进行此操作
			if isValidSeg(seg + ch) {
				helper(s, i+1, segI, seg+ch, ip, list)
			}

			// 决策 2: 结束当前段 seg，开始一个新的段（插入一个点）
			// 条件：
			// 1. len(seg) > 0: 当前段 seg 必须非空（不能有空段，例如 "1.."）
			// 2. segI < 3: 已经完成的段数必须少于 3 个（因为 IP 地址只有 4 段，不能在第 4 段之后再插入点）
			//    注意：当前 seg 后面会加点，意味着它作为一个完整段结束，所以 segI 变为 segI+1
			if len(seg) > 0 && segI < 3 {
				// 递归调用：i+1 (处理下一个字符)，segI+1 (段数加1)，ch (新段以当前字符开始)，ip+seg+"." (IP地址前缀添加当前完整段和点)
				helper(s, i+1, segI+1, ch, ip+seg+".", list)
			}
		}
	}
	helper(s, 0, 0, "", "", &list)

	return list
}

/*
1. 每一段的值在 0 到 255 之间。
2. 每一段的值不能有前导零，除非该段本身是 "0"。
*/
func isValidSeg(seg string) bool {
	n, err := strconv.ParseInt(seg, 10, 64)
	if err != nil {
		return false
	}

	return n <= 255 && (seg == "0" || seg[0] != '0')
}

func isValidSeg2(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	num := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
		num = num*10 + int(s[i]-'0')
		if num > 255 {
			return false
		}
	}
	return true
}

func isValidSeg3(seg string) bool {
	if len(seg) > 1 && seg[0] == '0' {
		return false // 如果有前导零，则不合法
	}
	num := 0
	for i := 0; i < len(seg); i++ {
		num = num*10 + int(seg[i]-'0')
	}
	return num >= 0 && num <= 255 // 数字范围在 0 到 255 之间
}
