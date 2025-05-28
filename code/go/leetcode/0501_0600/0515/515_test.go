/*
515.中 在每个树行中找最大值

给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。

示例1：

输入: root = [1,3,2,5,3,null,9]
输出: [1,3,9]
示例2：

输入: root = [1,2,3]
输出: [1,3]

提示：

二叉树的节点个数的范围是 [0,104]
-231 <= Node.val <= 231 - 1
*/
package leetcode

import (
	"fmt"
	"testing"

	. "al/leetcode"
)

func TestLargestValues(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	fmt.Println(largestValues(root))
}

func largestValues(root *TreeNode) (ans []int) {
	var dfs func(*TreeNode, int)

	// 用 curHeight 来标记遍历到的当前节点的高度
	// 令 ans下标 = 树高度-1
	dfs = func(node *TreeNode, curHeight int) {
		if node == nil {
			return
		}
		// 如果当前深度 curHeight 是第一次访问（也就是说还没有创建该层对应的 ans[curHeight]），则插入当前节点的值
		if curHeight == len(ans) {
			ans = append(ans, node.Val)
		} else {
			ans[curHeight] = max(ans[curHeight], node.Val)
		}
		dfs(node.Left, curHeight+1)
		dfs(node.Right, curHeight+1)
	}
	dfs(root, 0)
	return
}
