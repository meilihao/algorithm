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

package ago

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"al/helper"

	"github.com/bradleyjkemp/memviz"
)

func TestRebuildBinaryTree(t *testing.T) {
	RebuildBinaryTree()
}

func RebuildBinaryTree() {
	preOrder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	midOrder := []int{4, 7, 2, 1, 5, 3, 8, 6}

	if !(len(preOrder) == len(midOrder) && len(preOrder) > 0) {
		panic("invalid binary tree")
	}

	root := rebuildBinaryTree(preOrder, midOrder)

	PreOutput(root)
	fmt.Println("----")
	MidOutput(root)

	// 右斜树
	preOrder = []int{1, 2, 3, 4, 5}
	midOrder = []int{1, 2, 3, 4, 5}

	root = rebuildBinaryTree(preOrder, midOrder)

	buf := &bytes.Buffer{}
	memviz.Map(buf, root)
	os.WriteFile("a.dot", buf.Bytes(), 0600)
}

func rebuildBinaryTree(pre, mid []int) *helper.TreeNode[int] {
	root := &helper.TreeNode[int]{
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
		root.Left = rebuildBinaryTree(pre[1:1+ln], mid[:ln])
	}

	if ln < len(mid)-1 { // 右子树有内容
		root.Right = rebuildBinaryTree(pre[ln+1:], mid[ln+1:])
	}

	return root
}

// 前序遍历输出
func PreOutput(root *helper.TreeNode[int]) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	PreOutput(root.Left)
	PreOutput(root.Right)
}

// 中序遍历输出
func MidOutput(root *helper.TreeNode[int]) {
	if root == nil {
		return
	}

	MidOutput(root.Left)
	fmt.Println(root.Val)
	MidOutput(root.Right)
}
