// 208. 实现 Trie (前缀树)
package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Insert("apple")
	fmt.Println(obj.Search("apple"))
	fmt.Println(obj.Search("app"))
	fmt.Println(obj.StartsWith("app"))
	obj.Insert("app")
	fmt.Println(obj.Search("app"))
}

type Trie struct {
	val      byte
	children [26]*Trie
	isEnd    int // 使用int可顺便统计word的频次
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
