/*
47.中 全排列 II

给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],

	[1,2,1],
	[2,1,1]]

示例 2：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

提示：

1 <= nums.length <= 8
-10 <= nums[i] <= 10
*/
package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums) // 将所有相同的数字都排列在一起（例如 [1,1,2]）。这使得在去重时，只需比较相邻的元素即可
	fmt.Println("sorted:", nums)
	n := len(nums)
	perm := []int{}        // 用于在递归过程中构建当前的排列
	vis := make([]bool, n) // nums[i] 这个位置的数字已经在当前排列 perm 中被使用
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		// 如果 idx<n时, 要考虑第 idx 个位置填哪个数. 根据题目要求肯定不能填已经填过的数，因此很容易想到的一个处理手段是定义一个标记数组 vis 来标记已经填过的数
		// 解决重复问题，只要设定一个规则，保证在填第 idx 个数的时候重复数字只会被填入一次即可
		for i, v := range nums {
			//  vis[i]=true: 已使用
			// i > 0 && !vis[i-1] && v == nums[i-1]: 如果当前数字 v 等于前一个数字 nums[i-1]，并且 nums[i-1] 没有被使用过，跳过, 这是处理重复数字的核心逻辑
			// 因为nums已排序, 如果 nums[i-1] 相同的数字没有被使用过，那么 nums[i] 也不能被使用
			/*
				举例:
				假设 nums = [1,1,2]。排序后是 [1_a, 1_b, 2] (用下标区分位置).
				- 当 backtrack(0) 尝试选择 nums[0] (1_a) 时：
					- perm = [1_a]，vis[0] = true。
					- 递归 backtrack(1)。
				- 当 backtrack(0) 尝试选择 nums[1] (1_b) 时：
					- i=1, v=nums[1] (1_b)。vis[1] 是 false。
					- i > 0 是 true。nums[1] == nums[0] 是 true。
					- !vis[0] 是 true：这意味着 nums[0] (1_a) 没有被当前排列路径使用。如果 1_a 没用，我们直接用 1_b，那么生成的排列会与从 1_a 开始的排列重复（例如 [1_b, ...] 和 [1_a, ...]）。所以我们 continue，跳过 1_b。
				- 总结：如果遇到重复数字，我们只允许第一个（在 nums 中）出现的相同数字被选择。如果它已经被使用了（vis[i-1] 为 true），那么后面的相同数字就可以正常选择（因为这是在不同的递归分支中，nums[i-1] 是被上一层使用了，现在是轮到 nums[i] 被使用）。
				  这种剪枝确保了 [1_a, 1_b, 2] 和 [1_b, 1_a, 2] 这样的实质相同但位置不同的排列只生成一次。
			*/
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}
