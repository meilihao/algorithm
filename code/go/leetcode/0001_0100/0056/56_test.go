/*
56.中 合并区间

以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

提示：

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
*/
package leetcode

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(reflect.DeepEqual(merge(intervals), [][]int{{1, 6}, {8, 10}, {15, 18}}))

	intervals2 := [][]int{{1, 4}, {4, 5}}
	fmt.Println(reflect.DeepEqual(merge(intervals2), [][]int{{1, 5}}))
}

func merge(intervals [][]int) (ans [][]int) {
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] }) // 按照左端点从小到大排序

	for _, p := range intervals {
		m := len(ans)
		if m > 0 && p[0] <= ans[m-1][1] { // 下一个片段的开头与上一个片段的结尾重合, 可以合并
			ans[m-1][1] = max(ans[m-1][1], p[1]) // 更新右端点最大值
		} else { // 不相交，无法合并
			ans = append(ans, p) // 新的合并区间
		}
	}

	return
}
