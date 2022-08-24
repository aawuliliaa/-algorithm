package main

import "fmt"

type queue struct {
	data []interface{}
	size int
}

func newQueue(cp int) *queue {
	return &queue{
		size: 0,
		data: make([]interface{}, 0, cp),
	}

}
func (q *queue) Push(e interface{}) {
	q.data = append(q.data, e)
	q.size++
}
func (q *queue) Pop() interface{} {
	e := q.data[0]
	q.data = q.data[1:]
	q.size--
	return e
}
func (q *queue) Size() int {
	return q.size
}
func (q *queue) IsEmpty() bool {
	return q.size == 0
}

type MyStack struct {
	q1 *queue
	q2 *queue
}

func Constructor() MyStack {
	return MyStack{
		q1: newQueue(3),
		q2: newQueue(3),
	}
}

func (this *MyStack) Push(x int) {
	this.q1.Push(x)

}

func (this *MyStack) Pop() int {
	if this.q1.IsEmpty() {
		panic("queue is empty")
	}
	q1Size := this.q1.size

	for i := 1; i < q1Size; i++ {
		this.q2.Push(this.q1.Pop())
	}
	e := this.q1.Pop().(int)
	q2Size := this.q2.size
	for i := 1; i <= q2Size; i++ {
		this.q1.Push(this.q2.Pop())
	}

	return e

}

func (this *MyStack) Top() int {
	if this.q1.IsEmpty() {
		panic("queue is empty")
	}
	q1Size := this.q1.size

	for i := 1; i < q1Size; i++ {
		this.q2.Push(this.q1.Pop())
	}
	e := this.q1.Pop().(int)
	this.q2.Push(e)
	q2Size := this.q2.size
	for i := 1; i <= q2Size; i++ {
		this.q1.Push(this.q2.Pop())
	}
	return e
}

func (this *MyStack) Empty() bool {
	return this.q1.size == 0
}
func main() {
	//q := newQueue(3)
	//for i := 0; i < 10; i++ {
	//	q.Push(i)
	//}
	//for i := 0; i < 3; i++ {
	//	fmt.Println(q.Pop())
	//}
	myStack := Constructor()
	myStack.Push(1)
	myStack.Push(2)
	fmt.Println(myStack.Top())   // 返回 2
	fmt.Println(myStack.Top())  // 返回 2
	fmt.Println(myStack.Pop())  // 返回 2
	fmt.Println(myStack.q1)
	fmt.Println(myStack.Empty())// 返回 False

}
