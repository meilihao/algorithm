/*
220.困 存在重复元素 III

给你一个整数数组 nums 和两个整数 indexDiff 和 valueDiff 。

找出满足下述条件的下标对 (i, j)：

i != j,
abs(i - j) <= indexDiff
abs(nums[i] - nums[j]) <= valueDiff
如果存在，返回 true ；否则，返回 false 。

示例 1：

输入：nums = [1,2,3,1], indexDiff = 3, valueDiff = 0
输出：true
解释：可以找出 (i, j) = (0, 3) 。
满足下述 3 个条件：
i != j --> 0 != 3
abs(i - j) <= indexDiff --> abs(0 - 3) <= 3
abs(nums[i] - nums[j]) <= valueDiff --> abs(1 - 1) <= 0
示例 2：

输入：nums = [1,5,9,1,5,9], indexDiff = 2, valueDiff = 3
输出：false
解释：尝试所有可能的下标对 (i, j) ，均无法满足这 3 个条件，因此返回 false 。

提示：

2 <= nums.length <= 105
-109 <= nums[i] <= 109
1 <= indexDiff <= nums.length
0 <= valueDiff <= 109
*/
package leetcode

import (
	"fmt"
	"testing"
)

// 暴力解法: 固定num[i], 求 |i-j|<=indexDiff && |num[i]-num[j]|<=valueDiff, 时间复杂度O(n*2k)=>O(n*k)
func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	// nums := []int{1, 2, 3, 1}
	// fmt.Println(containsNearbyAlmostDuplicate(nums, 3, 0))

	nums2 := []int{1, 5, 9, 1, 5, 9}
	fmt.Println(containsNearbyAlmostDuplicate(nums2, 2, 3))
}

// 让在数值上相近的数（差值小于 w）能落入同一个或相邻的桶中
func getID(x, w int) int {
	if x >= 0 {
		return x / w
	}

	// 当 x<0 时，桶的 ID 是 (x+1)/w−1。这种计算方式是为了确保负数也能正确地按桶宽 w 分组，并且保持与正数相同的分组逻辑，使得 0 和 −1 不会因为正负号而跳到很远的桶里
	// 如果 w=3，那么 −1,−2,−3 都在桶 -1， −4,−5,−6 都在桶 -2，以此类推
	return (x+1)/w - 1
}

// 如果两个数字 x 和 y 满足 abs(x−y)≤t，那么它们要么在大小是t的同一个桶里，要么在相邻的桶里（ID - 1 或 ID + 1）
func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	// 每个桶（bucket）只保留一个元素的设计是出于算法正确性和效率的考虑
	// 如果桶内多个元素[a0, a1, ..., aj],  计算|i-j|<=indexDiff时, 肯定是`|i-最近元素的索引j|<=indexDiff`最符合条件
	mp := map[int]int{}
	for i, x := range nums {
		id := getID(x, t+1)
		if _, has := mp[id]; has { // 当前桶已有元素，必然满足差值 <= t
			return true
		}

		// 左/右边相邻桶中有元素，差值需要判断
		if y, has := mp[id-1]; has && abs(x-y) <= t {
			return true
		}
		if y, has := mp[id+1]; has && abs(x-y) <= t {
			return true
		}

		mp[id] = x

		// 移除超出滑动窗口大小的元素（假定i>j, 那么保证 i-j <= k）
		// nums[i-k] 这个元素已经超出了我们长度为 k 的滑动窗口。因此，需要将 nums[i-k] 从 mp 中删除，以保持窗口的有效性
		if i >= k {
			// 假定k=2, 那么i=2时桶中存在buckets=[0, 1], len(buckets)=k, 下一轮i=3时, 0就超出窗口了要舍弃
			fmt.Println(i, k, i-k)
			delete(mp, getID(nums[i-k], t+1))
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
