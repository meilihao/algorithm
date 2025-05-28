/*
539.中 最小时间差

给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。

示例 1：

输入：timePoints = ["23:59","00:00"]
输出：1
示例 2：

输入：timePoints = ["00:00","23:59","00:00"]
输出：0

提示：

2 <= timePoints.length <= 2 * 104
timePoints[i] 格式为 "HH:MM"
*/
package leetcode

import (
	"math"
	"sort"
	"testing"
)

func TestFindMinDifference(t *testing.T) {

}

func getMinutes(t string) int { // 转成分钟
	return (int(t[0]-'0')*10+int(t[1]-'0'))*60 + int(t[3]-'0')*10 + int(t[4]-'0')
}

// 将 timePoints 排序后，最小时间差必然出现在 timePoints 的两个相邻时间，或者 timePoints 的两个首尾时间中(因为24小时循环)
func findMinDifference(timePoints []string) int {
	sort.Strings(timePoints)
	ans := math.MaxInt32
	t0Minutes := getMinutes(timePoints[0])
	preMinutes := t0Minutes
	for _, t := range timePoints[1:] {
		minutes := getMinutes(t)
		ans = min(ans, minutes-preMinutes) // 相邻时间的时间差
		preMinutes = minutes
	}
	ans = min(ans, t0Minutes+1440-preMinutes) // 首尾时间的时间差, 1440 是一天的总分钟数 (24×60)
	return ans
}
