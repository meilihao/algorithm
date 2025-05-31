/*
39.中 组合总和

给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

对于给定的输入，保证和为 target 的不同组合数少于 150 个。

示例 1：

输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。
示例 2：

输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]
示例 3：

输入: candidates = [2], target = 1
输出: []

提示：

1 <= candidates.length <= 30
2 <= candidates[i] <= 40
candidates 的所有元素 互不相同
1 <= target <= 40
*/
package leetcode

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7

	fmt.Println(combinationSum(candidates, target)) // [[2,2,3],[7]]
}

func combinationSum(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 不选idx, 直接跳过
		dfs(target, idx+1) // 递归调用：考虑下一个数字，当前数字不选
		// 选择idx
		if target-candidates[idx] >= 0 { // 只有当选择当前数字后，target 仍然非负时才进行
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx) // 这是允许数字重复使用的关键所在, 在递归调用中，target 减去了 candidates[idx]，但 idx 参数保持不变。这意味着在下一次递归中，以后仍然可以从当前 candidates[idx] 开始选择（即再次选择 candidates[idx]），或者选择它后面的数字
			// 当 dfs(target-candidates[idx], idx) 调用返回时，表示以当前 candidates[idx] 开头的（或包含当前 candidates[idx] 的）所有分支已经探索完毕。为了探索其他可能性，需要将 candidates[idx] 从 comb 中移除，恢复到之前的状态
			comb = comb[:len(comb)-1] // 回溯：撤销选择，将当前数字从组合中移除
		}
	}
	dfs(target, 0)
	return
}
