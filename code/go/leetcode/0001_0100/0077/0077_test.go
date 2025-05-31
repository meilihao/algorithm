/*
77.中 组合

给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。

示例 1：

输入：n = 4, k = 2
输出：
[

	[2,4],
	[3,4],
	[2,3],
	[1,2],
	[1,3],
	[1,4],

]
示例 2：

输入：n = 1, k = 1
输出：[[1]]
*/
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

// 非递归（字典序法）实现组合型枚举
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

// 递归实现组合型枚举
func combine2(n int, k int) (ans [][]int) {
	temp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		// 剪枝：temp 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
		if len(temp)+(n-cur+1) < k {
			return
		}
		// 记录合法的答案
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		// 考虑选择当前位置
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		// 考虑不选择当前位置
		dfs(cur + 1)
	}
	dfs(1) // 数字是从1开始的
	return
}
