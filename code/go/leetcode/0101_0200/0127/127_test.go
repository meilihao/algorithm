/*
127.困 单词接龙

字典 wordList 中从单词 beginWord 到 endWord 的 转换序列 是一个按下述规格形成的序列 beginWord -> s1 -> s2 -> ... -> sk：

每一对相邻的单词只差一个字母。

	对于 1 <= i <= k 时，每个 si 都在 wordList 中。注意， beginWord 不需要在 wordList 中。

sk == endWord
给你两个单词 beginWord 和 endWord 和一个字典 wordList ，返回 从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0 。

示例 1：

输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
输出：5
解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。
示例 2：

输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
输出：0
解释：endWord "cog" 不在字典中，所以无法进行转换。

提示：

1 <= beginWord.length <= 10
endWord.length == beginWord.length
1 <= wordList.length <= 5000
wordList[i].length == beginWord.length
beginWord、endWord 和 wordList[i] 由小写英文字母组成
beginWord != endWord
wordList 中的所有字符串 互不相同
*/
package leetcode

import (
	"math"
	"testing"
)

func TestLadderLength(t *testing.T) {

}

// 把每个单词都抽象为一个点，如果两个单词可以只改变一个字母进行转换，那么说明他们之间有一条双向边。因此只需要把满足转换条件的点相连，就形成了一张图
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{} // 存储单词到其唯一ID的映射
	graph := [][]int{}         // 邻接列表表示的图

	// 将单词加入 wordId 映射和 graph (如果它不在的话)
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId) // 为新单词分配一个ID
			wordId[word] = id
			graph = append(graph, []int{}) // 为新ID在图中添加一个空的邻接列表
		}
		return id
	}

	// 两个单词 wordA 和 wordB 如果只相差一个字符，那么它们会共享一个相同的“通用状态”节点。
	// 例如，"hot" 和 "dot" 都会生成 "*ot" 这个通用状态节点。通过这个共同的通用状态节点，hot 和 dot 就间接连接起来了
	// 这种方法避免了 O(L*N^2) (L是单词长度，N是单词数量) 的暴力比较，而是将每个单词与其 L 个通用状态连接，总边数是 O(N⋅L)，大大提高了图的构建效率
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'                // 将当前字符替换为通配符 '*'，形成一个“通用状态”单词
			id2 := addWord(string(s)) // 将这个“通用状态”单词也加入 wordId 和 graph

			// 在原始单词ID和通用状态单词ID之间添加双向边
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)

			s[i] = b // 恢复当前字符，以便处理下一个位置
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0 // 如果 endWord 不在 wordList 中，则无法到达，返回 0
	}

	const inf int = math.MaxInt64    // 定义无穷大，表示未访问
	dist := make([]int, len(wordId)) // 存储从 beginId 到每个节点的最短距离
	for i := range dist {
		dist[i] = inf // 初始化所有距离为无穷大
	}
	dist[beginId] = 0       // 起点到自身的距离为 0
	queue := []int{beginId} // // BFS 队列，存储待访问的节点ID
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == endId {
			// 返回 (dist[endId] / 2) + 1
			// 为什么除以 2 再加 1？
			// 因为我们的图是“单词 <-> 通用状态 <-> 单词”这种交替结构。
			// 实际的单词转换是：单词A -> 单词B (距离为1)。
			// 在我们的图中，路径是：单词A (id1) -> 通用状态 (id2) -> 单词B (id3)。
			// 这样，一个单词转换 (如 "hot" -> "dot") 在图中对应的路径长度是 2 (hot -> *ot -> dot)。
			// 所以，图中的最短路径长度 (dist[endId]) 是实际单词转换步数的两倍。
			// 因此，需要将 dist[endId] 除以 2，然后加 1 (因为题目要求的是序列长度，包含头尾单词)。
			// 例如，hit -> hot -> dot -> dog -> cog，实际4步转换，序列长度5。
			// 在图中：id(hit) -> *it -> id(hot) -> *ot -> id(dot) -> *og -> id(dog) -> *og -> id(cog)。
			// dist[id(cog)] 应该为 8。那么 8/2 + 1 = 5。
			return dist[endId]/2 + 1
		}
		for _, w := range graph[v] { // 遍历当前节点 v 的所有邻居 w
			if dist[w] == inf { // 如果邻居 w 尚未被访问过
				dist[w] = dist[v] + 1    // 更新邻居 w 的距离
				queue = append(queue, w) // 将邻居 w 加入队列
			}
		}
	}
	return 0 // 如果队列为空，但仍未到达 endId，说明无法找到路径，返回 0
}
