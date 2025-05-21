/*
101.简 对称二叉树

给你一个二叉树的根节点 root ， 检查它是否轴对称。

示例 1：

输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：

输入：root = [1,2,2,null,3,null,3]
输出：false

提示：

树中节点数目在范围 [1, 1000] 内
-100 <= Node.val <= 100
*/
package demo

import (
	"testing"

	. "al/leetcode"
)

func TestIsSymmetric(t *testing.T) {

}

/*
对称条件:
1. root val相同
2. 每个子树的右子树与另一个树的左子树镜像对称 + 每个子树的左子树与另一个树的右子树镜像对称
*/
func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}
