/*
84.困 柱状图中最大的矩形

给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。

示例 1:

输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10
示例 2：

输入： heights = [2,4]
输出： 4

提示：

1 <= heights.length <=105
0 <= heights[i] <= 104
*/
package leetcode

import (
	"fmt"
	"math"
	"testing"
)

// other: 单调栈 + 常数优化
func TestAsteroidCollision(t *testing.T) {
	heights := []int{2, 1, 5, 6, 2, 3}

	//fmt.Println(largestRectangleArea0(heights) == 10)
	fmt.Println(largestRectangleArea(heights) == 10)
}

// 暴力解法O(n^2): 固定高度, 再向两边扩散即从i向两边遍历, 找到左边和右边第一个小于heights[i]的时候停下, 中间长度就是最长底边
// 优化的关键点在与优化找左右边界的效率
func largestRectangleArea0(heights []int) int {
	n := len(heights)

	var height, left, right int
	result := math.MinInt
	for i := 0; i < n; i++ {
		height = heights[i]
		left, right = i, i

		for left-1 >= 0 && heights[left-1] >= height {
			left--
		}

		for right+1 < n && heights[right+1] >= height {
			right++
		}

		result = max(result, height*(right-left+1))
	}

	return result
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	mono_stack := []int{}
	// 左->右: 找每个柱子左边第一个更小的柱子
	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	// 右->左: 找每个柱子右边第一个更小的柱子
	mono_stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			right[i] = n
		} else {
			right[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	ans := 0

	//fmt.Println(left)
	//fmt.Println(right)

	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}
