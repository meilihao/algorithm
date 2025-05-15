/*
03. 找出数组中重复的数字

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：
输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3
*/

package lcof

import (
	"sort"
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	numLists := [][]int{
		// []int{1, 3, 4, 2, 2},
		// []int{0, 0, 1, 3, 4, 5, 2},
		// []int{0, 1, 2, 3, 3, 5, 2},
		[]int{3, 4, 2, 1, 1, 0},
	}
	cases := []struct {
		f    func([]int) int
		want int
	}{
		{
			f:    findDuplicate,
			want: -1,
		},
		{
			f:    findDuplicate_map,
			want: -1,
		},
		{
			f:    findDuplicate_swap,
			want: -1,
		},
	}

	for x, nums := range numLists {
		for y, v := range cases {
			nums2 := make([]int, len(nums))
			copy(nums2, nums)

			if tmp := v.f(nums2); tmp == v.want {
				t.Errorf("idx: (%d, %d), err get: %d", x, y, tmp)
			}
		}
	}
}

// 排序 + 遍历
func findDuplicate(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	for i := 1; i < n; i++ {
		if nums[i-1] == nums[i] {
			return nums[i]
		}
	}

	return 0
}

// map
func findDuplicate_map(nums []int) int {
	vis := map[int]bool{}
	for i := 0; ; i++ {
		if vis[nums[i]] {
			return nums[i]
		}
		vis[nums[i]] = true
	}
}

// 看到题目中的`0～n-1`, 可以很容易联想到下标
// 巧妙地利用数组本身作为哈希表，通过交换操作将数字归位，在归位过程中检测重复
// 时间复杂度 O(N) ： 遍历数组使用 O(N) ，每轮遍历的判断和交换操作使用 O(1)
func findDuplicate_swap(nums []int) int {
	i := 0

	for i < len(nums) {
		if nums[i] == i { // 数字已在正确位置
			i += 1
			continue
		}

		if nums[nums[i]] == nums[i] { // 当出现重复是nums[i]=i, 因此nums[nums[i]] == nums[i]
			return nums[i]
		}
		// 此处画图好理解
		// 因为nums[i]!=i, 交换nums[nums[i]], nums[i], 此时nums[nums[i]]=nums[i]即idx=nums[i]处是nums[i]=i
		// 但原位置idx=i处是交换来的新值还不一定是nums[i]=i, 因此需要一直交换直到idx=i处也是nums[i] = i为止
		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
	}

	return -1
}
