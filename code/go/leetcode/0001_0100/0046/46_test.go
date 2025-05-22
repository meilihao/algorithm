/*
46.中 全排列

给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]


提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同
*/

package demo

import (
	"fmt"
	"testing"
)

// 关键是越界
func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	tmp := make([]int, 0, len(nums))
	m := make(map[int]bool, 3) // 判断tmp是否已使用该数

	var dfs func(int)
	dfs = func(x int) {
		if x == len(nums) {
			target := make([]int, len(nums))
			copy(target, tmp)
			res = append(res, target)
			return
		}

		for _, v := range nums {
			if !m[v] {
				tmp = append(tmp, v)
				m[v] = true
				dfs(x + 1)

				// 开始回溯
				tmp = tmp[:len(tmp)-1]
				m[v] = false
			}
		}
	}

	dfs(0)

	return res
}
