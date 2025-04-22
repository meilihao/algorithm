package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestEraseOverlapIntervals(t *testing.T) {
	// intervals := [][]int{
	// 	{1, 2},
	// 	{2, 3},
	// 	{3, 4},
	// 	{1, 3},
	// }

	intervals := [][]int{
		{1, 2},
		{2, 3},
	}

	fmt.Println(eraseOverlapIntervals(intervals))
}

/*
贪心
要求保证移除区间最少，使得剩下的区间互不重叠 = 如何使得剩下互不重叠区间的数目最多 = 总区间个数 - 不重叠区间的最多个数

时间复杂度：O(nlogn)，其中 n 是区间的数量。我们需要 O(nlogn) 的时间对所有的区间按照右端点进行升序排序，并且需要 O(n) 的时间进行遍历。由于前者在渐进意义下大于后者，因此总时间复杂度为 O(nlogn)。

空间复杂度：O(logn)，即为排序需要使用的栈空间
*/
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	fmt.Println(intervals)
	ans, right := 1, intervals[0][1] // ans不重叠区间个数;当前不重叠区间的右端点 right
	for _, p := range intervals[1:] {
		if p[0] >= right { // 不重合
			ans++
			right = p[1]
		}
	}
	return n - ans
}
