/*
LCR 184.中 设计自助结算系统

请设计一个自助结账系统，该系统需要通过一个队列来模拟顾客通过购物车的结算过程，需要实现的功能有：

get_max()：获取结算商品中的最高价格，如果队列为空，则返回 -1
add(value)：将价格为 value 的商品加入待结算商品队列的尾部
remove()：移除第一个待结算的商品价格，如果队列为空，则返回 -1
注意，为保证该系统运转高效性，以上函数的均摊时间复杂度均为 O(1)

示例 1：

输入:
["Checkout","add","add","get_max","remove","get_max"]
[[],[4],[7],[],[],[]]

输出: [null,null,null,7,4,7]
示例 2：

输入:
["Checkout","remove","get_max"]
[[],[],[]]

输出: [null,-1,-1]

提示：

1 <= get_max, add, remove 的总操作数 <= 10000
1 <= value <= 10^5
*/
package main

import (
	"fmt"
	"testing"
)

func TestCheckout(t *testing.T) {
	c := Constructor()

	p := &c
	p.Add(4)
	p.Add(7)
	fmt.Println(p.Get_max())
	p.Remove()
	fmt.Println(p.Get_max())
}

type Checkout struct {
	q []int // raw values
	p []int // 辅助栈: 左, 小-> 右, 大
}

func Constructor() Checkout {
	return Checkout{}
}

func (this *Checkout) Get_max() int {
	if len(this.q) == 0 {
		return -1
	}
	return this.p[0]
}

func (this *Checkout) Add(value int) {
	this.q = append(this.q, value)
	//fmt.Println("b", value, this.p)
	for len(this.p) > 0 && value > this.p[len(this.p)-1] {
		this.p = this.p[:len(this.p)-1]
	}
	this.p = append(this.p, value)
	//fmt.Println("a", value, this.p)
}

func (this *Checkout) Remove() int {
	if len(this.q) == 0 {
		return -1
	}
	if this.q[0] == this.p[0] {
		this.p = this.p[1:]
	}
	value := this.q[0]
	this.q = this.q[1:]
	return value
}
