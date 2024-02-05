// https://cs.opensource.google/go/go/+/refs/tags/go1.21.6:src/container/list/list.go
package biomerge

type Unit struct {
	Sector int64  // start sector
	Len    uint64 // data len
	Start  int    // write start index
	Eend   int    // write end index
}

type Element struct {
	next, prev *Element

	list *List

	Value *Unit
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev returns the previous list element or nil.
func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

type List struct {
	root Element
	len  int
}

// 初始化哨兵节点或者清空l.
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// 初始化链表的哨兵节点
func New() *List { return new(List).Init() }

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *List) Len() int { return l.len }

// Front returns the first element of list l or nil if the list is empty.
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Back() *Element {
	// 先判断链表是否为空
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// 前插v–》 新节点插入到root之后
func (l *List) PushFront(v *Unit) *Element {
	// &l.root==》哨兵节点[的内存地址]，然后在哨兵节点后面插入 新节点
	return l.insertValue(v, &l.root)
}

func (l *List) PushBack(v *Unit) *Element {
	return l.insertValue(v, l.root.prev) // 找到哨兵节点的前一个节点的内存地址【也就是当前链表的最后一个节点】
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List) insertValue(v *Unit, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// insert inserts e after at, increments l.len, and returns e.
func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

func (l *List) InsertBefore(v *Unit, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark.prev)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List) InsertAfter(v *Unit, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark)
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *List) Remove(e *Element) any {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}

// remove removes e from its list, decrements l.len
func (l *List) remove(e *Element) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
}
