/*
45：把数组排成最小的数

输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼
接出的所有数字中最小的一个。例如输入数组{3, 32, 321}，则打印出这3个数
字能排成的最小数字321323。
*/
package demo

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func TestMinNumber(t *testing.T) {
	nums := []int{3, 32, 321}
	fmt.Println(MinNumber(nums) == "321323")
}

func MinNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x { // 计算 sx 为大于 x 的最小的 10 的幂
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		//fmt.Println(x, y, sx, sy)
		// sy*x+y: 数学上模拟 xy 的拼接
		return sy*x+y < sx*y+x
	})
	//fmt.Println(nums)
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
