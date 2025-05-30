/*
373.中 查找和最小的 K 对数字

给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。

定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。

请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。

示例 1:

输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
输出: [1,2],[1,4],[1,6]
解释: 返回序列中的前 3 对数：

	[1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]

示例 2:

输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
输出: [1,1],[1,1]
解释: 返回序列中的前 2 对数：

	[1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]

提示:

1 <= nums1.length, nums2.length <= 105
-109 <= nums1[i], nums2[i] <= 109
nums1 和 nums2 均为 升序排列
1 <= k <= 104
k <= nums1.length * nums2.length
*/
package leetcode

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestSolution(t *testing.T) {
	s := Constructor([]int{1, 3, 2})
	for i := 0; i < 6; i++ {
		fmt.Println(s.PickIndex())
	}
}

/*
假设 w = [1, 3, 2]：

Constructor：

计算前缀和：pre = [1, 4, 6]

PickIndex：

- 生成 x ∈ [1, 6]：

  - x = 1 → sort.SearchInts([1,4,6], 1) → 返回 0
  - x = 2 → 查找 2 在 [1,4,6] 中的位置 → 1（因为 2 > 1 且 2 ≤ 4）
  - x = 3 → 同上 → 1
  - x = 4 → 返回 1
  - x = 5 → 查找 5 → 2（5 > 4 且 5 ≤ 6）
  - x = 6 → 返回 2

- 因此：

  - 索引 0 被选中的概率：1/6。
  - 索引 1 被选中的概率：3/6 = 1/2。
  - 索引 2 被选中的概率：2/6 = 1/3。
*/
type Solution struct {
	// **可以将w[i]想象线段长度, w就是所有线段拼成的长度, 随机取线段上的一点, 返回所在线段的下标**
	// 前缀和的作用: 将权重转换为累积和，便于将随机数映射到对应的索引
	// (pre[i]-pre[i-1])/total = w[i] / sum(w)
	pre []int
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{w}
}

func (s *Solution) PickIndex() int {
	x := rand.Intn(s.pre[len(s.pre)-1]) + 1 // 生成一个范围在 [1, 总和] 之间的随机整数 x. pre[len(pre)-1] 就是所有权重的总和
	// 由于前缀和数组是单调递增的，可以使用二分查找（sort.SearchInts）来找到第一个大于或等于 x 的前缀和的索引, 这个索引就是应该选择的随机索引
	return sort.SearchInts(s.pre, x)
}
