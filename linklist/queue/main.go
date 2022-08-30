package main

import "fmt"

type Node struct {
	e    interface{}
	next *Node
}
type LinkListQueue struct {
	head *Node
	tail *Node
	size int
}

func New() *LinkListQueue {
	return &LinkListQueue{}
}

func (l *LinkListQueue) Enqueue(e interface{}) {
	newNode := &Node{
		e:    e,
		next: nil,
	}
	if l.tail == nil {

		l.tail = newNode
		l.head = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode

	}
	l.size++
}
func (l *LinkListQueue) Dequeue() interface{} {
	if l.size == 0 {
		panic("empty queue")
	}

	item := l.head.e
	if l.size == 1 {

		l.head = nil
		l.tail = nil
	} else {
		head := l.head
		l.head = l.head.next
		head.next = nil
	}
	l.size--
	return item
}
func (l *LinkListQueue) GetFront() interface{} {
	if l.size == 0 {
		return nil
	}
	return l.head.e

}
func (l *LinkListQueue) String() string {
	str := ""
	node := l.head
	for node != nil {
		str = str + fmt.Sprintf("%s-->", node.e)
		node = node.next
	}
	return str

}
func main() {
	q := New()
	for i := 0; i < 3; i++ {
		q.Enqueue(i)
		fmt.Println(q)
	}
	fmt.Println(q.GetFront())
	q.Dequeue()
	fmt.Println(q)

}


