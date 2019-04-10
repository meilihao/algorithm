package main

import "fmt"

func main() {
	fmt.Println(calculate("9/3+5*6/3"))
}

type Stack struct {
	data  []int
	index int // 栈顶指针
	n     int // 容量
}

func NewStack(n int) *Stack {
	return &Stack{
		data:  make([]int, n),
		index: -1,
		n:     n,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.index == -1
}

func (s *Stack) Push(c int) {
	if s.index == s.n-1 {
		panic("full stack")
	}

	s.index++
	s.data[s.index] = c
}

func (s *Stack) Pop() int {
	if s.index == -1 {
		return 0
	}

	tmp := s.data[s.index]
	s.data[s.index] = 0 // 清理旧值
	s.index--

	return tmp
}

func (s *Stack) Peek() int {
	if s.index == -1 {
		return 0
	}

	return s.data[s.index]
}

// 使用数组栈, 容易爆栈
func calculate(s string) int {
	numStack := NewStack(16)
	opStack := NewStack(16)
	tmp := []byte(s)
	n := len(tmp)

	var num int
	for i, v := range tmp {
		if v == ' ' {
			continue
		}

		if v >= '0' && v <= '9' {
			num = num*10 + int(v-'0') // 碰到多位数字
		}

		if v == '+' || v == '-' || v == '*' || v == '/' || i == n-1 {
			numStack.Push(num)
			num = 0

			switch v {
			case '*', '/', '-', '+':
				if calcPriority(int(v), int(opStack.Peek())) {
					opStack.Push(int(v))
				} else {
					numStack.Push(operate(numStack.Pop(), numStack.Pop(), byte(opStack.Pop())))
					opStack.Push(int(v))
				}
			}

			fmt.Println("---num", numStack)
			fmt.Println("---op", opStack)
		}
	}

	for opStack.Peek() != 0 {
		numStack.Push(operate(numStack.Pop(), numStack.Pop(), byte(opStack.Pop())))
	}

	return numStack.Pop()
}

func getPrioriyt(op int) byte {
	if byte(op) == '*' || byte(op) == '/' {
		return 1
	}

	return 0
}

// true为直接压栈
func calcPriority(op1, op2 int) bool {
	if op2 == 0 {
		return true
	}

	return getPrioriyt(op1) > getPrioriyt(op2)
}

// stack是先进后出
func operate(a, b int, opt byte) int {
	switch opt {
	case '+':
		return b + a
	case '-':
		return b - a
	case '*':
		return b * a
	default: // '/'
		return b / a
	}
}

// 基于当前操作符的前一个操作符, 可判断前一个操作符的优先级
// 加减操作很简单，需要处理的就是乘除的情况
// 设置一个栈，如果是+就直接入栈，如果是-，就负数入栈
// 如果是乘或者除，就从栈里取出最后一个元素，运算之后入栈
// 1-2*3-6+1 => 1+(-2)*3+(-6)+1
func calculate1(s string) int {
	var num int
	n := len(s)
	st := []int{}
	var preSign byte = '+' // 当前操作符的前一个操作符/最后一个数字的前操作符

	for i := 0; i < n; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0') // 碰到多位数字

			if i != n-1 { // 假定最后一个是数字,因此s不能以空格结尾
				continue
			}
		}

		if s[i] == ' ' {
			continue
		}
		// 处理操作符
		switch preSign {
		case '+':
		case '-':
			num = -num
		case '*':
			num = st[len(st)-1] * num
			st = st[:len(st)-1]
		case '/':
			num = st[len(st)-1] / num
			st = st[:len(st)-1]
		}

		preSign = s[i]
		st = append(st, num)
		num = 0
	}

	for _, v := range st {
		num += v
	}
	return num
}

// 推荐
func calculate2(s string) int {
	var num int
	n := len(s)
	st := []int{}
	var preSign byte = '+'

	for i := 0; i < n; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0') // 碰到多位数字
		}

		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' || i == n-1 { // 最后一个字符是空格也行
			switch preSign {
			case '+':
			case '-':
				num = -num
			case '*':
				num = st[len(st)-1] * num
				st = st[:len(st)-1]
			case '/':
				num = st[len(st)-1] / num
				st = st[:len(st)-1]
			}

			preSign = s[i]
			st = append(st, num)
			num = 0
		}
	}

	for _, v := range st {
		num += v
	}
	return num
}
