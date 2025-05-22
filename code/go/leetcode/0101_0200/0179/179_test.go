/*
179.中 最大数

给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。

示例 1：

输入：nums = [10,2]
输出："210"
示例 2：

输入：nums = [3,30,34,5,9]
输出："9534330"

提示：

1 <= nums.length <= 100
0 <= nums[i] <= 109
*/
package demo

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	nums := []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(nums) == "9534330")
}

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10 // sx 和 sy 用于计算 10^k，其中 k 是 x 或 y 的位数. 这是为了通过数学运算来模拟字符串拼接。
		for sx <= x {    // 计算 sx 为大于 x 的最小的 10 的幂
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		fmt.Println(x, y, sx, sy, sy*x+y, sx*y+x)
		// sy*x+y: 数学上模拟 xy 的拼接
		return sy*x+y > sx*y+x
	})
	fmt.Println(nums)
	// 如果排序后的数组的第一个元素是 0，那么意味着所有数字都是 0（例如 [0, 0, 0]）。在这种情况下，组合成的最大数字应该是 "0"，而不是 "000"
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}
