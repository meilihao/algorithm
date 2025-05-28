/*
199.中 二叉树的右视图

给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例 1：

输入：root = [1,2,3,null,5,null,4]

输出：[1,3,4]

解释：

示例 2：

输入：root = [1,2,3,4,null,null,null,5]

输出：[1,3,4,5]

解释：

示例 3：

输入：root = [1,null,3]

输出：[1,3]

示例 4：

输入：root = []

输出：[]

提示:

二叉树的节点个数的范围是 [0,100]
-100 <= Node.val <= 100
*/
package leetcode

import (
	"fmt"
	"testing"

	. "al/leetcode"
)

func TestRightSideView(t *testing.T) {
	r1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
		},
	}

	fmt.Println(rightSideView(r1))

	r2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
		},
	}

	fmt.Println(rightSideView(r2))
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}

	q := []*TreeNode{root}
	var size int
	var level []*TreeNode
	var tmp *TreeNode

	for len(q) > 0 {
		size = len(q)
		level = nil

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
		res = append(res, level[len(level)-1].Val)
	}

	return res
}
