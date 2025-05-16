/*
08. 二叉树的下一个节点

给定一个二叉树和其中的一个结点，请找出中序遍历顺序的下一个结点并且返回。注意，树中的结点不仅包含左右子结点，同时包含指向父结点的指针
*/
package demo

import "testing"

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func TestGetNext(t *testing.T) {

}

/*
二叉树的中序遍历：{ [左子树], 根节点, [右子树] }

获取下一个节点分三种情况：
1. 如果该节点有右子树，那么下一个节点就是其右子树中最左边的节点；
1. 如果该节点没有右子树

	1. 且是其父节点的左子节点，那么下一个节点就是其父节点；
	2. 且是其父节点的右子节点，沿着父指针一直向上，直到找到一个是它父节点的左子节点的节点，如果这样的节点存在，那么这个节点的父节点即是所求

*/

func getNext(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	//node有右子树，则取右子树的最左结点
	if node.Right != nil {
		next := node.Right
		for next.Left != nil {
			next = next.Left
		}
		return next
	} else {
		//node没有右子树且存在父节点
		for node.Parent != nil {
			//node为父结点的左孩子，则取node的父结点
			if node.Parent.Left == node {
				return node.Parent
			} else {
				//node为父结点的右孩子，则取node.P遍历，直到node.P是node.P.P的左孩子为止
				node = node.Parent
			}
		}
	}
	return nil
}
