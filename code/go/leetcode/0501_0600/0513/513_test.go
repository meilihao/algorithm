/*
513.中 找树左下角的值

给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。

假设二叉树中至少有一个节点。

示例 1:

输入: root = [2,1,3]
输出: 1
示例 2:

输入: [1,2,3,4,null,5,6,null,null,7]
输出: 7

提示:

二叉树的节点个数的范围是 [1,104]
-231 <= Node.val <= 231 - 1
*/
package leetcode

import (
	"fmt"
	"testing"

	. "al/leetcode"
)

func TestFindBottomLeftValue(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	fmt.Println(findBottomLeftValue(root))
	fmt.Println("---")
	fmt.Println(findBottomLeftValue2(root))
}

// 从右到左遍历每一层的节点, 最后一个节点的值就是最底层最左边节点的值
// best
func findBottomLeftValue(root *TreeNode) (ans int) {
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
		ans = node.Val
	}
	return
}

// dfs
func findBottomLeftValue2(root *TreeNode) (curVal int) {
	curHeight := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}

		height++
		//由于是“先左后右”，每当达到某一深度的“第一个节点”，一定是该层最左边的
		dfs(node.Left, height)
		dfs(node.Right, height)
		//fmt.Println(node.Val, height, curHeight, height > curHeight)
		if height > curHeight {
			curHeight = height
			curVal = node.Val
		}
	}
	dfs(root, 0)
	return
}
