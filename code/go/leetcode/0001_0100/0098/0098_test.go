/*
98.中 验证二叉搜索树

给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。


示例 1：



输入：root = [2,1,3]
输出：true
示例 2：



输入：root = [5,1,4,null,null,3,6]
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。
*/

package leetcode

import (
	"math"
	"testing"

	. "al/leetcode"
)

// 思路:
// 1. 中序遍历, 遍历结果的array是否是升序. 空间:O(n), 时间最坏O(n)
// 2. 递归  root.Val > max(root.Left) && root.Val < min(root.Right) 空间:O(n), 时间最坏O(n)
func TestIsValidBST(t *testing.T) {
	// root := &TreeNode{
	// 	Val: 4,
	// 	Left: &TreeNode{
	// 		Val: 2,
	// 		Left: &TreeNode{
	// 			Val: 1,
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 3,
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 5,
	// 		Right: &TreeNode{
	// 			Val: 6,
	// 		},
	// 	},
	// }

	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
	}

	if !isValidBST3(root2) {
		t.Errorf("invalid al")
	}
}

func isValidBST(root *TreeNode) bool {
	min, max := -1<<63, 1<<63-1 // math.MinInt64, math.MaxInt64

	return recur(min, max, root)
}

func recur(min, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}

	// if min >= root.Val || root.Val >= max {
	// 	return false
	// }

	// return recur(min, root.Val, root.Left) && recur(root.Val, max, root.Right)

	// 效果同上
	return min < root.Val && root.Val < max &&
		recur(min, root.Val, root.Left) &&
		recur(root.Val, max, root.Right)
}

// 二叉搜索树「中序遍历」得到的值构成的序列一定是升序的
// 在中序遍历的时候实时检查当前节点的值是否大于前一个中序遍历到的节点的值即可
func isValidBST2(root *TreeNode) bool {
	stack := []*TreeNode{}
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

func isValidBST3(root *TreeNode) bool {
	pre := math.MinInt64

	return RecursionMiddleorderTraversal(root, &pre)
}

// 中序遍历
func RecursionMiddleorderTraversal(node *TreeNode, preValue *int) bool {
	//如果当前节点为nil，则结束遍历
	if node == nil {
		return true
	}

	// 先递归遍历左子树
	if tmp := RecursionMiddleorderTraversal(node.Left, preValue); !tmp {
		return false
	}
	if node.Val <= *preValue {
		return false
	}

	*preValue = node.Val
	return RecursionMiddleorderTraversal(node.Right, preValue)
}
