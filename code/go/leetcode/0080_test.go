package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates80(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates80(nums))
	fmt.Println(nums)
}

/*
定义两个指针 slow 和 fast 分别为慢指针和快指针，其中慢指针表示处理出的数组的长度，快指针表示已经检查过的数组的长度，
即 nums[fast] 表示待检查的元素，nums[slow−1] 为上一个应该被保留的元素所移动到的指定位置

因为相同元素最多出现两次而非一次, nums[slow−1]肯定不与nums[fast]相同, 所以需要检查上上个应该被保留的元素 nums[slow−2] 是否和当前待检查元素 nums[fast] 相同

时间复杂度：O(n)，其中 n 是数组的长度。我们最多遍历该数组一次。

空间复杂度：O(1)。我们只需要常数的空间存储若干变量
*/
func removeDuplicates80(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
