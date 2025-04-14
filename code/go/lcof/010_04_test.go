package lcof

import (
	"fmt"
	"testing"
)

func TestFibonacci2(t *testing.T) {
	fmt.Println(Fibonacci2(5))
}

/*
假设 n>=2，第一步有 n 种跳法：跳 1 级、跳 2 级、到跳 n 级
跳 1 级，剩下 n-1 级，则剩下跳法是 f(n-1)
跳 2 级，剩下 n-2 级，则剩下跳法是 f(n-2)
……
跳 n-1 级，剩下 1 级，则剩下跳法是 f(1)
跳 n 级，剩下 0 级，则剩下跳法是 f(0)

在 n>=2 的情况下：
f(n)=f(n-1)+f(n-2)+...+f(1)
因为 f(n-1)=f(n-2)+f(n-3)+...+f(1)
所以 f(n)=2*f(n-1) 又 f(1)=1,所以可得f(n)=2^(number-1)
*/
func Fibonacci2(n int) int {
	return 1 << (n - 1)
}
