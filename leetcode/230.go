// 230. 二叉搜索树中第K小的元素
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
		return kthSmallest(root.Right, k-left-1)
	}
}

func countLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countLeaf(root.Left) + countLeaf(root.Right) + 1
}
