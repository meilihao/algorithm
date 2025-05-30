/*
820.中 单词的压缩编码

单词数组 words 的 有效编码 由任意助记字符串 s 和下标数组 indices 组成，且满足：

words.length == indices.length
助记字符串 s 以 '#' 字符结尾
对于每个下标 indices[i] ，s 的一个从 indices[i] 开始、到下一个 '#' 字符结束（但不包括 '#'）的 子字符串 恰好与 words[i] 相等
给你一个单词数组 words ，返回成功对 words 进行编码的最小助记字符串 s 的长度 。

示例 1：

输入：words = ["time", "me", "bell"]
输出：10
解释：一组有效编码为 s = "time#bell#" 和 indices = [0, 2, 5] 。
words[0] = "time" ，s 开始于 indices[0] = 0 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"
words[1] = "me" ，s 开始于 indices[1] = 2 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"
words[2] = "bell" ，s 开始于 indices[2] = 5 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"
示例 2：

输入：words = ["t"]
输出：2
解释：一组有效编码为 s = "t#" 和 indices = [0] 。

提示：

1 <= words.length <= 2000
1 <= words[i].length <= 7
words[i] 仅由小写字母组成
*/
package leetcode

import (
	"fmt"
	"testing"
)

func TestMinimumLengthEncoding(t *testing.T) {
	words := []string{"time", "me", "bell"}
	fmt.Println(minimumLengthEncoding(words))
}

type tire struct {
	child [26]*tire
}

func buildTire(words []string) *tire {
	root := &tire{}

	for _, v := range words {
		node := root
		for n := len(v) - 1; n >= 0; n-- { // 合并后缀
			if node.child[v[n]-'a'] == nil {
				node.child[v[n]-'a'] = &tire{}
			}

			node = node.child[v[n]-'a']
		}
	}

	return root
}

func dfs(root *tire, l int, total []int) {
	isLeaf := true // 假设当前节点是叶子节点

	for _, c := range root.child {
		if c != nil {
			isLeaf = false     //  如果有子节点存在, 那么当前节点就不是叶子节点
			dfs(c, l+1, total) // 递归地向下遍历子节点，深度加1
		}
	}

	// 如果遍历完所有子节点后，isLeaf 仍然为 true
	// 说明当前节点是一个叶子节点，它代表一个需要独立编码的单词
	// 将当前路径的长度 (l) 加到总长度中
	// 这里的 l 包含了单词的长度和结尾的 '#'
	if isLeaf {
		total[0] += l
	}
}

/*
time+me:
(root)

	| e
	(node_e)
	|   | m
	|   (node_em) --- 'em' (原始单词 'me')
	|       | i
	|       (node_emi)
	|           | t
	|           (node_emit) --- 'emit' (原始单词 'time')
*/
func minimumLengthEncoding(words []string) int {
	root := buildTire(words)

	total := []int{0}   // 使用切片作为指针，在递归中修改
	dfs(root, 1, total) // 初始深度为1（表示第一个字符串长度0及其末尾的#）
	return total[0]
}

func minimumLengthEncoding2(words []string) int {
	// trie：倒序

	// hash
	memo := make(map[string]struct{})
	ans := 0
	for _, w := range words {
		memo[w] = struct{}{}
	}
	for _, w := range words {
		for i := 1; i < len(w); i++ {
			delete(memo, w[i:])
		}
	}
	for w, _ := range memo {
		ans += len(w) + 1
	}
	return ans
}
