// 102. 二叉树的层次遍历
// 思路:
// 1. bfs
// 2. dfs 但需要传递level(层数)
package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
func levelOrder(root *TreeNode) [][]int {
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
