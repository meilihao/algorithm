/*
905.简 按奇偶排序数组

给你一个整数数组 nums，将 nums 中的的所有偶数元素移动到数组的前面，后跟所有奇数元素。

返回满足此条件的 任一数组 作为答案。

示例 1：

输入：nums = [3,1,2,4]
输出：[2,4,3,1]
解释：[4,2,3,1]、[2,4,1,3] 和 [4,2,1,3] 也会被视作正确答案。
示例 2：

输入：nums = [0]
输出：[0]

提示：

1 <= nums.length <= 5000
0 <= nums[i] <= 5000
*/
package demo

import (
	"fmt"
	"testing"
)

func TestExchange(t *testing.T) {
	//fmt.Println(Exchange([]int{1, 2, 3, 4}))
	fmt.Println(Exchange([]int{1, 3, 2, 4, 5}))
}

/*
定义两个指针 i 和 j，i指向当前元素，j指向当前最后一个偶数的下一个位置即放下一个偶数的位置

从左到右遍历数组，当 nums[i] 是偶数时，将其与nums[j] 交换，然后j 向右移动一位. i 每次向右移动一位，直到遍历完整个数组

时间复杂度 O(n)，空间复杂度 O(1)
*/
func Exchange(nums []int) []int {
	j := 0
	for i, x := range nums {
		if x&1 == 0 {
			//fmt.Println(i, j, nums[i], nums[j])
			if i != j { // 跳过原地交换
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return nums
}
