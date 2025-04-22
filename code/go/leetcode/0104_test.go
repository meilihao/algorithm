// 104. 二叉树的最大深度
// 思路:
// 1. bfs
// 2. dfs
package leetcode

import "testing"

func TestMaxDepth104(t *testing.T) {

}

// 左子树和右子树的最大深度 l 和 r，那么该二叉树的最大深度即为max(l,r)+1
func maxDepth104(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth104(root.Left), maxDepth104(root.Right))
}

// bfs
func maxDepth104_2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}
