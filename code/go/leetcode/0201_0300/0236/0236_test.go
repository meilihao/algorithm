/*
236.中 二叉树的最近公共祖先

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

示例 1：

输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
示例 2：

输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出：5
解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
示例 3：

输入：root = [1,2], p = 1, q = 2
输出：1

提示：

树中节点数目在范围 [2, 105] 内。
-109 <= Node.val <= 109
所有 Node.val 互不相同 。
p != q
p 和 q 均存在于给定的二叉树中。
*/
package leetcode

import (
	"fmt"
	"testing"

	. "al/leetcode"
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
