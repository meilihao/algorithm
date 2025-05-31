/*
1122.简 数组的相对排序

给你两个数组，arr1 和 arr2，arr2 中的元素各不相同，arr2 中的每个元素都出现在 arr1 中。

对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。

示例 1：

输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
输出：[2,2,2,1,4,3,3,9,6,7,19]
示例  2:

输入：arr1 = [28,6,22,8,44,17], arr2 = [22,28,8,6]
输出：[22,28,8,6,17,44]

提示：

1 <= arr1.length, arr2.length <= 1000
0 <= arr1[i], arr2[i] <= 1000
arr2 中的元素 arr2[i]  各不相同
arr2 中的每个元素 arr2[i] 都出现在 arr1 中
*/
package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRelativeSortArray(t *testing.T) {
	arr1 := []int{28, 6, 22, 8, 44, 17}
	arr2 := []int{22, 28, 8, 6}
	fmt.Println(reflect.DeepEqual(relativeSortArray(arr1, arr2), []int{22, 28, 8, 6, 17, 44}))
}

// 计数排序
// 时间复杂度：O(m+n+upper)，其中 m 和 n 分别是数组 arr1, arr2的长度, upper 是数组 arr1中的最大值
func relativeSortArray(arr1 []int, arr2 []int) []int {
	upper := 0 // 统计 arr1 中最大值 (upper), 用于创建frequency 数组
	for _, v := range arr1 {
		if v > upper {
			upper = v
		}
	}
	frequency := make([]int, upper+1)
	for _, v := range arr1 {
		frequency[v]++ // 统计频率
	}

	ans := make([]int, 0, len(arr1))
	for _, v := range arr2 { // 按 arr2 顺序填充结果 (ans)
		for ; frequency[v] > 0; frequency[v]-- {
			ans = append(ans, v)
		}
	}
	for v, freq := range frequency {
		for ; freq > 0; freq-- { // 只要当前数字 v 还有剩余频率. 桶已排序可直接追加
			ans = append(ans, v)
		}
	}
	return ans
}
