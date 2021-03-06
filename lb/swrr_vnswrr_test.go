// 参考:
// - [Nginx 的 Smooth Weighted Round-Robin（SWRR）](https://tenfy.cn/2018/11/12/smooth-weighted-round-robin/)
// - [Tengine 的 Virtual Node Smooth Weighted Round-Robin（VNSWRR）](https://mp.weixin.qq.com/s/3KZ99d94yqRDxEByn7nGWg)
// - [kkdai/maglev](https://github.com/kkdai/maglev)
//
// 结论:
// 1. Node较少时, 这里的测试是N<=7时, SWRR有优势, N>7后VNSWRR有优势
// 1. VNSWRR的ns/op稳定在8.75左右, 几乎不随N变化

package lb

import (
	"math/rand"
	"testing"
	"time"
)

type List struct {
	Nodes   []*Node
	L       []*Node
	N       int
	TW      int
	Alloced int // 已初始化个数
	Start   int
}

type Node struct {
	Name    string
	Current int
	Weight  int // 机器权重
}

func NewList(ns []*Node) *List {
	l := &List{
		Nodes: ns,
		N:     len(ns),
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l.Start = r.Intn(l.N)

	for _, v := range l.Nodes {
		l.TW += v.Weight
	}

	l.L = make([]*Node, 0, l.TW)

	l.alloc()

	return l
}

func (l *List) alloc() {
	//fmt.Println("**************alloc")
	var best *Node

	start := l.Alloced
	limit := l.Alloced + l.N
	if limit > l.TW {
		limit = l.TW
	}

	for ; start < limit; start++ {
		if best = SmoothWrr(l.Nodes); best != nil {
			l.L = append(l.L, best)
			l.Alloced++
		}
	}
}

func (l *List) Pick() *Node {
	//fmt.Printf("----------%+v\n", *l)

	var index int

	if l.Alloced != l.TW && l.Start == l.Alloced {
		l.alloc()
	}

	index = l.Start % l.TW
	l.Start++

	return l.L[index]
}

// nginx/src/http/ngx_http_upstream_round_robin.c#ngx_http_upstream_get_peer
// 第一次会选择当前机器列表中权重最大的机器
func SmoothWrr(nodes []*Node) (best *Node) {
	if len(nodes) == 0 {
		return
	}
	total := 0
	for _, node := range nodes {
		if node == nil {
			continue
		}
		total += node.Weight
		node.Current += node.Weight
		if best == nil || node.Current > best.Current {
			best = node
		}
	}
	if best == nil {
		return
	}

	best.Current -= total
	return
}

var nodes = []*Node{
	&Node{"a", 0, 5},
	&Node{"b", 0, 1},
	&Node{"c", 0, 1},
	&Node{"d", 0, 3},
	&Node{"e", 0, 2},
	&Node{"f", 0, 4},
	&Node{"g", 0, 5},
	&Node{"h1", 0, 6},
	&Node{"h2", 0, 6},
	&Node{"h3", 0, 6},
	&Node{"h4", 0, 6},
	&Node{"h5", 0, 6},
}

func BenchmarkSWRR(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		best := SmoothWrr(nodes)
		if best != nil {
			//fmt.Println(best.Name)
		}
	}

	b.StopTimer()
}

func BenchmarkVNSWRR(b *testing.B) {
	l := NewList(nodes)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		best := l.Pick()
		if best != nil {
			//fmt.Println(best.Name)
		}

	}

	b.StopTimer()
}
