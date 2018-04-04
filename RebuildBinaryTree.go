/*
	preOrder := []int{1, 2, 4, 7, 3, 5, 6, 8} // 前序遍历
	midOrder := []int{4, 7, 2, 1, 5, 3, 8, 6} // 中序遍历

	思路:
		不断找根的过程
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

	root := rebuildBinaryTree(preOrder, midOrder)

	buf := &bytes.Buffer{}
	memviz.Map(buf, root)
	ioutil.WriteFile("a.dot", buf.Bytes(), 0600)
}

func rebuildBinaryTree(pre, mid []int) *node {
	root := &node{
		value: pre[0],
	}

	var ln int
	for i, v := range mid {
		if v == pre[0] {
			ln = i
		}
	}

	fmt.Println(root.value, 1+ln, ln)
	if 1+ln <= len(pre)-1 { // 左子树有内容
		root.left = rebuildBinaryTree(pre[1:1+ln], mid[:ln])
	} else {
		root.left = nil
	}

	if ln < len(mid)-1 { // 右子树有内容
		root.right = rebuildBinaryTree(pre[ln+2:], mid[ln+1:])
	} else {
		root.right = nil
	}

	return root
}
