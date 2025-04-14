package leetcode

import "testing"

func TestRemoveDuplicates(t *testing.T) {

}

/*
定义两个指针 slow 和 fast 分别为慢指针和快指针，其中慢指针表示处理出的数组的长度，快指针表示已经检查过的数组的长度，
即 nums[fast] 表示待检查的第一个元素，nums[slow−1] 为上一个应该被保留的元素所移动到的指定位置

时间复杂度：O(n)，其中 n 是数组的长度。我们最多遍历该数组一次。

空间复杂度：O(1)。我们只需要常数的空间存储若干变量
*/
func removeDuplicates(nums []int) int {
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
