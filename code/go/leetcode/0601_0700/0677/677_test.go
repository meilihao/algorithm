/*
677.中 键值映射

设计一个 map ，满足以下几点:

字符串表示键，整数表示值
返回具有前缀等于给定字符串的键的值的总和
实现一个 MapSum 类：

MapSum() 初始化 MapSum 对象
void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。如果键 key 已经存在，那么原来的键值对 key-value 将被替代成新的键值对。
int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。

示例 1：

输入：
["MapSum", "insert", "sum", "insert", "sum"]
[[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
输出：
[null, null, 3, null, 5]

解释：
MapSum mapSum = new MapSum();
mapSum.insert("apple", 3);
mapSum.sum("ap");           // 返回 3 (apple = 3)
mapSum.insert("app", 2);
mapSum.sum("ap");           // 返回 5 (apple + app = 3 + 2 = 5)

提示：

1 <= key.length, prefix.length <= 50
key 和 prefix 仅由小写英文字母组成
1 <= val <= 1000
最多调用 50 次 insert 和 sum
*/
package leetcode

import (
	"fmt"
	"testing"
)

// other: 将值仅存在末尾节点, 再prefix后的所有节点相加
func TestMapSum(t *testing.T) {
	s := Constructor()

	s.Insert("apple", 3)
	fmt.Println(s.Sum("ap"))
	s.Insert("app", 2)
	fmt.Println(s.Sum("ap"))
	s.Insert("apple", 4)
	fmt.Println(s.Sum("ap"))
}

type TrieNode struct {
	children [26]*TrieNode
	val      int // 存储所有以从根节点到当前节点的路径为前缀的单词的**累加值**
}

type MapSum struct {
	root *TrieNode
	cnt  map[string]int // 存储每个完整的键值对，用于处理键的更新
}

func Constructor() MapSum {
	return MapSum{&TrieNode{}, map[string]int{}}
}

func (m *MapSum) Insert(key string, val int) {
	delta := val
	if m.cnt[key] > 0 {
		delta -= m.cnt[key]
	}
	m.cnt[key] = val // 更新 cnt 中 key 对应的值为新值

	node := m.root
	for _, ch := range key {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{}
		}
		node = node.children[ch]
		node.val += delta // 更新当前节点（及所有前缀节点）的 val
	}
}

func (m *MapSum) Sum(prefix string) int {
	node := m.root
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return 0
		}
		node = node.children[ch]
	}
	return node.val
}
