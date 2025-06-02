/*
895.困 最大频率栈

设计一个类似堆栈的数据结构，将元素推入堆栈，并从堆栈中弹出出现频率最高的元素。

实现 FreqStack 类:

FreqStack() 构造一个空的堆栈。
void push(int val) 将一个整数 val 压入栈顶。
int pop() 删除并返回堆栈中出现频率最高的元素。
如果出现频率最高的元素不只一个，则移除并返回最接近栈顶的元素。

示例 1：

输入：
["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"],
[[],[5],[7],[5],[7],[4],[5],[],[],[],[]]
输出：[null,null,null,null,null,null,null,5,7,5,4]
解释：
FreqStack = new FreqStack();
freqStack.push (5);//堆栈为 [5]
freqStack.push (7);//堆栈是 [5,7]
freqStack.push (5);//堆栈是 [5,7,5]
freqStack.push (7);//堆栈是 [5,7,5,7]
freqStack.push (4);//堆栈是 [5,7,5,7,4]
freqStack.push (5);//堆栈是 [5,7,5,7,4,5]
freqStack.pop ();//返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,5,7,4]。
freqStack.pop ();//返回 7 ，因为 5 和 7 出现频率最高，但7最接近顶部。堆栈变成 [5,7,5,4]。
freqStack.pop ();//返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,4]。
freqStack.pop ();//返回 4 ，因为 4, 5 和 7 出现频率最高，但 4 是最接近顶部的。堆栈变成 [5,7]。

提示：

0 <= val <= 109
push 和 pop 的操作数不大于 2 * 104。
输入保证在调用 pop 之前堆栈中至少有一个元素。
*/
package main

import (
	"fmt"
	"testing"
)

type FreqStack struct {
	freq    map[int]int   // val->数字出现的次数
	group   map[int][]int // 频率(次数)->具有该频率的元素的映射. 在具有相同的频率的元素中，靠近栈顶的元素总是相对更新一些, 即尾部的元素最新
	maxfreq int           // 栈中任意元素的当前最大频率
}

func Constructor() FreqStack {
	return FreqStack{
		freq:  make(map[int]int, 0),
		group: make(map[int][]int, 0),
	}
}

func (this *FreqStack) Push(x int) {
	n := this.freq[x]
	n++
	this.freq[x] = n

	if ng := this.group[n]; ng != nil {
		ng = append(ng, x)

		this.group[n] = ng
	} else {
		this.group[n] = []int{x}
	}

	if n > this.maxfreq {
		this.maxfreq = n
	}
}

func (this *FreqStack) Pop() int {
	if this.maxfreq == 0 {
		return 0
	}

	ng := this.group[this.maxfreq]
	l := len(ng) - 1

	// 取出当前频率最大且最新的数
	target := ng[l]
	this.group[this.maxfreq] = ng[:l]

	if l == 0 {
		this.maxfreq--
	}

	n := this.freq[target]
	n--
	this.freq[target] = n

	return target
}

func TestConstructor(t *testing.T) {
	nums := []int{4, 0, 9, 3, 4, 2}

	solution := Constructor()
	for _, v := range nums {
		solution.Push(v)
	}

	fmt.Println("---0", solution)

	fmt.Println(solution.Pop())
	solution.Push(6)
	fmt.Println("---1", solution)

	fmt.Println(solution.Pop())
	solution.Push(1)

	fmt.Println("---2", solution)

	fmt.Println(solution.Pop())
	solution.Push(4)

	fmt.Println("---3", solution)
	fmt.Println(solution.Pop())

	fmt.Println("---4", solution)

	fmt.Println(solution.Pop())
	fmt.Println(solution.Pop())
	fmt.Println(solution.Pop())
	fmt.Println(solution.Pop())
	fmt.Println(solution.Pop())

	fmt.Println("---5", solution)
}
