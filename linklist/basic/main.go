package main

import "fmt"

type Node struct {
	e    interface{}
	next *Node
}

func (n *Node) String() string {
	return fmt.Sprint(n.e)
}

type linkList struct {
	head *Node
	size int
}

func (l *linkList) GetSize() int {
	return l.size
}
func (l *linkList) IsEmpty() bool {
	return l.size == 0
}
func (l *linkList) AddFirst(e interface{}) {
	node := &Node{
		e:    e,
		next: l.head,
	}
	l.head = node
	l.size++
}
func (l *linkList) Add(index int, e interface{}) {
	//头部加元素，不需要计算pre
	if index == 0 {
		l.AddFirst(e)
	} else {
		pre := l.head
		for i := 0; i < index-1; i++ {
			pre = pre.next
		}
		newNode := &Node{
			e:    e,
			next: pre.next,
		}
		pre.next = newNode
		l.size++
	}
}
func (l *linkList) AddLast(e interface{}) {
	l.Add(l.size, e)
}
