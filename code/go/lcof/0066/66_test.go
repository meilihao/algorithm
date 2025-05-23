/*
LCR 191.中 按规则计算统计结果

为了深入了解这些生物群体的生态特征，你们进行了大量的实地观察和数据采集。数组 arrayA 记录了各个生物群体数量数据，其中 arrayA[i] 表示第 i 个生物群体的数量。请返回一个数组 arrayB，该数组为基于数组 arrayA 中的数据计算得出的结果，其中 arrayB[i] 表示将第 i 个生物群体的数量从总体中排除后的其他数量的乘积。

示例 1：

输入：arrayA = [2, 4, 6, 8, 10]
输出：[1920, 960, 640, 480, 384]

提示：

所有元素乘积之和不会溢出 32 位整数
arrayA.length <= 100000
*/
/*
66：构建乘积数组

给定一个数组A[0, 1, …, n-1]，请构建一个数组B[0, 1, …, n-1]，其
中B中的元素B[i] =A[0]×A[1]×… ×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
*/
package main

import (
	"fmt"
	"testing"
)

// 如果可以使用除法b[j] = (A[0]*...A[n])/A[j]
func TestReverseWords(t *testing.T) {
	nums := []int{2, 4, 6, 8, 10}
	fmt.Println(statisticalResult(nums))
}

// B[i] =(A[0]×A[1]×… ×A[i-1])×(A[i+1]×…×A[n-1])
func statisticalResult(a []int) []int {
	n := len(a)
	b := make([]int, n)
	p := 1
	// 从左到右计算前缀积
	for i := 0; i < n; i++ {
		b[i] = p
		p *= a[i] // 更新 p，将其乘以 a[i]，以便在下一次迭代中，p 成为 a[i+1] 左侧所有元素的乘积
	}
	p = 1
	// 从右到左的后缀积
	// b[i] 此时已经包含了其左侧所有元素的乘积（来自第一次遍历）
	for i := n - 1; i >= 0; i-- {
		b[i] *= p
		p *= a[i] // 更新 p，将其乘以 a[i]，以便在下一次迭代中，p 成为 a[i-1] 右侧所有元素的乘积
	}
	return b
}
