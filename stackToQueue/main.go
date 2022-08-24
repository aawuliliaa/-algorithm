package main

import "fmt"

type stack struct {
	data []interface{}
	size int
}

func newStack(cp int) *stack {
	return &stack{
		data: make([]interface{}, 0, cp),
		size: 0,
	}
}
func (s *stack) Push(e interface{}) {
	s.data = append(s.data, e)
	s.size++
}
func (s *stack) Pop() interface{} {
	if s.Size() == 0 {
		panic("stack is empty")
	}
	e := s.data[s.size-1]
	s.data = s.data[:s.size-1]
	s.size--
	return e
}
func (s *stack) Size() int {
	return s.size
}
func (s *stack) IsEmpty() bool {
	return s.size == 0
}

type MyQueue struct {
	s1 *stack
	s2 *stack
}

func Constructor() MyQueue {
	return MyQueue{s1: newStack(3), s2: newStack(3)}
}

func (this *MyQueue) Push(x int) {
	this.s1.Push(x)
}

//从队列开头移除并返回
func (this *MyQueue) Pop() int {
	if this.s1.IsEmpty() {
		panic("empty queue")
	}
	if this.s1.size == 1 {
		return this.s1.Pop().(int)
	}
	//先挪到s2
	s1Size := this.s1.size
	for i := 1; i < s1Size; i++ {
		this.s2.Push(this.s1.Pop())
	}
	//获取s1中的最底下的元素，返回
	e := this.s1.Pop().(int)
	//再把s2中的元素挪回到s1
	s2Size := this.s2.size
	for i := 1; i <= s2Size; i++ {
		this.s1.Push(this.s2.Pop())
	}
	return e
}

//返回队列开头的元素
func (this *MyQueue) Peek() int {
	s1Size := this.s1.size
	for i := 1; i < s1Size; i++ {
		this.s2.Push(this.s1.Pop())
	}
	//获取s1中的最底下的元素，返回
	e := this.s1.Pop().(int)
	this.s1.Push(e)
	//再把s2中的元素挪回到s1
	s2Size := this.s2.size
	for i := 1; i <= s2Size; i++ {
		this.s1.Push(this.s2.Pop())
	}
	return e

}

func (this *MyQueue) Empty() bool {
	return this.s1.size == 0
}

func main() {

	q := Constructor()

	for _, i := range []int{1, 2, 3, 4, 5, 6, 7} {
		q.Push(i)
	}
	for i := 0; i < 3; i++ {
		fmt.Println(q.Pop())
	}
	fmt.Println(q.Peek())

}
