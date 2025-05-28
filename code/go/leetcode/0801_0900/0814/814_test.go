/*
814.中 二叉树剪枝

给你二叉树的根结点 root ，此外树的每个结点的值要么是 0 ，要么是 1 。

返回移除了所有不包含 1 的子树的原二叉树。

节点 node 的子树为 node 本身加上所有 node 的后代。

示例 1：

输入：root = [1,null,0,0,1]
输出：[1,null,0,null,1]
解释：
只有红色节点满足条件“所有不包含 1 的子树”。 右图为返回的答案。
示例 2：

输入：root = [1,0,1,0,0,0,1]
输出：[1,null,1,null,1]
示例 3：

输入：root = [1,1,0,1,1,0,1,0]
输出：[1,1,0,1,1,null,1]

提示：

树中节点的数目在范围 [1, 200] 内
Node.val 为 0 或 1
*/
package leetcode

import (
	"testing"

	. "al/leetcode"
)

func TestPruneTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	n1 := pruneTree(root)
	PrintTreeByBfs(n1)

	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 0,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	n2 := pruneTree(root2)
	PrintTreeByBfs(n2)

	root3 := &TreeNode{
		Val: 0,
		Right: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 0,
			},
		},
	}

	n3 := pruneTree(root3)
	PrintTreeByBfs(n3)
}

// best
// 本质: 递归的后续遍历
func pruneTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = pruneTree2(root.Left)
	root.Right = pruneTree2(root.Right)

	// 左子树为空，右子树为空，当前节点的值为 0，同时满足时，才表示以当前节点为根的原二叉树的所有节点都为 0，需要将这棵子树移除，返回空
	// root.Left == root.Right只能是同时为nil
	if root.Left == root.Right && root.Val == 0 {
		return nil
	}
	return root
}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	SumTree(root)

	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}

	return root
}

func SumTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lSum := SumTree(root.Left)
	rSum := SumTree(root.Right)
	if lSum == 0 {
		root.Left = nil
	}
	if rSum == 0 {
		root.Right = nil
	}

	return lSum + rSum + root.Val
}
