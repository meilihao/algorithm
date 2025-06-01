/*
873.中 最长的斐波那契子序列的长度

如果序列 X_1, X_2, ..., X_n 满足下列条件，就说它是 斐波那契式 的：

n >= 3
对于所有 i + 2 <= n，都有 X_i + X_{i+1} = X_{i+2}
给定一个严格递增的正整数数组形成序列 arr ，找到 arr 中最长的斐波那契式的子序列的长度。如果一个不存在，返回  0 。

（回想一下，子序列是从原序列 arr 中派生出来的，它从 arr 中删掉任意数量的元素（也可以不删），而不改变其余元素的顺序。例如， [3, 5, 8] 是 [3, 4, 5, 6, 7, 8] 的一个子序列）

示例 1：

输入: arr = [1,2,3,4,5,6,7,8]
输出: 5
解释: 最长的斐波那契式子序列为 [1,2,3,5,8] 。
示例 2：

输入: arr = [1,3,7,11,12,14,18]
输出: 3
解释: 最长的斐波那契式子序列有 [1,11,12]、[3,11,14] 以及 [7,11,18] 。

提示：

3 <= arr.length <= 1000
1 <= arr[i] < arr[i + 1] <= 10^9
*/
package leetcode

import (
	"testing"
)

func TestLenLongestFibSubseq(t *testing.T) {

}

func lenLongestFibSubseq(arr []int) (ans int) {
	n := len(arr)
	indices := make(map[int]int, n) // indices 存储数组元素值到其索引的映射, 用于快速查找某个值是否存在以及其索引
	for i, x := range arr {
		indices[x] = i
	}
	dp := make([][]int, n) // 创建一个二维数组 dp，dp[j][i] 表示以 arr[j] 和 arr[i] 结尾的斐波那契子序列的长度, 且arr[j] < arr[i]=> j < i, 因为arr是严格递增
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i, x := range arr { // 外层循环遍历 arr 数组
		// j: 斐波那契子序列的倒数第二个元素 arr[j] 的索引
		// 由于 arr[k] + arr[j] = arr[i]，且 arr[k] < arr[j]，那么 arr[j] 必须大于 arr[i] 的一半，因为如果 arr[j] <= x/2，那么 arr[k] < arr[j] <= x/2, arr[k] + arr[j] < arr[i], 这不符合斐波那契序列递增且前一个元素索引更小的特性
		// 倒序遍历 j 可以更快地找到满足条件的 k，或者更快地判断 arr[j]*2 > x 剪枝条件
		for j := n - 1; j >= 0 && arr[j]*2 > x; j-- {
			// 对于每个 i 和 j（j 从后向前遍历），检查是否存在 k 使得 arr[k] + arr[j] = arr[i]
			// 如果存在，则 dp[j][i] = dp[k][j] + 1（即当前子序列长度比之前的增加 1）
			if k, ok := indices[x-arr[j]]; ok { // 查找是否存在 x - arr[j]，即 k 是否存在
				dp[j][i] = max(dp[k][j]+1, 3) // 更新 dp[j][i]，即如果 k 存在，则将其长度加 1（至少长度为 3）
				ans = max(ans, dp[j][i])      //  更新最大答案
			}
		}
	}
	return
}
