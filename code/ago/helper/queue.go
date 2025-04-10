package helper

type Queue[T any] struct {
	items []T
	max   int
}

func NewQueue[T any](max int) *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0, max),
		max:   max,
	}
}

func (q *Queue[T]) Push(n T) {
	q.items = append(q.items, n)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Front() (T, bool) {
	if len(q.items) == 0 {
		var zeroValue T
		return zeroValue, false
	}

	return q.items[0], true
}

func (q *Queue[T]) Pop() (T, bool) {
	tmp, isExist := q.Front()

	if n := len(q.items); n > 0 {
		q.items = q.items[1:]
	}

	return tmp, isExist
}
