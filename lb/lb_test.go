package lb

import (
	"fmt"
	"hash/crc32"
	"math/rand"
	"testing"
	"time"
)

// 随机选择
// seed一样, 则随机数一样
// 多次执行`go test -v lb_test.go`, t均相同, 除非lb_test.go内容出现变化(不包括注释)
func Random(n int) int {
	t := time.Now().Unix()
	r := rand.New(rand.NewSource(t))
	return r.Intn(n)
}

func TestRandom(t *testing.T) {
	n := 50

	if i := Random(n); i < 0 || i >= n {
		t.Fatalf("out of range: %d, %d", n, i)
	} else {
		t.Logf("range: %d, %d", n, i)
	}
}

// 加权随机选择
func RandomWeight(ls []NodeWeight) int {
	tmp := []int{}

	for _, v := range ls {
		for i := 0; i < v.Weight; i++ {
			tmp = append(tmp, v.Index)
		}
	}

	n := Random(len(tmp))
	return tmp[n]
}

type NodeWeight struct {
	Index, Weight int
}

func TestRandomWeight(t *testing.T) {
	ls := []NodeWeight{
		{
			Index:  0,
			Weight: 1,
		},
		{
			Index:  1,
			Weight: 2,
		},
		{
			Index:  2,
			Weight: 3,
		},
	}

	n := len(ls)
	if i := RandomWeight(ls); i < 0 || i >= n {
		t.Fatalf("out of range: %d, %d", n, i)
	} else {
		t.Logf("range: %d, %d", n, i)
	}
}

// 轮转
type RoundRobin struct {
	now             int
	Num             int
	IsRandomAtFirst bool // 第一次随机以避免负载不均衡
}

func NewRoundRobin(num int, IsRandomAtFirst bool) *RoundRobin {
	t := &RoundRobin{
		Num:             num,
		IsRandomAtFirst: IsRandomAtFirst,
	}

	if t.IsRandomAtFirst {
		t.now = Random(t.Num)
	}

	return t
}

func (r *RoundRobin) Next() int {
	tmp := r.now

	r.now++
	if r.now >= r.Num {
		r.now = 0
	}

	return tmp
}

func TestRoundRobin(t *testing.T) {
	n := 5

	s := NewRoundRobin(n, true)
	x := s.Next()

	if y := s.Next(); y < 0 || y >= n || (x+1)%n != y {
		t.Fatalf("out of range: %d, %d, %d", n, x, y)
	} else {
		t.Logf("range: %d, %d, %d", n, x, y)
	}
}

// 加权轮转
type RoundRobinWeight struct {
	now             int
	ls              []int
	List            []NodeWeight
	IsRandomAtFirst bool // 第一次随机以避免负载不均衡
}

func NewRoundRobinWeight(ls []NodeWeight, IsRandomAtFirst bool) *RoundRobinWeight {
	t := &RoundRobinWeight{
		List:            ls,
		IsRandomAtFirst: IsRandomAtFirst,
	}

	for _, v := range t.List {
		for i := 0; i < v.Weight; i++ {
			t.ls = append(t.ls, v.Index)
		}
	}

	if t.IsRandomAtFirst {
		t.now = Random(len(t.ls))
	}

	return t
}

func (r *RoundRobinWeight) Next() int {
	tmp := r.now

	r.now++
	if r.now >= len(r.ls) {
		r.now = 0
	}

	return r.ls[tmp]
}

func TestRoundRobinWeight(t *testing.T) {
	ls := []NodeWeight{
		{
			Index:  0,
			Weight: 1,
		},
		{
			Index:  1,
			Weight: 2,
		},
		{
			Index:  2,
			Weight: 3,
		},
	}

	n := len(ls)

	s := NewRoundRobinWeight(ls, true)
	x := s.Next()

	if y := s.Next(); y < 0 || y >= n {
		t.Fatalf("out of range: %d, %d, %d", n, x, y)
	} else {
		t.Logf("range: %d, %d, %d", n, x, y)
	}
}

func TestHash(t *testing.T) {
	host := fmt.Sprintf("10.0.%d.%d", rand.Intn(255), rand.Intn(255))
	n := 5

	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(host), crcTable)
	i := int(hashVal) % n

	if i < 0 || i >= n {
		t.Fatalf("out of range: %d, %d", n, i)
	} else {
		t.Logf("range: %d, %d", n, i)
	}
}
