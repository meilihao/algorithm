package leetcode

import "fmt"

func PrintTreeByBfs(root *TreeNode) {
	if root == nil {
		fmt.Println("empty tree")
		return
	}

	res := [][]*TreeNode{}

	q := []*TreeNode{root}
	var size int
	var tmp *TreeNode

	for len(q) > 0 {
		size = len(q)
		level := make([]*TreeNode, 0, size)

		for i := 0; i < size; i++ {
			tmp = q[i]

			if tmp.Left != nil {
				q = append(q, tmp.Left)
			}
			if tmp.Right != nil {
				q = append(q, tmp.Right)
			}

			level = append(level, tmp)
		}

		q = q[size:]
		res = append(res, level)
	}

	for _, l := range res {
		for _, n := range l {
			if n != nil {
				fmt.Printf("%d ", n.Val)
			}
		}
		fmt.Println("")
	}
}
