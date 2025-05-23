/*
260.中 只出现一次的数字 III

给你一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。

你必须设计并实现线性时间复杂度的算法且仅使用常量额外空间来解决此问题。

示例 1：

输入：nums = [1,2,1,3,2,5]
输出：[3,5]
解释：[5, 3] 也是有效的答案。
示例 2：

输入：nums = [-1,0]
输出：[-1,0]
示例 3：

输入：nums = [0,1]
输出：[1,0]

提示：

2 <= nums.length <= 3 * 104
-231 <= nums[i] <= 231 - 1
除两个只出现一次的整数外，nums 中的其他数字都出现两次
*/
package demo

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber(nums))
	fmt.Println(singleNumber2(nums))
}

/*
	位运算

假设数组 nums 中只出现一次的元素分别是 x1和x2, nums 中出现两次的元素都会因为异或运算的性质 a⊕b⊕b=a 抵消掉，那么最终的结果就只剩下 x1 和 x2 的异或和=x
x 显然不会等于 0，因为如果 x=0，那么说明 x 就不是只出现一次的数字了
使用位运算 x & -x ( -x = ~x + 1)取出 x 的二进制表示中最低位那个 1（最低设置位,Lowest Set Bit，LSB）
这一位是 x1 和 x2 之间唯一不同的位。如果 lsb 的某一位是 1，那么在 x1 和 x2 中，一个的该位是 0，另一个的该位是 1

可以把 nums 中的所有元素分成两类，其中一类包含所有二进制表示的第 l 位为 0 的数，另一类包含所有二进制表示的第 l 位为 1 的数。可以发现：
- 对于任意一个在数组 nums 中出现两次的元素，该元素的两次出现会被包含在同一类中；
- 对于任意一个在数组 nums 中只出现了一次的元素，它们会被包含在不同类中

最后分组并抵消出现2次的, 结果即为所求指
*/
func singleNumber2(nums []int) (ans []int) {
	xorSum := 0 // 存储所有数字的异或和
	for _, num := range nums {
		xorSum ^= num
	}
	lsb := xorSum & -xorSum
	type1, type2 := 0, 0
	for _, num := range nums {
		if num&lsb > 0 {
			type1 ^= num
		} else {
			type2 ^= num
		}
	}
	return []int{type1, type2}
}

func singleNumber(nums []int) (ans []int) {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}
	for num, occ := range freq {
		if occ == 1 {
			ans = append(ans, num)
		}
	}
	return
}
