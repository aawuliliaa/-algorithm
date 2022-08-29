package main

type Node struct {
	e    interface{}
	next *Node
}
type LinkList struct {
	dummyHead *Node
	size      int
}

func (l *LinkList) GetSize() int {
	return l.size
}
func (l *LinkList) IsEmpty() bool {
	return l.size == 0
}
func (l *LinkList) add(index int, e interface{}) {
	if index < 0 || index > l.size {
		panic("index illegal")
	}
	pre := l.dummyHead
	for i := 0; i < index; i++ {
		pre = pre.next
	}
	pre.next = &Node{e: e, next: pre.next}
}
func (l *LinkList) addFirst(e interface{}) {
	l.add(0, e)
}
func (l *LinkList) addLast(e interface{}) {
	l.add(l.size, e)
}
func (l *LinkList) get(index int) interface{} {
	if index < 0 || index > l.size {
		panic("index illegal")
	}
	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}
