package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(removeDuplicates2(nil))
}

/*
当数组 nums 的长度大于 0 时，数组中至少包含一个元素，在删除重复元素之后也至少剩下一个元素，因此 nums[0] 保持原状即可，从下标 1 开始删除重复元素。

定义两个指针 fast 和 slow 分别为快指针和慢指针，快指针表示遍历数组到达的下标位置，慢指针表示下一个不同元素要填入的下标位置，初始时两个指针都指向下标 1。

假设数组 nums 的长度为 n。将快指针 fast 依次遍历从 1 到 n−1 的每个位置，对于每个位置，如果 nums[fast]!=nums[fast−1]，说明 nums[fast] 和之前的元素都不同，因此将 nums[fast] 的值复制到 nums[slow]，然后将 slow 的值加 1，即指向下一个位置。

遍历结束之后，从 nums[0] 到 nums[slow−1] 的每个元素都不相同且包含原数组中的每个不同的元素，因此新的长度即为 slow，返回 slow 即可

时间复杂度：O(n)，其中 n 是数组的长度。快指针和慢指针最多各移动 n 次

空间复杂度：O(1)。只需要使用常数的额外空间
*/
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow, fast := 1, 1
	for fast < n {
		if nums[fast] != nums[fast-1] { // 当前元素不同于之前元素, 即相连元素有变化
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// fast表示遍历数组到达的下标位置，slow表示最近一个不同元素的下标位置, 所以返回是slow+1
func removeDuplicates2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
