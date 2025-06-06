/*
953.简 验证外星语词典

某种外星语也使用英文小写字母，但可能顺序 order 不同。字母表的顺序（order）是一些小写字母的排列。

给定一组用外星语书写的单词 words，以及其字母表的顺序 order，只有当给定的单词在这种外星语中按字典序排列时，返回 true；否则，返回 false。

示例 1：

输入：words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
输出：true
解释：在该语言的字母表中，'h' 位于 'l' 之前，所以单词序列是按字典序排列的。
示例 2：

输入：words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
输出：false
解释：在该语言的字母表中，'d' 位于 'l' 之后，那么 words[0] > words[1]，因此单词序列不是按字典序排列的。
示例 3：

输入：words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
输出：false
解释：当前三个字符 "app" 匹配时，第二个字符串相对短一些，然后根据词典编纂规则 "apple" > "app"，因为 'l' > '∅'，其中 '∅' 是空白字符，定义为比任何其他字符都小（更多信息）。

提示：

1 <= words.length <= 100
1 <= words[i].length <= 20
order.length == 26
在 words[i] 和 order 中的所有字符都是英文小写字母。
*/
package leetcode

import (
	"testing"
)

func TestIsAlienSorted(t *testing.T) {

}

// 如果已排序, rule(words[i-1]) <= rule(words[i])
func isAlienSorted(words []string, order string) bool {
	index := [26]int{} // index[i] 表示字符 i 在字母表 order 的排序索引
	for i, c := range order {
		index[c-'a'] = i
	}

next:
	for i := 1; i < len(words); i++ { // 外层循环遍历 words[i-1] 和 words[i]
		for j := 0; j < len(words[i-1]) && j < len(words[i]); j++ { // 内层循环比较两个单词中的每个字母
			pre, cur := index[words[i-1][j]-'a'], index[words[i][j]-'a']
			if pre > cur {
				return false
			}
			if pre < cur { // 当前字母顺序正确, 后续字母不用比较了
				continue next
			}
		}
		if len(words[i-1]) > len(words[i]) {
			return false
		}
	}
	return true
}
