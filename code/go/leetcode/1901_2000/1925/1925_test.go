/*
1925.简 统计平方和三元组的数目

一个 平方和三元组 (a,b,c) 指的是满足 a2 + b2 = c2 的 整数 三元组 a，b 和 c 。

给你一个整数 n ，请你返回满足 1 <= a, b, c <= n 的 平方和三元组 的数目。

示例 1：

输入：n = 5
输出：2
解释：平方和三元组为 (3,4,5) 和 (4,3,5) 。
示例 2：

输入：n = 10
输出：4
解释：平方和三元组为 (3,4,5)，(4,3,5)，(6,8,10) 和 (8,6,10) 。

提示：

1 <= n <= 250
*/
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
