package demo

import (
	"al/leetcode"
	"testing"
)

type TreeNode = leetcode.TreeNode

func TestGenerateTrees(t *testing.T) {

}

/*
二叉搜索树关键的性质是根节点的值大于左子树所有节点的值，小于右子树所有节点的值，且左子树和右子树也同样为二叉搜索树.
因此在生成所有可行的二叉搜索树的时候，假设当前序列长度为 n，如果枚举根节点的值为 i，
那么根据二叉搜索树的性质可以知道左子树的节点值的集合为 [1…i−1]，右子树的节点值的集合为 [i+1…n].
而左子树和右子树的生成相较于原问题是一个序列长度缩小的子问题，因此可以想到用回溯的方法来解决这道题目

定义 generateTrees(start, end) 函数表示当前值的集合为 [start,end]，返回序列 [start,end] 生成的所有可行的二叉搜索树.
按照上文的思路，考虑枚举 [start,end] 中的值 i 为当前二叉搜索树的根，那么序列划分为了 [start,i−1] 和 [i+1,end] 两部分.
递归调用这两部分，即 generateTrees(start, i - 1) 和 generateTrees(i + 1, end)，获得所有可行的左子树和可行的右子树，
那么最后一步只要从可行左子树集合中选一棵，再从可行右子树集合中选一棵拼接到根节点上，并将生成的二叉搜索树放入答案数组即可
*/
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTrees := []*TreeNode{}
	// 枚举可行根节点
	for i := start; i <= end; i++ {
		// 获得所有可行的左子树集合
		leftTrees := helper(start, i-1)
		// 获得所有可行的右子树集合
		rightTrees := helper(i+1, end)
		// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currTree := &TreeNode{i, nil, nil}
				currTree.Left = left
				currTree.Right = right
				allTrees = append(allTrees, currTree)
			}
		}
	}
	return allTrees
}
