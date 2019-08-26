// 104. 二叉树的最大深度
// 思路:
// 1. bfs
// 2. dfs
package main

func main() {

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
