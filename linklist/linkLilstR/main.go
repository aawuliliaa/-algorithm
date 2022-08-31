package main

import (
	"fmt"
)

type Node struct {
	e    interface{}
	next *Node
}
type LinkListR struct {
	size int
	head *Node
}

func (l *LinkListR) IsEmpty() bool {
	return l.size == 0
}
func (l *LinkListR) verifyIndex(index int) error {
	if index < 0 || index > l.size {
		return fmt.Errorf("index out of range")
	}
	return nil

}
func (l *LinkListR) Add(index int, e interface{}) {
	err := l.verifyIndex(index)
	if err != nil {
		panic(err)
	}
	l.addR(l.head, index, e)

}
func (l *LinkListR) addR(node *Node, index int, e interface{}) *Node {
	if index == 0 {
		return &Node{e: e, next: node}
	}
	node.next = l.addR(node.next, index-1, e)
	return node
}
func (l *LinkListR) AddFirst(e interface{}) {
	l.Add(0, e)
}
func (l *LinkListR) AddLast(e interface{}) {
	l.Add(l.size, e)
}
func (l *LinkListR) Get(index int) interface{} {
	err := l.verifyIndex(index)
	if err != nil {
		panic(err)
	}
	return l.getR(l.head, index)
}
func (l *LinkListR) getR(node *Node, index int) interface{} {
	if index == 0 {
		return node.e
	}
	e := l.getR(node, index-1)
	return e
}
//func (l *LinkListR) GetFirst() interface{} {
//	return l.Get(0)
//}
//func (l *LinkListR) GetLast() interface{} {
//	return l.Get(l.size - 1)
//}
//func (l *LinkListR) Remove(index int)interface{}{
//
//}
//func (l *LinkListR) removeR(node *Node,index int)interface{}{
//	if index==0{
//
//	}
//}