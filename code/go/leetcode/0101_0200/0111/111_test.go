/*
111.简 二叉树的最小深度

给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。

示例 1：

输入：root = [3,9,20,null,null,15,7]
输出：2
示例 2：

输入：root = [2,null,3,null,4,null,5,null,6]
输出：5

提示：

树中节点数的范围在 [0, 105] 内
-1000 <= Node.val <= 1000
*/
package main

import (
	"testing"

	. "al/leetcode"
)

// 思路: 叶子节点更新min
// 1. bfs
// 2. dfs
func TestMinDepth(t *testing.T) {

}

func minDepth(root *TreeNode) int {
	switch {
	case root == nil:
		return 0
	case root.Left == nil: // 因为左分支没了, 最小深度取决于右分支
		return 1 + minDepth(root.Right)
	case root.Right == nil:
		return 1 + minDepth(root.Left)
	default:
		return 1 + min(minDepth(root.Left), minDepth(root.Right))
	}
}
