// 98. 验证二叉搜索树
// 思路:
// 1. 中序遍历, 遍历结果的array是否是升序. 空间:O(n), 时间最坏O(n)
// 2. 递归  root.Val > max(root.Left) && root.Val < min(root.Right) 空间:O(n), 时间最坏O(n)
package leetcode

import "testing"

func TestIsValidBST(t *testing.T) {

}

func isValidBST(root *TreeNode) bool {
	min, max := -1<<63, 1<<63-1 // math.MinInt64, math.MaxInt64

	return recur(min, max, root)
}

func recur(min, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}

	// if min >= root.Val || root.Val >= max {
	// 	return false
	// }

	// return recur(min, root.Val, root.Left) && recur(root.Val, max, root.Right)

	// 效果同上
	return min < root.Val && root.Val < max &&
		recur(min, root.Val, root.Left) &&
		recur(root.Val, max, root.Right)
}
