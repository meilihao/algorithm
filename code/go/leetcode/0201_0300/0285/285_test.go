/*
285.中 二叉搜索树中的中序后继

给定一棵二叉搜索树和其中的一个节点 p ，找到该节点在树中的中序后继。如果节点没有中序后继，请返回 null 。

节点 p 的后继是值比 p.val 大的节点中键值最小的节点，即按中序遍历的顺序节点 p 的下一个节点。

示例 1：

输入：root = [2,1,3], p = 1
输出：2
解释：这里 1 的中序后继是 2。请注意 p 和返回值都应是 TreeNode 类型。
示例 2：

输入：root = [5,3,6,2,4,null,null,1], p = 6
输出：null
解释：因为给出的节点没有中序后继，所以答案就返回 null 了。

提示：

树中节点的数目在范围 [1, 104] 内。
-105 <= Node.val <= 105
树中各节点的值均保证唯一。
*/
package leetcode

import (
	"testing"

	. "al/leetcode"
)

func TestInorderSuccessor(t *testing.T) {

}

// 利用二叉搜索树的性质: root.Left.Val < root.Val < root.Right.Val
// best
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var successor *TreeNode
	if p.Right != nil { // target > p.val , 在右子树中
		successor = p.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		return successor
	}

	node := root
	for node != nil {
		if node.Val > p.Val { // p在左子树中
			successor = node
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return successor
}

// 中序遍历
func inorderSuccessor2(root *TreeNode, p *TreeNode) *TreeNode {
	st := []*TreeNode{}
	var pre, cur *TreeNode = nil, root
	for len(st) > 0 || cur != nil {
		for cur != nil {
			st = append(st, cur)
			cur = cur.Left
		}
		cur = st[len(st)-1]
		st = st[:len(st)-1]
		if pre == p {
			return cur
		}
		pre = cur
		cur = cur.Right
	}
	return nil
}
