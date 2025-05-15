/*
287.中 寻找重复数

给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。

你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。


示例 1：

输入：nums = [1,3,4,2,2]
输出：2
示例 2：

输入：nums = [3,1,3,4,2]
输出：3
示例 3 :

输入：nums = [3,3,3,3,3]
输出：3


提示：

1 <= n <= 105
nums.length == n + 1
1 <= nums[i] <= n
nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次
*/

package lcof

import (
	"testing"
)

func TestFindDuplicate2(t *testing.T) {
	nums := []int{1, 1, 6, 3, 4, 5, 2}

	if tmp := findDuplicate(nums); tmp == -1 {
		t.Errorf(" err get: %d", tmp)
	}
}

/*
利用了数组元素的值作为索引来标记访问过的数字，通过将对应位置的数值取负来标记该数字已经被访问过

限制:
1. 数组元素必须在1到n之间（n为数组长度-1）; 如果数组中可能有0或负数，需要特殊处理
1. 必须恰好有一个重复数字（可以有多次重复）
*/
func findDuplicate(nums []int) int {
	for i := range nums { // 遍历数组中的每个元素
		k := nums[i] // 获取当前元素值
		if k < 0 {   // 如果当前元素是负数（可能已被标记）
			k = -k // 取其绝对值
		}
		if nums[k] < 0 { // 检查nums[k]是否为负数（表示之前已经被标记过）
			return k // 如果为负，说明k是重复的数字
		}
		nums[k] = -nums[k] // 将nums[k]标记为负数，表示数字k已经被访问过
	}
	return -1 // 如果没有找到重复（根据题意应该总是有重复）
}
