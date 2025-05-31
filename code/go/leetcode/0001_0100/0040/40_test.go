/*
40.中 组合总和 II

给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用 一次 。

注意：解集不能包含重复的组合。

示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
输出:
[
[1,2,2],
[5]
]

提示:

1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
*/
package leetcode

import (
	"sort"
	"testing"
)

func TestCombinationSum2(t *testing.T) {

}

func combinationSum2(candidates []int, target int) (ans [][]int) {
	// 排序和频率统计：确保相同的数字被一起处理，避免重复组合, 避免出现[1,1,6]和[1, 6, 1]
	sort.Ints(candidates)
	var freq [][2]int // [[num, 频次]]
	for _, num := range candidates {
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	var sequence []int
	var dfs func(pos, rest int)
	dfs = func(pos, rest int) {
		if rest == 0 { // 找到一个有效组合
			ans = append(ans, append([]int(nil), sequence...))
			return
		}
		if pos == len(freq) || rest < freq[pos][0] { // 遍历完所有不重复数字 或 剩余目标值小于当前数字
			return
		}

		dfs(pos+1, rest) // 递归调用：跳过当前数字，考虑 freq 中的下一个数字

		// 选择当前数字 freq[pos][0]（循环选择其次数）
		most := min(rest/freq[pos][0], freq[pos][1])
		for i := 1; i <= most; i++ { // 循环 i 次：选择 1 次、2 次...直到最多most次
			sequence = append(sequence, freq[pos][0])
			// pos+1: 这里递归调用时，pos 变成了 pos+1。这意味着一旦决定了 freq[pos][0] 的使用次数（i 次），就不会再在后续的递归中回头选择 freq[pos][0] 了,
			// 需要直接跳到下一个不重复数字 freq[pos+1][0]。这保证了每个组合的唯一性，避免了 [1,1,6] 和 [1,6,1] 这种实质相同的组合
			// pos+1: 是选择数字的逻辑, 和重复次数i没有关系
			dfs(pos+1, rest-i*freq[pos][0]) // 递归调用：target 减去 i*当前数字，考虑下一个数字
		}
		sequence = sequence[:len(sequence)-most] // 回溯：撤销本次循环中添加的所有当前数字
	}
	dfs(0, target)
	return
}
