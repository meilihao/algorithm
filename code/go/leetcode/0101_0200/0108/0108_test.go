/*
108.简 将有序数组转换为二叉搜索树

给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 平衡 二叉搜索树。

示例 1：

输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：

示例 2：

输入：nums = [1,3]
输出：[3,1]
解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。
*/
package leetcode

import (
	"testing"

	. "al/leetcode"
)

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
