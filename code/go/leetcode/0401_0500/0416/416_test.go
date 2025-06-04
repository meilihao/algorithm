/*
416.中 分割等和子集

给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

示例 1：

输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。
示例 2：

输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。

提示：

1 <= nums.length <= 200
1 <= nums[i] <= 100
*/
package leetcode

import (
	"math/big"
	"testing"
)

func TestCanPartition(t *testing.T) {

}

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum%2 != 0 { // 能分成两个子数组
		return false
	}

	target := sum / 2
	if max > target { // 如果数组中最大的数字比目标和还大，则不可能达到目标和
		return false
	}

	dp := make([]bool, target+1) // dp[j] 表示：是否存在 nums 中的某个子集，其元素和恰好等于 j
	dp[0] = true                 // 和为 0 总是可以达到的（通过选择空集）
	for i := 0; i < n; i++ {
		v := nums[i]

		// 从 target 向下遍历到 v。
		// 为什么要从大到小遍历 j？
		// 这是 0/1 背包问题的关键。
		// 如果从 dp[v] 向上遍历到 dp[target]，当计算 dp[j] 时，
		// dp[j-v] 可能已经使用了当前的 v 来更新，导致 v 被重复使用。
		// 而从大到小遍历时，dp[j-v] 的值是在当前 v 被考虑之前就确定的，
		// 确保了每个数字只被使用一次。
		for j := target; j >= v; j-- {
			// dp[j] = true 表示：当前和 j 可以通过两种方式达到：
			// 1. 不包含当前数字 v：这取决于 dp[j] 在考虑 v 之前的状态（即 dp[j] 本身）。
			// 2. 包含当前数字 v：这取决于 dp[j-v] 是否为 true。如果 dp[j-v] 为 true，
			//    那么加上当前的 v 就可以达到 j。
			dp[j] = dp[j] || dp[j-v]
		}
	}
	return dp[target]
}

func canPartition2(nums []int) bool {
	// bit状态叠加
	s := 0
	for _, x := range nums {
		s += x
	}
	if s%2 != 0 {
		return false
	}
	s /= 2 // 此时s 表示 target
	// dp 的第 k 位（从右往左数，最低位是第 0 位）为 1，表示存在一个子集，其元素和恰好等于 k
	// 初始时 dp 为 1（二进制表示为 ...0001）。这意味着只有和为 0 的子集（空集）是可达的，对应 dp[0] = true
	dp := big.NewInt(1) // 初始化 dp 为一个大整数，初始值为 1
	p := new(big.Int)   // 暂存dp左移的结果
	for _, x := range nums {
		// dp.Or(dp, p.Lsh(dp, uint(x)))
		// 这一行代码的核心逻辑是：
		// 对于当前数字 x，我们有两种选择：
		// 1. 不选择 x：那么能达到的和的状态保持不变，由当前 dp 存储。
		// 2. 选择 x：如果之前能达到和 sum' (即 dp 的 sum' 位为 1)，
		//    那么现在就能达到 sum' + x (即 dp 的 sum' + x 位应该为 1)。
		//
		// p.Lsh(dp, uint(x)) 的作用：
		// 将当前的 dp 值左移 x 位。如果 dp 的第 k 位是 1，
		// 那么左移 x 位后，新的 big.Int 的第 k+x 位将是 1。
		// 这就表示：如果之前能凑出和 k，现在加上 x 就能凑出和 k+x。
		//
		// dp.Or(dp, ...) 的作用：
		// 将原来的 dp (不选择 x 时的状态) 和 左移后的 dp (选择 x 时的状态) 进行按位或运算。
		// 这样做的好处是，如果某个和 k 既可以通过不选 x 达到，也可以通过选 x 达到，
		// 那么或运算会保留其为 1 的状态。
		// 确保了 dp 的第 k 位为 1，只要有任何一种方式能凑出和 k。

		dp.Or(dp, p.Lsh(dp, uint(x))) // 用或运算累计了原始不选的状态和或上选择当前的状态
	}

	// dp.Bit(s) 方法返回 big.Int 中第 s 位的值（0 或 1）
	return dp.Bit(s) == 1 // 判断 dp 的第 s 位是否为 1
}
