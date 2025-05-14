/*
567.中 字符串的排列

给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的 排列。如果是，返回 true ；否则，返回 false 。

换句话说，s1 的排列之一是 s2 的 子串 。

示例 1：

输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").
示例 2：

输入：s1= "ab" s2 = "eidboaoo"
输出：false
*/

package demo

import (
	"testing"
)

// 由于排列不会改变字符串中每个字符的个数，所以只有当两个字符串每个字符的个数均相等时，一个字符串才是另一个字符串的排列
func TestCheckInclusion(t *testing.T) {
	cs := []struct {
		s1, s2 string
		want   bool
	}{
		//{"ab", "eidbaooo", true},
		//{"ab", "eidboaoo", false},
		{"ab", "eadboaoo", false},
	}

	for i, v := range cs {
		if tmp := checkInclusion2(v.s1, v.s2); tmp != v.want {
			t.Errorf("i: %d, not match", i)
		}
	}
}

func checkInclusion(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m { // 长度检查
		return false
	}
	var cnt1, cnt2 [26]int //  cnt1 统计 s1中各个字符的个数，cnt2 统计当前遍历的子串中各个字符的个数. 窗口len(cnt1)==len(cnt2)

	for i, ch := range s1 {
		cnt1[ch-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}
	for i := n; i < m; i++ {
		// 前进一个窗口
		cnt2[s2[i]-'a']++
		cnt2[s2[i-n]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}

// best
func checkInclusion2(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for _, ch := range s1 {
		cnt[ch-'a']-- // cnt 数组存储的是 s1 中每个字符的欠账（还需要多少才能匹配) // 即使找到符合的序列cnt[<N>]均<=0, 因此需要处理 cnt[x] > 0的情况
	}
	left := 0
	for right, ch := range s2 { // 使用双指针 left 和 right 维护一个滑动窗口
		x := ch - 'a'
		cnt[x]++

		//fmt.Printf("b:  %c, %d-%d, %v\n", ch, left, right, cnt)

		// 非 s1 字符: 会立刻使 left 右移，窗口内始终只保留可能匹配 s1 的字符, 即出现其他字符, 放弃当前窗口. 如果此时[left,right]中包含s1字符, left 右移后其重新重新变成欠账
		// 对 s1 字符：只有当前窗口内该字符的出现次数超过 s1 中的次数时（即 cnt[x] > 0），才会触发循环
		for cnt[x] > 0 { // 如果 cnt[x] > 0，说明当前字符 ch 在 s2 中的出现次数超过了 s1 中的次数，需要移动 left 缩小窗口，直到 cnt[x] 不再大于 0（即去掉多余的字符）
			cnt[s2[left]-'a']--
			left++
		}

		// fmt.Printf("a: %c, %d-%d, %v\n", ch, left, right, cnt)
		if right-left+1 == n { // 检查当前窗口大小 right - left + 1 是否等于 n（即窗口大小是否等于 s1 的长度）
			return true
		}
	}
	return false
}
