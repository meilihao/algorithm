/*
402.中 移掉 K 位数字

给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。

示例 1 ：

输入：num = "1432219", k = 3
输出："1219"
解释：移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219 。
示例 2 ：

输入：num = "10200", k = 1
输出："200"
解释：移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
示例 3 ：

输入：num = "10", k = 2
输出："0"
解释：从原数字移除所有的数字，剩余为空就是 0 。

提示：

1 <= k <= num.length <= 105
num 仅由若干位数字（0 - 9）组成
除了 0 本身之外，num 不含任何前导零
*/
package main

import (
	"fmt"
	"testing"
)

func TestRemoveKdigits(t *testing.T) {
	fmt.Println(removeKdigits("1001", 3))
}

// O(len(num)+k)
func removeKdigits(num string, k int) string {
	n := len(num)
	if n <= k {
		return "0"
	}

	// 使用栈: 上一个元素与当前元素比较
	stack := make([]byte, n) // 使用n是因为压栈时一个元素也无法删除
	top := 0

	// 每次压入一个元素: 如果当前元素大于栈顶元素则先pop后再压入
	for i := 0; i <= n-1; i++ {
		for top > 0 && stack[top-1] > num[i] && k > 0 { // 使用for: 尽可能多地删除k个字符,比如"12340"
			top--
			k--
		}

		stack[top] = num[i]
		top++
	}

	if k > 0 { // 12345
		top -= k
	}

	fmt.Println(k, stack)
	stack = stack[:top]

	i, n := 0, len(stack)
	for i < n && stack[i] == '0' {
		i++
	}

	if i == n {
		return "0"
	}

	return string(stack[i:])
}

// best
func removeKdigits2(num string, k int) string {
	n := len(num)
	if n <= k {
		return "0"
	}

	stack := make([]byte, n)
	top := 0

	for i := range num {
		for top > 0 && stack[top-1] > num[i] && k > 0 {
			top--
			k--
		}

		if num[i] != '0' || top > 0 { // 禁止栈底是'0', 因为压入'0'则需另外步骤来处理前导零
			stack[top] = num[i]
			top++
		}
	}

	fmt.Println(top, k)
	for top > 0 && k > 0 {
		top--
		k--
	}

	if top == 0 {
		return "0"
	}

	return string(stack[:top])
}
