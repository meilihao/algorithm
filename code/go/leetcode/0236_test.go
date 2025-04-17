// 236. 二叉树的最近公共祖先
package leetcode

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	// p := root.Left
	// q := root.Right

	p := root.Left
	q := root.Left.Right.Right

	r := lowestCommonAncestor(root, p, q)

	fmt.Println(r.Val)
}

// 时间复杂度是 O(n)
// lowestCommonAncestor返回值是「最近公共祖先的候选项」
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		// if root != nil {
		// 	fmt.Println("m:", root.Val)
		// }
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)  // 在左分支找得到p/q
	r := lowestCommonAncestor(root.Right, p, q) // 在右分支找得到p/q

	// if l != nil {
	// 	fmt.Println("l:", l.Val)
	// }
	// if r != nil {
	// 	fmt.Println("r:", l.Val)
	// }

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
