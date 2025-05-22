/*
297.困 二叉树的序列化与反序列化

序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

示例 1：

输入：root = [1,2,3,null,null,4,5]
输出：[1,2,3,null,null,4,5]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]
示例 4：

输入：root = [1,2]
输出：[1,2]

提示：

树中结点数在范围 [0, 104] 内
-1000 <= Node.val <= 1000
*/
package demo

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	. "al/leetcode"
)

func TestCodec(t *testing.T) {
	t1 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
	}

	c := Constructor()
	s := c.serialize(t1)
	fmt.Println(s) // 5,4,11,7,null,null,2,null,null,null,null,
	fmt.Println("---")
	t2 := c.deserialize(s)
	fmt.Println(t2)
}

type Codec struct{}

func Constructor() (_ Codec) {
	return
}

func (Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteByte(',')
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

// serialize()最后输出的多余`,`不影响deserialize()的原因: 每个TreeNode调用2次build()即build()调用受到节点个数限制, 最后一个`""`不会用到
func (Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	//fmt.Println(sp, len(sp), sp[len(sp)-1] == "")
	var build func() *TreeNode
	build = func() *TreeNode {
		//fmt.Println(sp[0], sp, len(sp), sp[len(sp)-1] == "")
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()
}
