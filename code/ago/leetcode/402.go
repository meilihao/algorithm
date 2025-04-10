/*
	移掉K位数字: Remove K Digits
*/
package main

import "fmt"

func main() {
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

	fmt.Println(k,stack)
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
