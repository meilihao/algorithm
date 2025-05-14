/*
117.中 填充每个节点的下一个右侧节点指针 II

给定一个二叉树：

	struct Node {
	  int val;
	  Node *left;
	  Node *right;
	  Node *next;
	}

填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。

初始状态下，所有 next 指针都被设置为 NULL 。

示例 1：

输入：root = [1,2,3,4,5,null,7]
输出：[1,#,2,3,#,4,5,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化输出按层序遍历顺序（由 next 指针连接），'#' 表示每层的末尾。
示例 2：

输入：root = []
输出：[]
*/
package leetcode

import (
	"testing"

	. "al/leetcode"
)

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
