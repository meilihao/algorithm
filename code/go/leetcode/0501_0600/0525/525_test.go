/*
525.中 连续数组

给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。

示例 1：

输入：nums = [0,1]
输出：2
说明：[0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
示例 2：

输入：nums = [0,1,0]
输出：2
说明：[0, 1] (或 [1, 0]) 是具有相同数量 0 和 1 的最长连续子数组。
示例 3：

输入：nums = [0,1,1,1,1,1,0,0,0]
输出：6
解释：[1,1,1,0,0,0] 是具有相同数量 0 和 1 的最长连续子数组。

提示：

1 <= nums.length <= 105
nums[i] 不是 0 就是 1
*/
package leetcode

import (
	"fmt"
	"testing"
)

/*
注意:
1.  r右移不一定让窗口和变大, 同理l左移不一定让窗口和变小, 因此不适用滑动窗口
*/
func TestThreeSum(t *testing.T) {
	var f func(nums []int) int
	f = findMaxLength

	nums := []int{0, 1, 1, 1, 1, 1, 0, 0, 0}
	fmt.Println(f(nums) == 6)
}

// 与题560类似, 将0换成-1就变成求最长连续子数组且其和为0
// i<j && preSum[j] - preSum[i-1] = 0 (preSum[j] = preSum[i-1])
func findMaxLength(nums []int) (maxLength int) {
	mp := map[int]int{0: -1} // 前缀和0->**第一次出现的下标**, 这样`i-prevIndex`才可以尽量大
	counter := 0             // 前缀和
	for i, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter-- // counter 减 1（相当于 0 变成了 -1）, 因此counter 始终代表从 nums[0] 到 nums[i] (将 0 视为 -1 后) 的累加和
		}
		if prevIndex, has := mp[counter]; has {
			maxLength = max(maxLength, i-prevIndex)
		} else {
			mp[counter] = i
		}
	}
	return
}
