/*
03. 数组中重复的数字

找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：
输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3
*/

package lcof

import (
	"fmt"
	"sort"
	"testing"
)

/*
other:
1. map
*/
func TestFindDuplicate(t *testing.T) {
	nums := []int{1, 3, 4, 2, 2}

	fmt.Println(findDuplicate(nums))
	fmt.Println(findDuplicate_map(nums))
	fmt.Println(findDuplicate_fast(nums))
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

func findDuplicate_map(nums []int) int {
	vis := map[int]bool{}
	for i := 0; ; i++ {
		if vis[nums[i]] {
			return nums[i]
		}
		vis[nums[i]] = true
	}
}

/*
利用了数组元素的值作为索引来标记访问过的数字，通过将对应位置的数值取负来标记该数字已经被访问过

限制:
1. 数组元素必须在1到n之间（n为数组长度-1）
1. 必须恰好有一个重复数字（可以有多次重复）
1. 如果数组中可能有0或负数，需要特殊处理
*/
func findDuplicate_fast(nums []int) int {
	for i := range nums { // 遍历数组中的每个元素
		k := nums[i] // 获取当前元素值
		if k < 0 {   // 如果当前元素是负数（可能已被标记）
			k = -k // 取其绝对值
		}
		if nums[k] < 0 { // 检查nums[k]是否为负数（表示之前已经被标记过）
			return k // 如果为负，说明k是重复的数字
		}
		nums[k] = -nums[k] // 将nums[k]标记为负数，表示数字k已经被访问过
	}
	return 0 // 如果没有找到重复（根据题意应该总是有重复）
}
