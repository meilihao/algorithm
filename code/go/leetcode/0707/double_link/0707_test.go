package doubelink

type node struct {
	val        int
	next, prev *node
}

type MyLinkedList struct {
	head, tail *node
	size       int
}

func Constructor() MyLinkedList {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head
	return MyLinkedList{head, tail, 0}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	var curr *node
	if index+1 < l.size-index { // 判断正向遍历(index+1次)还是逆向遍历(l.size-index次)更快
		curr = l.head
		for i := 0; i <= index; i++ {
			curr = curr.next
		}
	} else {
		curr = l.tail
		for i := 0; i < l.size-index; i++ {
			curr = curr.prev
		}
	}
	return curr.val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

func (l *MyLinkedList) AddAtIndex(index, val int) {
	if index > l.size {
		return
	}
	index = max(0, index)
	var pred, succ *node      // pred=插入位置的pre node;  succ= 原插入位置的节点
	if index < l.size-index { // 正逆向找pred方便
		pred = l.head
		for i := 0; i < index; i++ {
			pred = pred.next
		}
		succ = pred.next
	} else {
		succ = l.tail
		for i := 0; i < l.size-index; i++ {
			succ = succ.prev
		}
		pred = succ.prev
	}
	l.size++
	toAdd := &node{val, succ, pred}
	pred.next = toAdd
	succ.prev = toAdd
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}
	var pred, succ *node // pred=删除位置的pre node;  succ= 删除位置的next节点
	if index < l.size-index {
		pred = l.head
		for i := 0; i < index; i++ {
			pred = pred.next
		}
		succ = pred.next.next
	} else {
		succ = l.tail
		for i := 0; i < l.size-index-1; i++ {
			succ = succ.prev
		}
		pred = succ.prev.prev
	}
	l.size--
	pred.next = succ
	succ.prev = pred
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
