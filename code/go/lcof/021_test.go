/*
21. 调整数组顺序使奇数位于偶数前面
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分

示例：

输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一
*/
package lcof

import (
	"fmt"
	"testing"
)

func TestExchange(t *testing.T) {
	//fmt.Println(Exchange([]int{1, 2, 3, 4}))
	fmt.Println(Exchange([]int{1, 3, 2, 4, 5}))
}

/*
定义两个指针 i 和 j，i指向当前元素，j指向当前最后一个奇数的下一个位置即放下一个奇数的位置

从左到右遍历数组，当 nums[i] 是奇数时，将其与nums[j] 交换，然后j 向右移动一位. i 每次向右移动一位，直到遍历完整个数组

时间复杂度 O(n)，空间复杂度 O(1)
*/
func Exchange(nums []int) []int {
	j := 0
	for i, x := range nums {
		if x&1 == 1 { // 是奇数
			//fmt.Println(i, j, nums[i], nums[j])
			if i != j { // 跳过原地交换
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return nums
}
