/*
437.中 路径总和 III

给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

示例 1：

输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
输出：3
解释：和等于 8 的路径有 3 条，如图所示。
示例 2：

输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：3

提示:

二叉树的节点个数的范围是 [0,1000]
-109 <= Node.val <= 109
-1000 <= targetSum <= 1000
*/
package leetcode

import (
	"testing"

	. "al/leetcode"
)

func TestPathSum(t *testing.T) {

}

// preSum[j]-preSum[i-1] = target
func pathSum(root *TreeNode, targetSum int) (ans int) {
	// 存储前缀和及其出现的次数。初始时，preSum 包含 {0: 1}，表示前缀和为 0 的路径出现了 1 次（即空路径）
	preSum := map[int64]int{0: 1}

	var dfs func(*TreeNode, int64)
	dfs = func(node *TreeNode, curr int64) {
		if node == nil {
			return
		}
		curr += int64(node.Val)
		ans += preSum[curr-int64(targetSum)] // 检查curr-int64(targetSum)对于的前缀和是否存在, 存在的话刚好加上相应的次数
		preSum[curr]++                       // 当前前缀和加入哈希表

		// 递归遍历左右子树
		dfs(node.Left, curr)
		dfs(node.Right, curr)

		preSum[curr]-- // 在回溯时，将当前路径和 curr 的计数减一，以避免影响其他路径的计算
		return
	}
	dfs(root, 0)
	return
}
