package main

import "fmt"

type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func New(cp int) *LoopQueue {
	return &LoopQueue{
		data:  make([]interface{}, cp),
		front: 0,
		tail:  0,
		size:  0,
	}
}
func (lq *LoopQueue) Enqueue(e interface{}) {
	if lq.size == len(lq.data) {
		lq.Resize(lq.getCapacity() * 2)
	}
	lq.data[lq.tail] = e
	lq.tail = (lq.tail + 1) % len(lq.data)
	lq.size++

}
func (lq *LoopQueue) Dequeue() interface{} {
	if lq.IsEmpty() {
		panic("data is empty")
	}
	e := lq.data[lq.front]
	lq.data[lq.front] = nil
	lq.front = (lq.front + 1) % len(lq.data)
	lq.size--
	if lq.size == lq.getCapacity()/4 {
		lq.Resize(lq.getCapacity() / 2)
	}
	return e
}
func (lq *LoopQueue) Resize(cp int) {
	newData := make([]interface{}, cp)
	for i := 0; i < lq.size; i++ {
		newData[i] = lq.data[(i+lq.front)%len(lq.data)]
	}
	lq.data = newData
	lq.front = 0
	lq.tail = lq.size
}
func (lq *LoopQueue) GetSize() int {
	return lq.size

}
func (lq *LoopQueue) getCapacity() int {
	return len(lq.data)
}
func (lq *LoopQueue) IsEmpty() bool {
	return lq.size == 0
}
func main() {
	lq := New(3)
	for _, i := range []int{3, 5, 7, 3, 4, 5} {
		lq.Enqueue(i)
		fmt.Println(lq)
	}
	for i := 0; i < 3; i++ {
		lq.Dequeue()
		fmt.Println(lq)
	}

}
