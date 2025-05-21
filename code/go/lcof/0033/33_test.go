/*
LCR 152.中 验证二叉搜索树的后序遍历序列

请实现一个函数来判断整数数组 postorder 是否为二叉搜索树的后序遍历结果。

示例 1：

输入: postorder = [4,9,6,5,8]
输出: false
解释：从上图可以看出这不是一颗二叉搜索树
示例 2：

输入: postorder = [4,6,5,9,8]
输出: true
解释：可构建的二叉搜索树如上图

提示：

数组长度 <= 1000
postorder 中无重复数字
*/
package demo

import (
	"testing"

	. "al/leetcode"
)

func TestVerifyTreeOrder(t *testing.T) {

}

/*
后序遍历定义： [ 左子树 | 右子树 | 根节点 ] ，即遍历顺序为 “左、右、根”

二叉搜索树定义： 左子树中所有节点的值 < 根节点的值；右子树中所有节点的值 > 根节点的值；其左、右子树也分别为二叉搜索树
*/
func verifyTreeOrder(postorder []int) bool {
	return recur(postorder, 0, len(postorder)-1)
}

func recur(postorder []int, l, r int) bool {
	if l >= r {
		return true
	}

	p := l
	for postorder[p] < postorder[r] { // postorder[r]为root
		p++
	}
	m := p
	for (postorder[p]) > postorder[r] {
		p++
	}
	/*
		p=j ： 判断 此树 是否正确
		recur(i,m−1) ： 判断 此树的左子树 是否正确
		recur(m,j−1) ： 判断 此树的右子树 是否正确
	*/
	return p == r && recur(postorder, l, m-1) && recur(postorder, m, r-1)
}
