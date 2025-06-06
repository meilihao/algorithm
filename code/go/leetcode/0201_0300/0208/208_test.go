/*
208.中 实现 Trie (前缀树)

Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补全和拼写检查。

请你实现 Trie 类：

Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

示例：

输入
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出
[null, null, true, false, true, null, true]

解释
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True

提示：

1 <= word.length, prefix.length <= 2000
word 和 prefix 仅由小写英文字母组成
insert、search 和 startsWith 调用次数 总计 不超过 3 * 104 次
*/
package main

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	obj := Constructor()
	obj.Insert("apple")
	fmt.Println(obj.Search("apple"))
	fmt.Println(obj.Search("app"))
	fmt.Println(obj.StartsWith("app"))
	obj.Insert("app")
	fmt.Println(obj.Search("app"))
}

// 每个字符是一个节点
type Trie struct {
	val      byte
	children [26]*Trie
	isEnd    int // 表示该节点是否为字符串的结尾, 使用int可顺便统计word的频次
}

func Constructor() Trie {
	return Trie{}
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
