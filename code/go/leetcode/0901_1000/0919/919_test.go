/*
919.中 完全二叉树插入器

完全二叉树 是每一层（除最后一层外）都是完全填充（即，节点数达到最大）的，并且所有的节点都尽可能地集中在左侧。

设计一种算法，将一个新节点插入到一棵完全二叉树中，并在插入后保持其完整。

实现 CBTInserter 类:

CBTInserter(TreeNode root) 使用头节点为 root 的给定树初始化该数据结构；
CBTInserter.insert(int v)  向树中插入一个值为 Node.val == val的新节点 TreeNode。使树保持完全二叉树的状态，并返回插入节点 TreeNode 的父节点的值；
CBTInserter.get_root() 将返回树的头节点。

示例 1：

输入
["CBTInserter", "insert", "insert", "get_root"]
[[[1, 2]], [3], [4], []]
输出
[null, 1, 2, [1, 2, 3, 4]]

解释
CBTInserter cBTInserter = new CBTInserter([1, 2]);
cBTInserter.insert(3);  // 返回 1
cBTInserter.insert(4);  // 返回 2
cBTInserter.get_root(); // 返回 [1, 2, 3, 4]

提示：

树中节点数量范围为 [1, 1000]
0 <= Node.val <= 5000
root 是完全二叉树
0 <= val <= 5000
每个测试用例最多调用 insert 和 get_root 操作 104 次
*/
package leetcode

import (
	"testing"

	. "al/leetcode"
)

func TestCBTInserter(t *testing.T) {

}

type CBTInserter struct {
	root      *TreeNode
	candidate []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	q := []*TreeNode{root}
	// 完全二叉树的特性：从上往下，从左往右填满。所以某个节点如果还有子节点空位（左或右），它就有可能作为接下来插入的父节点
	candidate := []*TreeNode{} // 保存当前可以插入子节点的节点列表

	// bfs
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left == nil || node.Right == nil {
			candidate = append(candidate, node)
		}
	}
	return CBTInserter{root, candidate}
}

func (c *CBTInserter) Insert(val int) int {
	child := &TreeNode{Val: val}
	node := c.candidate[0]
	if node.Left == nil {
		node.Left = child
	} else {
		node.Right = child
		c.candidate = c.candidate[1:] // 当前节点插满，出队
	}
	c.candidate = append(c.candidate, child)
	return node.Val
}

func (c *CBTInserter) Get_root() *TreeNode {
	return c.root
}
