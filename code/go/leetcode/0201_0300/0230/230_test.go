/*
230.中 二叉搜索树中第 K 小的元素

给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（从 1 开始计数）。

示例 1：

输入：root = [3,1,4,null,2], k = 1
输出：1
示例 2：

输入：root = [5,3,6,2,4,null,null,1], k = 3
输出：3

提示：

树中的节点数为 n 。
1 <= k <= n <= 104
0 <= Node.val <= 104
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 树相关的题目，第一眼就想到递归求解，左右子树分别遍历.
// 联想到二叉搜索树的性质，root 大于左子树，小于右子树，
// 如果左子树的节点数目等于 K-1，那么 root 就是结果，否则如果左子树节点数目小于 K-1，那么结果必然在右子树，否则就在左子树.
// 因此在搜索的时候同时返回节点数目，跟 K 做对比，就能得出结果了。
func kthSmallest(root *TreeNode, k int) int {
	left := countLeaf(root.Left)

	if left == k-1 {
		return root.Val
	} else if left > k-1 {
		return kthSmallest(root.Left, k)
	} else {
		return kthSmallest(root.Right, k-left-1) // k-(left+1)即仅检查右子树, 1是root
	}
}

func countLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countLeaf(root.Left) + countLeaf(root.Right) + 1
}

// todo stack
