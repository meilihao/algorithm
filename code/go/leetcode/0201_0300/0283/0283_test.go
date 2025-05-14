/*
283.简 移动零

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:

输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
示例 2:

输入: nums = [0]
输出: [0]

提示:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1
*/
package leetcode

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	//nums := []int{0, 1, 0, 3, 12}
	nums := []int{2, 1, 0, 3, 12, 0}

	moveZeroes(nums)
	//moveZeroes2(nums)

	fmt.Println(nums)
}

/*
使用双指针，左指针指向当前已经处理好的序列的尾部，右指针指向待处理的当前元素

右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移

左指针左边均为非零数；
右指针左边直到左指针处均为零
因此每次交换，都是将左指针的零与右指针的非零数交换，且非零数的相对顺序并未改变

时间复杂度：O(n)，其中 n 为序列长度。每个位置至多被遍历两次。

空间复杂度：O(1)。只需要常数的空间存放若干变量
*/
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums) // nums[left]放待处理的非零元素
	for right < n {
		if nums[right] != 0 {
			fmt.Println("b:", right, left, nums)

			nums[left], nums[right] = nums[right], nums[left]
			left++

			fmt.Println("a:", right, left, nums)
		}
		right++
	}
}

// 直接将非0的数字左移即可，再填充剩下的格子为0
func moveZeroes2(nums []int) {
	cur := 0 // 放待处理的非零元素
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[cur] = nums[i]
			cur++
		}
	}
	for ; cur < len(nums); cur++ {
		nums[cur] = 0
	}
}
