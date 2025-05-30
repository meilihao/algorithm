package leetcode

import (
	"testing"
)

func TestMagicDictionary2(t *testing.T) {

}

type trie struct {
	children   [26]*trie // 存储子节点，数组索引对应 'a' 到 'z'
	isFinished bool      // 标记当前节点是否是某个单词的结尾
}

type MagicDictionary2 struct {
	*trie
}

func Constructor2() MagicDictionary2 {
	return MagicDictionary2{&trie{}}
}

func (d *MagicDictionary2) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		cur := d.trie
		for _, c := range word {
			c -= 'a'
			if cur.children[c] == nil {
				cur.children[c] = &trie{}
			}
			cur = cur.children[c]
		}
		cur.isFinished = true
	}
}

func (d *MagicDictionary2) Search(word string) bool {
	return dfs(d.trie, word, 0, 0)
}

// i int: 当前正在处理 word 中字符的索引。
// edit int: 到目前为止，已经进行的字符修改（编辑）次数
func dfs(root *trie, word string, i, edit int) bool {
	if root == nil {
		return false
	}

	if root.isFinished && i == len(word) && edit == 1 {
		return true
	}

	if i < len(word) && edit <= 1 {
		found := false // 用于标记是否在当前路径下找到了满足条件的单词

		var j byte
		for j = 0; j < 26 && !found; j++ {
			next := edit
			if word[i]-'a' != j {
				next = edit + 1
			}
			found = dfs(root.children[j], word, i+1, next)
		}

		return found
	}

	return false
}
