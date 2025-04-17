package leetcode

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {
	n := 4
	k := 2

	fmt.Println(combine(n, k))
}

func combine(n int, k int) (ans [][]int) {
	// 初始化
	// 将 temp 中 [0, k - 1] 每个位置 i 设置为 i + 1，即 [0, k - 1] 存 [1, k]
	// 末尾加一位 n + 1 (达不到的数)作为哨兵
	temp := []int{}
	for i := 1; i <= k; i++ {
		temp = append(temp, i)
	}
	temp = append(temp, n+1)

	fmt.Println("tmp:", temp)

	for j := 0; j < k; {
		comb := make([]int, k)
		fmt.Println("j:", j, temp)
		copy(comb, temp[:k])
		ans = append(ans, comb)

		// 寻找第一个 temp[j] + 1 != temp[j + 1] 的位置 t
		// 需要把 [0, t - 1] 区间内的每个位置重置成 [1, t]
		for j = 0; j < k && temp[j]+1 == temp[j+1]; j++ {
			temp[j] = j + 1
		}
		// j 是第一个 temp[j] + 1 != temp[j + 1] 的位置
		temp[j]++
	}
	return
}
