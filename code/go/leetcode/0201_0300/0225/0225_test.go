/*
225.简 用队列实现栈

请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。

实现 MyStack 类：

void push(int x) 将元素 x 压入栈顶。
int pop() 移除并返回栈顶元素。
int top() 返回栈顶元素。
boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。

注意：

你只能使用队列的标准操作 —— 也就是 push to back、peek/pop from front、size 和 is empty 这些操作。
你所使用的语言也许不支持队列。 你可以使用 list （列表）或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。

示例：

输入：
["MyStack", "push", "push", "top", "pop", "empty"]
[[], [1], [2], [], [], []]
输出：
[null, null, null, 2, 2, false]

解释：
MyStack myStack = new MyStack();
myStack.push(1);
myStack.push(2);
myStack.top(); // 返回 2
myStack.pop(); // 返回 2
myStack.empty(); // 返回 False

提示：

1 <= x <= 9
最多调用100 次 push、pop、top 和 empty
每次调用 pop 和 top 都保证栈不为空
*/
package leetcode

import (
	"fmt"
	"testing"
)

// 思路: 使用两个队列. 这两个队列中始终有一个是空的,另一个非空. push添加元素到非空队列中，pop把非空队列中前面的元素都转移到另一个队列中，只剩最后一个元素，再把最后一个元素pop出来. 这样这一个队列是空的，另一个队列又非空了.
// 基本思想如上图所示：
// 1. 添加数据，直接在queue栈添加
// 2. 移除数据，将queue栈的前n-1个元素添加到help栈，将第n个元素移除(作为栈顶元素弹出)，然后修改引用
// ![](/misc/img/20180527105637787.png)
func TestMyStack(t *testing.T) {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Top())   // 返回 2
	fmt.Println(s.Pop())   // 返回 2
	fmt.Println(s.Empty()) // 返回 false
}

// MyStack 是用 Queue 实现的 栈
type MyStack struct {
	queue *Queue // 非空队列
	help  *Queue // 辅助队列, pop时将queue中除最后一个数据外的其他数据转移到这里, 再将queue的最后一个数据弹出即可, 最后交换queue和help
}

// Constructor Initialize your data structure here.
func Constructor() MyStack {
	return MyStack{queue: NewQueue(), help: NewQueue()}
}

// Push Push element x onto stack.
func (ms *MyStack) Push(x int) {
	ms.queue.Push(x)
}

// Pop Removes the element on top of the stack and returns that element.
func (ms *MyStack) Pop() int {
	if ms.queue.Len() == 0 {
		return -1
	}

	for ms.queue.Len() > 1 {
		ms.help.Push(ms.queue.Pop())
	}

	res := ms.queue.Pop()

	ms.queue, ms.help = ms.help, ms.queue

	return res
}

// Top Get the top element.
func (ms *MyStack) Top() int {
	if ms.queue.Len() == 0 {
		return -1
	}

	for ms.queue.Len() > 1 {
		ms.help.Push(ms.queue.Pop())
	}

	res := ms.queue.Pop()
	ms.help.Push(res) // 重新加入队列

	ms.queue, ms.help = ms.help, ms.queue

	return res
}

// Empty Returns whether the stack is empty.
func (ms *MyStack) Empty() bool {
	return ms.queue.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

// Queue 是用于存放 int 的队列
type Queue struct {
	nums []int
}

// NewQueue 返回 *kit.Queue
func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}

// Push 把 n 放入队列
func (q *Queue) Push(n int) {
	q.nums = append(q.nums, n)
}

// Pop 从 q 中取出最先进入队列的值
func (q *Queue) Pop() int {
	res := q.nums[0]
	q.nums = q.nums[1:]
	return res
}

// Len 返回 q 的长度
func (q *Queue) Len() int {
	return len(q.nums)
}

// IsEmpty 反馈 q 是否为空
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
