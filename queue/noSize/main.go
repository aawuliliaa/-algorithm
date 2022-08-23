package main

import "fmt"

type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
}

func New(cp int) *LoopQueue {
	return &LoopQueue{
		data:  make([]interface{}, cp+1),
		front: 0,
		tail:  0,
	}
}
func (lq *LoopQueue) Enqueue(e interface{}) {
	if (lq.tail+1)%lq.getCapacity() == lq.front {
		lq.Resize(lq.getCapacity() * 2)
	}
	lq.data[lq.tail] = e
	lq.tail = (lq.tail + 1) % len(lq.data)

}
func (lq *LoopQueue) Dequeue() interface{} {
	if lq.IsEmpty() {
		panic("data is empty")
	}
	e := lq.data[lq.front]
	lq.data[lq.front] = nil
	lq.front = (lq.front + 1) % len(lq.data)
	fmt.Println("size",lq.GetSize())
	fmt.Println("capacity",lq.getCapacity())
	if lq.GetSize() == lq.getCapacity()/4 {
		lq.Resize(lq.getCapacity() / 2)
	}
	return e
}
func (lq *LoopQueue) Resize(cp int) {
	newData := make([]interface{}, cp)
	for i := 0; i < lq.GetSize(); i++ {
		newData[i] = lq.data[(i+lq.front)%len(lq.data)]
	}
	lq.data = newData
	lq.front = 0
	lq.tail = lq.GetSize()
}
func (lq *LoopQueue) GetSize() int {
	if lq.tail >= lq.front {
		return lq.tail - lq.front
	} else {
		return (lq.tail - lq.front) + len(lq.data)
	}

}
func (lq *LoopQueue) getCapacity() int {
	return len(lq.data) - 1
}
func (lq *LoopQueue) IsEmpty() bool {
	return lq.GetSize() == 0
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
