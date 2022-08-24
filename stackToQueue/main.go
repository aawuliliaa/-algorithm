package main

import "fmt"

type stack struct {
	data []interface{}
	size int
}

func newStack(cp int) *stack {
	return &stack{
		data: make([]interface{}, 0,cp),
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

	s.data[s.size-1] = nil
	s.size--
	return e
}
func (s *stack) Size() int {
	return s.size
}
func (s *stack) IsEmpty() bool {
	return s.size == 0
}
func main() {
	s := newStack(3)
	for _,i :=range []int{1,2,3,4,5,6,7}{
		s.Push(i)
	}
	fmt.Println(s)
	for i:=0;i<3;i++{
		fmt.Println(s.Pop())
	}
}
