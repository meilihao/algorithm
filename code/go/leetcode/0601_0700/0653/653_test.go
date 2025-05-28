/*
653.简单 两数之和 IV - 输入二叉搜索树

给定一个二叉搜索树 root 和一个目标结果 k，如果二叉搜索树中存在两个元素且它们的和等于给定的目标结果，则返回 true。

示例 1：

输入: root = [5,3,6,2,4,null,7], k = 9
输出: true
示例 2：

输入: root = [5,3,6,2,4,null,7], k = 28
输出: false

提示:

二叉树的节点个数的范围是  [1, 104].
-104 <= Node.val <= 104
题目数据保证，输入的 root 是一棵 有效 的二叉搜索树
-105 <= k <= 105
*/
package leetcode

import (
	"testing"

	. "al/leetcode"
)

func TestFindTarget(t *testing.T) {

}

func findTarget(root *TreeNode, k int) bool {
	set := map[int]struct{}{}
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if _, ok := set[k-node.Val]; ok {
			return true
		}
		set[node.Val] = struct{}{}
		return dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}
