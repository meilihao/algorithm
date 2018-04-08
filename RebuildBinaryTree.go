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

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/bradleyjkemp/memviz"
)

type node struct {
	value int
	left  *node
	right *node
}

func main() {
	preOrder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	midOrder := []int{4, 7, 2, 1, 5, 3, 8, 6}

	if !(len(preOrder) == len(midOrder) && len(preOrder) > 0) {
		panic("invalid binary tree")
	}

	root := rebuildBinaryTree(preOrder, midOrder)

	PreOutput(root)
	fmt.Println("----")
	MidOutput(root)

	buf := &bytes.Buffer{}
	memviz.Map(buf, root)
	ioutil.WriteFile("a.dot", buf.Bytes(), 0600)
}

func rebuildBinaryTree(pre, mid []int) *node {
	root := &node{
		value: pre[0],
	}

	// 其实len(pre)必定等于len(mid)
	if len(pre) == 1 {
		if len(pre) == len(mid) && pre[0] == mid[0] {
			return root
		} else {
			panic("invalid put")
		}
	}

	ln := -1
	for i, v := range mid {
		if v == pre[0] {
			ln = i
			break
		}
	}
	if ln == -1 { // 必定应该找到root对应在mid中的节点
		panic("invalid put")
	}

	if ln>0 { // 左子树有内容
		root.left = rebuildBinaryTree(pre[1:1+ln], mid[:ln])
	}

	if ln < len(mid)-1 { // 右子树有内容
		root.right = rebuildBinaryTree(pre[ln+1:], mid[ln+1:])
	} 

	return root
}

// 前序遍历输出
func PreOutput(root *node) {
	if root == nil {
		return
	}

	fmt.Println(root.value)
	PreOutput(root.left)
	PreOutput(root.right)
}

// 中序遍历输出
func MidOutput(root *node) {
	if root == nil {
		return
	}

	MidOutput(root.left)
	fmt.Println(root.value)
	MidOutput(root.right)
}
