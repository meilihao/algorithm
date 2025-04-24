package leetcode

import "testing"

func TestIsBalanced(t *testing.T) {

}

func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 { // 不平衡
		return -1
	}
	return max(leftHeight, rightHeight) + 1 // 高度
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
