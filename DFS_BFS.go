package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

func main() {
	root := buildTestTree()

	depthFirstSearch(root)
	fmt.Println("-----")
	breadthFirstSearch(root)
}

// DFS(Deep First Search）深度优先搜索
// 这里是非递归写法, 也可使用递归实现(常用,因为不用自行维护一个stack)
func depthFirstSearch(root *node) {
	st := newStack(10)
	var tmp *node

	st.push(root)
	for !st.isEmpty() {
		tmp = st.pop()
		fmt.Println(tmp.value)

		if tmp.right != nil {
			st.push(tmp.right)
		}

		if tmp.left != nil {
			st.push(tmp.left)
		}
	}
}

// BFS(Breath First Search）广度优先搜索
// bfs没有递归写法
func breadthFirstSearch(root *node) {
	q := newQueue(10)
	var tmp *node

	q.push(root)
	for !q.isEmpty() {
		tmp = q.pop()
		fmt.Println(tmp.value)

		if tmp.left != nil {
			q.push(tmp.left)
		}
		if tmp.right != nil {
			q.push(tmp.right)
		}
	}
}

func buildTestTree() *node {
	root := newNode(1)
	root.left = newNode(2)
	root.right = newNode(3)

	root.left.left = newNode(4)
	root.left.right = newNode(5)
	root.right.left = newNode(6)
	root.right.right = newNode(7)

	return root
}

func newNode(i int) *node {
	return &node{
		value: i,
	}
}

type stack struct {
	data   []*node
	curTop int
	maxCap int
}

func newStack(maxCap int) *stack {
	return &stack{
		data:   make([]*node, maxCap),
		maxCap: maxCap,
	}
}

func (s *stack) push(n *node) {
	if s.curTop == s.maxCap {
		panic("full stack")
	}

	s.data[s.curTop] = n
	s.curTop++
}

func (s *stack) pop() *node {
	tmp := s.top()

	if s.curTop > 0 {
		s.data[s.curTop-1] = nil
		s.curTop--
	}

	return tmp
}

func (s *stack) top() *node {
	if s.curTop == 0 {
		return nil
	}

	return s.data[s.curTop-1]
}

func (s *stack) isEmpty() bool {
	return s.curTop == 0
}

type queue struct {
	data []*node
}

func newQueue(maxCap int) *queue {
	return &queue{
		data: make([]*node, 0, maxCap),
	}
}

func (q *queue) push(n *node) {
	q.data = append(q.data, n)
}

func (q *queue) isEmpty() bool {
	return len(q.data) == 0
}

func (q *queue) front() *node {
	if len(q.data) == 0 {
		return nil
	}

	return q.data[0]
}

func (q *queue) pop() *node {
	tmp := q.front()

	if n := len(q.data); n > 0 {
		q.data = q.data[1:]
	}

	return tmp
}
