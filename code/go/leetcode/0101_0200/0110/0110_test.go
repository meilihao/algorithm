/*
110.简 平衡二叉树

给定一个二叉树，判断它是否是 平衡二叉树

示例 1：

输入：root = [3,9,20,null,null,15,7]
输出：true
示例 2：

输入：root = [1,2,2,3,3,null,null,4,4]
输出：false
示例 3：

输入：root = []
输出：true
*/
package leetcode

import (
	"testing"

	. "al/leetcode"
)

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
