package leetcode

import "testing"

func TestSortedArrayToBST(t *testing.T) {
	a := []int{-10, -3, 0, 5, 9}
	sortedArrayToBST(a)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 { // 叶子节点
		return nil
	}
	mid := (0 + len(nums) - 1) / 2           // 中间元素下标
	cur := new(TreeNode)                     // 创建当前节点
	cur.Val = nums[mid]                      // 获取当前节点值
	cur.Left = sortedArrayToBST(nums[0:mid]) // 获取左子树和右子树
	cur.Right = sortedArrayToBST(nums[mid+1:])
	return cur
}
