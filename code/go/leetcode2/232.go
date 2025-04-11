// 232. 用栈实现队列
// 思路: 数学的负负得正
// leetcode 环境代码不包含`import "fmt"`报错
// ![](/misc/img/20180527092623978.png)
package main

import "fmt"

func main() {
	queue := Constructor()

	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())  // 返回 1
	fmt.Println(queue.Pop())   // 返回 1
	fmt.Println(queue.Empty()) // 返回 false
}

// MyQueue defines Input queue by two stacks
// push 均进入Input
// peek/pop 将Input转入Output即可
type MyQueue struct {
	Input, Output *stack
}

// Constructor Initialize your data structure here.
func Constructor() MyQueue {
	return MyQueue{
		Input:  newStack(),
		Output: newStack(),
	}
}

// Push element x to the back of queue.
func (queue *MyQueue) Push(x int) {
	queue.Input.push(x)
}

// Pop Removes the element from in front of queue and returns that element.
func (queue *MyQueue) Pop() int {
	if queue.Output.isEmpty() {
		//优化: 栈a中留一个元素供pop,可以少一次操作
		for queue.Input.len() > 1 {
			queue.Output.push(queue.Input.pop())
		}
		return queue.Input.pop()
	}
	return queue.Output.pop()
}

// Peek Get the front element.
func (queue *MyQueue) Peek() int {
	if queue.Output.isEmpty() {
		if queue.Input.isEmpty() {
			return -1
		}
		return queue.Input.nums[0] // 即查看Input栈的栈底元素
	}
	return queue.Output.nums[queue.Output.len()-1]
}

// Empty Returns whether the queue is empty.
func (queue *MyQueue) Empty() bool {
	return queue.Input.isEmpty() && queue.Output.isEmpty()
}

// stack defines stack
type stack struct {
	nums []int
}

// newStack creates a empty stack
func newStack() *stack {
	return &stack{
		nums: []int{},
	}
}

func (s *stack) push(n int) {
	s.nums = append(s.nums, n)
}

func (s *stack) pop() int {
	if s.isEmpty() {
		return -1
	}
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

func (s *stack) len() int {
	return len(s.nums)
}
func (s *stack) isEmpty() bool {
	return s.len() == 0
}
