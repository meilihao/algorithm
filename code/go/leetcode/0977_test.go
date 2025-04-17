package leetcode

import (
	"fmt"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	//nums := []int{-7, -3, 2, 3, 11}
	nums := []int{-7, -3, -1, -1}
	//nNums := sortedSquares(nums)
	nNums := sortedSquares2(nums)
	fmt.Println(nNums)
}

/*
使用两个指针分别指向位置 0 和 n−1，每次比较两个指针对应的数，选择较大的那个逆序放入答案并移动指针。这种方法无需处理某一指针移动至边界的情况

时间复杂度：O(n)，其中 n 是数组 nums 的长度。

空间复杂度：O(1)
*/
func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	i, j := 0, n-1
	for pos := n - 1; pos >= 0; pos-- { // pos是放置结果数据位置的指针
		if v, w := nums[i]*nums[i], nums[j]*nums[j]; v > w {
			ans[pos] = v
			i++
		} else {
			ans[pos] = w
			j--
		}

		//fmt.Println(i, j, ans)
	}
	return ans
}

func sortedSquares2(nums []int) []int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	res := make([]int, len(nums))
	idx := len(nums) - 1
	left := 0
	right := len(nums) - 1

	var x, y int
	for left <= right {
		x = abs(nums[left])
		y = abs(nums[right])

		if x > y {
			res[idx] = x * x
			left++
			idx--
		} else {
			res[idx] = y * y
			right--
			idx--
		}
	}
	return res
}
