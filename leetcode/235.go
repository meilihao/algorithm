// 235. 二叉搜索树的最近公共祖先
package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return helper(root, p.Val, q.Val)
}

func helper(root *TreeNode, p, q int) *TreeNode {
	if p < root.Val && q < root.Val { // p,q 均在左子树
		return helper(root.Left, p, q)
	} else if root.Val < p && root.Val < q { // p,q 均在右子树
		return helper(root.Right, p, q)
	}
	return root // p,q 在root两边
}
