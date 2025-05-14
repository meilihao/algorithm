/*
212.困 单词搜索 II

给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words， 返回所有二维网格上的单词 。

单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。

示例 1：

输入：board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
输出：["eat","oath"]
示例 2：

输入：board = [["a","b"],["c","d"]], words = ["abcb"]
输出：[]

提示：

m == board.length
n == board[i].length
1 <= m, n <= 12
board[i][j] 是一个小写英文字母
1 <= words.length <= 3 * 104
1 <= words[i].length <= 10
words[i] 由小写英文字母组成
words 中的所有字符串互不相同
*/
package main

import (
	"fmt"
	"testing"

	. "al/leetcode"
)

func TestFindWords(t *testing.T) {
	words := []string{"aab", "bbaabaabaaaaabaababaaaaababb", "aabbaaabaaabaabaaaaaabbaaaba", "babaababbbbbbbaabaababaabaaa", "bbbaaabaabbaaababababbbbbaaa", "babbabbbbaabbabaaaaaabbbaaab", "bbbababbbbbbbababbabbbbbabaa", "babababbababaabbbbabbbbabbba", "abbbbbbaabaaabaaababaabbabba", "aabaabababbbbbbababbbababbaa", "aabbbbabbaababaaaabababbaaba", "ababaababaaabbabbaabbaabbaba", "abaabbbaaaaababbbaaaaabbbaab", "aabbabaabaabbabababaaabbbaab", "baaabaaaabbabaaabaabababaaaa", "aaabbabaaaababbabbaabbaabbaa", "aaabaaaaabaabbabaabbbbaabaaa", "abbaabbaaaabbaababababbaabbb", "baabaababbbbaaaabaaabbababbb", "aabaababbaababbaaabaabababab", "abbaaabbaabaabaabbbbaabbbbbb", "aaababaabbaaabbbaaabbabbabab", "bbababbbabbbbabbbbabbbbbabaa", "abbbaabbbaaababbbababbababba", "bbbbbbbabbbababbabaabababaab", "aaaababaabbbbabaaaaabaaaaabb", "bbaaabbbbabbaaabbaabbabbaaba", "aabaabbbbaabaabbabaabababaaa", "abbababbbaababaabbababababbb", "aabbbabbaaaababbbbabbababbbb", "babbbaabababbbbbbbbbaabbabaa"}
	// boardStr := `[
	//   ['o','a','a','n'],
	//   ['e','t','a','e'],
	//   ['i','h','k','r'],
	//   ['i','f','l','v']
	// ]`
	// boardStr := `
	// [
	// 	["a","a"]
	// ]`
	boardStr := `
	[
		["b","a","a","b","a","b"],
		["a","b","a","a","a","a"],
		["a","b","a","a","a","b"],
		["a","b","a","b","b","a"],
		["a","a","b","b","a","b"],
		["a","a","b","b","b","a"],
		["a","a","b","a","a","b"]
	]`

	f := String2Byte
	board := String2TwoDimensionalByteArray(boardStr, ",", f, 0, 0)
	PrintTwoDimensionalByteArray(board)

	//_ = words
	fmt.Println(findWords(board, words))
}

func findWords(board [][]byte, words []string) []string {
	if len(words) == 0 || len(board) == 0 || len(board[0]) == 0 {
		return nil
	}

	trie := &Trie{}
	for i := range words {
		trie.Insert(words[i])
	}

	var results []string

	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if trie.children[int(board[i][j]-'a')] == nil { // 起始字母不在trie中
				continue
			}

			dfs(board, i, j, trie, &results)
		}
	}

	return results
}

func dfs(board [][]byte, i int, j int, trie *Trie, results *[]string) {
	c := board[i][j]

	//fmt.Println(i, j, string(c))
	if c == '@' || trie.children[int(c-'a')] == nil { // 已走过或c对应的节点在trie不存在
		return
	}

	trie = trie.children[int(c-'a')]
	if trie.word != "" {
		// Found one
		//fmt.Println(*results)
		*results = append(*results, trie.word)
		trie.word = "" //表示已找到, 避免words含重复内容
		//return 不能return而是需要继续查找, 因为: 比如aaa匹配到, 继续查找可能匹配到aaab,即它们有相同的前缀
	}

	board[i][j] = '@' // 表示已被搜索

	if i > 0 { // 左
		dfs(board, i-1, j, trie, results)
	}

	if i < len(board)-1 { //右
		dfs(board, i+1, j, trie, results)
	}

	if j > 0 { // 上
		dfs(board, i, j-1, trie, results)
	}

	if j < len(board[0])-1 { //下
		dfs(board, i, j+1, trie, results)
	}

	// 还原现场
	board[i][j] = c
}

type Trie struct {
	val      byte
	children [26]*Trie
	isEnd    int    // 使用int可顺便统计word的频次
	word     string // 插入词
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if word == "" {
		return
	}

	root := this

	ss := []byte(word)
	var idx int

	for _, v := range ss {
		idx = int(v - 'a')

		if root.children[idx] == nil {
			root.children[idx] = &Trie{val: v}
		}

		root = root.children[idx]
	}

	root.isEnd++
	root.word = word
}

func (this *Trie) Search(word string) bool {
	if word == "" {
		return false
	}

	root := this
	ss := []byte(word)
	var idx int

	for _, v := range ss {
		idx = int(v - 'a')

		if root.children[idx] == nil {
			return false
		}

		root = root.children[idx]
	}

	return root.isEnd > 0
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return false
	}

	root := this
	ss := []byte(prefix)
	var idx int

	for _, v := range ss {
		idx = int(v - 'a')

		if root.children[idx] == nil {
			return false
		}

		root = root.children[idx]
	}

	return true
}
