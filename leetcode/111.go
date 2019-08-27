// 111. 二叉树的最小深度
// 思路: 叶子节点更新min
// 1. bfs
// 2. dfs
package main

func main() {

}

func minDepth(root *TreeNode) int {
	switch {
	case root == nil:
		return 0
	case root.Left == nil: // 因为左分支没了, 最小深度取决于右分支
		return 1 + minDepth(root.Right)
	case root.Right == nil:
		return 1 + minDepth(root.Left)
	default:
		return 1 + min(minDepth(root.Left), minDepth(root.Right))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
