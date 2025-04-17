package leetcode

import "testing"

func TestConnect(t *testing.T) {
	connect(nil)
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root} // 队列
	for len(q) > 0 {
		tmp := q
		q = nil
		for i, node := range tmp {
			if i+1 < len(tmp) { // 有水平方向的其他节点
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}
