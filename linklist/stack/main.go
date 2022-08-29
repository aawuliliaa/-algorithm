package main

type LinkListStack struct {
	linkList *LinkList
}

func NewLinkListStack() *LinkListStack {
	return &LinkListStack{linkList: NewLinkList()}
}
func (l *LinkListStack) GetSize() int {
	return l.linkList.size
}
func (l *LinkListStack) IsEmpty() bool {
	return l.linkList.IsEmpty()
}
func (l *LinkListStack) Push(e interface{})  {
	l.linkList.AddFirst(e)
}
func (l *LinkListStack) Pop()  {
	l.linkList.RemoveFirst()
}
func (l *LinkListStack) Peek()  {
	l.linkList.GetFirst()
}
