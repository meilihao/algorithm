package lcof

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	fmt.Println(Fibonacci(5))
}

/*
F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
*/
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	a, b := 1, 1
	result := 0
	for i := 3; i <= n; i++ {
		result = a + b
		a, b = b, result
	}
	return result
}
