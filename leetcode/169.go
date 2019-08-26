// 169. 求众数
// 思路:
// 1. map
// 2. sort后遍历或者根据题意众数必定存在即直接nums[nums.length/2]即可
// 3. Boyer-Moore 投票算法 : 把众数记为 +1+1+1 ，把其他数记为 −1-1−1 ，将它们全部加起来，显然和大于 0
package main

import "fmt"

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}

	fmt.Println(majorityElement(nums))
}

// 寻找数组中超过一半的数字，这意味着数组中其他数字出现次数的总和都是比不上这个数字出现的次数.

// 即如果把 该众数记为 +1 ，把其他数记为 −1 ，将它们全部加起来，和是大于 0 的.

// 所以可以这样操作：

//     设置两个变量 candidate 和 count，candidate 用来保存数组中遍历到的某个数字，count 表示当前数字的出现次数
//     遍历整个数组
//     如果数字与之前 candidate 保存的数字相同，则 count 加 1
//     如果数字与之前 candidate 保存的数字不同，则 count 减 1
//     如果出现次数 count 变为 0 ，candidate 进行变化，保存为当前遍历的那个数字，并且同时把 count 重置为 1
//     遍历完数组中的所有数字即可得到结果
func majorityElement(nums []int) int {
	candidate, count := 0, 0 // 可优化: 一开始 candidate 保存为数组中的第一个数字，count 为 1, 从第二个数开始遍历

	for _, v := range nums {
		if count == 0 {
			candidate = v
		}

		if v == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
