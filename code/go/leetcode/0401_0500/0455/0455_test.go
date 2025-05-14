/*
455.简 分发饼干

假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。

对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j] 。如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是满足尽可能多的孩子，并输出这个最大数值。

示例 1:

输入: g = [1,2,3], s = [1,1]
输出: 1
解释:
你有三个孩子和两块小饼干，3 个孩子的胃口值分别是：1,2,3。
虽然你有两块小饼干，由于他们的尺寸都是 1，你只能让胃口值是 1 的孩子满足。
所以你应该输出 1。
示例 2:

输入: g = [1,2], s = [1,2,3]
输出: 2
解释:
你有两个孩子和三块小饼干，2 个孩子的胃口值分别是 1,2。
你拥有的饼干数量和尺寸都足以让所有孩子满足。
所以你应该输出 2。

提示：

1 <= g.length <= 3 * 104
0 <= s.length <= 3 * 104
1 <= g[i], s[j] <= 231 - 1
*/
package leetcode

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	g := []int{1, 2}
	s := []int{1, 2, 3}

	r := findContentChildren(g, s)
	fmt.Println(r)
}

/*
时间复杂度：O(mlogm+nlogn)，其中 m 和 n 分别是数组 g 和 s 的长度。对两个数组排序的时间复杂度是 O(mlogm+nlogn)，遍历数组的时间复杂度是 O(m+n)，因此总时间复杂度是 O(mlogm+nlogn)。

空间复杂度：O(logm+logn)，其中 m 和 n 分别是数组 g 和 s 的长度。空间复杂度主要是排序的额外空间开销
*/
func findContentChildren(g []int, s []int) (ans int) {
	sort.Ints(g)
	sort.Ints(s)

	m, n := len(g), len(s)
	for i, j := 0, 0; i < m && j < n; i++ { // i=小孩索引, j=饼干索引
		for j < n && g[i] > s[j] { // 饼干不够大
			j++
		}
		// 饼干够大或已越界
		if j < n {
			ans++
			j++
		}
	}
	return
}

// best
func findContentChildren2(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	m, n := len(g), len(s)
	i, j := 0, 0

	for i < m && j < n {
		if s[j] >= g[i] {
			i++
		}
		j++
	}

	return i
}

// best
func findContentChildren3(g []int, s []int) int {
	slices.Sort(g) //胃口
	slices.Sort(s) //饼干
	//遍历饼干
	ans := 0 //可以满足到第几个胃口的下标
	for _, x := range s {
		if ans < len(g) && x >= g[ans] {
			ans++
		}
	}
	return ans
}
