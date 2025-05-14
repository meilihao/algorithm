/*
739.中 每日温度

给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

示例 1:

输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]
示例 2:

输入: temperatures = [30,40,50,60]
输出: [1,1,1,0]
示例 3:

输入: temperatures = [30,60,90]
输出: [1,1,0]
*/
package leetcode

import (
	"fmt"
	"testing"
)

// 期望j>i & answer[j] > answer[i]
func TestDailyTemperatures(t *testing.T) {
	//temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	temperatures := []int{73, 74, 73, 73, 72, 72, 72, 72}

	fmt.Println(dailyTemperatures(temperatures))
}

// 使用单调栈:从栈底到栈顶的下标对应的温度列表中的温度依次递减
/*
时间复杂度：O(n)，其中 n 是温度列表的长度。正向遍历温度列表一遍，对于温度列表中的每个下标，最多有一次进栈和出栈的操作

空间复杂度：O(n)，其中 n 是温度列表的长度。需要维护一个单调栈存储温度列表中的下标
*/
func dailyTemperatures739(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i++ {
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			// 处理符合条件的answer[j]
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}

// 容易超时, 比如temperatures成员全是99的很长数组
/*
时间复杂度：O(nm)，其中 n 是温度列表的长度，m 是数组 next 的长度，在本题中温度不超过 100，所以 m 的值为 100。反向遍历温度列表一遍，对于温度列表中的每个值，都要遍历数组 next 一遍。

空间复杂度：O(m)，其中 m 是数组 next 的长度。除了返回值以外，需要维护长度为 m 的数组 next 记录每个温度第一次出现的下标位置
*/
func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)

	var j int
	for i := 0; i < length; i++ {
		for j = i + 1; j < length; j++ {
			if temperatures[j] > temperatures[i] {
				break
			}
		}

		if j == length {
			ans[i] = 0
		} else {
			ans[i] = j - i
		}
	}
	return ans
}
