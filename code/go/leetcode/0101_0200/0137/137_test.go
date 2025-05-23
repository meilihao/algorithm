/*
137.中 只出现一次的数字 II

给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。

你必须设计并实现线性时间复杂度的算法且使用常数级空间来解决此问题。

示例 1：

输入：nums = [2,2,3,2]
输出：3
示例 2：

输入：nums = [0,1,0,1,0,1,99]
输出：99

提示：

1 <= nums.length <= 3 * 104
-231 <= nums[i] <= 231 - 1
nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
*/
package demo

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	nums := []int{2, 2, 3, 2}
	fmt.Println(singleNumber(nums))
	fmt.Println(singleNumber2(nums))
}

// 如果出现3次, 那么该二进制位所有值相加能被3整除
// 第 i 个二进制位就是数组中所有元素的第 i 个二进制位之和除以 3 的余数
// 时间复杂度：O(nlogC)，其中 n 是数组的长度，C 是元素的数据范围, 本题C=32
func singleNumber2(nums []int) int {
	ans := int32(0)
	// 从 0 到 31 遍历一个 32 位整数的每一个二进制位
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1 // 提取了 num 的第 i 位
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}

func singleNumber(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		b = (b ^ num) &^ a
		a = (a ^ num) &^ b
	}
	return b
}
