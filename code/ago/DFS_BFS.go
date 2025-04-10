package ago

import (
	"fmt"
	"testing"

	"ago/helper"
)

type node2 struct {
	value int
	left  *node2
	right *node2
}

func TestDBFS(t *testing.T) {
	DBFS()
}

func DBFS() {
	root := buildTestTree()

	depthFirstSearch(root)
	fmt.Println("-----")
	breadthFirstSearch(root)
}

// DFS(Deep First Search）深度优先搜索
// 这里是非递归写法, 也可使用递归实现(常用,因为不用自行维护一个stack)
func depthFirstSearch[T int](root *helper.TreeNode[T]) {
	st := helper.NewStack[int](10)
	var tmp *helper.TreeNode[T]

	st.Push(root)
	for !st.IsEmpty() {
		tmp = st.Pop()
		fmt.Println(tmp.Val)

		if tmp.Right != nil {
			st.Push(tmp.Right)
		}

		if tmp.Left != nil {
			st.Push(tmp.Left)
		}
	}
}

// BFS(Breath First Search）广度优先搜索
// bfs没有递归写法
func breadthFirstSearch[T int](root *helper.TreeNode[T]) {
	q := helper.NewQueue[int](10)
	var tmp *helper.TreeNode[T]

	q.Push(root)
	for !q.IsEmpty() {
		tmp, _ = q.Pop()
		fmt.Println(tmp.Val)

		if tmp.Left != nil {
			q.Push(tmp.Left)
		}
		if tmp.Right != nil {
			q.Push(tmp.Right)
		}
	}
}

func buildTestTree[T int]() *helper.TreeNode[T] {
	root := newNode(1)
	root.Left = newNode(2)
	root.Right = newNode(3)

	root.Left.Left = newNode(4)
	root.Left.Right = newNode(5)
	root.Right.Left = newNode(6)
	root.Right.Right = newNode(7)

	return root
}

func newNode(i int) *helper.TreeNode[int] {
	return &helper.TreeNode[int]{
		Val: i,
	}
}
