/*
	Two Sum 两数之和

	for 算法效率
*/
package main

import "fmt"
import "sort"

func main() {
	nums := []int{3, 2, 4} //{2, 7, 11, 15}
	target := 6

	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSum21(nums, target))
	fmt.Println(twoSum22(nums, target))
	fmt.Println(twoSum3(nums, target))
}

// 尝试所有组合, 类似与冒泡排序的过程
// O(n^2),
func twoSum(nums []int, target int) []int {
	var a []int
	n := len(nums)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				a = append(a, i, j)
			}
		}
	}

	return a
}

// 用哈希表
// O(n)
func twoSum21(nums []int, target int) []int {
	var a []int
	n := len(nums)

	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}

	for i := 0; i < n-1; i++ { // 使用n-1,是因为数据已全部在m中
		if j, ok := m[target-nums[i]]; ok && i != j { // i != j, 原因也是数据全在m中
			a = append(a, i, j)

			return a
		}
	}

	return nil
}

// 用哈希表, 比twoSum21少用一次for循环,但返回的下标不是有序了, 符合条件的下标必定在其另一个值的后面
// O(n)
func twoSum22(nums []int, target int) []int {
	var a []int
	n := len(nums)
	m := make(map[int]int)

	for i := 0; i < n; i++ { // 使用n,是因为数据未全部在m中
		if j, ok := m[target-nums[i]]; ok { // j是之前存入m的i,因此 j < i
			a = append(a, j, i)

			return a
		}

		m[nums[i]] = i
	}

	return nil
}

// 优化执行时间
func twoSum221(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}

		m[v] = i
	}

	return nil
}

func twoSum222(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, v := range nums {
		if j, ok := m[v]; ok {
			return []int{j, i}
		}

		m[target-nums[i]] = i
	}

	return nil
}

// 使用夹逼方法:原理是因为数组是有序的，那么假设当前结果比target大，那么左端序号右移只会使两个数的和更大，反之亦然. 所以每次只会有一个选择，从而实现线性就可以求出结果.
// 返回结果和原下标不同(因为排序过)
// O(nlogn+n)=O(nlogn)，空间复杂度取决于排序算法
func twoSum3(nums []int, target int) []int {
	sort.Ints(nums)

	l, r, tmp := 0, len(nums)-1, 0

	for l < r {
		tmp = nums[l] + nums[r]

		if tmp == target {
			return []int{l, r}
		} else if tmp > target {
			r--
		} else {
			l++
		}
	}

	return nil
}
