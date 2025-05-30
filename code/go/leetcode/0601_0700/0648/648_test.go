/*
648.中 单词替换

在英语中，我们有一个叫做 词根(root) 的概念，可以词根 后面 添加其他一些词组成另一个较长的单词——我们称这个词为 衍生词 (derivative)。例如，词根 help，跟随着 继承词 "ful"，可以形成新的单词 "helpful"。

现在，给定一个由许多 词根 组成的词典 dictionary 和一个用空格分隔单词形成的句子 sentence。你需要将句子中的所有 衍生词 用 词根 替换掉。如果 衍生词 有许多可以形成它的 词根，则用 最短 的 词根 替换它。

你需要输出替换之后的句子。

示例 1：

输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
输出："the cat was rat by the bat"
示例 2：

输入：dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
输出："a a b c"

提示：

1 <= dictionary.length <= 1000
1 <= dictionary[i].length <= 100
dictionary[i] 仅由小写字母组成。
1 <= sentence.length <= 106
sentence 仅由小写字母和空格组成。
sentence 中单词的总量在范围 [1, 1000] 内。
sentence 中每个单词的长度在范围 [1, 1000] 内。
sentence 中单词之间由一个空格隔开。
sentence 没有前导或尾随空格。
*/
package leetcode

import (
	"strings"
	"testing"
)

func TestReplaceWords(t *testing.T) {

}

// 时间复杂度：O(d+s)。其中 d 是 dictionary 的字符数，s 是 sentence 的字符数
func replaceWords(dictionary []string, sentence string) string {
	type trie map[rune]trie
	root := trie{}
	for _, s := range dictionary { // 将dictionary 中所有词根构建一棵字典树，并用特殊符号标记结尾
		cur := root
		for _, c := range s {
			if cur[c] == nil {
				cur[c] = trie{}
			}
			cur = cur[c]
		}
		cur['#'] = trie{}
	}

	words := strings.Split(sentence, " ")
	for i, word := range words {
		cur := root
		for j, c := range word {
			if cur['#'] != nil {
				words[i] = word[:j]
				break
			}
			// 没有对应的词根, 即当前单词的前缀不能构成字典树中的任何词根。这意味着当前单词不能被字典树中的任何词根取代
			if cur[c] == nil {
				break
			}
			cur = cur[c]
		}
	}
	return strings.Join(words, " ")
}
