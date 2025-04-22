package leetcode

import (
	"fmt"
	"testing"
)

func TestCountTriples(t *testing.T) {
	n := 5
	fmt.Println(countTriples(n))
}

func countTriples(n int) int {
	tmp := make([]int, n+1)
	mp := make(map[int]struct{}, n*2) // c^2 是否存在

	var lable struct{}
	for i := 1; i <= n; i++ {
		tmp[i] = i * i
		mp[tmp[i]] = lable
	}

	res := 0
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if _, ok := mp[tmp[i]+tmp[j]]; ok {
				res++
			}
		}
	}
	return res * 2
}
