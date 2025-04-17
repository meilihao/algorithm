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
