/*
105.中 从前序与中序遍历序列构造二叉树

给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

示例 1:

输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]
示例 2:

输入: preorder = [-1], inorder = [-1]
输出: [-1]
*/
package leetcode

import (
	"bytes"
	"fmt"
	"testing"

	. "al/leetcode"

	"github.com/bradleyjkemp/memviz"
)

/*
		preOrder := []int{1, 2, 4, 7, 3, 5, 6, 8} // 前序遍历
		midOrder := []int{4, 7, 2, 1, 5, 3, 8, 6} // 中序遍历

		思路:
			不断找根的过程

		其他: 中后也可构建唯一的二叉树; 但前后序不可以,比如下面两个二叉树,它们的前后序都一致,但树不唯一, 这是因为无法确定左右子树:
		2          2
	    /            \
	   1              1
	  /                \
	 3                  3
*/
func TestRebuildTree(t *testing.T) {
	RebuildTree()
}

func RebuildTree() {
	preOrder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	midOrder := []int{4, 7, 2, 1, 5, 3, 8, 6}

	if !(len(preOrder) == len(midOrder) && len(preOrder) > 0) {
		panic("invalid binary tree")
	}

	root := rebuildTree(preOrder, midOrder)

	PreOutput(root)
	fmt.Println("----")
	MidOutput(root)

	// 右斜树
	preOrder = []int{1, 2, 3, 4, 5}
	midOrder = []int{1, 2, 3, 4, 5}

	root = rebuildTree(preOrder, midOrder)

	buf := &bytes.Buffer{}
	memviz.Map(buf, root)
	//os.WriteFile("a.dot", buf.Bytes(), 0600)
}

// pre=[root+left+right]
// mid=[left+root+right]
func rebuildTree(pre, mid []int) *TreeNode {
	root := &TreeNode{
		Val: pre[0],
	}

	// 其实len(pre)必定等于len(mid)
	if len(pre) == 1 {
		if len(pre) == len(mid) && pre[0] == mid[0] {
			return root
		} else {
			panic("invalid put")
		}
	}

	ln := -1 // root节点对应在mid中的位置, 也可用于推断pre中左右子树的个数
	for i, v := range mid {
		if v == pre[0] {
			ln = i
			break
		}
	}
	if ln == -1 { // 必定应该找到root对应在mid中的节点
		panic("invalid put")
	}

	if ln > 0 { // 左子树有内容
		root.Left = rebuildTree(pre[1:1+ln], mid[:ln]) // pre[1:1+ln] = pre[1:len(mid[:ln])+1] // 从1开始是跳过root
	}

	if ln < len(mid)-1 { // 右子树有内容
		root.Right = rebuildTree(pre[ln+1:], mid[ln+1:]) //pre[ln+1:] = pre[len(mid)+1] // pre[ln+1:]剩余节点皆是right
	}

	return root
}

// 解法2
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree2(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree2(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

// 前序遍历输出
func PreOutput(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	PreOutput(root.Left)
	PreOutput(root.Right)
}

// 中序遍历输出
func MidOutput(root *TreeNode) {
	if root == nil {
		return
	}

	MidOutput(root.Left)
	fmt.Println(root.Val)
	MidOutput(root.Right)
}
