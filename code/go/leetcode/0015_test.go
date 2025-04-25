// 15. 三数之和
// 思路:
// 1. a+b = -c , 两层嵌套 + map => O(n^2)
// 2. sort + find(夹逼) => O(n^2) 但不需要其他空间, 但sort改了输入
package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	//nums := []int{0, 0, 0, 0}
	nums := []int{-1, 0, 1, 2, -1, -4}
	//nums := []int{-2, 0, 1, 1, 2}

	fmt.Println(threeSum(nums))
}

// // c = -a-b
// // res 去重麻烦废弃
// func threeSum(nums []int) [][]int {
// 	if len(nums) < 3 {
// 		return nil
// 	}

// 	res := make([][]int, 0)

// 	for i, a := range nums[:len(nums)-2] { // 查找a, 必定是nums[:len(nums)-2]个, 因为其他两个是b,c
// 		if i >= 1 && a == nums[i-1] { // 跳过重复的a
// 			continue
// 		}

// 		m := make(map[int]bool)

// 		for _, b := range nums[i+1:] { // 在a的右边查找b
// 			if m[b] { // 已出现过-c
// 				res = append(res, []int{a, -a - b, b})
// 			} else {
// 				m[-a-b] = true // 存储-c
// 			}
// 		}
// 	}

// 	// todo res去重

// 	return res
// }

// todo sort + find
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)

	res := make([][]int, 0)

	var l, r, s int
	for k := range nums[:len(nums)-2] { // 固定最小数的索引k, k的范围是nums[:len(nums)-2]
		if nums[k] > 0 {
			break // nums[j] > nums[i] > nums[k] > 0
		}
		if k > 0 && nums[k] == nums[k-1] { // 跳过重复的数字, 因为会导致结果重复，所以应该跳过. 即nums[k]固定, nums[l]和nums[r]变化而已
			continue
		}

		l, r = k+1, len(nums)-1 // k是最小数的索引, 那么其他两数应均在k的右边

		for l < r {
			s = nums[k] + nums[l] + nums[r]

			if s < 0 {
				l++
			} else if s > 0 {
				r--
			} else {
				res = append(res, []int{nums[k], nums[l], nums[r]}) // 匹配成功

				for l < r && nums[l] == nums[l+1] { // 跳过重复值
					l++
				}
				for l < r && nums[r] == nums[r-1] { // 跳过重复值
					r--
				}

				// 开始尝试下一个组合
				l++
				r--
			}
		}
	}

	return res
}
