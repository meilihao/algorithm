// 98. 验证二叉搜索树
// 思路:
// 1. 中序遍历, 遍历结果的array是否是升序. O(n)
// 2. 递归  root.Val > max(root.Left) && root.Val < min(root.Right) O(n)
package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	min, max := -1<<63, 1<<63-1 // math.MinInt64, math.MaxInt64

	return recur(min, max, root)
}

func r(min, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}

	if min >= root.Val || root.Val >= max {
		return false
	}

	return r(min, root.Val, root.Left) && r(root.Val, max, root.Right)
}
