package helper

import "fmt"

type Stack[T any] struct {
	items []T
	max   int
}

// 创建一个新的堆栈
func NewStack[T any](max int) *Stack[T] {
	if max > 0 {
		return &Stack[T]{
			max:   max,
			items: make([]T, 0, max),
		}
	} else {
		return &Stack[T]{}
	}
}

// 将元素推入堆栈
func (s *Stack[T]) Push(element T) bool {
	if s.max > 0 && len(s.items) == s.max {
		return false
	}

	s.items = append(s.items, element)

	return true
}

// 从堆栈弹出元素
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zeroValue T // 如果堆栈为空，返回T的零值
		return zeroValue, false
	}
	lastIndex := len(s.items) - 1
	element := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return element, true
}

// 获取堆栈顶部的元素（不弹出）
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	return s.items[len(s.items)-1], true
}

// 检查堆栈是否为空
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Print() {
	for _, item := range s.items {
		fmt.Print(item)
		fmt.Print(" ")
	}
	fmt.Println("")
}
