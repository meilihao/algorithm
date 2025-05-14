/*
232.简 用栈实现队列

请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：

实现 MyQueue 类：

void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false
说明：

你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。

示例 1：

输入：
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
输出：
[null, null, null, 1, 1, false]

解释：
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false

提示：

1 <= x <= 9
最多调用 100 次 push、pop、peek 和 empty
假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）
*/
package leetcode

import (
	"fmt"
	"testing"
)

// 思路: 数学的负负得正
// leetcode 环境代码不包含`import "fmt"`报错
// ![](/misc/img/20180527092623978.png)
func Test232(t *testing.T) {
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
		//优化: 栈a中留一个元素供pop, 最后一个元素直接pop(), 这样可以少一次操作
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
