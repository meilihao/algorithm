package leetcode

import (
	"fmt"
	"testing"
)

// 求nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2[j] 的 下一个更大元素即求nums2 中每个元素右侧下一个更大的元素
func TestNextGreaterElement(t *testing.T) {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}

	// nums1 := []int{2, 4}
	// nums2 := []int{1, 2, 3, 4}

	//fmt.Println(nextGreaterElement(nums1, nums2))
	fmt.Println(nextGreaterElement2(nums1, nums2))
}

// best
// 因为 nums1 是 nums2 的子集，所以可以先遍历一遍 nums2，并构造单调递增栈，求出 nums2 中每个元素右侧下一个更大的元素。然后将其存储到哈希表中。然后再遍历一遍
// nums1，从哈希表中取出对应结果，存放到答案数组中
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return nil
	}

	st := []int{}
	m := map[int]int{} // 元素右侧下一个更大的元素

	for _, v := range nums2 {
		for len(st) > 0 && v > st[len(st)-1] {
			m[st[len(st)-1]] = v
			st = st[:len(st)-1]
		}

		st = append(st, v)
	}

	var target int
	var exist bool
	for i, v := range nums1 {
		if target, exist = m[v]; exist {
			nums1[i] = target
		} else {
			nums1[i] = -1
		}
	}

	return nums1
}

// 从栈底到栈顶的元素是单调递减的=单调递增栈
func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return nil
	}

	st := []int{}
	m := map[int]int{}

	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(st) > 0 && num >= st[len(st)-1] {
			st = st[:len(st)-1]
		}
		fmt.Println(st)
		if len(st) > 0 { // num < st[len(st)-1], 符合条件
			m[num] = st[len(st)-1]
		} // else { // 找不到
		// 	m[num] = -1
		// }
		st = append(st, num)
	}

	var target int
	var exist bool
	for i := range nums1 {
		if target, exist = m[i]; exist {
			nums1[i] = target
		} else {
			nums1[i] = -1
		}
	}

	return nums1
}
