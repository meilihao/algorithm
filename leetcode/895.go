package main

import (
	"fmt"
)

type FreqStack struct {
	freq map[int]int // 统计数字出现的次数
	group map[int][]int // 频率(次数)到具有该频率的元素的映射. 在具有相同的频率的元素中，靠近栈顶的元素总是相对更新一些, 即尾部的元素最新
	maxfreq int // 栈中任意元素的当前最大频率
}

func Constructor() FreqStack {
    return FreqStack{
		freq: make(map[int]int,0),
		group: make(map[int][]int,0),
	}
}

func (this *FreqStack) Push(x int)  {
	n:=this.freq[x]
	n++
	this.freq[x] = n

	if ng:=this.group[n];ng!=nil{
		ng=append(ng,x)

		this.group[n] = ng
	}else{
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

	ng:=this.group[this.maxfreq]
	l:=len(ng)-1

	target:=ng[l]
	this.group[this.maxfreq] = ng[:l]

	if l==0 {
		this.maxfreq--
	}

	n:=this.freq[target]
	n--
	this.freq[target] = n 

	return target
}

func main() {
	nums := []int{4,0,9,3,4,2}

	solution := Constructor()
	for _,v:=range nums{
		solution.Push(v)
	}

	fmt.Println("---0",solution)

	fmt.Println(solution.Pop())
	solution.Push(6)
	fmt.Println("---1",solution)

	fmt.Println(solution.Pop())
	solution.Push(1)

	fmt.Println("---2",solution)

	fmt.Println(solution.Pop())
	solution.Push(4)

	fmt.Println("---3",solution)
	fmt.Println(solution.Pop())

	fmt.Println("---4",solution)

	fmt.Println(solution.Pop())
	fmt.Println(solution.Pop())
	fmt.Println(solution.Pop())
	fmt.Println(solution.Pop())
	fmt.Println(solution.Pop())

	fmt.Println("---5",solution)
}
