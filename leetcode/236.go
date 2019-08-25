// 236. 二叉树的最近公共祖先
package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度是 O(n)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	// 根据题意(p、q 为不同节点且均存在于给定的二叉树中)， l 和 r 不可能同时为 nil
	if l != nil && r != nil {
		// 此时 p 和 q 分别在 root.Left 和 root.Right 中
		return root
	}
	if l == nil {
		// 此时 p 和 q 在 root.Right 中
		return r
	}
	// 此时 p 和 q 在 root.Left 中
	return l
}
