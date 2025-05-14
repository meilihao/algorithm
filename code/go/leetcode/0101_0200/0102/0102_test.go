/*
102.中 二叉树的层序遍历

给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

示例 1：

输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]
*/

package leetcode

import (
	"testing"

	. "al/leetcode"
)

// 思路:
// 1. bfs
// 2. dfs 但需要传递level(层数)
func TestLevelOrder(t *testing.T) {

}

// bfs
func levelOrder(root *TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	enqueue(root)

	var qSize int
	var n *TreeNode
	for len(queue) > 0 {
		qSize = len(queue) // 当前level的节点数量
		var list []int

		for qSize > 0 {
			n = dequeue()
			list = append(list, n.Val)
			if n.Left != nil {
				enqueue(n.Left)
			}

			if n.Right != nil {
				enqueue(n.Right)
			}

			qSize--
		}

		res = append(res, list)
	}

	return res
}

var queue []*TreeNode

func enqueue(n *TreeNode) {
	queue = append(queue, n)
}

func dequeue() *TreeNode {
	if len(queue) == 0 {
		return nil
	}

	n := queue[0]
	queue = queue[1:]
	return n
}

// dfs
func levelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		// 出现了新的 level
		if level >= len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}
